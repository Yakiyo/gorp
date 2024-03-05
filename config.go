package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hugolgst/rich-go/client"
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

func (c *Config) asActivity() client.Activity {
	act := client.Activity{
		Details:    c.Details,
		State:      c.State,
		Timestamps: &client.Timestamps{},
	}
	if c.StartTime != "" {
		time := asTime(c.StartTime)
		act.Timestamps.Start = &time
	}
	if c.EndTime != "" {
		time := asTime(c.EndTime)
		act.Timestamps.End = &time
	}
	// if theres at least 1 image
	if len(c.Images) > 0 {
		act.LargeImage = c.Images[0].Name
		act.LargeText = c.Images[0].Tooltip
	}
	// theres 2 images
	if len(c.Images) == 2 {
		act.SmallImage = c.Images[1].Name
		act.SmallText = c.Images[1].Tooltip
	}

	for _, button := range c.Buttons {
		act.Buttons = append(act.Buttons, &client.Button{
			Label: button.Label,
			Url:   button.Url,
		})
	}
	return act
}

// reads content from `path` into a `Config`
func readConfig(path string) (Config, error) {
	conf := Config{}
	content, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}
	err = toml.Unmarshal(content, &conf)
	if err != nil {
		return conf, err
	}
	err = conf.Validate()
	return conf, err
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

func asTime(s string) time.Time {
	if s == "now" {
		return time.Now()
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Now()
	}
	return time.Unix(n, 0)
}
