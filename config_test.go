package main

import (
	"os"
	"reflect"
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestConfig(t *testing.T) {
	config := Config{}
	content, err := os.ReadFile("config.toml")
	if err != nil {
		t.Fatal(err)
	}
	err = toml.Unmarshal(content, &config)
	if err != nil {
		t.Fatal(err)
	}
	expected := Config{
		Id:        "01234567890123456789",
		State:     "Custom Rich Presence",
		Details:   "Using gorp",
		StartTime: "now",
		EndTime:   "",
		Buttons: [2]button{
			{
				Label: "first button",
				Url:   "https://example.com",
			},
			{
				Label: "second button",
				Url:   "https://example.com",
			},
		},
		Images: [2]image{
			{
				Name:    "image-a",
				Tooltip: "The larger image",
			},
			{
				Name:    "image-b",
				Tooltip: "The smaller image",
			},
		},
	}
	if !reflect.DeepEqual(config, expected) {
		t.Error("Parsed config did not match expected value")
		t.Error("received", config)
		t.Error("expected", expected)
		t.Fail()
	}
}
