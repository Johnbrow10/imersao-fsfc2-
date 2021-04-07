package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	route2 "github.com/codeedu/imersaofsfc2-simulator/application/route"
	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// Produce is responsible to publish the positions of each request
// Example of a json request:
// {"clientId":"1","routeId":"1"}
// {"clientId":"2","routeId":"2"}
// {"clientId":"3","routeId":"3"}
func Produce(msg *ckafka.Message) {
	//  Criando um Produce com a mansagem que ja foi capturado no infra de kafka
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	// Ultilizando UnMarshal ele converte de Json para Structs
	// E esse Json vai preencher a nova rota que foi feita em route.go
	// Esse "&" comercial e usado para preencher a variavel com algo antes dele
	json.Unmarshal(msg.Value, &route)
	// Agoracarregando as positions do struct
	route.LoadPositions()
	//  Entao irei exportar todas as posições dessa route
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	// Com um for sem incremento percorre as positions e publica as mensagens enviadas aos topicos
	for _, p := range positions {
		// publica na propriedade da env
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		// Coloca um tempo para nao ser enviado muito rapido
		time.Sleep(time.Millisecond * 500)
	}
}
