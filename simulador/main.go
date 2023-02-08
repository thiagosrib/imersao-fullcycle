package main

import (
	"fmt"
	kafka2 "github.com/thiagosrib/imersaofullcycle2-simulator/application/kafka"
	"github.com/thiagosrib/imersaofullcycle2-simulator/infra/kafka"
	// ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	// // para enviar msg, é utilizado o códito abaido
	// // para ficar lendo a msg pelo console do kafka, entrar no container do kafka (kafka-kafka-1) e digita: 
	// //     kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-direction
	// //     kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal
	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("ola", "route.new-direction", producer)

	// // é necessário esse loop infinito para que dê tempo de o kafka receber a msg
	// for {
	// 	_ = 1
	// }

	// para receber msg, é utilizado o códito abaido
	// // para enviar uma msg direto pelo console do kafka executar o comando abaixo:
	// //    kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}