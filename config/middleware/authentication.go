package middleware

import (
	"context"
	"fmt"
	"strings"

	firebase "firebase.google.com/go"
)

// Auth - Auth recieves the JWT from the front-end
type Auth struct {
	Token string
}

// AuthHandler - object that handles authentication
type AuthHandler struct {
	Model *Auth
}

// VerifyTokenAndReturnUID - verifies the firebase token and returns the UID
func (handler AuthHandler) VerifyTokenAndReturnUID(app *firebase.App) string {
	// split 'bearer' and token
	stringToken := strings.Fields(handler.Model.Token)
	handler.Model.Token = stringToken[1]

	// initialize firebase client for auth
	client, err := app.Auth(context.Background())
	if err != nil {
		fmt.Println("Error getting auth client: ", err.Error())
	}

	// verify the token
	ctx := context.Background()
	token, err := client.VerifyIDToken(ctx, handler.Model.Token)
	if err != nil {
		fmt.Println("Error verifying ID token: ", err.Error())
	}

	return token.UID
}
