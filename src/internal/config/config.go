package config

import (
	"github.com/spf13/viper"
)

// App's general configuration.
type Config struct {
	Port int
	Environment
	RunWithDapr bool
}

// This function reads app's configuration from a config file.
func ReadCfgFromFile(path string) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	return Config{
		Port:        viper.GetInt("port"),
		Environment: Environment(viper.GetString("environment")),
		RunWithDapr: viper.GetBool("run_with_dapr"),
	}, nil
}
