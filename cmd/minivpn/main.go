package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.1-step0"

func main() {
	cfg := flag.String("config", "./configs/example.toml", "config path")
	flag.Parse()
	if _, err := os.Stat(*cfg); err != nil {
		fmt.Printf("mini-vpn %s start (config %s not found)\n", version, *cfg)
	} else {
		fmt.Printf("mini-vpn %s start with %s\n", version, *cfg)
	}
}
