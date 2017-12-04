package reckon

import (
	"core/common/utils/log"
	"core/common/service"
	"core/common/model"
	"core/common/utils/convert"
	"errors"
)

type ReckonHandler struct {
	handicap *model.Handicap
	service  *service.Mservice
}

func NewReckonHandler() *ReckonHandler {
	return &ReckonHandler{}
}

func (reckonHandler *ReckonHandler) Init() *ReckonHandler {
	reckonHandler.service = service.NewMservice()
	return reckonHandler
}

func (reckonHandler *ReckonHandler) Destroy() {
	if reckonHandler.service != nil {
		reckonHandler.service.Mgo.Close()
	}
}

func (reckonHandler *ReckonHandler) Run(msgs *model.Msgs) {
	reckonHandler.requestReckon(msgs)
	log.Info("[Reckon " + convert.ToStr(msgs.Han_id) + "] starting")
}

func (reckonHandler *ReckonHandler) requestReckon(msgs *model.Msgs) {
	reckonHandler.service.GetHandicapById(msgs.Han_id, &reckonHandler.handicap)
	if reckonHandler.handicap.Han_id > 0 {
		if !reckonHandler.service.HandicapHashIsValid(reckonHandler.handicap) {
			log.Error2Exit(errors.New("handicap["+convert.ToStr(msgs.Han_id)+"] HandicapHashIsValid faile"), "error")
		}
	} else {
		log.Error2Exit(errors.New("handicap["+convert.ToStr(msgs.Han_id)+"] not found in db"), "error")
	}
}
