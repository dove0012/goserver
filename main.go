package main

import (
	"core"
	"core/work"
	"core/common/utils/log"
)

const APP_VER = "1.0.0"

func main() {
	app := core.NewApp()
	app.Name = "goserver"
	app.Usage = "run services"
	app.Version = APP_VER
	app.Servers = []*core.Work{
		work.Reckon,
	}
	app.Run()
	log.Info("App is over")
}
