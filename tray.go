package main

import (
	_ "embed"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/getlantern/systray"
)

// global declarations to different menu items
var errMenu, reloadB, quitB *systray.MenuItem

func initTray() {
	systray.SetTemplateIcon(icon, icon)
	systray.SetTitle("gorp")
	systray.SetTooltip("Rich Presence Client for Discord")
	reloadB = systray.AddMenuItem("Reload", "Reload config")
	quitB = systray.AddMenuItem("Quit", "Stop application")
	errMenu = systray.AddMenuItem("Errors", "Error detected")
	errMenu.Hide()
}

func onReady() {
	go func() {
		for {
			select {
			case <-quitB.ClickedCh:
				systray.Quit()
			case <-reloadB.ClickedCh:
				conf, err := readConfig(configPath)
				if err != nil {
					log.Error("Error when reloading config", "err", err)
					setError(err)
					break
				} else {
					errMenu.Hide()
				}
				config = conf
				configChan <- 0
			}
		}
	}()
}

func onExit() {
	log.Info("Shutting down app")
	// TODO: shutdown rpc client
}

func setError(err error) {
	errMenu.SetTitle(fmt.Sprintf("error: %v", err))
	errMenu.Show()
}
