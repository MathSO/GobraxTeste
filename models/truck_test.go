package models

import (
	"testing"
)

func TestTruckCreationSuccessAndDelete(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	testes := []Truck{
		{
			Model: "VOLVO FH-540 GLOBETROTTER 6X4",
			Plate: "BBB1234",
		},
		{
			Model: "VOLVO FH-540 6x4",
			Plate: "AAA4444",
		},
	}

	for _, d := range testes {
		if err := InsertTruck(conn, &d); err != nil {
			t.Fatal(err)
		}

		DeleteTruck(conn, d.ID)
	}
}

func TestTruckCreationFailure(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := Truck{
		Model: "VOLVO FH-540 GLOBETROTTER 6X4",
		Plate: "BBB1234",
	}
	b := Truck{
		Model: "VOLVO FH-540 GLOBETROTTER 6X4",
		Plate: "BBB1234",
	}

	if err := InsertTruck(conn, &a); err != nil {
		t.Fatal(err)
	}
	defer DeleteTruck(conn, a.ID)

	if err := InsertTruck(conn, &b); err == nil {
		t.Fail()
	}
}

func TestTruckUpdateSuccess(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := Truck{
		Model: "VOLVO FH-540 GLOBETROTTER 6X4",
		Plate: "BBB1234",
	}

	if err := InsertTruck(conn, &a); err != nil {
		t.Fatal(err)
	}
	defer DeleteTruck(conn, a.ID)

	a.Model = "VOLVO FH-540 6x4"

	if err := UpdateTruck(conn, a); err != nil {
		t.Fatal(err)
	}
}

func TestTruckUpdateFailure(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := Truck{
		ID:    "12345",
		Model: "VOLVO FH-540 GLOBETROTTER 6X4",
		Plate: "BBB1234",
	}

	if err := UpdateTruck(conn, a); err == nil {
		t.Fail()
	}
}
