package users

import "github.com/go-fuego/fuego"

type UsersResources struct {
	UsersService UsersService
}

type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UsersCreate struct {
	Name string `json:"name"`
}

type UsersUpdate struct {
	Name string `json:"name"`
}

func (rs UsersResources) Routes(s *fuego.Server) {
	usersGroup := fuego.Group(s, "/users")

	fuego.Get(usersGroup, "/", rs.getAllUsers)
	fuego.Post(usersGroup, "/", rs.postUsers)

	fuego.Get(usersGroup, "/{id}", rs.getUsers)
	fuego.Put(usersGroup, "/{id}", rs.putUsers)
	fuego.Delete(usersGroup, "/{id}", rs.deleteUsers)
}

func (rs UsersResources) getAllUsers(c fuego.ContextNoBody) (string, error) {
	return "hey", nil
}

func (rs UsersResources) postUsers(c *fuego.ContextWithBody[UsersCreate]) (Users, error) {
	body, err := c.Body()
	if err != nil {
		return Users{}, err
	}

	new, err := rs.UsersService.CreateUsers(body)
	if err != nil {
		return Users{}, err
	}

	return new, nil
}

func (rs UsersResources) getUsers(c fuego.ContextNoBody) (Users, error) {
	id := c.PathParam("id")

	return rs.UsersService.GetUsers(id)
}

func (rs UsersResources) putUsers(c *fuego.ContextWithBody[UsersUpdate]) (Users, error) {
	id := c.PathParam("id")

	body, err := c.Body()
	if err != nil {
		return Users{}, err
	}

	new, err := rs.UsersService.UpdateUsers(id, body)
	if err != nil {
		return Users{}, err
	}

	return new, nil
}

func (rs UsersResources) deleteUsers(c *fuego.ContextNoBody) (any, error) {
	return rs.UsersService.DeleteUsers(c.PathParam("id"))
}
