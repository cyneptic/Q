package main

import (
	"log"
	"os"

	"github.com/cyneptic/letsgo_smspanel_mockapi/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	controller.AddMessageRoutes(e)

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
