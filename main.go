package main

import (
	"cli"
	"goserver/ser"
	"utils"
)

const APP_VER = "1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "goserver"
	app.Usage = "run services"
	app.Version = APP_VER
	app.Servers = []*cli.Server{
		ser.Reckon,
	}
	app.Run()
	utils.Log.Info("App is over")
}
