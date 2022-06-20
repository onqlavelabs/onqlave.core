package utils

import (
	"context"
	"encoding/json"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"
	"onqlave.io/onqlave.core/pkg/domainevents"
)

const (
	PubSubMain = "pubsub"
)

type publisher[T domainevents.DomainEvent] struct {
}

type PublisherService[T domainevents.DomainEvent] interface {
	Publish(context context.Context, event T) error
}

func (s *publisher[T]) Publish(context context.Context, event T) error {
	fmt.Printf("Initialising to %s topic\n", event.Topic())
	client, err := dapr.NewClient()
	if err != nil {
		return err
	}
	if client == nil {
		return fmt.Errorf("dapr is not running properly. please check the configurations")
	}
	fmt.Printf("Publishing to %s topic\n", event.Topic())

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	if err := client.PublishEvent(context, PubSubMain, event.Topic(), data); err != nil {
		return err
	}
	//client.Close()
	fmt.Printf("Event published to %s topic\n", event.Topic())
	return nil
}

func NewPublisherService[T domainevents.DomainEvent]() PublisherService[T] {
	return &publisher[T]{}
}
