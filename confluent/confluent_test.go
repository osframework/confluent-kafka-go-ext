package confluent

import (
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	k "github.com/osframework/confluent-kafka-go-ext/kafka"
	"github.com/osframework/confluent-kafka-go-ext/kafka/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

const GoodConfigFile = "testdata/kafka.properties"
const EmptyConfigFile = "testdata/empty.properties"

type MockAdminClientCreator struct {
	AdminClient k.AdminClient
}

func (c *MockAdminClientCreator) NewAdminClientFromProducer(p k.Producer) (a k.AdminClient, err error) {
	return c.AdminClient, nil
}

type NoAdminClientCreator struct{}

func (c *NoAdminClientCreator) NewAdminClientFromProducer(p k.Producer) (a k.AdminClient, err error) {
	return nil, errors.New("Expected AdminClient creation failure")
}

func TestReadConfluentCloudConfig(t *testing.T) {
	a := assert.New(t)

	configMap := ReadConfluentCloudConfig(GoodConfigFile)
	a.Contains(configMap, "bootstrap.servers", "Did not find expected property")
}

func TestReadConfluentCloudConfig_EmptyFile(t *testing.T) {
	a := assert.New(t)
	configMap := ReadConfluentCloudConfig(EmptyConfigFile)
	a.Empty(configMap, "Expected empty configuration map")
}

func TestReadConfluentCloudConfig_EmptyFilePath(t *testing.T) {
	a := assert.New(t)
	a.Panics(func() {
		_ = ReadConfluentCloudConfig("")
	}, "Expected panic on empty file path")
}

func TestCreateBasicConfigMap(t *testing.T) {
	a := assert.New(t)

	configMap := ReadConfluentCloudConfig(GoodConfigFile)
	a.Contains(configMap, "bootstrap.servers", "Did not find expected property")

	kafkaConfigMap := createBasicConfigMap(configMap)
	a.NotNil(kafkaConfigMap)
	value, err := kafkaConfigMap.Get("bootstrap.servers", "")
	a.Nil(err)
	a.NotEmpty(value, "Did not find expected property")
}

func TestNewProducer(t *testing.T) {
	a := assert.New(t)

	configMap := ReadConfluentCloudConfig(GoodConfigFile)
	a.Contains(configMap, "bootstrap.servers", "Did not find expected property")

	kafkaProducer, err := NewProducer(configMap)
	a.Nil(err)
	a.NotNil(kafkaProducer, "Expected constructed Kafka producer")
}

func TestNewConsumer_MissingProperties(t *testing.T) {
	a := assert.New(t)

	configMap := ReadConfluentCloudConfig(GoodConfigFile)
	a.Contains(configMap, "bootstrap.servers", "Did not find expected property")

	_, err := NewConsumer(configMap)
	a.NotNil(err, "Expected error on missing consumer properties")
}

func TestNewConsumer(t *testing.T) {
	a := assert.New(t)

	configMap := ReadConfluentCloudConfig(GoodConfigFile)
	a.Contains(configMap, "bootstrap.servers", "Did not find expected property")
	configMap["group.id"] = "test_group"
	configMap["auto.offset.reset"] = "latest"

	kafkaConsumer, err := NewConsumer(configMap)
	a.Nil(err)
	a.NotNil(kafkaConsumer, "Expected constructed Kafka consumer")
}

func TestCreateTopics(t *testing.T) {
	mockProducer := &mocks.Producer{}
	mockAdminClient := &mocks.AdminClient{}
	creator := new(MockAdminClientCreator)
	creator.AdminClient = mockAdminClient

	topics := make([]string, 1)
	topics[0] = "test.topic"

	topicResults := make([]kafka.TopicResult, 1)
	topicResults[0] = kafka.TopicResult{
		Topic: topics[0],
		Error: kafka.Error{},
	}

	mockAdminClient.On("CreateTopics", mock.Anything, mock.AnythingOfType("[]kafka.TopicSpecification"), mock.AnythingOfType("kafka.AdminOptionOperationTimeout")).Return(topicResults, nil)
	mockAdminClient.On("Close")

	a := assert.New(t)

	configMap := ReadConfluentCloudConfig(GoodConfigFile)
	a.Contains(configMap, "bootstrap.servers", "Did not find expected property")

	createTopics(mockProducer, topics, configMap, creator)
	mockAdminClient.AssertExpectations(t)
}

func TestCreateTopics_NoAdminClient(t *testing.T) {
	mockProducer := &mocks.Producer{}
	creator := new(NoAdminClientCreator)

	topics := make([]string, 1)
	topics[0] = "test.topic"

	a := assert.New(t)

	configMap := ReadConfluentCloudConfig(GoodConfigFile)
	a.Contains(configMap, "bootstrap.servers", "Did not find expected property")

	a.Panics(func() {
		createTopics(mockProducer, topics, configMap, creator)
	}, "Expected panic on missing AdminClient")
}

func TestCreateTopics_AdminClientFails(t *testing.T) {
	mockProducer := &mocks.Producer{}
	mockAdminClient := &mocks.AdminClient{}
	creator := new(MockAdminClientCreator)
	creator.AdminClient = mockAdminClient

	topics := make([]string, 1)
	topics[0] = "test.topic"

	error := errors.New("Expected call failure")

	mockAdminClient.On("CreateTopics", mock.Anything, mock.AnythingOfType("[]kafka.TopicSpecification"), mock.AnythingOfType("kafka.AdminOptionOperationTimeout")).Return(nil, error)
	mockAdminClient.On("Close")

	a := assert.New(t)

	configMap := ReadConfluentCloudConfig(GoodConfigFile)
	a.Contains(configMap, "bootstrap.servers", "Did not find expected property")

	a.Panics(func() {
		createTopics(mockProducer, topics, configMap, creator)
	}, "Expected panic on missing AdminClient")
}
