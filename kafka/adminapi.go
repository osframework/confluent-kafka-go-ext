package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

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
