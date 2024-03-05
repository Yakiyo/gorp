package main

import (
	"os"

	"github.com/charmbracelet/log"
	"fyne.io/systray"
	"github.com/hugolgst/rich-go/client"
)

var rch = make(chan bool)

func main() {
	go initClient(rch)
	systray.Run(onReady, client.Logout)
}

func onReady() {
	systray.SetIcon(icon)
	systray.SetTitle("gorp")
	systray.SetTooltip("Rich Presence Client for Discord")
	reloadB := systray.AddMenuItem("Reload", "Reload config")
	quitB := systray.AddMenuItem("Quit", "Stop application")
	reloadB.SetTooltip("Reload config")
	quitB.SetTooltip("Stop the application")

	defer systray.Quit()
	// for initializing the config for the first time
	rch <- true
	for {
		select {
		case <-quitB.ClickedCh:
			systray.Quit()
		case <-reloadB.ClickedCh:
			rch <- true
		}
	}
}

func initClient(ch chan (bool)) {
	if !pathExists("config.toml") {
		log.Error("missing config.toml file in current directory")
		os.Exit(1)
	}
	logged := false
	for {
		<-ch
		c, err := readConfig("config.toml")
		if err != nil {
			log.Error("error when reading config", "err", err)
			continue
		}
		if logged {
			client.Logout()
		}
		client.Login(c.Id)
		err = client.SetActivity(c.asActivity())
		if err != nil {
			log.Error("error setting activity", "err", err)
			continue
		}
	}
}
