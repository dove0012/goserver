package main

import (
	"core"
	"core/ser"
	"utils/log"
)

const APP_VER = "1.0.0"

func main() {
	app := core.NewApp()
	app.Name = "goserver"
	app.Usage = "run services"
	app.Version = APP_VER
	app.Servers = []*core.Server{
		ser.Reckon,
	}
	app.Run()
	log.Info("App is over")
}
