package main

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
