package model

type ConfigDb struct {
	host      string
	user      string
	password  string
	database  string
	charset   string
	parseTime bool
	loc       string
	logMode   bool
}

var (
	cfgDB = &ConfigDb{
		host:      "127.0.0.1",
		user:      "user",
		password:  "your-pass",
		database:  "your-db",
		charset:   "utf8mb4",
		parseTime: true,
		loc:       "Local",
		logMode:   true,
	}
)
