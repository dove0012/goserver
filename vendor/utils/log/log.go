package log

import (
	"fmt"
	"runtime"
	"utils/time"
)

func Error2Exit(err error, msg string) {
	if err != nil {
		fmt.Printf(time.NowFormat()+"  %s: %s\n", msg, err)
		runtime.Goexit()
	}
}

func Info(msg string) {
	fmt.Printf(time.NowFormat()+"  log info: %s\n", msg)
}

func Warn(msg string) {
	fmt.Printf(time.NowFormat()+"  log warning: %s\n", msg)
}

func Debug(msg string) {
	fmt.Printf(time.NowFormat()+"  log debug: %s\n", msg)
}

func Error(err error, msg string) {
	if err != nil {
		fmt.Printf(time.NowFormat()+"  %s: %s\n", msg, err)
	}
}
