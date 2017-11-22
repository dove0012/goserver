package ser

import (
	"core"
	"fmt"
	"utils/amqp"
	"utils/config"
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

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	for d := range msgs {
		go func() {
			fmt.Printf("Received a message: %s", d.Body)
			fmt.Printf("Done")
			d.Ack(false)
		}()
	}
}
