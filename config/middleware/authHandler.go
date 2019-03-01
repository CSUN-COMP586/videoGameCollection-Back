package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/jinzhu/gorm"
	"github.com/videogamelibrary/businesslogic"
	"github.com/videogamelibrary/models"
)

// VerifyToken - verifies the jwt and returns the account related to jwt
func VerifyToken(r *http.Request, app *firebase.App, conn *gorm.DB) (bool, uint, error) {
	// verify token and return uid
	idToken := r.Header.Get("authorization")
	authModel := Auth{Token: idToken}
	authHandler := AuthHandler{Model: &authModel}
	UID := authHandler.VerifyTokenAndReturnUID(app)

	// verify uid and return verification status and account model
	account := models.Account{}
	accountHandler := businesslogic.AccountHandler{Model: &account}
	verifyStatus, err := accountHandler.VerifyUID(conn, UID)
	if err != nil {
		fmt.Println("Verify UID error from authHandler.go: ", err.Error())
		return false, 0, err
	}

	return verifyStatus, accountHandler.Model.ID, nil
}

// HandleFalseVerification - Returns an error header response for unverified UIDa
func HandleFalseVerification(verifyStatus bool, w http.ResponseWriter, err error) {
	response := make(map[string]interface{})

	if verifyStatus != true {
		response["message"] = "UID not verified." + err.Error()
		responseJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error encoding UID verification failure: " + err.Error())
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}
}
