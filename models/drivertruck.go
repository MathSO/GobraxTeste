package models

import (
	"database/sql"
	"fmt"
)

const DriveTruckTableName = `driver_truck`

type driverTruck struct {
	DriverID string `json:"driver_id"`
	TruckID  string `json:"truck_id"`
}

func NewDriverTruck(driverID string, truckID string) Model {
	return &driverTruck{
		DriverID: driverID,
		TruckID:  truckID,
	}
}

func (dt *driverTruck) Insert(mysqlConnection *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO `%s` (`id_driver`, `id_truck`) VALUES (?, ?)", DriveTruckTableName)
	_, err := mysqlConnection.Exec(query, dt.DriverID, dt.TruckID)

	return err
}

func (dt *driverTruck) Load(mysqlConnection *sql.DB, ids ...string) error {
	if len(ids) < 2 {
		return fmt.Errorf("driver_truck load missing id")
	}

	query := fmt.Sprintf("SELECT `id_driver`, `id_truck` FROM `%s` WHERE `id_driver` = ? AND `id_truck` = ?", DriveTruckTableName)
	err := mysqlConnection.QueryRow(query, ids[0], ids[1]).Scan(&dt.DriverID, &dt.TruckID)

	return err
}

func (dt driverTruck) Update(mysqlConnection *sql.DB) error {
	return fmt.Errorf("can not update table")
}

func (dt driverTruck) Delete(mysqlConnection *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM `%s` WHERE `id_driver` = ? AND `id_truck` = ?", DriveTruckTableName)
	_, err := mysqlConnection.Exec(query, dt.DriverID, dt.TruckID)
	return err
}

func (d driverTruck) GetID() []string {
	return []string{d.DriverID, d.TruckID}
}
