package main

import (
	"flag"
	"github.com/Shopify/sarama"
	"kafka-server-client/client"
	"kafka-server-client/server"
	"log"
	"os"
)

var (
	brokers = []string { "192.168.99.100:9092" }
	topic = "chart1"
	topics = []string { topic, "chart2" }
)

func main() {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	act := flag.String("act", "consumer", "producer or consumer")
	flag.Parse()

	switch *act {
	case "producer":
		server.Start(brokers, topic)
	case "consumer":
		client.Start(brokers, topics...)
	}
}