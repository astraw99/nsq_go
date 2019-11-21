package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1000)

	config := nsq.NewConfig()
	c, _ := nsq.NewConsumer("testTopic", "ch", config)
	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %s", message.Body)
		wg.Done()
		return nil
	}))

	// 1.直连nsqd
	// err := c.ConnectToNSQD("127.0.0.1:4150")

	// 2.通过 nsqlookupd 服务发现
	err := c.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Panic(err)
	}
	wg.Wait()
}
