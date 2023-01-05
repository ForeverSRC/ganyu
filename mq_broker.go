package ganyu

import (
	"context"
	"hash/fnv"

	"github.com/ForeverSRC/ganyu/pkg/domain"
	"github.com/ForeverSRC/ganyu/pkg/logger"
	redisstorage "github.com/ForeverSRC/ganyu/pkg/storage/redis"
)

type Broker interface {
	CreateTopic(ctx context.Context, topic domain.Topic) error
	Push(ctx context.Context, topic, payload, partitionKey string) error
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

func (b *defaultBroker) Push(ctx context.Context, topic, payload, partitionKey string) error {
	partition := b.partitionRoute(ctx, topic, partitionKey)
	return b.repo.Push(ctx, topic, payload, partition)
}

func (b *defaultBroker) partitionRoute(ctx context.Context, topic, partitionKey string) int {
	h := fnv.New64()
	_, err := h.Write([]byte(partitionKey))
	if err != nil {
		return 0
	}

	count, err := b.repo.GetTopicPartitions(ctx, topic)
	if err != nil {
		b.logger.Error("get partition count of topic %s error %v", topic, err)
		return 0
	}

	return int(h.Sum64() % uint64(count))

}
