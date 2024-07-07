package main

import (
	"fmt"
	"managym-api/app/users"
	"managym-api/database"
	"managym-api/utils"
	"os"

	"github.com/go-fuego/fuego"
)

// This API rewrite features user authentication and application management for
// the Managym app.
func main() {
	utils.LoadEnv()
	server := fuego.NewServer(fuego.WithAddr(fmt.Sprint("localhost:", os.Getenv("PORT"))))
	database.Connect()

	users.UsersResources{}.Routes(server)

	server.Run()
}
