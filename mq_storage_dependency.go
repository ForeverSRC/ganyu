package ganyu

import (
	"context"

	"github.com/ForeverSRC/ganyu/pkg/domain"
)

type TopicRepository interface {
	CreateTopic(ctx context.Context, topic domain.Topic) error
	GetTopicPartitions(ctx context.Context, topic string) (int, error)
}

type MQRepository interface {
	Push(ctx context.Context, topic, payload string, partition int) error
	Pop(ctx context.Context, topic string, partition int) (string, error)
}
type RepositorySet interface {
	TopicRepository
	MQRepository
}
