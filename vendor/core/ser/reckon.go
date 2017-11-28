package ser

import (
	"core"
	"fmt"
	"core/common/utils/amqp"
	"core/common/utils/config"
	"core/common/utils/log"
	"core/common/utils/json"
	"core/common/utils/time"
	"core/common/utils/convert"
	"core/ser/reckon"
	"core/common/model"
	"errors"
)

const NAME_RECKON = "reckon"

var Reckon = &core.Server{
	Name:       NAME_RECKON,
	Usage:      "Bigame reckon server",
	Reboot:     true,
	RebootTime: 1,
	Action:     runReckon,
}

func runReckon() {
	cfg := config.NewCfg("config.ini")
	cfg.Section = NAME_RECKON

	mq := amqp.NewAmqp()
	mq.Url = cfg.GetString("mq_url")
	mq.Qd.Name = NAME_RECKON
	msgs := mq.Receive()
	defer mq.Close()

	log.Info("[*] Waiting for messages. To exit press CTRL+C")
	for d := range msgs {
		go func() {
			log.Info(fmt.Sprintf("Received a message: %s", d.Body))
			startTime := time.NowUnixMilli()
			handicap := &model.Handicap{}
			json.Unmarshal(d.Body, &handicap)
			defer func() {
				log.TimeConsuming(startTime, "[handicap "+convert.ToStr(handicap.Han_id)+"] is over")
				d.Ack(false)
			}()
			if handicap.Han_id > 0 {
				reckon.NewReckon().Run(handicap)
			} else {
				log.Error(errors.New("[handicap "+convert.ToStr(handicap.Han_id)+"] is not gt zero error"), "")
			}
		}()
	}
}
