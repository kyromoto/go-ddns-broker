package producerdomain

import "fmt"

type Producer struct {
	_uuid UUID
	name  string
}

type Consumer struct {
	_uuid UUID
	name  string
}

func (c *Consumer) execute() error {

}

type ProducerRepository interface {
	GetProducerByUuid(uuid UUID) (error, Producer)
}

type ConsumerRepository interface {
	GetConsumersByProducerUuid(uuid UUID) (error, []Consumer)
}

type UpdateProducerIpUC interface {
	UpdateIp(ip string, producer Producer) error
}

type updateProducerIpUseCaseImpl struct {
	producerRepository *ProducerRepository
	consumerRepository *ConsumerRepository
}

func (uc *updateProducerIpUseCaseImpl) UpdateIp(ip string, producer Producer) error {
	err, consumers := uc.consumerRepository.GetConsumersByProducerUuid(producer.uuid)

	if err != nil {
		return err
	}

	for _, consumer := range consumers {
		go consumer.execute()
	}
}

func NewUpdateIpProducerUC(
	producerRepository *ProducerRepository,
	consumerRepository *ConsumerRepository
) UpdateProducerIpUC {
	return &updateProducerIpUseCaseImpl{
		producerRepository: producerRepository,
		consumerRepository: consumerRepository
	}
}
