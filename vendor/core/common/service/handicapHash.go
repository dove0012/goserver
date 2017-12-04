package service

import (
	"gopkg.in/mgo.v2/bson"
	"core/common/utils/log"
	"core/common/utils/convert"
	"crypto/sha1"
	"core/common/model"
	"crypto/hmac"
	"encoding/hex"
	"fmt"
)

func (s *Mservice) GetHandicapHashById(han_id int64, result interface{}) {
	err := s.Mgo.S.DB(DB_DM_ADMIN).C(C_HANDICAP_HASH).Find(bson.M{"han_id": han_id}).One(result)
	log.Error(err, "GetHandicapById["+convert.ToStr(han_id)+"] error")
}

func (s *Mservice) HandicapHashIsValid(handicap *model.Handicap) bool {
	handicapHash := &model.HandicapHash{}
	s.GetHandicapHashById(handicap.Han_id, handicapHash)
	str := fmt.Sprintf("%d%d%d%d%s%d",
		handicap.Han_id,
		handicap.Han_type,
		handicap.Return_rate,
		handicap.Result,
		handicap.Resulttime.Format("2006-01-02 15:04:05"),
		handicap.Status,
	)
	return s.HashString(str) == handicapHash.Hash1
}

func (s *Mservice) HashString(str string) string {
	h := hmac.New(sha1.New, []byte(model.HANDICAP_HASH_KEY))
	h.Write([]byte(str))
	str = hex.EncodeToString(h.Sum(nil))
	log.Info("--------------" + str + "--------------")
	return str
}
