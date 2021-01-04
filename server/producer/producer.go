package producer

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

func New(brokers []string, config ...*sarama.Config) sarama.SyncProducer {
	var conf *sarama.Config
	if len(config) > 0 {
		conf = config[0]
	}
	if len(brokers) == 0 {
		log.Fatal("producer brokers are empty")
	}
	kafka, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		e := fmt.Sprintf("Kafka raise error: %v", err)
		log.Fatal(e)
	}

	return kafka
}

func SendMessage(producer sarama.SyncProducer, topic string, event interface{}) error {
	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(b),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Partition: %d, offset: %d\n", p, o)
	fmt.Printf("Message: %v\n", event)

	return nil
}

func Start(producer sarama.SyncProducer, topic string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-->")
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(text, "\n", "")
		args := strings.Split(text, "###")
		cmd := args[0]

		switch cmd {
		case "send":
			if len(args) == 2 {
				mess := args[1]
				event := events.NewMessageEvent(mess)
				if err := SendMessage(producer, topic, event); err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Println("Not enough arguments")
			}
		default:
			fmt.Printf("unknown command, %s\n", cmd)
		}
	}
}