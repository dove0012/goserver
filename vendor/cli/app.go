package cli

import (
	"os"
	"path/filepath"
	"sync"
	"runtime"
	"time"
)

type App struct {
	Name    string
	Usage   string
	Version string
	Servers []Server
	Wg      sync.WaitGroup
}

func NewApp() *App {
	return &App{
		Name:    filepath.Base(os.Args[0]),
		Usage:   "A new cli application",
		Version: "0.0.0",
	}
}

func (app *App) Run() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for _, server := range app.Servers {
		app.Wg.Add(1)
		app.runSer(server)
	}
	app.Wg.Wait()
}

func (app *App) rebootSer(name string) {
	for _, server := range app.Servers {
		if name == server.Name {
			app.runSer(server)
		}
	}
}

func (app *App) runSer(server Server)  {
	go func() {
		if server.Reboot {
			defer func() {
				time.Sleep(time.Second * server.RebootTime)
				app.rebootSer(server.Name)
			}()
		} else {
			defer app.Wg.Done()
		}
		switch action := server.Action.(type) {
		case func(app *App):
			action(app)
		case func():
			action()
		default:
			panic("server main func error.")
		}
	}()
}
