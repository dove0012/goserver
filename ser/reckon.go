package ser

import (
	"cli"
	"fmt"
	"amqp"
	"utils"
)

const NAME_RECKON = "reckon"

var Reckon = &cli.Server{
	Name:       NAME_RECKON,
	Usage:      "Bigame reckon server",
	Reboot:     true,
	RebootTime: 1,
	Action:     runReckon,
}

func runReckon() {
	cfg, err := cli.NewCfg("config.ini")
	utils.FailOnError(err, "goconfig NewCfg error")
	cfg.Section = NAME_RECKON
	mq_url, err := cfg.GetString("mq_url")
	utils.FailOnError(err, "goconfig GetString error")

	mq := amqp.NewAmqp()
	mq.Url = mq_url
	mq.Qd.Name = NAME_RECKON
	msgs, err := mq.Receive()
	defer mq.Close()
	utils.FailOnError(err, "mq receive error")

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	for d := range msgs {
		go func() {
			fmt.Printf("Received a message: %s", d.Body)
			fmt.Printf("Done")
			d.Ack(false)
		}()
	}
}
