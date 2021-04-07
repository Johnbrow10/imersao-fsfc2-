package main

import (
	"fmt"
	"log"

	kafka2 "github.com/codeedu/imersaofsfc2-simulator/application/kafka"
	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

// Ela ultiliza um pacote .DontEnv para verificar se o env estar errado ou nao assim para continuar
// e ultilizar a funcao main
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

// Funcao para inicializar o programa
// Chamando o consumer para consumir mensagem
func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
