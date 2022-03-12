# learn-kafka
Multiple brokers of Kafka in Go 

**Kafka helpfull commands**

To Create a topic

    ./bin/kafka-topics.sh --bootstrap-server localhost:9092,localhost:9093,localhost:9094,localhost:9095,localhost:9096,localhost:9097 --partitions 1 --replication-factor 6 --create MultiBrokerTopic 

To describe the topic

    ./bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092,localhost:9093,localhost:9094,localhost:9095,localhost:9096,localhost:9097 --all-groups --describe

To run consumer

    ./bin/kafka-console-consumer.sh -bootstrap-server localhost:9092,localhost:9093,localhost:9094,localhost:9095,localhost:9096,localhost:9097 --from-beginning --topic MultiBrokerTopic

    ./bin/kafka-console-consumer.sh -bootstrap-server localhost:9092,localhost:9093,localhost:9094,localhost:9095,localhost:9096,localhost:9097 --from-beginning --topic MultiBrokerTopic --from-beginning

To run producer

    /bin/kafka-console-producer.sh --broker-list localhost:9092,localhost:9093,localhost:9094,localhost:9095,localhost:9096,localhost:9097 --topic MultiBrokerTopic

Producer to feed data from file

    /bin/kafka-console-producer.sh --broker-list localhost:9092,localhost:9093,localhost:9094,localhost:9095,localhost:9096,localhost:9097 --topic MultiBrokerTopic < /home/rajat/Downloads/data.txt