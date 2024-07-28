package users

import (
	"bytes"
	"encoding/json"
	"io"
	"managym/database"
	"managym/utils"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pashagolub/pgxmock/v4"
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

	mock := database.MockConnection()

	mockedUser := CreateUserRequestBody{
		FirstName: "foo",
		LastName:  "bar",
		Password:  "12345",
		Email:     "foo@bar.com",
	}

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(utils.AnyString{}, mockedUser.Email, utils.AnyString{}, mockedUser.FirstName, mockedUser.LastName).
		WillReturnRows(pgxmock.NewRows([]string{"id", "email", "first_name", "last_name", "created_at", "updated_at", "is_active"}).AddRow("", mockedUser.Email, mockedUser.FirstName, mockedUser.LastName, time.Now(), time.Now(), true))

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

	mock := database.MockConnection()
	mockedUser := CreateUserRequestBody{
		FirstName: "foo",
		LastName:  "bar",
		Password:  "12345",
		Email:     "foo@bar.com",
	}

	data, _ := json.Marshal(mockedUser)

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(utils.AnyString{}, mockedUser.Email, utils.AnyString{}, mockedUser.FirstName, mockedUser.LastName).
		WillReturnRows(pgxmock.NewRows([]string{"id", "email", "first_name", "last_name", "created_at", "updated_at", "is_active"}).AddRow("", mockedUser.Email, mockedUser.FirstName, mockedUser.LastName, time.Now(), time.Now(), true))

	request := httptest.NewRequest(fiber.MethodPost, "/users", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	r, _ := app.Test(request)
	io.ReadAll(r.Body)

	request = httptest.NewRequest(fiber.MethodPost, "/users", bytes.NewReader(data))
	request.Header.Add("Content-Type", "application/json")
	response, _ := app.Test(request)

	responseInBytes, _ := io.ReadAll(response.Body)

	var expected User
	json.Unmarshal(responseInBytes, &expected)

	assert.Equal(fiber.StatusForbidden, response.StatusCode)
	assert.Empty(expected.CreatedAt)
}
