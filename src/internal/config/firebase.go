package config

import (
	firebase "firebase.google.com/go/v4"
)

// Function for obtaining Firebase's configuration.
func GetFirebaseConfig() *firebase.Config {
	return &firebase.Config{
		ProjectID: "ocr-microservice-project",
	}
}
