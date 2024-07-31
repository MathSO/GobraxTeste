package models

import (
	"database/sql"
	"fmt"

	"github.com/MathSO/GobraxTeste/factory"
	"github.com/google/uuid"
)

const TruckTableName = `truck`

type Truck struct {
	ID    string `json:"id" insert:"id"`
	Model string `json:"model" insert:"model"`
	Plate string `json:"plate" insert:"plate"`
}

func InsertTruck(mysqlConnection *sql.DB, t *Truck) error {
	t.ID = uuid.NewString()
	_, err := factory.Insert(mysqlConnection, *t, TruckTableName)
	return err
}

func RetrieveTruck(mysqlConnection *sql.DB, id string) (Truck, error) {
	var t Truck

	query := fmt.Sprintf("SELECT * FROM `%s` WHERE `id` = ?", TruckTableName)
	err := mysqlConnection.QueryRow(query, id).Scan(&t.ID, &t.Model, &t.Plate)

	return t, err
}

func UpdateTruck(mysqlConnection *sql.DB, d Truck) error {
	query := fmt.Sprintf("UPDATE `%s` SET `model` = ?, `plate` = ? WHERE `id` = ?", TruckTableName)
	result, err := mysqlConnection.Exec(query, d.Model, d.Plate, d.ID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return sql.ErrNoRows
	}

	return err
}

func DeleteTruck(mysqlConnection *sql.DB, id string) error {
	query := fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", TruckTableName)
	_, err := mysqlConnection.Exec(query, id)
	return err
}
