package tredis

import (
	"reflect"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

/**
Redigo: 整个redigo使用和mysql driver使用非常类似，与redis库比，发送是通用的，但需要自己做返回值判断，多了一层心智负担；
	包括拨号连接的选项支持链式操作，整体下来，感觉redigo还是非常冗余，没有redis直观简单
1. 连接通过`Dial`拨号创建客户端，实现Conn接口与redis通信
2. 应用停止时候必须调用`conn.Close()`关闭连接，防止泄露
3. 通过Do(cmd string, args ...interface{})(reply interface{}, err error)作为通用Redis命令执行方法,cmd为官方命令，比如
	n, err := conn.Do("SET", "key", "hello");
	- Do的会做Go类型转换，都会转成字符串：比如`[]byte`,`string`原样发送，`nil`会转`""`，`bool`会转字符串"1"或"0"，`float64`和`int`会通过`strconv`包转换
	- 回包会将redis类型转成go类型，比如error为`redis.Error`，integer为`int64`，字符串转为string或[]byte或nil，数组转为[]interface{}或nil，需要通过类型断言或者reply组手转成Go类型
4. 管道支持`Send()`,`Flush()`,`Receive()`, Do()方法兼容上述几个方法，即写命令，刷新输出，收到所有Do回包
	- 如果无错，则返回最后一个回包
	- 如果Do的cmd为空字符串，则会直接走flush和收包操作
5. redigo的Conn支持一个并发`Send()`、`Flush()`以及一个并发`Receive()`，没有其他并发支持`Do()`和`Close()`
6. Pub和Sub
7. 返回值助手，redigo很大一部分函数都是组手函数，即通过包装器方式将`Do()`或`Receive()`的返回值和错误包装处理，即返回转换后的Go类型值和错误
	- 若无错，则组手函数返回指定类型的值
	- 若有错，则组手函数返回错误
	- 若返回值是数组，则可以通过`redis.Scan()`
8. Error，通过`conn.Err()`检测不可恢复的错误(比如网络错误)，如果返回非空，应该关闭conn连接池
9. Dial创建Conn连接实例，看doc文档，可以发现Redigo还有很大一块内容就是连接的管理，比如`Dial()`,`DialContext()`,`DialTimeout()`...，包括拨号连接的选项支持链式操作，
	这里整体下来，感觉还是非常冗余，没有redis直观简单
*/

var c redis.Conn

func init() {
	c = dial()
}

func dial() redis.Conn {
	passwd := redis.DialPassword("your-pass")
	connectTimeout := redis.DialConnectTimeout(time.Second)
	db := redis.DialDatabase(1)
	conn, err := redis.Dial("tcp", ":6379", passwd, connectTimeout, db)
	if err != nil {
		panic(err)
	}
	return conn
}

func TestRedigoString(t *testing.T) {
	// 这里的心智负较高，需要将cmd、组手函数断言返回接口
	s, err := redis.String(c.Do("set", "key", "100"))
	assert.Nil(t, err)
	t.Logf("set return:%s", s)

	// +1 断言返回string
	s1, err := redis.String(c.Do("get", "key"))
	assert.Nil(t, err)
	t.Logf("get return:%s, type s1:%v", s1, reflect.TypeOf(s1))

	// +2 断言返回int64，如果value无法无法被转换，则会报错(比如set key 100hello)
	i, err := redis.Int64(c.Do("get", "key"))
	assert.Nil(t, err)
	t.Logf("get return:%d, type i: %v", i, reflect.TypeOf(i))
}

func TestRedigoHash(t *testing.T) {
	s, err := redis.String(c.Do("HMSET", "myhash", "f1", 100, "f2", "200"))
	assert.Nil(t, err)
	t.Logf("hmset return:%s", s)

	m, err := redis.StringMap(c.Do("HGETALL", "myhash"))
	assert.Nil(t, err)
	t.Logf("hgetall %+v, m[f1]:%s, m[f2]:%s", m, m["f1"], m["f2"])

	// hscan
	values, err := redis.Values(c.Do("hmget", "myhash", "f1", "f2", "f3"))
	assert.Nil(t, err)
	t.Logf("redis values of hgetall:%+v", values)

	type dataVal []interface{}
	var f1, f2, f3 string
	scanval, err := redis.Scan(values)
	assert.Nil(t, err)
	t.Logf("redis.Scan(val) %+v, f1:%s, f2:%s, f3:%s", scanval, f1, f2, f3)

	scanval, err = redis.Scan(values, &f1, &f2, &f3)
	assert.Nil(t, err)
	t.Logf("redis.Scan(val, &f1...)  %+v, f1:%s, f2:%s, f3:%s", scanval, f1, f2, f3)

	var dest interface{}
	err = redis.ScanSlice(values, dest)
	assert.Nil(t, err)
	t.Logf("redis.ScanSlice(val, dest):%v", dest)
}

func TestRedigoKeys(t *testing.T) {
	exists, err := redis.Bool(c.Do("EXISTS", "foo"))
	assert.Nil(t, err)
	t.Logf("exist foo return:%t", exists)
}

func TestRedigoPipeline(t *testing.T) {
	// -- 通过Receive收包
	c.Send("SET", "foo1", "100")
	c.Send("SET", "foo2", "200")
	c.Send("SET", "foo3", "300")
	c.Send("GET", "foo1")
	c.Send("GET", "foo2")
	c.Send("GET", "foo3")
	c.Flush()
	c.Receive()            // reply from SET
	v1, err := c.Receive() // reply from GET
	v2, err := c.Receive() // reply from GET
	v3, err := c.Receive() // reply from GET
	assert.Nil(t, err)
	t.Logf("c.Receive() return:%v, %v, %v", v1, v2, v3)

	// -- 通过Do发包
	// c.Send("MULTI")
	// c.Send("INCR", "foo")
	// c.Send("INCR", "bar")
	// r, err := c.Do("EXEC")
	// fmt.Println(r) // prints [1, 1]
}
