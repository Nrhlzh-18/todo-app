package main

import (
	"github.com/Nrhlzh-18/todo-app/config"
	"github.com/Nrhlzh-18/todo-app/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	env := config.NewEnvironment()
	e := echo.New()
	db := config.NewDatabase(env, "mysql")

	// cors
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	// }))
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	routes.NewRoutes(e, db)
 
	// start
	e.Logger.Fatal(e.Start(":" + env.Get("PORT")))
}
