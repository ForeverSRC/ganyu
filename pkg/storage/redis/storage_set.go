package redis

type StorageSet struct {
	*TopicRedisRepository
	*MQRedisRepository
}

func NewRedisStorageSet(opts PoolOptions) *StorageSet {
	client := NewRedisPool(opts)
	return &StorageSet{
		TopicRedisRepository: NewTopicRedisRepository(client),
		MQRedisRepository:    NewMQRedisRepository(client),
	}
}
