package models

import (
	"database/sql"
	"testing"
)

func InsertDriver(conn *sql.DB) (Model, error) {
	driver := NewDriver("", "jonas", "12453532125")
	err := driver.Insert(conn)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func InsertTruck(conn *sql.DB) (Model, error) {
	truck := NewTruck("", "Volvo", "AAA1234")
	err := truck.Insert(conn)
	if err != nil {
		return nil, err
	}

	return truck, nil
}

func TestDriverTruckCreationSuccessAndDelete(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	driver, err := InsertDriver(conn)
	if err != nil {
		t.Skip()
	}
	defer driver.Delete(conn)

	truck, err := InsertTruck(conn)
	if err != nil {
		t.Skip()
	}
	defer truck.Delete(conn)

	testes := []Model{
		NewDriverTruck(driver.GetID()[0], truck.GetID()[0]),
		NewDriverTruck(driver.GetID()[0], truck.GetID()[0]),
	}

	for _, d := range testes {
		if err := d.Insert(conn); err != nil {
			t.Fatal(err)
		}

		d.Delete(conn)
	}
}

func TestDriverTruckCreationFailure(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := NewDriverTruck(
		"Joao",
		"53242343224",
	)

	if err := a.Insert(conn); err == nil {
		t.Fail()
	}
	defer a.Delete(conn)
}

func TestDriverTruckUpdateFailure(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	driver, err := InsertDriver(conn)
	if err != nil {
		t.Skip()
	}
	defer driver.Delete(conn)

	truck, err := InsertTruck(conn)
	if err != nil {
		t.Skip()
	}
	defer truck.Delete(conn)

	a := NewDriverTruck(
		driver.GetID()[0],
		truck.GetID()[0],
	)

	if err := a.Insert(conn); err != nil {
		t.Fail()
	}
	defer a.Delete(conn)

	err = SetAttribute(a, "DriverID", "53242433224")
	if err != nil {
		t.Fatal(err)
	}

	if err := a.Update(conn); err == nil {
		t.Fail()
	}
}
