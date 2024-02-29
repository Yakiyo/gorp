package main

import (
	_ "embed"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/getlantern/systray"
	"github.com/hugolgst/rich-go/client"
)

// global declarations to different menu items
// var errMenu, reloadB, quitB *systray.MenuItem

var errCh = make(chan error)

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
	systray.SetIcon(icon)
	systray.SetTitle("gorp")
	systray.SetTooltip("Rich Presence Client for Discord")
	reloadB := systray.AddMenuItem("Reload", "Reload config")
	quitB := systray.AddMenuItem("Quit", "Stop application")
	errMenu := systray.AddMenuItem("Errors", "Error detected")
	errMenu.Hide()
	connect()

	go func() {
		for {
			select {
			case <-quitB.ClickedCh:
				systray.Quit()
			case err := <-errCh:
				errMenu.SetTitle(fmt.Sprintf("error: %v", err))
				errMenu.Show()
			case <-reloadB.ClickedCh:
				conf, err := readConfig(configPath)
				if err != nil {
					log.Error("Error when reloading config", "err", err)
					// setError(err)
					break
				} else {
					errMenu.Hide()
				}
				config = conf
				err = client.SetActivity(config.asActivity())
				if err != nil {
					log.Error("Error when setting activity", "err", err)
					errCh <- err
				}
			}
		}
	}()
}

func onExit() {
	log.Info("Shutting down app")
	client.Logout()
}

// func setError(err error) {
// 	errMenu.SetTitle(fmt.Sprintf("error: %v", err))
// 	errMenu.Show()
// }
