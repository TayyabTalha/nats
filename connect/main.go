package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	// Default port is used 4222
	// nc, err := nats.Connect("localhost")
	// Use default nats://localhost:4222 & setting 10 sec time out
	nc, err := nats.Connect(nats.DefaultURL, nats.Timeout(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
}
