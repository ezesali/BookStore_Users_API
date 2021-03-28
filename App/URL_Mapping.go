package App

import (
	UsersController "github.com/ezesali/BookStore_Users_API/Controllers/Users"
)

func MapsURLs() {

	router.GET("/users/search", UsersController.FindUser)

	router.GET("/users/:uid", UsersController.GetUser)

	router.POST("/users", UsersController.CreateUser)

}
