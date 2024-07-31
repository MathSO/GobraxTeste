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

	var drivers = make([]models.Driver, 0)
	for rows.Next() {
		var d models.Driver

		err := rows.Scan(&d.ID, &d.Name, &d.CNH)
		if err != nil {
			return echo.ErrInternalServerError
		}

		drivers = append(drivers, d)
	}

	return ctx.JSON(200, drivers)
}

func CreateDriver(ctx echo.Context) error {
	var d models.Driver
	if err := json.NewDecoder(ctx.Request().Body).Decode(&d); err != nil {
		return echo.ErrBadRequest
	}

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = models.InsertDriver(mysqlConnection, &d)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "CNH já cadastrada na base"})
	}

	return ctx.JSON(200, d)
}

func GetDriver(ctx echo.Context) error {
	id := ctx.Param(`id`)

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	d, err := models.RetrieveDriver(mysqlConnection, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "Driver não encontrado na base"})
		}

		return echo.ErrBadRequest
	}

	return ctx.JSON(200, d)
}

func UpdateDriver(ctx echo.Context) error {
	var d models.Driver
	if err := json.NewDecoder(ctx.Request().Body).Decode(&d); err != nil {
		return echo.ErrBadRequest
	}
	d.ID = ctx.Param(`id`)

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = models.UpdateDriver(mysqlConnection, d)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "Driver não encontrado na base"})
		}

		return echo.ErrBadRequest
	}

	return ctx.JSON(200, d)
}

func DeleteDriver(ctx echo.Context) error {
	id := ctx.Param(`id`)

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = models.DeleteDriver(mysqlConnection, id)
	if err != nil {
		if err != sql.ErrNoRows {
			return echo.ErrBadRequest
		}
	}

	return ctx.JSON(200, map[string]any{"message": "Driver deletado com sucesso"})
}
