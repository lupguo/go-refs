package dbs

import (
	"context"

	"git.code.oa.com/trpc-go/trpc-go/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type dbInfra struct {
	DB *gorm.DB
}

// NewDbInfra 初始化MysqlDB实例
func NewDbInfra() *dbInfra {
	// get config from rainbow
	cfg, err := rainbowcfg.GetMysqlConfig()
	if err != nil {
		log.Fatalf("NewDbInfra() got err: %s", err)
	}
	// new mysql instance
	gormDB, err := database.NewMysqlDB(cfg)
	if err != nil {
		log.Fatalf("NewMysqlDB() got err: %s", err)
	}

	return &dbInfra{
		DB: gormDB,
	}
}

// SelectAuthorizeRecord 查询老师认证授权记录
func (d *dbInfra) SelectAuthorizeRecord(ctx context.Context, uid uint64) (*entity.AuthorizeRecord, error) {
	record := &entity.AuthorizeRecord{}
	err := d.DB.WithContext(ctx).
		Model(&entity.AuthorizeRecord{}).
		Where("uid=?", uid).
		Last(record).Error
	if err != nil {
		err := errors.Wrapf(err, "db select authorize record")
		log.ErrorContext(ctx, err)
		return nil, err
	}
	return record, nil
}

// InsertAuthorizeRecord 插入认证授权记录
func (d *dbInfra) InsertAuthorizeRecord(ctx context.Context, record *entity.AuthorizeRecord) error {
	err := d.DB.WithContext(ctx).Create(record).Error
	if err != nil {
		err := errors.Wrapf(err, "db create authorize record")
		log.ErrorContext(ctx, err)
		return err
	}
	return nil
}
