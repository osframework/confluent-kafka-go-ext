/*
 * Copyright 2020 OSFramework Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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

	topics := []string{"test.topic"}
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

	err := createTopics(mockProducer, topics, configMap, func(p k.Producer) (k.AdminClient, error) {
		return mockAdminClient, nil
	})
	a.Nil(err)
	mockAdminClient.AssertExpectations(t)
}

func TestCreateTopics_NoAdminClient(t *testing.T) {
	mockProducer := &mocks.Producer{}

	topics := make([]string, 1)
	topics[0] = "test.topic"

	a := assert.New(t)

	configMap := ReadConfluentCloudConfig(GoodConfigFile)
	a.Contains(configMap, "bootstrap.servers", "Did not find expected property")

	err := createTopics(mockProducer, topics, configMap, func(p k.Producer) (k.AdminClient, error) {
		return nil, errors.New("expected error on AdminClient creation")
	})
	a.NotNil(err)
	a.Equal("failed to create new admin client from producer: expected error on AdminClient creation", err.Error())
}

func TestCreateTopics_AdminClientFails(t *testing.T) {
	mockProducer := &mocks.Producer{}
	mockAdminClient := &mocks.AdminClient{}

	topics := make([]string, 1)
	topics[0] = "test.topic"

	err := errors.New("expected call failure")

	mockAdminClient.On("CreateTopics", mock.Anything, mock.AnythingOfType("[]kafka.TopicSpecification"), mock.AnythingOfType("kafka.AdminOptionOperationTimeout")).Return(nil, err)
	mockAdminClient.On("Close")

	a := assert.New(t)

	configMap := ReadConfluentCloudConfig(GoodConfigFile)
	a.Contains(configMap, "bootstrap.servers", "Did not find expected property")

	err2 := createTopics(mockProducer, topics, configMap, func(p k.Producer) (k.AdminClient, error) {
		return mockAdminClient, nil
	})
	a.NotNil(err2)
	a.Equal("admin client request error: expected call failure", err2.Error())
	mockAdminClient.AssertExpectations(t)
}
