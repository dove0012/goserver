package ser

import (
	"core"
	"fmt"
	"utils/amqp"
	"utils/config"
	"utils/log"
	"utils/json"
	"utils/time"
	"core/ser/reckon"
	"core/ser/reckon/model"
	"strconv"
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
			startTime := time.NowUnixMilli()
			log.Info(fmt.Sprintf("Received a message: %s", d.Body))
			handicap := model.Handicap{}
			json.Unmarshal(d.Body, &handicap)
			if handicap.Han_id > 0 {
				defer func() {
					log.TimeConsuming(startTime, "[handicap "+strconv.Itoa(handicap.Han_id)+"] is over")
					d.Ack(false)
				}()
				reckon.NewReckon().Run(handicap.Han_id)
			} else {
				log.Error(errors.New("[handicap "+strconv.Itoa(handicap.Han_id)+"] is not gt zero error"), "")
			}
		}()
	}
}
