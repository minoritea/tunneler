package main

import (
	"flag"
	"log"
)

func main() {
	var confpath string
	flag.StringVar(&confpath, "c", "./tunneler.toml", "Path to config file(shorthand)")
	flag.StringVar(&confpath, "config", "./tunneler.toml", "Path to config file")
	var quiet bool
	flag.BoolVar(&quiet, "q", false, "Don't show the config content")
	flag.Parse()

	config, err := LoadConfig(confpath, !quiet)
	if err != nil {
		log.Println("Failed to load a config file:", confpath, err)
		return
	}
	execute(config)
}
