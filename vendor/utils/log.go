package utils

import (
	"fmt"
	"runtime"
)

func FailOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
		runtime.Goexit()
	}
}
