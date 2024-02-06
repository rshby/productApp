package main

import (
	"fmt"
	"productApp/app/config"
)

func main() {
	// load config
	cfg := config.NewConfigApp()

	appCfg := cfg.GetConfig()
	fmt.Println(appCfg)
}
