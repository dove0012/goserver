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
	reckonHandler.service.GetHandicapById(msgs.Han_id, &reckonHandler.handicap)
}

func (reckonHandler *ReckonHandler) requestReckon(msgs *model.Msgs) {
	if reckonHandler.handicap.Han_id > 0 {
		log.Info("[Reckon " + convert.ToStr(msgs.Han_id) + "] starting")
	} else {
		log.Error2Exit(errors.New("handicap["+convert.ToStr(msgs.Han_id)+"] not found in db"), "error")
	}
}
