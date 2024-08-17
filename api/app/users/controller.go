package users

import (
	"managym/utils"
	"managym/validation"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UsersResources struct {
	UsersService UsersService
}

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	IsActive  bool   `json:"is_active"`
}

type CreateUserRequestBody struct {
	FirstName string `json:"first_name" validate:"required,gte=3,max=100"`
	LastName  string `json:"last_name" validate:"max=100"`
	Email     string `json:"email" validate:"required,email,max=100"`
	Password  string `json:"password" validate:"max=100"`
}

func (rs UsersResources) MountRoutesInto(app *fiber.App) {
	usersGroup := app.Group("/users")

	usersGroup.Post("/", rs.createUser)
}

func (rs UsersResources) createUser(context *fiber.Ctx) error {
	var body CreateUserRequestBody
	context.BodyParser(&body)

	errs := validation.Validate(body)
	if len(errs) > 0 {
		return utils.ThrowError(context, fiber.StatusBadRequest, strings.Join(errs, ","))
	}

	createdUser, err := rs.UsersService.CreateUser(body)
	if err.StatusCode != 0 {
		return utils.ThrowError(context, err.StatusCode, err.Message)
	}

	return context.Status(fiber.StatusCreated).JSON(createdUser)
}
