// Package confluent provides functions for creating Kafka consumers and
// producers connected to Confluent Cloud.
package confluent

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/magiconair/properties"
	log "github.com/sirupsen/logrus"
	"time"
)

// Create a new Kafka consumer, using the specified configuration settings. An
// error will be returned if the given configuration does not provide a consumer
// group ID, or if the consumer cannot connect to the specified host for any
// reason.
func NewConsumer(config map[string]string) (*kafka.Consumer, error) {
	configMap := createBasicConfigMap(config)
	err := configMap.SetKey("group.id", config["group.id"])
	err = configMap.SetKey("auto.offset.reset", config["auto.offset.reset"])
	if nil != err {
		return nil, fmt.Errorf("failed to configure Kafka consumer: %w", err)
	}
	return kafka.NewConsumer(&configMap)
}

// Create a new Kafka producer, using the specified configuration settings. An
// error will be returned if the producer cannot connect to the specified host
// for any reason.
func NewProducer(config map[string]string) (*kafka.Producer, error) {
	configMap := createBasicConfigMap(config)
	return kafka.NewProducer(&configMap)
}

// Read the Confluent Cloud configuration settings from the file at the given
// path. This function will panic if settings cannot be fully read from the file
// for any reason.
func ReadConfluentCloudConfig(configFile string) map[string]string {
	configMap := make(map[string]string)
	properties.ErrorHandler = properties.PanicHandler
	props := properties.MustLoadFile(configFile, properties.UTF8)
	var ok bool
	for idx, key := range props.Keys() {
		if configMap[key], ok = props.Get(key); !ok {
			log.Errorf("did not load configuration[%d] value for '%s'", idx, key)
		}
	}
	return configMap
}

func createBasicConfigMap(properties map[string]string) kafka.ConfigMap {
	configMap := kafka.ConfigMap{}
	for key, value := range properties {
		_ = configMap.SetKey(key, value)
	}
	return configMap
}

// CreateTopic creates a topic using the Admin Client API.
func CreateTopic(producer *kafka.Producer, topic string) {

	adminClient, err := kafka.NewAdminClientFromProducer(producer)
	if err != nil {
		log.Fatalf("Failed to create new admin client from producer: %s", err)
	}

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for adminClient result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create topics on cluster.
	// Set Admin options to wait up to 60s for the operation to finish on the remote cluster
	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		log.Fatalf("ParseDuration(60s): %s", err)
	}
	results, err := adminClient.CreateTopics(
		ctx,
		// Multiple topics can be created simultaneously
		// by providing more TopicSpecification structs here.
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 3}},
		// Admin options
		kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		log.Fatalf("Admin Client request error: %v\n", err)
	}
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError && result.Error.Code() != kafka.ErrTopicAlreadyExists {
			log.Fatalf("Failed to create topic: %v\n", result.Error)
		}
		log.Infof("Created topic: %v\n", result)
	}

	adminClient.Close()
}
