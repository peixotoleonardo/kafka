
# Kafka

## Quickstart

Start containers:

```bash
docker compose up -d
```

Create a topic with 3 partitions:

```bash
./kafka-topics.sh --create --topic test --bootstrap-server localhost:19092 --partitions 3
```
### Producing Events

Before we try, we need to consume the topic `test`. We can do this with script `kafka-console-consumer.sh`
inner any Kafka container.

```bash
./kafka-console-consumer.sh --topic test --bootstrap-server localhost:19092
```

#### Golang

```bash
# sends an event with the string 'teste'
KAFKA_TOPIC=test KAFKA_BROKERS=localhost:29092,localhost:39092,localhost:49092 go run golang/cmd/producer
```
