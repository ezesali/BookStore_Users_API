package UsersController

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ezesali/BookStore_Users_API/Domain/Users"
	UsersService "github.com/ezesali/BookStore_Users_API/Services"
	"github.com/ezesali/BookStore_Users_API/Utils/Errors"
)

// Response Error
var resError Errors.ResError

func CreateUser(c *gin.Context) {

	fmt.Println("Starting Create Controller")

	//Declare User to be create
	var user Users.User

	//Looking for request raw body
	bytes, err := ioutil.ReadAll(c.Request.Body)

	//Handle request error
	if err != nil {
		fmt.Println(err.Error())

		resError.Message = "An error has ocurred"
		resError.Description = strings.ToUpper(err.Error())
		resError.Status = http.StatusBadRequest

		c.JSON(resError.Status, resError)
		return
	}

	//Handle JSON error
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())

		resError.Message = "Invalid JSON"
		resError.Description = strings.ToUpper(err.Error())
		resError.Status = http.StatusBadRequest

		c.JSON(resError.Status, resError)
		return
	}

	// Insert DB and return result
	result, saveErr := UsersService.CreateUser(user)

	//Handle error from DB
	if saveErr != nil {
		fmt.Println(saveErr)

		c.JSON(saveErr.Status, saveErr)
		return
	}

	// Result OK - User created
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)

}

func GetUser(c *gin.Context) {

	user, getErr := UsersService.GetUser(c.Param("uid"))

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func FindUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "Not implement")

}
