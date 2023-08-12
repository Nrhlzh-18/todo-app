package main

import (
	"github.com/Nrhlzh-18/todo-app/config"
	"github.com/Nrhlzh-18/todo-app/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// repo "github.com/Nrhlzh-18/todo-app/app/schedule/repository"
)

func main() {
	env := config.NewEnvironment()
	e := echo.New()
	db := config.NewDatabase(env, "mysql")
	// cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// scheduleRows, err := repo.GetSchedules(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// helpers.SendEmailForTomorrowSchedules(scheduleRows)

	// Routes
	routes.NewRoutes(e, db)

	// start
	e.Logger.Fatal(e.Start(":" + env.Get("PORT")))
}
