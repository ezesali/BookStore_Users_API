package Users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	Cloudant "github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/core"
	"github.com/ezesali/BookStore_Users_API/DataSources/MySQL/UsersDB"
	"github.com/ezesali/BookStore_Users_API/Utils/Errors"
	UUID "github.com/google/uuid"
)

/*const (
	queryInsertUser = ("INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);")
)
*/

var (
	resError Errors.ResError
)

func (user *User) Save() *Errors.ResError {

	fmt.Println("Starting User Saving")

	//Search for all docs in DB
	resultQuery, _, err := UsersDB.Client.PostAllDocs(
		UsersDB.Client.NewPostAllDocsOptions("users"),
	)
	if err != nil {

		resError.Message = "Cannot get all Documents"
		resError.Description = err.Error()
		resError.Status = http.StatusBadRequest

		return &resError

	}

	//If i have docs
	if len(resultQuery.Rows) > 1 {

		// Iterate document IDs
		for document := 0; document < len(resultQuery.Rows); document++ {

			// Get document data
			resGetDoc, _, err := UsersDB.Client.GetDocument(
				UsersDB.Client.NewGetDocumentOptions("users", *resultQuery.Rows[document].ID),
			)

			if err != nil {

				resError.Message = "Cannot get Document: " + *resultQuery.Rows[document].ID
				resError.Description = err.Error()
				resError.Status = http.StatusBadRequest

				return &resError
			}

			// if Email is already created
			if resGetDoc.GetProperty("email") == user.Email {

				resError.Message = "Cannot insert document"
				resError.Description = "Email was already created "
				resError.Status = http.StatusBadRequest

				return &resError

			}
		}
	}

	// New Document ID
	id, err := UUID.NewUUID()

	if err != nil {

		resError.Message = "Cannot create ID Document"
		resError.Description = err.Error()
		resError.Status = http.StatusBadRequest

		return &resError
	}

	//Get ID string
	idData, _ := id.MarshalText()

	UserDoc := Cloudant.Document{
		ID: core.StringPtr("users:" + string(idData)),
	}

	// Format Sysdate
	user.DateCreated = time.Now().Format("02/01/2006 03:04:05")

	// Setting Cloudant Document
	UserDoc.SetProperty("first_name", user.FirstName)
	UserDoc.SetProperty("last_name", user.LastName)
	UserDoc.SetProperty("email", user.Email)
	UserDoc.SetProperty("date_created", user.DateCreated)

	UserDocStr, _ := json.MarshalIndent(UserDoc, "", "    ")

	fmt.Println("User Object: ", string(UserDocStr))

	// Insert document to Users DB
	result, _, err := UsersDB.Client.PostDocument(
		UsersDB.Client.NewPostDocumentOptions("users").SetDocument(&UserDoc),
	)

	if err != nil {

		resError.Message = "Cannot insert document"
		resError.Description = err.Error()
		resError.Status = http.StatusBadRequest

		return &resError
	}

	resStr, _ := json.MarshalIndent(result, "", " ")

	fmt.Println("Result: ", string(resStr))

	return nil

	//MYSQL IMPLEMENTATION
	/*stmt, err := UsersDB.Client.Prepare(queryInsertUser)

		if err != nil {

			resError.Message = "An error has ocurred"
			resError.Description = err.Error()
			resError.Status = http.StatusInternalServerError
		}

		// Defer means when returns execute this
		defer stmt.Close()

		user.DateCreated = time.Now().Format("02/01/2006T15:05:05Z")

		insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

		if err != nil {

			resError.Message = "An error has ocurred"
			resError.Description = err.Error()
			resError.Status = http.StatusInternalServerError
		}

		userId, err := insertResult.LastInsertId()

		if err != nil {

			resError.Message = "An error has ocurred"
			resError.Description = err.Error()
			resError.Status = http.StatusInternalServerError
		}

		user.Id = userId
		return nil
	}*/
}

func (user *User) Get() *Errors.ResError {

	fmt.Println("GETTTT")

	/*if err := UsersDB.Client.Ping(); err != nil {

		panic(err)

	}*/

	/*result := Users_DB[user.Id]

	if result == nil {

		resError.Message = "An error has ocurred"
		resError.Description = fmt.Sprintf("User %d not found", user.Id)
		resError.Status = http.StatusNotFound

		return &resError
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated*/

	return nil

}
