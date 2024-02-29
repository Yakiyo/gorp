package main

import (
	"github.com/charmbracelet/log"
	"github.com/hugolgst/rich-go/client"
)

// `connect` handles things related to the rich presence,
// which includes, creating the connection, reading config,
// reloading config, etc.
//
// If any error occurs midway, it just set's the error and returns
// early from the function
func connect() {
	var err error
	config, err = readConfig(configPath)
	if err != nil {
		log.Error("Error reading config file", "path", configPath, "err", err)
		setError(err)
	}
	log.Info("Read config", "config", config)

	err = client.Login(config.Id)
	if err != nil {
		log.Error("Error when creating connection", "err", err)
		setError(err)
		return
	}
}
