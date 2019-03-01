package middleware

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// App the firebase application variable to be passed around other files
var FireApp *firebase.App

func init() {
	var err error
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICE_ACCOUNT_PATH"))
	FireApp, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
}
