package config

type Config struct {
	Kafka Kafka
}

type Kafka struct {
	Topic string
	Host  string
}
