package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MathSO/GobraxTeste/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initiate handler
	handler := echo.New()
	handler.Pre(middleware.CORS())

	addDriversRoutes(handler)
	addTrucksRoutes(handler)

	// Gracefull shutdown
	go func() {
		sig := make(chan os.Signal, 1)

		signal.Notify(sig, syscall.SIGTERM)

		<-sig
		fmt.Println("Shutting down server")

		shutdownCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
		defer cancelFunc()

		handler.Shutdown(shutdownCtx)
	}()

	// Start server
	if err := handler.Start(":8080"); err != nil {
		panic(err)
	}
}

// addDriversRoutes adds CRUD routes for drivers
func addDriversRoutes(handler *echo.Echo) {
	handler.GET("/drivers", handlers.ListDrivers)

	driver := handler.Group("/driver")

	driver.POST("", handlers.CreateDriver)
	driver.GET("/:id", handlers.GetDriver)
	driver.PUT("/:id", handlers.UpdateDriver)
	driver.DELETE("/:id", handlers.DeleteDriver)
}

// addTrucksRoutes adds CRUD routes for trucks
func addTrucksRoutes(handler *echo.Echo) {
	handler.GET("/trucks", handlers.ListTrucks)

	truck := handler.Group("/truck")

	truck.POST("", handlers.CreateTruck)
	truck.GET("/:id", handlers.GetTruck)
	truck.PUT("/:id", handlers.UpdateTruck)
	truck.DELETE("/:id", handlers.DeleteTruck)
}
