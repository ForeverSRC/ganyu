package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ForeverSRC/ganyu/pkg/domain"

	redisclient "github.com/go-redis/redis/v8"
)

type TopicRedisRepository struct {
	client redisclient.UniversalClient
}

func NewTopicRedisRepository(client redisclient.UniversalClient) *TopicRedisRepository {
	return &TopicRedisRepository{
		client: client,
	}
}

func (tr *TopicRedisRepository) CreateTopic(ctx context.Context, topic domain.Topic) error {
	ok, err := tr.client.SetNX(ctx, tr.topicMetaPartitionNumsKey(topic.Name), fmt.Sprintf("%d", topic.Partitions), 0).Result()
	if err != nil {
		return err
	}

	if !ok {
		return nil
	}

	ctxt, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	pipeline := tr.client.TxPipeline()
	pipeline.HSet(ctxt, topicsMetaKey, topic.Name, fmt.Sprintf("%d", topic.Partitions))

	ptKey := tr.partitionsKey(topic.Name)
	for i := 0; i < topic.Partitions; i++ {
		pipeline.LPush(ctxt, ptKey, fmt.Sprintf("%d", i))
	}

	_, err = pipeline.Exec(ctxt)

	return err

}

func (tr *TopicRedisRepository) GetTopicPartitions(ctx context.Context, topic string) (int, error) {
	r, err := tr.client.Get(ctx, tr.topicMetaPartitionNumsKey(topic)).Result()
	if err != nil {
		return -1, err
	}

	num, err := strconv.Atoi(r)
	if err != nil {
		return -1, err
	}

	return num, nil
}

func (tr *TopicRedisRepository) topicMetaPartitionNumsKey(topic string) string {
	return fmt.Sprintf(topicMetaPartitionNumsKeyFormat, topic)
}

func (tr *TopicRedisRepository) partitionsKey(topic string) string {
	return fmt.Sprintf(topicPartitionsKeyFormat, topic)
}
