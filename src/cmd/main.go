package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"user-service/internal/config"
	"user-service/internal/transport"

	firebase "firebase.google.com/go/v4"
	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	// Read app's config
	cfg, err := config.ReadCfgFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	firebaseApp, err := firebase.NewApp(context.Background(), config.GetFirebaseConfig())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	firebaseApp.Auth(context.Background())

	daprClient, err := dapr.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer daprClient.Close()

	handler := transport.Initalize(cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
