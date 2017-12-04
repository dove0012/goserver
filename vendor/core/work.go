package core

import "time"

type Work struct {
	Name   string
	Usage  string
	Before BeforeFunc
	After  AfterFunc
	Reboot bool
	RebootTime time.Duration
	Action interface{}
}
