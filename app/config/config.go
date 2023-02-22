package config

import (
	"sync"

	"github.com/spf13/viper"
)

var cfg *Config
var loadConfigOnce = &sync.Once{}

func Load() {
	loadConfigOnce.Do(load)
}

func load() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	return cfg
}
