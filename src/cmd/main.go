package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/internal/config"
	"user-service/internal/transport"
)

func main() {
	// Read app's config
	cfg, err := config.ReadCfgFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	handler := transport.Initalize(cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
