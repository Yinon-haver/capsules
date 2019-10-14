package config

import (
	"encoding/json"
	"github.com/capsules-web-server/logger"
	"os"
)

type configuration struct {
	Port		int
	DBUrl		string
}

var config configuration

func Init() {
	var configPath string
	if mode := os.Getenv("MODE"); mode == "RELEASE" {
		configPath = "config/release.json"
	} else {
		configPath = "config/debug.json"
	}

	file, err := os.Open(configPath)
	if err != nil {
		logger.Fatal("open config file failed", err)
	}
	decoder := json.NewDecoder(file)
	config = configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		logger.Fatal("get configuration from file failed", err)
	}
}

func GetDBUrl() string {
	return config.DBUrl
}

func GetPort() int {
	return config.Port
}

