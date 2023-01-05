package ganyu

import "context"

type Producer struct {
	Broker Broker
}

func NewProducer(broker Broker) *Producer {
	return &Producer{
		Broker: broker,
	}
}

func (p *Producer) Produce(ctx context.Context, topic, message string, partitionKey string) error {
	pk := message
	if len(partitionKey) > 0 {
		pk = partitionKey
	}

	err := p.Broker.Push(ctx, topic, message, pk)
	return err
}
