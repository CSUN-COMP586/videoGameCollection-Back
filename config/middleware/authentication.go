package middleware

// var JwtAuth = func(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		noAuth := []string{"/account/register", "/account/login"} // list of non-authenticated routes
// 		requestPath := r.URL.Path                                 // current route

// 		// serve the request if it doesn't need authentication
// 		for _, value := range noAuth {
// 			if value == requestPath {
// 				next.ServeHTTP(w, r)
// 				return
// 			}
// 		}

// 		response := make(map[string]string)
// 		tokenHeader := r.Header.Get("Authentication")

// 		// if authentication token is missing
// 		if tokenHeader == "" {
// 			response["message"] = "403 error not authorized, missing authentication token."
// 			responseJSON, err := json.Marshal(response)
// 			if err != nil {
// 				fmt.Println("encoding error for 403 message")
// 				log.Fatal(err)
// 				return
// 			}

// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.WriteHeader(http.StatusForbidden)
// 			w.Write(responseJSON)
// 			return
// 		}

// 		// if token is malformed
// 		splitted := strings.Split(tokenHeader, " ")
// 		if len(splitted) != 2 {
// 			response["message"] = "Invalid malformed token."
// 			responseJSON, err := json.Marshal(response)
// 			if err != nil {
// 				fmt.Println("encoding error for malformed message")
// 				log.Fatal(err)
// 				return
// 			}

// 			w.WriteHeader(http.StatusForbidden)
// 			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 			w.Write(responseJSON)
// 		}

// 		tokenPart := splitted[1]
// 		tk := businesslogic.Token{}

// 		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(os.Get)
// 		})
// 	})
// }
