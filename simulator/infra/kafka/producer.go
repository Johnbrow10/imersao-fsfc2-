package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewKafkaProducer creates a ready to go kafka.Producer instance
//  Vai retornar um Producer do kafka responsavel para publicar mensagens do kafka quando quisermos usar
func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
	}
	//  Criando um novo Producer para poder ultilizar
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	//  retornar caso tiver tudo certo
	return p
}

// Publish is simple function created to publish new message to kafka
// Publicar topics mensagens com o producter que estar sempre criando
func Publish(msg string, topic string, producer *ckafka.Producer) error {
	// Postar uma nova mensagem
	message := &ckafka.Message{
		// o topico Partiotion que iremos usar para mandar mensagem e qual a partition que iremos usar
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		//  o valor a mensagem em si
		Value: []byte(msg),
	}
	//  Um erro caso esteja com problema nessa operação
	err := producer.Produce(message, nil)
	// E assim verificar se tem erro no retorno
	if err != nil {
		return err
	}
	//  caso nao retornar erro ele continua em branco
	return nil
}
