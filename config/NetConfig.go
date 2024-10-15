package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

type CarConfig struct {
	Ip        string
	Port      string
	Broadcast string
}

type AllConfig struct {
	Car CarConfig
}

var Config AllConfig

func init() {
	configPath := flag.String("config", "config/config.yaml", "Path to config file")
	flag.Parse()
	viper.SetConfigFile(*configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("LoadConfig.ReadInConfig error: %v", err)
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("LoadConfig.Unmarshal error: %v", err)
	}
}
