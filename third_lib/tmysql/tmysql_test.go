package tmysql

import (
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	dns := `utuser:Secret@123@(9.134.233.187:3306)/conn_pool?timeout=3s&readTimeout=3s&writeTimeout=3s`
	setting := &DBSetting{
		ConnMaxLifeTime: 1 * time.Hour,
		ConnMaxIdleTime: 30 * time.Second,
		MaxOpenConns:    100,
		MaxIdleConns:    10,
	}
	bchSetting := &BchSetting{
		WriteClientNum: 1,
		WriteInterval:  10 * time.Millisecond,
		ReadClientNum:  1,
		ReadInternal:   10 * time.Millisecond,
	}
	if err := StartServer(dns, setting, bchSetting); err != nil {
		t.Errorf("StartServer() error = %v", err)
	}
}
