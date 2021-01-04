package client

import (
	"kafka-server-client/client/consumer"
	"kafka-server-client/configuration"
)

func Start(brokers []string, topic ...string) {
	cons := consumer.New(brokers, configuration.NewKafkaConfiguration())
	consumer.Start(cons, topic...)
}