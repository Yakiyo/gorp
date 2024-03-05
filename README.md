# gorp

A simple rich presence client for discord. This is a rewrite of [crp](https://github.com/Yakiyo/crp). While both of
them do the same thing, gorp is a tray application while crp is a console client. When running crp, it opens up a 
terminal window, but gorp doesn't open up any windows, but rather shows up in the system tray. This lets the app
run in the background without cluttering the taskbar.

Currently prebuilt binaries are available only for windows. This is due to the [systray](https://github.com/fyne-io/systray) 
library being hard to cross-compile. Before running gorp, make sure there is a `config.toml` file located in the same
directory as the app. A demo config file with full documentation can be found at [`./config.toml`](config.toml).

## Installation
Just download the archive from [release](https://github.com/Yakiyo/gorp/releases) section and unzip it. Also copy and
paste [`./config.toml`](config.toml) file into the same folder as the app. You can edit it as you want. After that,
just run the app. You should see it in the system tray. For stopping the app, click on the app icon on the tray and press
quit. Press reload to reload the config file.
