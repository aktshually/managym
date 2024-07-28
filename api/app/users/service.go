package users

import (
	"context"
	"crypto/sha256"
	"fmt"
	"managym/database"
	"managym/utils"
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

	createdUser, err := database.Manager.CreateUser(context.Background(), database.CreateUserParams{
		ID:        userId.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  encryptedPassword.CipherText,
	})
	if err != nil {
		fmt.Println(err)
		return User{}, utils.APIError{
			StatusCode: fiber.StatusForbidden,
			Message:    "an user with this email address already exists",
		}
	}

	return User{
		Id:        createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt.Time.Unix(),
		UpdatedAt: createdUser.UpdatedAt.Time.Unix(),
		IsActive:  createdUser.IsActive.Bool,
	}, utils.APIError{}
}
