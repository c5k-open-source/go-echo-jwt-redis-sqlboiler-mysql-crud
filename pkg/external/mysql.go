package external

import (
	"database/sql"
	"fmt"

	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/config"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewWriterDB(DBName string) *sql.DB {

	if !utils.ValidateDBName(DBName) {
		panic("DB name failed")
	}
	mysqlURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8", config.Config.MysqlWriterUser, config.Config.MysqlWriterPass, config.Config.MysqlWriterHost, DBName)
	db, err := sql.Open("mysql", mysqlURL)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	if config.Config.DebugBoiler {
		boil.DebugMode = true
	}

	return db
}

func NewReaderDB(DBName string) *sql.DB {

	if !utils.ValidateDBName(DBName) {
		panic("DB name failed")
	}
	mysqlURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8", config.Config.MysqlReaderUser, config.Config.MysqlReaderPass, config.Config.MysqlReaderHost, DBName)
	db, err := sql.Open("mysql", mysqlURL)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	if config.Config.DebugBoiler {
		boil.DebugMode = true
	}

	return db
}
