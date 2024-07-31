package models

import (
	"database/sql"
	"testing"

	"github.com/go-sql-driver/mysql"
)

func getMysqlConnection() (*sql.DB, error) {
	config := mysql.NewConfig()
	config.Addr = "172.17.236.222"
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

	testes := []Driver{
		{
			Name: "Joao",
			CNH:  "53242343224",
		},
		{
			Name: "Joao",
			CNH:  "53342525235",
		},
	}

	for _, d := range testes {
		if err := InsertDriver(conn, &d); err != nil {
			t.Fatal(err)
		}

		DeleteDriver(conn, d.ID)
	}
}

func TestDriverCreationFailure(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := Driver{
		Name: "Joao",
		CNH:  "53242343224",
	}
	b := Driver{
		Name: "Joao",
		CNH:  "53242343224",
	}

	if err := InsertDriver(conn, &a); err != nil {
		t.Fail()
	}
	defer DeleteDriver(conn, a.ID)

	if err := InsertDriver(conn, &b); err == nil {
		t.Fail()
	}
}

func TestDriverUpdateSuccess(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := Driver{
		Name: "Joao",
		CNH:  "53242343224",
	}

	if err := InsertDriver(conn, &a); err != nil {
		t.Fail()
	}
	defer DeleteDriver(conn, a.ID)

	a.CNH = "53242433224"

	if err := UpdateDriver(conn, a); err != nil {
		t.Fatal(err)
	}
}

func TestDriverUpdateFailure(t *testing.T) {
	conn, err := getMysqlConnection()
	if err != nil {
		t.Skip()
	}

	a := Driver{
		ID:   "12345",
		Name: "Joao",
		CNH:  "53242343224",
	}

	if err := UpdateDriver(conn, a); err == nil {
		t.Fail()
	}
}
