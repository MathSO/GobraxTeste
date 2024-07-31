package models

import (
	"database/sql"
	"testing"

	"github.com/go-sql-driver/mysql"
)

func getMysqlConnection() (*sql.DB, error) {
	config := mysql.NewConfig()
	config.Addr = "10.1.0.207"
	config.User = "root"
	config.Passwd = "example"
	config.DBName = "gobrax"
	config.Net = `tcp`

	return sql.Open("mysql", config.FormatDSN())
}

func TestDriverCreationSuccessAndDelete(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	testes := []Model{
		NewDriver("", "Joao", "53242343224"),
		NewDriver("", "Joao", "53342525235"),
	}

	for _, d := range testes {
		if err := d.Insert(conn); err != nil {
			t.Fatal(err)
		}

		d.Delete(conn)
	}
}

func TestDriverCreationFailure(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := NewDriver(
		"",
		"Joao",
		"53242343224",
	)

	b := NewDriver(
		"",
		"Joao",
		"53242343224",
	)

	if err := a.Insert(conn); err != nil {
		t.Fail()
	}
	defer a.Delete(conn)

	if err := b.Insert(conn); err == nil {
		t.Fail()
	}
}

func TestDriverUpdateSuccess(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := NewDriver(
		"",
		"Joao",
		"53242343224",
	)

	if err := a.Insert(conn); err != nil {
		t.Fail()
	}
	defer a.Delete(conn)

	SetAttribute(a, "CNH", "53242433224")

	if err := a.Update(conn); err != nil {
		t.Fatal(err)
	}
}

func TestDriverUpdateFailure(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := NewDriver(
		"12345",
		"Joao",
		"53242343224",
	)

	if err := a.Update(conn); err == nil {
		t.Fail()
	}
}
