package config

import (
	"github.com/kelseyhightower/envconfig"
)

var Config = generate()

type config struct {
	JWTAccessSecret   string `envconfig:"JWT_ACCESS_SECRET"`
	JWTRefreshSecret  string `envconfig:"JWT_REFRESH_SECRET"`
	RedisSecurityHost string `envconfig:"REDIS_SECURITY_HOST"`
	RedisHost         string `envconfig:"REDIS_HOST"`
	MysqlReaderHost   string `envconfig:"MYSQL_READER_HOST"`
	MysqlReaderUser   string `envconfig:"MYSQL_READER_USERNAME"`
	MysqlReaderPass   string `envconfig:"MYSQL_READER_PASSWORD"`
	MysqlWriterHost   string `envconfig:"MYSQL_WRITER_HOST"`
	MysqlWriterUser   string `envconfig:"MYSQL_WRITER_USERNAME"`
	MysqlWriterPass   string `envconfig:"MYSQL_WRITER_PASSWORD"`
	PasswordHashSalt  string `envconfig:"PASSWORD_HASH_SALT"`
	DebugBoiler       bool   `envconfig:"DEBUG_BOILER"`
	SlackEnabled      bool   `envconfig:"SLACK_ENABLED"`
	SlackMainWebHook  string `envconfig:"SLACK_MAIN_WEB_HOOK"`
	SlackMainChannel  string `envconfig:"SLACK_MAIN_WEB_CHANNEL"`
}

func generate() *config {
	c := &config{
		JWTAccessSecret:   "JWTAccessSecret",
		JWTRefreshSecret:  "JWTRefreshSecret",
		RedisSecurityHost: "127.0.0.1:6379",
		RedisHost:         "127.0.0.1:6379",
		MysqlReaderHost:   "127.0.0.1:3306",
		MysqlReaderUser:   "root",
		MysqlReaderPass:   "",
		MysqlWriterHost:   "127.0.0.1:3306",
		MysqlWriterUser:   "root",
		MysqlWriterPass:   "",
		PasswordHashSalt:  "veryverysalty",
		DebugBoiler:       true,
		SlackEnabled:      true,
		SlackMainWebHook:  "",
		SlackMainChannel:  "canberk-test",
	}

	c.loadENV()
	return c
}

func (c *config) loadENV() {
	if err := envconfig.Process("", c); err != nil {
		panic(err)
	}
}
