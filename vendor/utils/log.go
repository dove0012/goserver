package utils

import (
	"fmt"
	"runtime"
)

var Log = &Mlog{}

type Mlog struct{}

func (log *Mlog) Error2Exit(err error, msg string) {
	if err != nil {
		fmt.Printf(Time.NowFormat()+"  %s: %s\n", msg, err)
		runtime.Goexit()
	}
}

func (log *Mlog) Info(msg string) {
	fmt.Printf(Time.NowFormat()+"  log info: %s\n", msg)
}

func (log *Mlog) Warn(msg string) {
	fmt.Printf(Time.NowFormat()+"  log warning: %s\n", msg)
}

func (log *Mlog) Debug(msg string) {
	fmt.Printf(Time.NowFormat()+"  log debug: %s\n", msg)
}

func (log *Mlog) Error(err error, msg string) {
	if err != nil {
		fmt.Printf(Time.NowFormat()+"  %s: %s\n", msg, err)
	}
}
