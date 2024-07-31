package factory

import (
	"database/sql"
	"fmt"
	"strconv"
)

// Generic get all paginated
func GetAll(mysqlConnection *sql.DB, tableName, page, perPage string) (*sql.Rows, error) {
	pg, err := strconv.Atoi(page)
	if err != nil || pg <= 0 {
		pg = 1
	}

	pp, err := strconv.Atoi(perPage)
	if err != nil || pg <= 0 {
		pp = 30
	}
	if pp > 100 {
		pp = 100
	}

	sql := fmt.Sprintf("SELECT * FROM `%s` LIMIT ? OFFSET ?;", tableName)

	return mysqlConnection.Query(sql, pp, pp*(pg-1))
}
