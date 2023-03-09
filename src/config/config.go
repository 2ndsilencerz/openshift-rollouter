package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Viper *viper.Viper
}

func (config *Config) load() {
	config.Viper = viper.GetViper()
	confLocation := os.Getenv("CONFIG_LOCATION")
	confName := os.Getenv("CONFIG_NAME")
	confType := os.Getenv("CONFIG_TYPE")
	if len(confName) <= 0 {
		log.Println("Loading default config")
		confName = "config"
	}
	if len(confLocation) > 0 {
		log.Println("Location defined, adding confLocation")
		config.Viper.AddConfigPath(confLocation)
	} else {
		config.Viper.AddConfigPath(".")
	}
	if len(confType) > 0 {
		config.Viper.SetConfigType(confType)
	} else {
		config.Viper.SetConfigType("yaml")
	}
	config.Viper.SetConfigName(confName)
	err := config.Viper.ReadInConfig()
	if err != nil {
		log.Println("Error loading config" + err.Error())
	}
	//log.Println(config.Viper.AllKeys())
}

func NewConfig() *Config {
	config := Config{}
	config.load()
	return &config
}

func (config *Config) GetString(key string) string {
	config.load()
	return viper.GetString(key)
}

func (config *Config) GetInt(key string) int {
	config.load()
	return viper.GetInt(key)
}
