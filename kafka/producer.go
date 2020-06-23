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

package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Extracted Kafka producer interface for purposes of composition and mocking
type Producer interface {
	String() string
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
	Events() chan kafka.Event
	Logs() chan kafka.LogEvent
	ProduceChannel() chan *kafka.Message
	Len() int
	Flush(timeoutMs int) int
	Close()
	Purge(flags int) error
	GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error)
	QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (low, high int64, err error)
	OffsetsForTimes(times []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error)
	GetFatalError() error
	TestFatalError(code kafka.ErrorCode, str string) kafka.ErrorCode
	SetOAuthBearerToken(oauthBearerToken kafka.OAuthBearerToken) error
	SetOAuthBearerTokenFailure(errstr string) error
	InitTransactions(ctx context.Context) error
	BeginTransaction() error
	SendOffsetsToTransaction(ctx context.Context, offsets []kafka.TopicPartition, consumerMetadata *kafka.ConsumerGroupMetadata) error
	CommitTransaction(ctx context.Context) error
	AbortTransaction(ctx context.Context) error

	GetTarget() *kafka.Producer
}

// Simple decorator implementation of Producer interface
type ProducerImpl struct {
	Target *kafka.Producer
}

func (p ProducerImpl) GetTarget() *kafka.Producer {
	return p.Target
}

func (p ProducerImpl) String() string {
	return p.Target.String()
}

func (p ProducerImpl) Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error {
	return p.Target.Produce(msg, deliveryChan)
}

func (p ProducerImpl) Events() chan kafka.Event {
	return p.Target.Events()
}

func (p ProducerImpl) Logs() chan kafka.LogEvent {
	return p.Target.Logs()
}

func (p ProducerImpl) ProduceChannel() chan *kafka.Message {
	return p.Target.ProduceChannel()
}

func (p ProducerImpl) Len() int {
	return p.Target.Len()
}

func (p ProducerImpl) Flush(timeoutMs int) int {
	return p.Target.Flush(timeoutMs)
}

func (p ProducerImpl) Close() {
	p.Target.Close()
}

func (p ProducerImpl) Purge(flags int) error {
	return p.Target.Purge(flags)
}

func (p ProducerImpl) GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error) {
	return p.Target.GetMetadata(topic, allTopics, timeoutMs)
}

func (p ProducerImpl) QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (low, high int64, err error) {
	return p.Target.QueryWatermarkOffsets(topic, partition, timeoutMs)
}

func (p ProducerImpl) OffsetsForTimes(times []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error) {
	return p.Target.OffsetsForTimes(times, timeoutMs)
}

func (p ProducerImpl) GetFatalError() error {
	return p.Target.GetFatalError()
}

func (p ProducerImpl) TestFatalError(code kafka.ErrorCode, str string) kafka.ErrorCode {
	return p.Target.TestFatalError(code, str)
}

func (p ProducerImpl) SetOAuthBearerToken(oauthBearerToken kafka.OAuthBearerToken) error {
	return p.Target.SetOAuthBearerToken(oauthBearerToken)
}

func (p ProducerImpl) SetOAuthBearerTokenFailure(errstr string) error {
	return p.Target.SetOAuthBearerTokenFailure(errstr)
}

func (p ProducerImpl) InitTransactions(ctx context.Context) error {
	return p.Target.InitTransactions(ctx)
}

func (p ProducerImpl) BeginTransaction() error {
	return p.Target.BeginTransaction()
}

func (p ProducerImpl) SendOffsetsToTransaction(ctx context.Context, offsets []kafka.TopicPartition, consumerMetadata *kafka.ConsumerGroupMetadata) error {
	return p.Target.SendOffsetsToTransaction(ctx, offsets, consumerMetadata)
}

func (p ProducerImpl) CommitTransaction(ctx context.Context) error {
	return p.Target.CommitTransaction(ctx)
}

func (p ProducerImpl) AbortTransaction(ctx context.Context) error {
	return p.Target.AbortTransaction(ctx)
}
