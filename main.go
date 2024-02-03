package main

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	_ "github.com/getlantern/systray"
	_ "github.com/hugolgst/rich-go/client"
	homedir "github.com/mitchellh/go-homedir"
	_ "github.com/pelletier/go-toml/v2"
)

var (
	gerr error
	// config Config
	hdir string
)

func main() {
	if gerr != nil {
		return
	}
}

func init() {
	dir, err := homedir.Dir()
	if err != nil {
		gerr = err
		return
	}
	hdir = dir
	logfile := filepath.Join(hdir, ".gorp", "gorp.log")
	log.SetOutput(os.NewFile(uintptr(os.ModePerm), logfile))
}
