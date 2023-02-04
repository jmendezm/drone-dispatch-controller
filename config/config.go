package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DatabasePath string `json:"db_path"`
	ShowLogs     bool   `json:"show_logs"`
	LogLevel     int    `json:"log_level"`
}

func New(configFilePath string) *Config {
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalln("can't read config file: ", err)
	}
	var c Config
	err = json.Unmarshal(file, &c)
	if err != nil {
		log.Fatalln("error parsing config file: ", err)
	}
	return &c
}
