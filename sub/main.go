package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	nc.Subscribe("x", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})
	defer nc.Close()

	for {

	}
}
