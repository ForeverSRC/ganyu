package broker

import (
	"context"

	"github.com/ForeverSRC/ganyu/pkg/domain"
)

type TopicRepository interface {
	CreateTopic(ctx context.Context, topic domain.Topic) error
}

type RepositorySet interface {
	TopicRepository
}
