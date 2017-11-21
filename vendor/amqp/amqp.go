package amqp

import (
	rabbitmq "github.com/streadway/amqp"
	"errors"
	"utils"
)

type Amqp struct {
	Url  string
	conn *rabbitmq.Connection
	ch   *rabbitmq.Channel
	Qd   QueueDeclare
	C    Consume
}

type QueueDeclare struct {
	Name      string
	Durable   bool
	AutoAck   bool
	Exclusive bool
	NoWait    bool
	Arguments map[string]interface{}
}

type Consume struct {
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Arguments map[string]interface{}
}

func NewAmqp() *Amqp {
	return &Amqp{}
}

func (amqp *Amqp) Close() {
	if amqp.conn != nil {
		amqp.conn.Close()
	}
}

func (amqb *Amqp) initMq() {
	var err error
	if amqb.Url == "" {
		utils.Log.FailOnError(errors.New("amqp url is empty"), "error")
	}
	amqb.conn, err = rabbitmq.Dial(amqb.Url)
	utils.Log.FailOnError(err, "rabbitmq.Dial error")
	amqb.ch, err = amqb.conn.Channel()
	utils.Log.FailOnError(err, "conn.Channel error")
}

func (amqp *Amqp) Receive() (<-chan rabbitmq.Delivery) {
	amqp.initMq()

	q, err := amqp.ch.QueueDeclare(
		amqp.Qd.Name,      // name
		amqp.Qd.Durable,   // durable
		amqp.Qd.AutoAck,   // delete when usused
		amqp.Qd.Exclusive, // exclusive
		amqp.Qd.NoWait,    // no-wait
		amqp.Qd.Arguments, // arguments
	)
	utils.Log.FailOnError(err, "ch.QueueDeclare error")

	msgs, err := amqp.ch.Consume(
		q.Name,           // queue
		amqp.C.Consumer,  // consumer
		amqp.C.AutoAck,   // auto-ack
		amqp.C.Exclusive, // exclusive
		amqp.C.NoLocal,   // no-local
		amqp.C.NoWait,    // no-wait
		amqp.C.Arguments, // args
	)
	utils.Log.FailOnError(err, "ch.Consume error")

	return msgs
}
