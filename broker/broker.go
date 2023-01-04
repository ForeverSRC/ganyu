package broker

import (
	"context"

	"github.com/ForeverSRC/ganyu/pkg/logger"

	"github.com/ForeverSRC/ganyu/pkg/domain"

	redisstorage "github.com/ForeverSRC/ganyu/pkg/storage/redis"
)

type Broker interface {
	CreateTopic(ctx context.Context, topic domain.Topic) error
}

type defaultBroker struct {
	*BrokerOpts
}

type BrokerOpts struct {
	repo   RepositorySet
	logger logger.Logger
}

type Option func(opt *BrokerOpts)

func NewDefaultBroker(cfg redisstorage.PoolOptions, opts ...Option) Broker {
	brokerOpts := &BrokerOpts{
		repo: redisstorage.NewRedisStorageSet(cfg),
	}

	for _, op := range opts {
		op(brokerOpts)
	}

	if brokerOpts.logger == nil {
		brokerOpts.logger = logger.DefaultLogger()
	}

	return newDefaultBroker(brokerOpts)
}

func newDefaultBroker(opts *BrokerOpts) *defaultBroker {
	b := &defaultBroker{
		BrokerOpts: opts,
	}

	return b
}

func (b *defaultBroker) CreateTopic(ctx context.Context, topic domain.Topic) error {
	return b.repo.CreateTopic(ctx, topic)
}
