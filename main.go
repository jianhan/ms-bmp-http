package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/api/option"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		opt := option.WithCredentialsFile("./admin_sdk.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			panic(err)
		}
		client, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) > 1 {
			reqToken = strings.Trim(splitToken[1], " ")
			spew.Dump(reqToken)
			token, err := client.VerifyIDToken(r.Context(), reqToken)
			if err != nil {
				log.Fatalf("error verifying ID token: %v\n", err)
			}

			log.Printf("Verified ID token: %v\n", token)
		}
		setupResponse(&w, r)
		fmt.Fprintf(w, "Hello")
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
