package handlers

import "github.com/labstack/echo/v4"

func ListDrivers(ctx echo.Context) error {
	return ctx.String(200, "List drivers")
}

func CreateDriver(ctx echo.Context) error {
	return ctx.String(200, "Create a driver")
}

func GetDriver(ctx echo.Context) error {
	return ctx.String(200, "Get a driver")
}

func UpdateDriver(ctx echo.Context) error {
	return ctx.String(200, "Update a driver")
}

func DeleteDriver(ctx echo.Context) error {
	return ctx.String(200, "Delete a driver")
}
