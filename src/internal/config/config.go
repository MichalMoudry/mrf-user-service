package config

import (
	"github.com/spf13/viper"
)

type FirebaseConfig struct {
	Type                string
	PrivateKeyId        string
	PrivateKey          string
	ClientEmail         string
	ClientId            string
	AuthProviderCertUrl string
	ClientCertUrl       string
}

type Config struct {
	Port        int
	RunWithDapr bool
	FirebaseConfig
}

// This function reads app's configuration from a config file.
func ReadCfgFromFile(path string) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	return Config{
		Port:        viper.GetInt("port"),
		RunWithDapr: viper.GetBool("run_with_dapr"),
		FirebaseConfig: FirebaseConfig{
			Type:                viper.GetString("firebase.config.type"),
			PrivateKeyId:        viper.GetString("firebase.config.private_key_id"),
			PrivateKey:          viper.GetString("firebase.config.private_key"),
			ClientEmail:         viper.GetString("firebase.config.client_email"),
			ClientId:            viper.GetString("firebase.config.client_id"),
			AuthProviderCertUrl: viper.GetString("firebase.config.auth_provider_cert_url"),
			ClientCertUrl:       viper.GetString("firebase.config.client_cert_url"),
		},
	}, nil
}
