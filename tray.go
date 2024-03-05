package main

import (
	_ "embed"

	"github.com/charmbracelet/log"
	"github.com/getlantern/systray"
	"github.com/hugolgst/rich-go/client"
)

// func initTray() {
// 	// systray.SetIcon(icon)
// 	systray.SetTitle("gorp")
// 	systray.SetTooltip("Rich Presence Client for Discord")
// 	reloadB = systray.AddMenuItem("Reload", "Reload config")
// 	quitB = systray.AddMenuItem("Quit", "Stop application")
// 	errMenu = systray.AddMenuItem("Errors", "Error detected")
// 	errMenu.Hide()
// }

func onReady() {
	connect := func(config Config) error {
		err := client.Login(config.Id)
		if err != nil {
			return err
		}
		err = client.SetActivity(config.asActivity())
		return err
	}
	config, err := readConfig(configPath)
	if err != nil {
		log.Error("Error reading config file", "path", configPath, "err", err)
		return
	}
	log.Info("Read config", "config", config)

	err = connect(config)
	if err != nil {
		log.Error("error initializing client", "err", err)
		return
	}

	systray.SetIcon(icon)
	systray.SetTitle("gorp")
	systray.SetTooltip("Rich Presence Client for Discord")
	reloadB := systray.AddMenuItem("Reload", "Reload config")
	quitB := systray.AddMenuItem("Quit", "Stop application")

loop:
	for {
		select {
		case <-quitB.ClickedCh:
			systray.Quit()
		case <-reloadB.ClickedCh:
			config, err := readConfig(configPath)
			if err != nil {
				log.Error("Error when reloading config", "err", err)
				break loop
			}
			// logout first, since this might be a different rich presence client
			client.Logout()

			err = connect(config)
			if err != nil {
				log.Error("error creating connection", "err", err)
				break loop
			}
		}
	}
	defer systray.Quit()
}

func onExit() {
	log.Info("Shutting down app")
	client.Logout()
}
