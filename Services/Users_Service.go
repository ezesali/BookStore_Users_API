package UsersService

import (
	"fmt"

	"github.com/ezesali/BookStore_Users_API/Domain/Users"
	"github.com/ezesali/BookStore_Users_API/Utils/Errors"
)

func CreateUser(user Users.User) (*Users.User, *Errors.ResError) {

	fmt.Println("Starting Create User Service")

	if err := user.ValidateCreate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(email string) (*Users.User, *Errors.ResError) {

	fmt.Println("Starting Get User Service")

	//Email from parameter URL
	result := &Users.User{Email: email}

	if err := result.ValidateGet(); err != nil {
		return nil, err
	}

	// Get service
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}
