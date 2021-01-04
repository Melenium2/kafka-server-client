package server

import (
	"kafka-server-client/configuration"
	"kafka-server-client/server/producer"
)

func Start(brokers []string, topic string) {
	prod := producer.New(brokers, configuration.NewKafkaConfiguration())
	producer.Start(prod, topic)
}