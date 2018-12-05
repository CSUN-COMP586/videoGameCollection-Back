package middleware

// type Auth struct {
// 	Sub  uint
// 	Role string
// }

// type AuthHandler struct {
// 	Model *Auth
// }

// // CreateToken creates and returns a jwt token
// func (handler AuthHandler) CreateToken() string {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub":  handler.Model.Sub,
// 		"exp":  time.Now().Add(time.Minute * 15),
// 		"role": handler.Model.Role,
// 		"csrf": os.Getenv("CSRF_KEY"),
// 	})
// 	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_SIGN_KEY")))
// 	if err != nil {
// 		fmt.Println("signed string error")
// 		log.Fatal(err)
// 	}

// 	return tokenString
// }
