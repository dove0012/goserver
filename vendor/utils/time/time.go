package time

import (
	t "time"
)

func NowFormat() string {
	return t.Now().Format("2006-01-02 15:04:05")
}

func NowUnix() int64 {
	return t.Now().Unix()
}
