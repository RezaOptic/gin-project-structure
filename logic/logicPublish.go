package logic

import (
	"context"
	"github.com/RezaOptic/gin-project-structure/service"
	"gitlab.com/RezaOptic/go-utils/logger"
)

// PublishInterface interface
type PublishInterface interface {
	PublishDataToMessageBroker(keyPrefix string, data []byte) error
}

// Publish struct
type Publish struct {
	Context           context.Context
	Self              PublishInterface
}

// NewPublishLogic function for create Publish logic
func NewPublishLogic(ctx context.Context) PublishInterface {
	x := Publish{Context: ctx}
	x.Self = &x
	return &x
}

// PublishDataToMessageBroker method for publish data to message broker
func (p *Publish) PublishDataToMessageBroker(keyPrefix string, data []byte) error {
	err := service.Sd.Publish(keyPrefix, data)
	if err != nil {
		logger.ZSLogger.Errorf("error on publish data to nats with error :%s", err)
		return err
	}
	return nil
}
