package App

import (
	UsersController "github.com/ezesali/BookStore_Users_API/Controllers/Users"
)

func MapsURLs() {

	//Map all users
	router.GET("/findusers/search", UsersController.FindUser)

	//Map specific user
	router.GET("/users/:uid", UsersController.GetUser)

	//Create a user
	router.POST("/users", UsersController.CreateUser)

}
