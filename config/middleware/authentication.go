package middleware

import (
	"context"
	"log"
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
	stringToken := strings.Fields(handler.Model.Token)
	handler.Model.Token = stringToken[1]

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal("Error getting auth client: %v\n", err)
	}

	ctx := context.Background()
	token, err := client.VerifyIDToken(ctx, handler.Model.Token)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	return token.UID
}

// func (handler AuthHandler) DecodeToken(token string) {
// 	stringToken := strings.Fields(token)
// 	parsedToken, err := jwt.Parse(stringToken[1], func(token *jwt.Token) (interface{}, error) {
// 		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 		// 	return nil, fmt.Errorf("There was an error")
// 		// }
// 		return "hello friends", nil
// 	})
// 	if err != nil {
// 		fmt.Println("Error parsing token")
// 		log.Fatal(err)
// 	}
// 	handler.Model.UID = parsedToken.Claims.(jwt.MapClaims)["user_id"].(string)
// }

// func (handler AuthHandler) Authenticate(conn *gorm.DB) (bool, error) {
// 	account := businesslogic.Account{}

// 	if conn.Where(&businesslogic.Account{UID: handler.Model.UID}).Find(&account).RecordNotFound() != false {
// 		err := errors.New("Invalid user identification")
// 		return false, err
// 	}
// 	fmt.Println(account)

// 	return true, nil
// }
