package models

import (
	"testing"
)

func TestTruckCreationSuccessAndDelete(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	testes := []Model{
		NewTruck("", "VOLVO FH-540 GLOBETROTTER 6X4", "BBB1234"),
		NewTruck("", "VOLVO FH-540 6x4", "AAA4444"),
	}

	for _, d := range testes {
		if err := d.Insert(conn); err != nil {
			t.Fatal(err)
		}

		d.Delete(conn)
	}
}

func TestTruckCreationFailure(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := NewTruck(
		"",
		"VOLVO FH-540 GLOBETROTTER 6X4",
		"BBB1234",
	)

	b := NewTruck(
		"",
		"VOLVO FH-540 GLOBETROTTER 6X4",
		"BBB1234",
	)

	if err := a.Insert(conn); err != nil {
		t.Fatal(err)
	}
	defer a.Delete(conn)

	if err := b.Insert(conn); err == nil {
		t.Fail()
	}
}

func TestTruckUpdateSuccess(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := NewTruck("", "VOLVO FH-540 GLOBETROTTER 6X4", "BBB1234")

	if err := a.Insert(conn); err != nil {
		t.Fatal(err)
	}
	defer a.Delete(conn)

	SetAttribute(a, "Brand", "VOLVO FH-540 6x4")

	if err := a.Update(conn); err != nil {
		t.Fatal(err)
	}
}

func TestTruckUpdateFailure(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := NewTruck(
		"12345",
		"VOLVO FH-540 GLOBETROTTER 6X4",
		"BBB1234",
	)

	if err := a.Update(conn); err == nil {
		t.Fail()
	}
}

func TestTruckSetAttribute(t *testing.T) {
	truck := NewTruck("", "", "")

	err := SetAttribute(truck, "ID", "12345153")
	if err != nil {
		t.Fatal(err)
	}

	err = SetAttribute(truck, "Brand", "Caminhao")
	if err != nil {
		t.Fatal(err)
	}

	err = SetAttribute(truck, "Plate", "aaa1234")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(truck)
}
