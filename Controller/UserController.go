package Controller

import (
	"FiberGo/Database"
	"FiberGo/Model/Entity"
	"FiberGo/Model/Request"
	"FiberGo/Utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

func UserControllerGetAll(ctx *fiber.Ctx) error {
	var users []Entity.User
	result := Database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func UserControllerCreate(ctx *fiber.Ctx) error {
	user := new(Request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	//VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "fail",
			"error":   errValidate.Error(),
		})
	}
	NewUser := Entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	hashedPassword, err := Utils.HasingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	NewUser.Password = hashedPassword

	errCreateUser := Database.DB.Create(&NewUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    NewUser,
	})
}

func UserControllerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user Entity.User
	err := Database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserControllerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(Request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user Entity.User

	userId := ctx.Params("id")
	//CHECK AVAILABLE USER
	err := Database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//UPDATE USER DATA
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	errUpdate := Database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserControllerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(Request.UserEmailRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user Entity.User
	var isEmailUserExist Entity.User

	userId := ctx.Params("id")
	//CHECK AVAILABLE USER
	err := Database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//CHECK AVAILABLE EMAIL
	errCheckEmail := Database.DB.First(&isEmailUserExist, "email = ?", userRequest.Email).Error
	if errCheckEmail == nil {
		return ctx.Status(402).JSON(fiber.Map{
			"message": "email already used.",
		})
	}

	//UPDATE USER DATA
	user.Email = userRequest.Email
	errUpdate := Database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserControllerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user Entity.User

	//CHECK AVAILABLE USER
	err := Database.DB.Debug().First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	errDelete := Database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})
}
