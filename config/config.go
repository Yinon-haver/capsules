package config

import (
	"encoding/json"
	"github.com/capsules-web-server/logger"
	"os"
	"strconv"
)

type configuration struct {
	Port					int
	DBUrl					string
	BroadcastChannelSize	int
}

var config configuration

func Init() {
	mode := os.Getenv("MODE")

	var configPath string
	if mode == "RELEASE" {
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

	if mode == "RELEASE" {
		portEnv := os.Getenv("PORT")
		port, err := strconv.Atoi(portEnv)
		if err != nil {
			logger.Fatal("illegal port", err)
		}
		config.Port = port
	}
}

func GetDBUrl() string {
	return config.DBUrl
}

func GetPort() int {
	return config.Port
}

func GetBroadcastChannelSize() int {
	return config.BroadcastChannelSize
}

