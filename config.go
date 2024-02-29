package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// a config to a rich presence instance
type Config struct {
	// app ID
	Id string
	// app state
	State string
	// details of the current state of app
	Details string
	// starting timestamp (optional)
	StartTime string
	// ending timestamp (optional)
	EndTime string
	// clickable buttons,
	// can only have two of them
	Buttons [2]button
	// images to show, the first item
	// is always the large image
	Images [2]image
}

// a clickable button, `label` is the text shown
// and `url` is the link to where it points
type button struct {
	Label string
	Url   string
}

// an image to show, `name` is the snowflake of the
// asset, while `tooltip` is the text to show on hover
type image struct {
	Name    string
	Tooltip string
}

// Validate the config
func (c *Config) Validate() error {
	reqs := map[string]string{
		"id":      c.Id,
		"details": c.Details,
		"state":   c.State,
	}
	for k, v := range reqs {
		if v == "" {
			return fmt.Errorf("required field `%v` is missing", k)
		}
	}
	// find if any button has a `label` but not a url
	invalid, button := satisfies[button](c.Buttons[:], func(b button) bool {
		return b.Label != "" && b.Url == ""
	})
	if invalid {
		return fmt.Errorf("button element has label `%v` but is missing a url, which is required", button.Label)
	}

	// find if any image has a tooltip value but no name
	invalid, image := satisfies[image](c.Images[:], func(i image) bool {
		return i.Tooltip != "" && i.Name == ""
	})
	if invalid {
		return fmt.Errorf("image element has tooltip `%v` but is missing a name, which is required", image.Tooltip)
	}
	return nil
}

// check if any element of `arr` satisfies `f`
func satisfies[T any](arr []T, f func(T) bool) (bool, *T) {
	for _, v := range arr {
		if f(v) {
			return true, &v
		}
	}
	return false, nil
}

// reads content from `path` into a `Config`
func readConfig(path string) (Config, error) {
	conf := Config{}
	content, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}
	err = toml.Unmarshal(content, &conf)
	return conf, err
}
