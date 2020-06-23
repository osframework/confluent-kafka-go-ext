package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

// Extracted Kafka consumer interface for purposes of composition and mocking
type Consumer interface {
	String() string
	Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error
	SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) (err error)
	Unsubscribe() (err error)
	Assign(partitions []kafka.TopicPartition) (err error)
	Unassign() (err error)
	Commit() ([]kafka.TopicPartition, error)
	CommitMessage(m *kafka.Message) ([]kafka.TopicPartition, error)
	CommitOffsets(offsets []kafka.TopicPartition) ([]kafka.TopicPartition, error)
	StoreOffsets(offsets []kafka.TopicPartition) (storedOffsets []kafka.TopicPartition, err error)
	Seek(partition kafka.TopicPartition, timeoutMs int) error
	Poll(timeoutMs int) (event kafka.Event)
	Events() chan kafka.Event
	Logs() chan kafka.LogEvent
	ReadMessage(timeout time.Duration) (*kafka.Message, error)
	Close() (err error)
	GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error)
	QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (low, high int64, err error)
	GetWatermarkOffsets(topic string, partition int32) (low, high int64, err error)
	OffsetsForTimes(times []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error)
	Subscription() (topics []string, err error)
	Assignment() (partitions []kafka.TopicPartition, err error)
	Committed(partitions []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error)
	Position(partitions []kafka.TopicPartition) (offsets []kafka.TopicPartition, err error)
	Pause(partitions []kafka.TopicPartition) (err error)
	Resume(partitions []kafka.TopicPartition) (err error)
	SetOAuthBearerToken(oauthBearerToken kafka.OAuthBearerToken) error
	SetOAuthBearerTokenFailure(errstr string) error
	GetConsumerGroupMetadata() (*kafka.ConsumerGroupMetadata, error)
}

// Simple decorator implementation of Consumer interface
type ConsumerImpl struct {
	Target *kafka.Consumer
}

func (c ConsumerImpl) String() string {
	return c.Target.String()
}

func (c ConsumerImpl) Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error {
	return c.Target.Subscribe(topic, rebalanceCb)
}

func (c ConsumerImpl) SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) (err error) {
	return c.Target.SubscribeTopics(topics, rebalanceCb)
}

func (c ConsumerImpl) Unsubscribe() (err error) {
	return c.Target.Unsubscribe()
}

func (c ConsumerImpl) Assign(partitions []kafka.TopicPartition) (err error) {
	return c.Target.Assign(partitions)
}

func (c ConsumerImpl) Unassign() (err error) {
	return c.Target.Unassign()
}

func (c ConsumerImpl) Commit() ([]kafka.TopicPartition, error) {
	return c.Target.Commit()
}

func (c ConsumerImpl) CommitMessage(m *kafka.Message) ([]kafka.TopicPartition, error) {
	return c.Target.CommitMessage(m)
}

func (c ConsumerImpl) CommitOffsets(offsets []kafka.TopicPartition) ([]kafka.TopicPartition, error) {
	return c.Target.CommitOffsets(offsets)
}

func (c ConsumerImpl) StoreOffsets(offsets []kafka.TopicPartition) (storedOffsets []kafka.TopicPartition, err error) {
	return c.Target.StoreOffsets(offsets)
}

func (c ConsumerImpl) Seek(partition kafka.TopicPartition, timeoutMs int) error {
	return c.Target.Seek(partition, timeoutMs)
}

func (c ConsumerImpl) Poll(timeoutMs int) (event kafka.Event) {
	return c.Target.Poll(timeoutMs)
}

func (c ConsumerImpl) Events() chan kafka.Event {
	return c.Target.Events()
}

func (c ConsumerImpl) Logs() chan kafka.LogEvent {
	return c.Target.Logs()
}

func (c ConsumerImpl) ReadMessage(timeout time.Duration) (*kafka.Message, error) {
	return c.Target.ReadMessage(timeout)
}

func (c ConsumerImpl) Close() (err error) {
	return c.Target.Close()
}

func (c ConsumerImpl) GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error) {
	return c.Target.GetMetadata(topic, allTopics, timeoutMs)
}

func (c ConsumerImpl) QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (low, high int64, err error) {
	return c.Target.QueryWatermarkOffsets(topic, partition, timeoutMs)
}

func (c ConsumerImpl) GetWatermarkOffsets(topic string, partition int32) (low, high int64, err error) {
	return c.Target.GetWatermarkOffsets(topic, partition)
}

func (c ConsumerImpl) OffsetsForTimes(times []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error) {
	return c.Target.OffsetsForTimes(times, timeoutMs)
}

func (c ConsumerImpl) Subscription() (topics []string, err error) {
	return c.Target.Subscription()
}

func (c ConsumerImpl) Assignment() (partitions []kafka.TopicPartition, err error) {
	return c.Target.Assignment()
}

func (c ConsumerImpl) Committed(partitions []kafka.TopicPartition, timeoutMs int) (offsets []kafka.TopicPartition, err error) {
	return c.Target.Committed(partitions, timeoutMs)
}

func (c ConsumerImpl) Position(partitions []kafka.TopicPartition) (offsets []kafka.TopicPartition, err error) {
	return c.Target.Position(partitions)
}

func (c ConsumerImpl) Pause(partitions []kafka.TopicPartition) (err error) {
	return c.Target.Pause(partitions)
}

func (c ConsumerImpl) Resume(partitions []kafka.TopicPartition) (err error) {
	return c.Target.Resume(partitions)
}

func (c ConsumerImpl) SetOAuthBearerToken(oauthBearerToken kafka.OAuthBearerToken) error {
	return c.Target.SetOAuthBearerToken(oauthBearerToken)
}

func (c ConsumerImpl) SetOAuthBearerTokenFailure(errstr string) error {
	return c.Target.SetOAuthBearerTokenFailure(errstr)
}

func (c ConsumerImpl) GetConsumerGroupMetadata() (*kafka.ConsumerGroupMetadata, error) {
	return c.Target.GetConsumerGroupMetadata()
}
