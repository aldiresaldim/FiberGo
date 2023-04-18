package Route

import (
	"FiberGo/Config"
	"FiberGo/Controller"
	"FiberGo/Middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", Config.ProjectRootPath+"/Public/Asset")

	r.Post("/login", Controller.LoginHandler)

	r.Get("/user", Middleware.Auth, Controller.UserControllerGetAll)
	r.Get("/user/:id", Controller.UserControllerGetById)
	r.Post("/user", Controller.UserControllerCreate)
	r.Put("/user/:id", Controller.UserControllerUpdate)
	r.Put("/user/:id", Controller.UserControllerDelete)
	r.Put("/user/update_email", Controller.UserControllerUpdateEmail)
}
