package configuration

import "github.com/Shopify/sarama"

func NewKafkaConfiguration() *sarama.Config {
	conf := sarama.NewConfig()
	conf.ClientID = "kafka-server-client"
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Consumer.Return.Errors = true
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_1_0

	return conf
}