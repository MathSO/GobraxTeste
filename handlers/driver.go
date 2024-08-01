package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MathSO/GobraxTeste/factory"
	"github.com/MathSO/GobraxTeste/models"
	"github.com/labstack/echo/v4"
)

func ListDrivers(ctx echo.Context) error {
	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	rows, err := factory.GetAll(mysqlConnection, models.DriverTableName, ctx.QueryParam("page"), ctx.QueryParam("per_page"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	var drivers = make([]models.Model, 0)
	for rows.Next() {
		var id, name, cnh string

		err := rows.Scan(&id, &name, &cnh)
		if err != nil {
			return echo.ErrInternalServerError
		}

		drivers = append(drivers, models.NewDriver(id, name, cnh))
	}

	return ctx.JSON(200, drivers)
}

func CreateDriver(ctx echo.Context) error {
	var d = models.NewDriver("", "", "")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&d); err != nil {
		return echo.ErrBadRequest
	}

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = d.Insert(mysqlConnection)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "CNH já cadastrada na base"})
	}

	return ctx.JSON(200, d)
}

func GetDriver(ctx echo.Context) error {
	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	var d = models.NewDriver("", "", "")
	err = d.Load(mysqlConnection, ctx.Param(`id`))
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "Motorista não encontrado na base"})
		}

		return echo.ErrBadRequest
	}

	return ctx.JSON(200, d)
}

func UpdateDriver(ctx echo.Context) error {
	var d = models.NewDriver("", "", "")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&d); err != nil {
		return echo.ErrBadRequest
	}

	models.SetAttribute(d, "ID", ctx.Param(`id`))

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = d.Update(mysqlConnection)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "Motorista não encontrado na base"})
		}

		return echo.ErrBadRequest
	}

	return ctx.JSON(200, d)
}

func DeleteDriver(ctx echo.Context) error {
	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	var d = models.NewDriver(ctx.Param(`id`), "", "")
	err = d.Delete(mysqlConnection)
	if err != nil {
		if err != sql.ErrNoRows {
			return echo.ErrBadRequest
		}
	}

	return ctx.JSON(200, map[string]any{"message": "Motorista deletado com sucesso"})
}
