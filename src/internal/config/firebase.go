package config

import firebase "firebase.google.com/go/v4"

func GetFirebaseConfig() *firebase.Config {
	return &firebase.Config{
		ProjectID: "",
	}
}
