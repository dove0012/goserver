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
	mq := amqp.NewAmqp()
	mq.Url = "amqp://guest:guest@192.168.186.129:5672"
	mq.Qd.Name = NAME_RECKON
	msgs, err := mq.Receive()
	defer mq.Close()

	if err != nil {
		utils.FailOnError(err, "mq receive error")
	} else {
		fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		for d := range msgs {
			go func() {
				fmt.Printf("Received a message: %s", d.Body)
				fmt.Printf("Done")
				d.Ack(false)
			}()
		}
	}
}
