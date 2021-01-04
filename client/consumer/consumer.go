package consumer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"kafka-server-client/events"
	"log"
	"os"
	"strings"
)

func New(brokers []string, config ...*sarama.Config) sarama.Consumer {
	var conf *sarama.Config
	if len(config) > 0 {
		conf = config[0]
	}
	cons, err := sarama.NewConsumer(brokers, conf)
	if err != nil {
		log.Println("NEW FUNC")
		log.Fatal(err)
	}

	return cons
}

func Consume(consumer sarama.PartitionConsumer) {
	for {
		select {
		case err := <- consumer.Errors():
			fmt.Println(err)
		case msg := <-consumer.Messages():
			message := msg.Value
			var logs events.Event
			if err := json.Unmarshal(message, &logs); err != nil {
				fmt.Printf("can not unmarshal value: %v", string(message))
			} else {
				switch logs.Type {
				case "send":
					var event events.MessageEvent
					if err := json.Unmarshal(message, &event); err != nil {
						log.Println(err)
					}
					fmt.Printf("new message: %s", event.Message)
				default:
					fmt.Println("unknown command")
				}
			}
		}
	}
}

func Start(consumer sarama.Consumer, topic ...string) {
	defer consumer.Close()

	go ConsumeAllTopics(topic, consumer)

	fmt.Println("press ENTER to exit\n")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("terminating...")
}

func ConsumeAllTopics(topics []string, master sarama.Consumer) {
	for _, topic := range topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}
		partitions, _ := master.Partitions(topic)
		// this only consumes partition no 1, you would probably want to consume all partitions
		consumer, err := master.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)
		if nil != err {
			fmt.Printf("Topic %v Partitions: %v", topic, partitions)
			log.Fatal(err)
		}
		fmt.Println(" Start consuming topic ", topic)
		go func(topic string, consumer sarama.PartitionConsumer) {
			for {
				select {
				case consumerError := <-consumer.Errors():
					fmt.Println("consumerError: ", consumerError.Err)

				case msg := <-consumer.Messages():
					fmt.Println("Got message on topic ", topic, string(msg.Value))
				}
			}
		}(topic, consumer)
	}
}
