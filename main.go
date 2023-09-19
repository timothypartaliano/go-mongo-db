package main

import (
	"ugc-2/config"
	"ugc-2/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
    config.InitMongoDB()

    e := echo.New()

    e.POST("/employees", handlers.CreateEmployee)
    e.GET("/employees/:id", handlers.GetEmployee)
    e.PUT("/employees/:id", handlers.UpdateEmployee)
    e.DELETE("/employees/:id", handlers.DeleteEmployee)

    e.Start(":8080")
}