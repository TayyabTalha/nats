package main

import (
	"log"
	"strings"

	"github.com/nats-io/nats"
)

func main() {
	servers := []string{"nats://localhost:1222",
		"nats://localhost:1223",
		"nats://demo.nats.io",
	}

	nc, err := nats.Connect(strings.Join(servers, ","))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
}
