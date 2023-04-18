package Controller

import (
	"FiberGo/Database"
	"FiberGo/Model/Entity"
	"FiberGo/Model/Request"
	"FiberGo/Utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func LoginHandler(ctx *fiber.Ctx) error {
	LoginRequest := new(Request.LoginRequest)
	if err := ctx.BodyParser(LoginRequest); err != nil {
		return err
	}

	//VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(LoginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//CHECK AVAILABLE USER
	var user Entity.User
	err := Database.DB.First(&user, "email = ?", LoginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//CHECK VALIDATION PASSWORD
	isValid := Utils.CheckPasswordHash(LoginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//GENERATE JWT
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerateToken := Utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})

	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
