package models

import (
	"database/sql"
	"fmt"

	"github.com/MathSO/GobraxTeste/factory"
	"github.com/google/uuid"
)

const DriverTableName = `driver`

type Driver struct {
	ID   string `json:"id" insert:"id"`
	Name string `json:"name" insert:"name"`
	CNH  string `json:"cnh" insert:"cnh"`
}

func InsertDriver(mysqlConnection *sql.DB, d *Driver) error {
	d.ID = uuid.NewString()
	_, err := factory.Insert(mysqlConnection, *d, DriverTableName)

	return err
}

func RetrieveDriver(mysqlConnection *sql.DB, id string) (Driver, error) {
	var d Driver

	query := fmt.Sprintf("SELECT * FROM `%s` WHERE `id` = ?", DriverTableName)
	err := mysqlConnection.QueryRow(query, id).Scan(&d.ID, &d.Name, &d.CNH)

	return d, err
}

func UpdateDriver(mysqlConnection *sql.DB, d Driver) error {
	query := fmt.Sprintf("UPDATE `%s` SET `name` = ?, `cnh` = ? WHERE `id` = ?", DriverTableName)
	result, err := mysqlConnection.Exec(query, d.Name, d.CNH, d.ID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return sql.ErrNoRows
	}

	return err
}

func DeleteDriver(mysqlConnection *sql.DB, id string) error {
	query := fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", DriverTableName)
	_, err := mysqlConnection.Exec(query, id)
	return err
}
