package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"user-service/internal/config"
	"user-service/internal/transport"
	"user-service/internal/transport/model"

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
	firebaseAuth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Printf("error initializing firebase auth:\n%v\n", err)
	}

	var daprClient dapr.Client
	if cfg.RunWithDapr {
		daprClient, err = dapr.NewClient()
		if err != nil {
			log.Fatal(err)
		}
		defer daprClient.Close()
	}

	handler := transport.Initalize(
		cfg.Port,
		model.NewServiceCollection(daprClient, firebaseAuth),
	)
	err = http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
