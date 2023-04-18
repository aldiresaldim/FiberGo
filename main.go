package main

import (
	"FiberGo/Database"
	"FiberGo/Migration"
	"FiberGo/Route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//INITIAL DATABASE
	Database.DatabaseInit()

	//RUN MIGRATION
	Migration.RunMigration()

	app := fiber.New()

	//INITIAL INIT
	Route.RouteInit(app)

	app.Listen(":201")
}
