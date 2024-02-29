package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/getlantern/systray"
	_ "github.com/hugolgst/rich-go/client"
	homedir "github.com/mitchellh/go-homedir"
)

var (
	// home directory
	hdir string
	// the path to the config file being used
	configPath string

	// the active config
	config Config

	configChan = make(chan int)
)

func main() {
	err := Init()
	if err != nil {
		log.Error("initialization error", "err", err)
	}
	initTray()

	go connect()

	systray.Run(onReady, onExit)
}

func Init() error {
	dir, err := homedir.Dir()
	if err != nil {
		return err
	}
	hdir = dir
	logfile := filepath.Join(hdir, ".gorp", "gorp.log")
	// if log file exists, truncate it
	if pathExists(logfile) {
		err := os.Truncate(logfile, 0)
		// this is not fatal, so we just log it, no need for exiting the application
		if err != nil {
			log.Error("error when truncating previous log file", "err", err)
		}
	}
	log.SetOutput(os.NewFile(uintptr(os.ModePerm), logfile))

	// config path at ~/.gorp/config.toml
	configPath = filepath.Join(hdir, ".gorp", "config.toml")
	// if the first one doesn't exist, try at ~/.config/gorp.toml
	if !pathExists(configPath) {
		configPath = filepath.Join(hdir, ".config", "gorp.toml")
		// if even that doesn't exist, try in current directory
		if !pathExists(configPath) {
			configPath = "./config.toml"
		}
	}
	// if theres still no config file, return error
	if !pathExists(configPath) {
		return errors.New("unable to find config file")
	}
	log.Info("detected config file", "path", configPath)
	return nil
}
