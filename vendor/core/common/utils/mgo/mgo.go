package mgo

import (
	m "gopkg.in/mgo.v2"
	"core/common/utils/log"
)

type Mmgo struct {
	S *m.Session
}

func NewMmgo() *Mmgo {
	var mgo = &Mmgo{}
	var err error
	mgo.S, err = m.Dial("mongodb://192.168.10.142:27017/")
	log.Error2Exit(err, "NewMongodbSession Dial error")
	return mgo
}

func (mgo *Mmgo) Close() {
	if mgo.S != nil {
		mgo.S.Close()
	}
}
