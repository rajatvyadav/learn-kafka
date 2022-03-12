package main

import (
	"io/ioutil"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func main() {
	print("Hello producer!")

	// Read environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		print("Error loading .env file")
	}

	var (
		bootstrapServers = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
		kafkaTopic       = os.Getenv("KAFKA_TOPIC_MULTI_BROKER_TOPIC")
	)

	// More configuration can passed to configMap map
	// link https://github.com/edenhill/librdkafka/tree/master/CONFIGURATION.md
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		print("producer: ", err)
		return
	}

	defer producer.Close()

	// Read json data from file and send to topic
	if byteValue, err := ReadJsonFile("/home/rajat/Desktop/workspace/src/learn-kafka/data/MOCK_DATA.json"); err != nil {
		print("Error while reading data from json file: ", err)
		return
	} else {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
			Value:          byteValue,
		}, nil)
	}

}

// Read json file from path given in parameters
// return the array of type EmployeeDetails and error
func ReadJsonFile(filePath string) ([]byte, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}
