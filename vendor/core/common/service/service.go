package service

import (
	"core/common/utils/mgo"
)

const DB_DM_ADMIN = "dm_admin"
const DB_DM_DATA = "dm_data"

const C_HANDICAP = "handicap"
const C_HANDICAP_HASH = "handicap_hash"

type Mservice struct {
	Mgo *mgo.Mmgo
}

func NewMservice() *Mservice {
	s := &Mservice{}
	s.Mgo = mgo.NewMmgo()
	return s
}
