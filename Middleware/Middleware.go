package Middleware

import (
	"FiberGo/Utils"
	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims, err := Utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	//ctx.Locals("userInfo", claims)

	return ctx.Next()
}

func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}
