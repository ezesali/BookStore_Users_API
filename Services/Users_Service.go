package UsersService

import (
	"github.com/ezesali/BookStore_Users_API/Domain/Users"
	"github.com/ezesali/BookStore_Users_API/Utils/Errors"
)

func CreateUser(user Users.User) (*Users.User, *Errors.ResError) {
	return &user, nil

	//To implement
}
