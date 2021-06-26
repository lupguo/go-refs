package model

import (
	"fmt"
	"testing"
)

func TestConn2(t *testing.T) {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		cfgDB.user,
		cfgDB.password,
		cfgDB.host,
		cfgDB.database,
		cfgDB.charset,
		cfgDB.parseTime,
		cfgDB.loc,
	)
	t.Log(dsn)
}
