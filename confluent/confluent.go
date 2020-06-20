// Package confluent provides functions for creating Kafka consumers and
// producers connected to Confluent Cloud.
package confluent

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/magiconair/properties"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	AdminOperationTimeout = "admin.operation.timeout"
	AutoOffsetReset       = "auto.offset.reset"
	GroupId               = "group.id"
	NumPartitions         = "topic.partitions"
	ReplicationFactor     = "topic.replication.factor"

	DefaultAdminOperationTimeout = "60s"
	DefaultNumPartitions         = "1"
	DefaultReplicationFactor     = "3"
)

// Create a new Kafka consumer, using the specified configuration settings. An
// error will be returned if the given configuration does not provide a consumer
// group ID, or if the consumer cannot connect to the specified host for any
// reason.
func NewConsumer(config map[string]string) (*kafka.Consumer, error) {
	configMap := createBasicConfigMap(config)
	settingsToValidate := [2]string{GroupId, AutoOffsetReset}
	for _, setting := range settingsToValidate {
		v, err := configMap.Get(setting, "")
		if nil != err {
			return nil, fmt.Errorf("failed to configure Kafka consumer: %v", err)
		} else if "" == v {
			return nil, fmt.Errorf("missing setting for Kafka consumer: %s", setting)
		}
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
	topics := make([]string, 1)
	topics[0] = topic

	topicCreationConfig := make(map[string]string)
	topicCreationConfig[AdminOperationTimeout] = DefaultAdminOperationTimeout
	topicCreationConfig[NumPartitions] = DefaultNumPartitions
	topicCreationConfig[ReplicationFactor] = DefaultReplicationFactor

	CreateTopics(producer, topics, topicCreationConfig)
}

// CreateTopics creates one or more topics using the Admin Client API.
func CreateTopics(producer *kafka.Producer, topics []string, config map[string]string) {
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
	if _, ok := config[AdminOperationTimeout]; !ok {
		config[AdminOperationTimeout] = DefaultAdminOperationTimeout
		log.Warnf("Set '%s' to default: %s", AdminOperationTimeout, DefaultAdminOperationTimeout)
	}
	maxDur, err := time.ParseDuration(config[AdminOperationTimeout])
	if err != nil {
		log.Fatalf("ParseDuration(%s): %s", config[AdminOperationTimeout], err)
	}

	if _, ok := config[NumPartitions]; !ok {
		config[NumPartitions] = DefaultNumPartitions
		log.Warnf("Set '%s' to default: %s", NumPartitions, DefaultNumPartitions)
	}
	numPartitions, err := strconv.Atoi(config[NumPartitions])
	if err != nil {
		log.Fatalf("ParseInt(%s): %s", config[NumPartitions], err)
	}

	if _, ok := config[ReplicationFactor]; !ok {
		config[ReplicationFactor] = DefaultReplicationFactor
		log.Warnf("Set '%s' to default: %s", ReplicationFactor, DefaultReplicationFactor)
	}
	replicationFactor, err := strconv.Atoi(config[ReplicationFactor])
	if err != nil {
		log.Fatalf("ParseInt(%s): %s", config[ReplicationFactor], err)
	}

	topicSpecs := make([]kafka.TopicSpecification, len(topics))
	for idx, topic := range topics {
		topicSpecs[idx] = kafka.TopicSpecification{
			Topic:             topic,
			NumPartitions:     numPartitions,
			ReplicationFactor: replicationFactor,
		}
	}

	results, err := adminClient.CreateTopics(
		ctx,
		// Multiple topics can be created simultaneously
		// by providing more TopicSpecification structs here.
		topicSpecs,
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
