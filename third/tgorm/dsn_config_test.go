package tgorm

import (
	"encoding/json"
	"testing"

	"github.com/go-sql-driver/mysql"
)

func TestDsnConfig(t *testing.T) {
	dbConfig := &mysql.Config{
		User:   "root",
		Passwd: "Secrect123",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "edu_account_bind",
		Params: map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "True",
			"loc":       "Local",
		},
	}
	marshal, err := json.MarshalIndent(dbConfig, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", marshal)
	t.Log(dbConfig.FormatDSN())
}

func TestParseConfig(t *testing.T) {
	confStr := `{
        	"User": "root",
        	"Passwd": "Secrect123.",
        	"Net": "tcp",
        	"Addr": "127.0.0.1:3306",
        	"DBName": "edu_account_bind",
        	"Params": {
        		"charset": "utf8mb4",
        		"parseTime": "True",
        		"loc": "Local"
        	}
        }
`
	dbConfig := new(mysql.Config)
	err := json.Unmarshal([]byte(confStr), dbConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", dbConfig)
	t.Logf("%s", dbConfig.FormatDSN())
}

func TestName(t *testing.T) {

}
