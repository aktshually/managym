package users

import (
	"crypto/sha256"
	"managym-api/database"
	"managym-api/database/models"
	"managym-api/utils"
	"math/rand"

	"github.com/bwmarrin/snowflake"
	"github.com/gofiber/fiber/v2"
	p "github.com/wuriyanto48/go-pbkdf2"
)

type IUsersService interface {
	CreateUser(CreateUserRequestBody) (User, error)
	FindUserByEmail(email string) (User, error)
}

type UsersService struct {
	IUsersService
}

func (usersService *UsersService) CreateUser(user CreateUserRequestBody) (User, utils.APIError) {
	node, err := snowflake.NewNode(int64(rand.Intn(1023)))
	if err != nil {
		return User{}, utils.APIError{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "could not create user",
		}
	}

	userId := node.Generate()
	encryptedPassword := p.NewPassword(sha256.New, 12, 32, 1500).HashPassword(user.Password)

	createdUser := models.User{
		Id:        userId.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  encryptedPassword.CipherText,
	}

	tx := database.DatabaseManager.Model(&models.User{}).Create(&createdUser)

	if tx.Error != nil {
		return User{}, utils.APIError{
			StatusCode: fiber.StatusForbidden,
			Message:    "an user with this email address already exists",
		}
	}

	return User{
		Id:        createdUser.Id,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
		IsActive:  createdUser.IsActive,
	}, utils.APIError{}
}
