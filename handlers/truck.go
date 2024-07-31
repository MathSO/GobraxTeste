package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MathSO/GobraxTeste/factory"
	"github.com/MathSO/GobraxTeste/models"
	"github.com/labstack/echo/v4"
)

func ListTrucks(ctx echo.Context) error {
	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	rows, err := factory.GetAll(mysqlConnection, models.TruckTableName, ctx.QueryParam("page"), ctx.QueryParam("per_page"))
	if err != nil {
		return echo.ErrInternalServerError
	}

	var trucks = make([]models.Truck, 0)
	for rows.Next() {
		var t models.Truck

		err := rows.Scan(&t.ID, &t.Model, &t.Plate)
		if err != nil {
			return echo.ErrInternalServerError
		}

		trucks = append(trucks, t)
	}

	return ctx.JSON(200, trucks)
}

// CreateTruks creates a Truks
func CreateTruck(ctx echo.Context) error {
	var t models.Truck
	if err := json.NewDecoder(ctx.Request().Body).Decode(&t); err != nil {
		return echo.ErrBadRequest
	}

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = models.InsertTruck(mysqlConnection, &t)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "Placa já cadastrada na base"})
	}

	return ctx.JSON(200, t)
}

// GetTruks returns a Truks
func GetTruck(ctx echo.Context) error {
	id := ctx.Param(`id`)

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	t, err := models.RetrieveTruck(mysqlConnection, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "Caminhão não encontrado na base"})
		}

		return echo.ErrBadRequest
	}

	return ctx.JSON(200, t)
}

// UpdateTruks updates a Truks
func UpdateTruck(ctx echo.Context) error {
	var t models.Truck
	if err := json.NewDecoder(ctx.Request().Body).Decode(&t); err != nil {
		return echo.ErrBadRequest
	}
	t.ID = ctx.Param(`id`)

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = models.UpdateTruck(mysqlConnection, t)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]any{"message": "Caminhão não encontrado na base"})
		}

		return echo.ErrBadRequest
	}

	return ctx.JSON(200, t)
}

// DeleteTruks deletes a Truks
func DeleteTruck(ctx echo.Context) error {
	id := ctx.Param(`id`)

	mysqlConnection, err := factory.GetConnection()
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = models.DeleteTruck(mysqlConnection, id)
	if err != nil {
		if err != sql.ErrNoRows {
			return echo.ErrBadRequest
		}
	}

	return ctx.JSON(200, map[string]any{"message": "Caminhão deletado com sucesso"})
}
