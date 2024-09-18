package util

import "github.com/spf13/viper"

// Config struct stores all configurations of the application.
// The values are read by viper from a config file or environment var.
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// Read configuration from config file or overwrite with environment variable if available
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") //json or xml

	// Overwrite with environment variables if exist
	viper.AutomaticEnv()
	
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}