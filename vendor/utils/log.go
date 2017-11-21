package utils

import (
	"fmt"
	"runtime"
)

var Log = &Mlog{}

type Mlog struct {}

func (log *Mlog) Error2Exit(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
		runtime.Goexit()
	}
}

func (log *Mlog)Info(msg string)  {
	fmt.Printf("log info: %s\n", msg)
}

func (log *Mlog)Warn(msg string)  {
	fmt.Printf("log warning: %s\n", msg)
}

func (log *Mlog)Debug(msg string)  {
	fmt.Printf("log debug: %s\n", msg)
}

func (log *Mlog)Error(msg string)  {
	fmt.Printf("log error: %s\n", msg)
}
