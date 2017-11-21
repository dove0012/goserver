package utils

import (
	"fmt"
	"runtime"
)

var Log = &Mlog{}

type Mlog struct {}

func (log *Mlog)FailOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
		runtime.Goexit()
	}
}
