package models

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

const TruckTableName = `truck`

type truck struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Plate string `json:"plate"`
}

func NewTruck(id, brand, plate string) Model {
	return &truck{
		ID:    id,
		Brand: brand,
		Plate: plate,
	}
}

func (t *truck) Insert(mysqlConnection *sql.DB) error {
	t.ID = uuid.NewString()

	query := fmt.Sprintf("INSERT INTO `%s` (`id`, `brand`, `plate`) VALUES (?, ?, ?)", TruckTableName)
	_, err := mysqlConnection.Exec(query, t.ID, t.Brand, t.Plate)

	return err
}

func (t *truck) Load(mysqlConnection *sql.DB, id ...string) error {
	query := fmt.Sprintf("SELECT `id`, `brand`, `plate` FROM `%s` WHERE `id` = ?", TruckTableName)
	err := mysqlConnection.QueryRow(query, id[0]).Scan(&t.ID, &t.Brand, &t.Plate)

	return err
}

func (t truck) Update(mysqlConnection *sql.DB) error {
	query := fmt.Sprintf("UPDATE `%s` SET `brand` = ?, `plate` = ? WHERE `id` = ?", TruckTableName)
	result, err := mysqlConnection.Exec(query, t.Brand, t.Plate, t.ID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (t truck) Delete(mysqlConnection *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM `%s` WHERE `id` = ?", TruckTableName)
	_, err := mysqlConnection.Exec(query, t.ID)
	return err
}

func (t truck) GetID() []string {
	return []string{t.ID}
}
