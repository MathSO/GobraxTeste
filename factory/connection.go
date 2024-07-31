package factory

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

var mysqlConnection *sql.DB

func GetConnection() (*sql.DB, error) {
	if mysqlConnection == nil || mysqlConnection.Ping() != nil {
		config := mysql.NewConfig()
		config.Addr = os.Getenv(`MYSQL_HOSTNAME`)
		config.User = os.Getenv(`MYSQL_USERNAME`)
		config.Passwd = os.Getenv(`MYSQL_PASSWORD`)
		config.DBName = os.Getenv(`MYSQL_DATABASE`)
		config.Net = `tcp`

		var err error
		mysqlConnection, err = sql.Open("mysql", config.FormatDSN())

		return mysqlConnection, err
	}

	return mysqlConnection, nil
}
