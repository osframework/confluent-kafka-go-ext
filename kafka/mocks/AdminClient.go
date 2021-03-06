// Code generated by mockery v2.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
	mock "github.com/stretchr/testify/mock"
)

// AdminClient is an autogenerated mock type for the AdminClient type
type AdminClient struct {
	mock.Mock
}

// AlterConfigs provides a mock function with given fields: ctx, resources, options
func (_m *AdminClient) AlterConfigs(ctx context.Context, resources []kafka.ConfigResource, options ...kafka.AlterConfigsAdminOption) ([]kafka.ConfigResourceResult, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, resources)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []kafka.ConfigResourceResult
	if rf, ok := ret.Get(0).(func(context.Context, []kafka.ConfigResource, ...kafka.AlterConfigsAdminOption) []kafka.ConfigResourceResult); ok {
		r0 = rf(ctx, resources, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kafka.ConfigResourceResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []kafka.ConfigResource, ...kafka.AlterConfigsAdminOption) error); ok {
		r1 = rf(ctx, resources, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *AdminClient) Close() {
	_m.Called()
}

// ClusterID provides a mock function with given fields: ctx
func (_m *AdminClient) ClusterID(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerID provides a mock function with given fields: ctx
func (_m *AdminClient) ControllerID(ctx context.Context) (int32, error) {
	ret := _m.Called(ctx)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context) int32); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePartitions provides a mock function with given fields: ctx, partitions, options
func (_m *AdminClient) CreatePartitions(ctx context.Context, partitions []kafka.PartitionsSpecification, options ...kafka.CreatePartitionsAdminOption) ([]kafka.TopicResult, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, partitions)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []kafka.TopicResult
	if rf, ok := ret.Get(0).(func(context.Context, []kafka.PartitionsSpecification, ...kafka.CreatePartitionsAdminOption) []kafka.TopicResult); ok {
		r0 = rf(ctx, partitions, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kafka.TopicResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []kafka.PartitionsSpecification, ...kafka.CreatePartitionsAdminOption) error); ok {
		r1 = rf(ctx, partitions, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTopics provides a mock function with given fields: ctx, topics, options
func (_m *AdminClient) CreateTopics(ctx context.Context, topics []kafka.TopicSpecification, options ...kafka.CreateTopicsAdminOption) ([]kafka.TopicResult, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, topics)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []kafka.TopicResult
	if rf, ok := ret.Get(0).(func(context.Context, []kafka.TopicSpecification, ...kafka.CreateTopicsAdminOption) []kafka.TopicResult); ok {
		r0 = rf(ctx, topics, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kafka.TopicResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []kafka.TopicSpecification, ...kafka.CreateTopicsAdminOption) error); ok {
		r1 = rf(ctx, topics, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTopics provides a mock function with given fields: ctx, topics, options
func (_m *AdminClient) DeleteTopics(ctx context.Context, topics []string, options ...kafka.DeleteTopicsAdminOption) ([]kafka.TopicResult, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, topics)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []kafka.TopicResult
	if rf, ok := ret.Get(0).(func(context.Context, []string, ...kafka.DeleteTopicsAdminOption) []kafka.TopicResult); ok {
		r0 = rf(ctx, topics, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kafka.TopicResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, ...kafka.DeleteTopicsAdminOption) error); ok {
		r1 = rf(ctx, topics, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeConfigs provides a mock function with given fields: ctx, resources, options
func (_m *AdminClient) DescribeConfigs(ctx context.Context, resources []kafka.ConfigResource, options ...kafka.DescribeConfigsAdminOption) ([]kafka.ConfigResourceResult, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, resources)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []kafka.ConfigResourceResult
	if rf, ok := ret.Get(0).(func(context.Context, []kafka.ConfigResource, ...kafka.DescribeConfigsAdminOption) []kafka.ConfigResourceResult); ok {
		r0 = rf(ctx, resources, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kafka.ConfigResourceResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []kafka.ConfigResource, ...kafka.DescribeConfigsAdminOption) error); ok {
		r1 = rf(ctx, resources, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetadata provides a mock function with given fields: topic, allTopics, timeoutMs
func (_m *AdminClient) GetMetadata(topic *string, allTopics bool, timeoutMs int) (*kafka.Metadata, error) {
	ret := _m.Called(topic, allTopics, timeoutMs)

	var r0 *kafka.Metadata
	if rf, ok := ret.Get(0).(func(*string, bool, int) *kafka.Metadata); ok {
		r0 = rf(topic, allTopics, timeoutMs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kafka.Metadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*string, bool, int) error); ok {
		r1 = rf(topic, allTopics, timeoutMs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetOAuthBearerToken provides a mock function with given fields: oauthBearerToken
func (_m *AdminClient) SetOAuthBearerToken(oauthBearerToken kafka.OAuthBearerToken) error {
	ret := _m.Called(oauthBearerToken)

	var r0 error
	if rf, ok := ret.Get(0).(func(kafka.OAuthBearerToken) error); ok {
		r0 = rf(oauthBearerToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetOAuthBearerTokenFailure provides a mock function with given fields: errstr
func (_m *AdminClient) SetOAuthBearerTokenFailure(errstr string) error {
	ret := _m.Called(errstr)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(errstr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *AdminClient) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
