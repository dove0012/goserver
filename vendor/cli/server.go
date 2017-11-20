package cli

import "time"

type Server struct {
	Name   string
	Usage  string
	Before BeforeFunc
	After  AfterFunc
	Reboot bool
	RebootTime time.Duration
	Action interface{}
}
