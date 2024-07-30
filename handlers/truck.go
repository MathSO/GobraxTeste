package handlers

import "github.com/labstack/echo/v4"

// ListTruks returns a list of Truks
func ListTrucks(ctx echo.Context) error {
	return ctx.String(200, "List Truks")
}

// CreateTruks creates a Truks
func CreateTruck(ctx echo.Context) error {
	return ctx.String(200, "Create a Truks")
}

// GetTruks returns a Truks
func GetTruck(ctx echo.Context) error {
	return ctx.String(200, "Get a Truks")
}

// UpdateTruks updates a Truks
func UpdateTruck(ctx echo.Context) error {
	return ctx.String(200, "Update a Truks")
}

// DeleteTruks deletes a Truks
func DeleteTruck(ctx echo.Context) error {
	return ctx.String(200, "Delete a Truks")
}
