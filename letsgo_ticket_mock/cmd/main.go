package main

import (
	controllers "letsgo-flight-provider/controller"
	repositories "letsgo-flight-provider/infrastructure/repository"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	_ = repositories.NewGormDatabase()

	e := echo.New()
	controllers.AddFlightRoutes(e)

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
