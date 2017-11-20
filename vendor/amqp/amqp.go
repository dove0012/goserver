package amqp

import (
	rabbitmq "github.com/streadway/amqp"
	"errors"
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

func (amqb *Amqp) initMq() error {
	var err error
	if amqb.Url == "" {
		return errors.New("amqp url is empty")
	}
	amqb.conn, err = rabbitmq.Dial(amqb.Url)
	if err != nil {
		return err
	}
	amqb.ch, err = amqb.conn.Channel()
	if err != nil {
		return err
	}
	return nil
}

func (amqp *Amqp) Receive() (<-chan rabbitmq.Delivery, error) {
	err := amqp.initMq()
	if err != nil {
		return nil, err
	}

	q, err := amqp.ch.QueueDeclare(
		amqp.Qd.Name,      // name
		amqp.Qd.Durable,   // durable
		amqp.Qd.AutoAck,   // delete when usused
		amqp.Qd.Exclusive, // exclusive
		amqp.Qd.NoWait,    // no-wait
		amqp.Qd.Arguments, // arguments
	)

	msgs, err := amqp.ch.Consume(
		q.Name,           // queue
		amqp.C.Consumer,  // consumer
		amqp.C.AutoAck,   // auto-ack
		amqp.C.Exclusive, // exclusive
		amqp.C.NoLocal,   // no-local
		amqp.C.NoWait,    // no-wait
		amqp.C.Arguments, // args
	)

	return msgs, nil
}
