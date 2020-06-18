package confluent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const GoodConfigFile = "testdata/kafka.properties"
const EmptyConfigFile = "testdata/empty.properties"

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
