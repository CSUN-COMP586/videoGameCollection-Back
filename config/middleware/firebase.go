package middleware

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// FireApp - the firebase application variable to be passed around other files
var FireApp *firebase.App

// Initialize firebase account path and create new firebase client
func init() {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICE_ACCOUNT_PATH"))
	var err error

	FireApp, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
}
