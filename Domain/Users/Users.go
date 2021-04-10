package Users

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/ezesali/BookStore_Users_API/Utils/Errors"
)

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) ValidateCreate() *Errors.ResError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {

		// Response Error
		var resError Errors.ResError

		resError.Message = "An error has ocurred"
		resError.Description = "Email is require for creation"
		resError.Status = http.StatusBadRequest

		return &resError
	}

	reEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !reEmail.MatchString(user.Email) {

		resError.Message = "An error has ocurred"
		resError.Description = "Invalid Email"
		resError.Status = http.StatusBadRequest

		return &resError
	}

	return nil
}

func (user *User) ValidateGet() *Errors.ResError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {

		// Response Error
		var resError Errors.ResError

		resError.Message = "An error has ocurred"
		resError.Description = "Email is require to get User"
		resError.Status = http.StatusBadRequest

		return &resError
	}

	reEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !reEmail.MatchString(user.Email) {

		resError.Message = "An error has ocurred"
		resError.Description = "Invalid Email to get User"
		resError.Status = http.StatusBadRequest

		return &resError
	}

	return nil
}
