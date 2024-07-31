package models

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

const DriverTableName = `driver`

type driver struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	CNH  string `json:"cnh"`
}

func NewDriver(id, name, cnh string) Model {
	return &driver{
		ID:   id,
		Name: name,
		CNH:  cnh,
	}
}

func (d *driver) Insert(mysqlConnection *sql.DB) error {
	d.ID = uuid.NewString()

	query := fmt.Sprintf("INSERT INTO `%s` (`id`, `name`, `cnh`) VALUES (?, ?, ?)", DriverTableName)
	_, err := mysqlConnection.Exec(query, d.ID, d.Name, d.CNH)

	return err
}

func (d *driver) Load(mysqlConnection *sql.DB, id ...string) error {
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE `id` = ?", DriverTableName)
	err := mysqlConnection.QueryRow(query, id[0]).Scan(&d.ID, &d.Name, &d.CNH)

	return err
}

func (d driver) Update(mysqlConnection *sql.DB) error {
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

func (d driver) Delete(mysqlConnection *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", DriverTableName)
	_, err := mysqlConnection.Exec(query, d.ID)
	return err
}

func (d driver) GetID() []string {
	return []string{d.ID}
}
