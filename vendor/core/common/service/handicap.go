package service

import (
	"gopkg.in/mgo.v2/bson"
	"core/common/utils/log"
	"core/common/utils/convert"
)

func (s *Mservice) GetHandicapById(han_id int64, result interface{}) {
	err := s.Mgo.S.DB(DB_DM_DATA).C(C_HANDICAP).Find(bson.M{"han_id": han_id}).One(result)
	log.Error(err, "GetHandicapById["+convert.ToStr(han_id)+"] error")
}
