package users

type UsersService interface {
	GetUsers(id string) (Users, error)
	CreateUsers(UsersCreate) (Users, error)
	GetAllUsers() ([]Users, error)
	UpdateUsers(id string, input UsersUpdate) (Users, error)
	DeleteUsers(id string) (any, error)
}
