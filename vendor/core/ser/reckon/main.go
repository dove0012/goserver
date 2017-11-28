package reckon

import (
	"core/common/utils/log"
	"core/common/service"
	"core/common/model"
	"core/common/utils/convert"
	"fmt"
)

type ReckonHandler struct {
	handicap *model.Handicap
	service *service.Mservice
}

func NewReckon() *ReckonHandler {
	return &ReckonHandler{}
}

func (reckonHandler *ReckonHandler) Run(handicap *model.Handicap) {
	log.Info("[Reckon " + convert.ToStr(handicap.Han_id) + "] starting")
	reckonHandler.service = service.NewMservice()
	defer reckonHandler.service.Mgo.Close()
	reckonHandler.requestReckon(handicap)
}

func (reckonHandler *ReckonHandler) requestReckon(handicap *model.Handicap) {
	reckonHandler.service.GetHandicapById(handicap.Han_id, &reckonHandler.handicap)
	fmt.Printf("--------handicap:%#v", reckonHandler.handicap)
}
