package redis

type StorageSet struct {
	*TopicRedisRepository
}

func NewRedisStorageSet(opts PoolOptions) *StorageSet {
	client := NewRedisPool(opts)
	return &StorageSet{
		TopicRedisRepository: NewTopicRedisRepository(client),
	}
}
