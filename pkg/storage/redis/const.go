package redis

const (
	prefix = "gy:redis:"

	metaPrefix = prefix + "meta:"
	mqPrefix   = prefix + "mq:"

	topicDomain       = "topic"
	topicDomainPlural = topicDomain + "s"

	partitionDomain       = "pt"
	partitionDomainPlural = partitionDomain + "s"

	// topicsMetaKey stores all the topics and its partitions
	// kind key:hset
	topicsMetaKey = metaPrefix + topicDomainPlural

	// topicMetaPartitionNumsKeyFormat stores the partition numbers of a topic
	// kind key:value
	topicMetaPartitionNumsKeyFormat = metaPrefix + topicDomain + ":%s"

	// topicPartitionsKeyFormat stores a list for all the partitions of the topic
	// kind key:list
	topicPartitionsKeyFormat = mqPrefix + topicDomain + ":%s" + ":" + partitionDomainPlural

	readyQueueKeyFormat = mqPrefix + topicDomain + ":%s:" + partitionDomain + ":%d:" + "ready"
	unACKQueueKeyFormat = mqPrefix + topicDomain + ":%s:" + partitionDomain + ":%d:" + "unack"
)
