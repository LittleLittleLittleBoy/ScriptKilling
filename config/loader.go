package config

import (
	"log"

	"github.com/spf13/viper"
)

func Load(path string) (Config, error) {

	log.Printf("Start load config, file path is {%s}", path)
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	var config Config

	if err := v.ReadInConfig(); err != nil {
		return config, err
	}

	if err := v.Unmarshal(&config); err != nil {
		return config, err
	}

	log.Printf("End load config")

	return config, nil
}
