package config

import (
	"os"
	"user-service/internal/config/errors"

	"github.com/spf13/viper"
)

type Config struct {
	Port         int
	RunWithDapr  bool
	FirbaseCreds string
}

// This function reads app's configuration from a config file.
func ReadCfgFromFile(path string) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	firebaseCreds := os.Getenv("FIREBASE_CREDS")
	if firebaseCreds == "" {
		return Config{}, errors.ErrFirebaseCredentialsMissing
	}

	return Config{
		Port:         viper.GetInt("port"),
		RunWithDapr:  viper.GetBool("run_with_dapr"),
		FirbaseCreds: firebaseCreds,
	}, nil
}
