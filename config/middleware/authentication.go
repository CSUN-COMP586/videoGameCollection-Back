package middleware

import (
	"context"
	"fmt"
	"strings"

	firebase "firebase.google.com/go"
)

type Auth struct {
	Token string
}

type AuthHandler struct {
	Model *Auth
}

func (handler AuthHandler) VerifyTokenAndReturnUID(app *firebase.App) string {
	stringToken := strings.Fields(handler.Model.Token) // split 'beaer' and token
	handler.Model.Token = stringToken[1]

	client, err := app.Auth(context.Background()) // initialize firebase client for auth
	if err != nil {
		fmt.Println("Error getting auth client: ", err.Error())
	}

	ctx := context.Background()
	token, err := client.VerifyIDToken(ctx, handler.Model.Token) // verify the token
	if err != nil {
		fmt.Println("Error verifying ID token: ", err.Error())
	}

	return token.UID // return UID
}
