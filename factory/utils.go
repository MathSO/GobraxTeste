package factory

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

func Insert(conn *sql.DB, value any, tableName string) (sql.Result, error) {
	columns, values := getColumnsAndValues(value)
	query := fmt.Sprintf("INSERT INTO `%s`(%s) VALUES(?,?,?)", tableName, strings.Join(columns, ","))

	return conn.Exec(query, values...)
}

func getColumnsAndValues(value any) (columns []string, values []any) {
	reflectType := reflect.TypeOf(value)
	reflectValue := reflect.ValueOf(value)

	columns = make([]string, 0)
	values = make([]any, 0)

	switch reflectType.Kind() {
	case reflect.Struct:
		for i := 0; i < reflectType.NumField(); i++ {
			field := reflectType.Field(i)

			if col := field.Tag.Get(`insert`); col != "" {
				val := reflectValue.Field(i).Interface()

				values = append(values, val)
				columns = append(columns, fmt.Sprintf("`%s`", col))
			}
		}
	default:
		panic("value must be struct kind")
	}

	return columns, values
}
