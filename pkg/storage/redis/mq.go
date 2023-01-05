package redis

import (
	"context"
	"fmt"

	redisclient "github.com/go-redis/redis/v8"
)

type MQRedisRepository struct {
	client redisclient.UniversalClient
}

func NewMQRedisRepository(client redisclient.UniversalClient) *MQRedisRepository {
	return &MQRedisRepository{
		client: client,
	}
}

func (mqr *MQRedisRepository) Push(ctx context.Context, topic, payload string, partition int) error {
	return mqr.client.LPush(ctx, mqr.readyQueueKey(topic, partition), payload).Err()
}

func (mqr *MQRedisRepository) Pop(ctx context.Context, topic string, partition int) (string, error) {
	res, err := mqr.client.RPopLPush(ctx, mqr.readyQueueKey(topic, partition), mqr.unackQueueKey(topic, partition)).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (mqr *MQRedisRepository) readyQueueKey(topic string, partition int) string {
	return fmt.Sprintf(readyQueueKeyFormat, topic, partition)
}

func (mqr *MQRedisRepository) unackQueueKey(topic string, partition int) string {
	return fmt.Sprintf(unACKQueueKeyFormat, topic, partition)
}
