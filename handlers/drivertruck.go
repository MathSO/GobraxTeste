package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MathSO/GobraxTeste/factory"
	"github.com/MathSO/GobraxTeste/models"
	"github.com/labstack/echo/v4"
)

func ListDriversTruck(ctx echo.Context) error {
	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	rows, err := factory.GetAll(mysqlConnection, models.DriveTruckTableName, ctx.QueryParam("page"), ctx.QueryParam("per_page"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	var driverTrucks = make([]models.Model, 0)
	for rows.Next() {
		var idDriver, idTruck string

		err := rows.Scan(&idDriver, &idTruck)
		if err != nil {
			return echo.ErrInternalServerError
		}

		driverTrucks = append(driverTrucks, models.NewDriverTruck(idDriver, idTruck))
	}

	return ctx.JSON(200, driverTrucks)
}

func CreateDriverTruck(ctx echo.Context) error {
	var d = models.NewDriverTruck("", "")
	if err := json.NewDecoder(ctx.Request().Body).Decode(&d); err != nil {
		return echo.ErrBadRequest
	}

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = d.Insert(mysqlConnection)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "Motorista já cadastrado ao caminhão"})
	}

	return ctx.JSON(200, d)
}

func GetDriverTruck(ctx echo.Context) error {
	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	var d = models.NewDriverTruck("", "")
	err = d.Load(mysqlConnection, ctx.Param(`driver_id`), ctx.Param(`truck_id`))
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "Relação caminhão com motorista não encontrada"})
		}

		return echo.ErrBadRequest
	}

	return ctx.JSON(200, d)
}

func DeleteDriverTruck(ctx echo.Context) error {
	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	var d = models.NewDriverTruck(ctx.Param(`driver_id`), ctx.Param(`truck_id`))
	err = d.Delete(mysqlConnection)
	if err != nil {
		if err != sql.ErrNoRows {
			return echo.ErrBadRequest
		}
	}

	return ctx.JSON(200, map[string]any{"message": "Relação caminhão com motorista deletado"})
}
