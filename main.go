package main

import "github.com/nekonenene/youtube_description_updater/src/updater"

func main() {
	updater.ParseParameters()
	updater.ReplaceDescription()
}
