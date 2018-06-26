package router

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	auth0 "github.com/auth0-community/go-auth0"
	"gopkg.in/square/go-jose.v2"
)

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Decode secret
		secret, err := base64.URLEncoding.DecodeString(os.Getenv("AUTH0_CLIENT_SECRET"))
		if err != nil {
			fmt.Println("Error while decoding client:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
		}

		// get configuration
		configuration := auth0.NewConfiguration(
			auth0.NewKeyProvider(secret),
			[]string{os.Getenv("AUTH0_CLIENT_ID")},
			os.Getenv("AUTH0_ISSUER"),
			jose.HS256)

		// token validation
		token, err := auth0.NewValidator(configuration, nil).ValidateRequest(r)
		if err != nil {
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
