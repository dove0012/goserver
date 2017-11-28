package service

import (
	"core/common/utils/mgo"
)

const DB_DM_DATA = "dm_data"
const C_HANDICAP = "handicap"

type Mservice struct {
	Mgo *mgo.Mmgo
}

func NewMservice() *Mservice {
	s := &Mservice{}
	s.Mgo = mgo.NewMmgo()
	return s
}
