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

// Package kafka provides interfaces extracted from the core API struct types
// in the Confluent Kafka Golang client; see https://github.com/confluentinc/confluent-kafka-go
package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Extracted Kafka admin client interface for purposes of composition and mocking
type AdminClient interface {
	ClusterID(ctx context.Context) (clusterID string, err error)
	ControllerID(ctx context.Context) (controllerID int32, err error)
	CreateTopics(ctx context.Context, topics []kafka.TopicSpecification, options ...kafka.CreateTopicsAdminOption) (result []kafka.TopicResult, err error)
	DeleteTopics(ctx context.Context, topics []string, options ...kafka.DeleteTopicsAdminOption) (result []kafka.TopicResult, err error)
	CreatePartitions(ctx context.Context, partitions []kafka.PartitionsSpecification, options ...kafka.CreatePartitionsAdminOption) (result []kafka.TopicResult, err error)
	AlterConfigs(ctx context.Context, resources []kafka.ConfigResource, options ...kafka.AlterConfigsAdminOption) (result []kafka.ConfigResourceResult, err error)
	DescribeConfigs(ctx context.Context, resources []kafka.ConfigResource, options ...kafka.DescribeConfigsAdminOption) (result []kafka.ConfigResourceResult, err error)
	GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error)
	String() string
	SetOAuthBearerToken(oauthBearerToken kafka.OAuthBearerToken) error
	SetOAuthBearerTokenFailure(errstr string) error
	Close()
}
