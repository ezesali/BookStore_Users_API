package App

import "github.com/ezesali/BookStore_Users_API/Controllers/UsersController"

func MapsURLs() {

	router.GET("/ping", UsersController.Ping())

}
