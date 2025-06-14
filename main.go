package main

import "fmt"

type sendstrategy interface {
	send(m *Message)
}

type email struct {
}

func (l *email) send(c *Message) {
	fmt.Println("sent by email strategy")
}

type sms struct {
}

func (l *sms) send(c *Message) {
	fmt.Println("Sent by sms strategy")
}

type carrierpigeon struct {
}

func (l *carrierpigeon) send(c *Message) {
	fmt.Println("sent by carrier pigeon strategy")
}

type Message struct {
	text         string
	sendstrategy sendstrategy
}

func initMessage(e sendstrategy) *Message {
	return &Message{
		text:         "test",
		sendstrategy: e,
	}
}

func (c *Message) setSendAlgo(e sendstrategy) {
	c.sendstrategy = e
}

func (c *Message) send() {
	c.sendstrategy.send(c)
}

func main() {
	lfu := &sms{}
	Message := initMessage(lfu)

	Message.send()
	Message.send()
	Message.send()

	lru := &email{}
	Message.setSendAlgo(lru)
	Message.send()

	fifo := &carrierpigeon{}
	Message.setSendAlgo(fifo)

	Message.send()
}
