package main

import (
	"fmt"
	_ "embed"

	"github.com/charmbracelet/log"
	"github.com/getlantern/systray"
)

func onReady() {
	systray.SetTemplateIcon(icon, icon)
	systray.SetTitle("gorp")
	systray.SetTooltip("Rich Presence Client for Discord")
	reloadB := systray.AddMenuItem("Reload", "Reload config")
	quitB := systray.AddMenuItem("Quit", "Stop application")
	go func ()  {
		<-quitB.ClickedCh
		fmt.Println("clicked quit")
		systray.Quit()
	}()

	go func() {
		<-reloadB.ClickedCh
		fmt.Println("clicked reload")
		var err error
		config, err = readConfig(configPath)
		if err != nil {
			log.Error("Error when reloading config", "err", err)
		}
	}()
}

func onExit() {
	log.Info("Shutting down app")
}