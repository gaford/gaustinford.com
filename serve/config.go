package main

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

var configFile = "/etc/gaford/gaustinford.com/gaustinford.com.json"

// Config is the global configuration struct.
var Config struct {
	Domain string `json:"domain"`
}

func init() {
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.WithFields(log.Fields{
			"configFile": configFile,
			"error":      err,
		}).Panic("Configuration failed to load!")
	}

	err = json.Unmarshal(config, &Config)
	if err != nil {
		log.WithField("error", err).
			Panic("Configuration failed to unmarshal!")
	}

	log.WithField("configFile", configFile).
		Info("Configuration loaded.")
}
