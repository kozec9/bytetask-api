package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MONGOURL  string `mapstructure:"MONGO_URL"`
	MONGONAME string `mapstructure:"MONGO_NAME"`
	PORT      string `mapstructure:"PORT"`
	// PASETOKEY string `mapstructure:"PASETO_KEY"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read Config Env", err)
		return
	}

	// Unmarshal convert Json to Go object
	// Marshal convert Go object to Json
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
	return
}
