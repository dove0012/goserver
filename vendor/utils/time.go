package utils

import (
	"time"
)

var Time = &Mtime{}

type Mtime struct{}

func (t *Mtime) NowFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (t *Mtime) NowUnix() int64 {
	return time.Now().Unix()
}
