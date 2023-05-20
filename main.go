package main

import (
    "echo-mongo-api/configs"
    "echo-mongo-api/routes" 
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
          return c.JSON(200, &echo.Map{"data": "Hola holita"})
	})
	
	configs.ConnectDB()
	routes.UserRoute(e)
    e.Logger.Fatal(e.Start(":6000"))
}