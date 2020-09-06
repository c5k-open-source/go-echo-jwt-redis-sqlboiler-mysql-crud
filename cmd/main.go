package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/controllers"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/external"
	"github.com/c5k-open-source/go-echo-jwt-redis-sqlboiler-mysql-crud/pkg/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	// Test Route
	e.GET("/wowz", func(c echo.Context) error {
		s := external.NewSlackNotifier()
		s.Notify(fmt.Sprintf("API up and running"))
		return c.String(http.StatusOK, utils.GetIP(c))
	})

	pg := e.Group("p/api/panel/v1")
	// Middlewares
	pg.Use(middleware.Logger())  // Logger
	pg.Use(middleware.Recover()) // Recover
	pg.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.GET},
	}))

	pg.POST("/login", controllers.Login)
	pg.POST("/partner", controllers.CreatePartner)

	// Save routes into a json
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		fmt.Printf(err.Error())
	}
	ioutil.WriteFile("routes.json", data, 0644)

	// Start the server
	e.Logger.Fatal(e.Start(":8000"))
}
