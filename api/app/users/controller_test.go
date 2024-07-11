package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"managym-api/database"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Remove("test.db")

	os.Exit(code)
}

func TestCreateUserSuccess(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	UsersResources{}.MountRoutesInto(app)

	database.MockConnection()

	mockedUser := CreateUserRequestBody{
		FirstName: "foo",
		LastName:  "bar",
		Password:  "12345",
		Email:     "foo@bar.com",
	}

	data, _ := json.Marshal(mockedUser)

	request := httptest.NewRequest(fiber.MethodPost, "/users", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	response, _ := app.Test(request)

	responseInBytes, _ := io.ReadAll(response.Body)

	var expected User
	json.Unmarshal(responseInBytes, &expected)

	assert.Equal(fiber.StatusCreated, response.StatusCode)
	assert.NotEmpty(expected.CreatedAt)
}

func TestCreateUserNoRequiredData(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	UsersResources{}.MountRoutesInto(app)

	database.MockConnection()
	mUser := CreateUserRequestBody{}

	data, _ := json.Marshal(mUser)

	request := httptest.NewRequest(fiber.MethodPost, "/users", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	response, _ := app.Test(request)

	responseInBytes, _ := io.ReadAll(response.Body)

	var expected User
	json.Unmarshal(responseInBytes, &expected)

	assert.Equal(fiber.StatusBadRequest, response.StatusCode)
	assert.Empty(expected.CreatedAt)
}

func TestCreateUserNoBody(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	UsersResources{}.MountRoutesInto(app)

	database.MockConnection()

	request := httptest.NewRequest(fiber.MethodPost, "/users", nil)
	request.Header.Add("Content-Type", "application/json")
	response, _ := app.Test(request)

	responseInBytes, _ := io.ReadAll(response.Body)

	var expected User
	json.Unmarshal(responseInBytes, &expected)

	assert.Equal(fiber.StatusBadRequest, response.StatusCode)
	assert.Empty(expected.CreatedAt)
}

func TestCreateUserThatAlreadyExists(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	UsersResources{}.MountRoutesInto(app)

	database.MockConnection()
	mockedUser := CreateUserRequestBody{
		FirstName: "foo",
		LastName:  "bar",
		Password:  "12345",
		Email:     "foo@bar.com",
	}

	data, _ := json.Marshal(mockedUser)

	request := httptest.NewRequest(fiber.MethodPost, "/users", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	r, _ := app.Test(request)
	rb, _ := io.ReadAll(r.Body)
	fmt.Println(string(rb))

	request = httptest.NewRequest(fiber.MethodPost, "/users", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	response, _ := app.Test(request)

	responseInBytes, _ := io.ReadAll(response.Body)

	fmt.Println(string(responseInBytes))
	var expected User
	json.Unmarshal(responseInBytes, &expected)

	assert.Equal(fiber.StatusForbidden, response.StatusCode)
	assert.Empty(expected.CreatedAt)
}
