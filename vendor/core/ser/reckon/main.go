package reckon

import (
	"utils/log"
	"strconv"
)

type ReckonHandler struct {
	Hand_id int64
}

func NewReckon() *ReckonHandler {
	return &ReckonHandler{}
}

func (ReckonHandler *ReckonHandler) Run(han_id int64) {
	log.Info("[Reckon " + strconv.FormatInt(han_id, 10) + "] starting")
}
