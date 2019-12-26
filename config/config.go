package config

import (
	"encoding/json"
	"github.com/capsules-web-server/logger"
	"os"
	"strconv"
)

type configuration struct {
	IsReleaseMode bool
	Port          int
	DBUrl         string
}

var config configuration

func Init() {
	mode := os.Getenv("MODE")

	var configPath string
	if mode == "RELEASE" {
		config.IsReleaseMode = true
		configPath = "config/release.json"
		logger.Info("configuration init entered to release mode")
		logger.Info("GetIsReleaseMode return:", GetIsReleaseMode())
	} else {
		config.IsReleaseMode = false
		configPath = "config/debug.json"
		logger.Info("configuration init entered to debug mode")
		logger.Info("GetIsReleaseMode return:", GetIsReleaseMode())
	}

	file, err := os.Open(configPath)
	if err != nil {
		logger.Fatal("open config file failed", err)
	}
	decoder := json.NewDecoder(file)
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

	logger.Info("config:", config)
}

func GetDBUrl() string {
	return config.DBUrl
}

func GetPort() int {
	return config.Port
}

func GetIsReleaseMode() bool {
	logger.Info("config.IsReleaseMode is", config.IsReleaseMode)
	return config.IsReleaseMode
}
