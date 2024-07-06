package tmysql

import (
	"database/sql"
	_ "database/sql"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type DBSetting struct {
	ConnMaxLifeTime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
}

type BchSetting struct {
	WriteClientNum int
	WriteInterval  time.Duration
	ReadClientNum  int
	ReadInternal   time.Duration
}

func StartServer(dns string, setting *DBSetting, bchSetting *BchSetting) error {
	// 1. 服务启动
	dbCfg, err := mysql.ParseDSN(dns)
	if err != nil {
		return errors.Wrap(err, "parse dns got err")
	}
	log.Infof("db cfg: %+v", dbCfg)

	// 2. DB连接实例
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return errors.Wrapf(err, "sql open got err,")
	}

	// 3. Mysql实例+配置
	// See "Important settings" section.
	db.SetConnMaxLifetime(setting.ConnMaxLifeTime)
	db.SetMaxOpenConns(setting.MaxOpenConns)
	db.SetMaxIdleConns(setting.MaxIdleConns)
	db.SetConnMaxIdleTime(setting.ConnMaxIdleTime)

	// 4. 模拟多协程读、写用户(10w写入,50w读取)
	go srvStart(db, bchSetting)

	// 5. 启动协程，定期输出监控数据 db.Stat
	go srvMonitor(db)

	select {}
}

// 服务监控
func srvMonitor(db *sql.DB) {
	for {
		select {
		case <-time.Tick(1 * time.Second):
			log.Infof("db status: %+v", db.Stats())
			log.Infof("%s", strings.Repeat("-", 50))
			log.Infof("go routine num: %d", runtime.NumGoroutine())
		}
	}
}

// 模拟启动并发读、写协程
func srvStart(db *sql.DB, bchCfg *BchSetting) {
	// 读
	for i := 0; i < bchCfg.ReadClientNum; i++ {
		go func() {
			GetUser(bchCfg, db, uuid.New().String())
		}()
	}

	// 写
	for i := 0; i < bchCfg.WriteClientNum; i++ {
		go func() {
			NewUser(bchCfg, db, randUser())
		}()
	}

}

func NewUser(bcfg *BchSetting, db *sql.DB, u *User) {
	for {
		select {
		case <-time.Tick(bcfg.WriteInterval):
			// log.Infof("insert name[%s]", u.Name)
			_, err := db.Exec("insert into user (name) values(?)", u.Name)
			if err != nil {
				log.Errorf("insert db exec got err: %s", err)
			}
		}

	}
}

func GetUser(bcfg *BchSetting, db *sql.DB, s string) {
	for {
		select {
		case <-time.Tick(bcfg.ReadInternal):
			// rows, err := db.Query("select * from user where name like ?", "%"+s) : 人为改成为慢查询
			rows, err := db.Query("select * from user where name like ?", s)
			if err != nil {
				log.Errorf("read db query got err: %s", err)
				continue
			}
			rows.Close() // 不关闭的话，会有资源死锁问题
		}
	}
}

type User struct {
	ID   int
	Name string
}

func randUser() *User {
	return &User{
		Name: uuid.New().String(),
	}
}
