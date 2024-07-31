package models

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Model interface {
	Insert(*sql.DB) error
	Load(conn *sql.DB, id string) error
	Update(conn *sql.DB) error
	Delete(*sql.DB) error
	GetID() []string
}

func SetAttribute(m Model, name string, value any) error {
	rM := reflect.ValueOf(m)

	s := rM.Elem()

	field := s.FieldByName(name)
	if field.IsValid() && field.CanSet() {
		field.Set(reflect.ValueOf(value))
	} else {
		return fmt.Errorf("field %s can not set or not found", name)
	}

	return nil
}
