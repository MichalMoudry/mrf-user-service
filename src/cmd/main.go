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
	"google.golang.org/api/option"
)

func main() {
	log.Println("Hello from user service! ʕ•ᴥ•ʔ")

	// Read app's config
	cfg, err := config.ReadCfgFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Firebase init
	firebaseApp, err := firebase.NewApp(
		context.Background(),
		config.GetFirebaseConfig(),
		option.WithCredentialsJSON([]byte(cfg.FirbaseCreds)),
	)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	firebaseAuth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Printf("error initializing firebase auth:\n%v\n", err)
	}

	// Dapr init
	var daprClient dapr.Client
	if cfg.RunWithDapr {
		daprClient, err = dapr.NewClient()
		if err != nil {
			log.Fatal(err)
		}
		defer daprClient.Close()
	}

	// Web server init
	handler := transport.Initalize(
		cfg.Port,
		model.NewServiceCollection(daprClient, firebaseAuth),
		firebaseAuth,
	)
	log.Printf("Starting a web server on port %d\n", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
