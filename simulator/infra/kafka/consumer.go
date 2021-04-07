package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumer holds all consumer logic and settings of Apache Kafka connections/
// Also has a Message channel which is a channel where the messages are going to be pushed

// criando um struct para receber mensagens topics do Kafka
type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

// NewKafkaConsumer creates a new KafkaConsumer struct with its message channel as dependency
//  Toda vez que criar um consumer ja tem esse construct ja feito
func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

//  Funcao para ficar consumindo algo no KAFKA e ele fica no loop infinito
func (k *KafkaConsumer) Consume() {
	// parametros para colocar variaveis de ambiente
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}
	// agora com configMap criar um consumer
	c, err := ckafka.NewConsumer(configMap)
	// erro mesmo nao podendo ter erro
	if err != nil {
		log.Fatalf("error consuming kafka message:" + err.Error())
	}
	//  e agora criadno um consumer iremos criar topics com variaveis de ambiente
	topics := []string{os.Getenv("KafkaReadTopic")}
	// e os topics irao ser sobrecristos com as mensagens que iremos criar
	c.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")
	// com loop infinito
	for {
		msg, err := c.ReadMessage(-1)
		//  se nao receber nenhum erro vai ter mensagem do kafka enviada para um canal que iremos usar depois no programa todo
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
