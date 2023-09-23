package config

import "github.com/spf13/viper"

type Config struct {
	Port        int
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
		RunWithDapr: viper.GetBool("run_with_dapr"),
	}, nil
}
