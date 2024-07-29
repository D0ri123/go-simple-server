package main

import (
	"flag"
	"go-server/init/cmd"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	cmd.NewCmd(*configPathFlag)
}
