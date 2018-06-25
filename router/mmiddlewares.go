package router

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	auth0 "github.com/auth0-community/go-auth0"
	"gopkg.in/square/go-jose.v2"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Creates a configuration with the Auth0 information
		secret, _ := base64.URLEncoding.DecodeString(os.Getenv("AUTH0_CLIENT_SECRET"))
		secretProvider := auth0.NewKeyProvider(secret)
		audience := os.Getenv("AUTH0_CLIENT_ID")
		configuration := auth0.NewConfiguration(secretProvider, []string{audience}, "https://mydomain.eu.auth0.com/", jose.HS256)
		validator := auth0.NewValidator(configuration, nil)

		token, err := validator.ValidateRequest(r)

		if err != nil {
			fmt.Println("Token is not valid:", token)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
