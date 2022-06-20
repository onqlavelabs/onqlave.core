package utils

import (
	"fmt"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	"github.com/labstack/echo/v4"
)

type Subscription struct {
	PubsubName string            `json:"pubsubname"`
	Topic      string            `json:"topic"`
	Metadata   map[string]string `json:"metadata,omitempty"`
	Routes     Routes            `json:"routes"`
}

type Routes struct {
	Rules   []Rule `json:"rules,omitempty"`
	Default string `json:"default,omitempty"`
}

type Rule struct {
	Match string `json:"match"`
	Path  string `json:"path"`
}

type Subscriber struct {
	subscriptions []Subscription
}

func (s *Subscriber) Register(subscriptions []Subscription) {
	if len(subscriptions) > 0 {
		s.subscriptions = append(s.subscriptions, subscriptions...)
	}
}

func (s *Subscriber) DaprHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.subscriptions)
}

func NewSubscriber() *Subscriber {
	return &Subscriber{}
}

type Order struct {
	Id string `json:"orderId"`
}

func (s *Subscriber) TestSubscription(c echo.Context) error {
	request := new(common.TopicEvent)
	if err := c.Bind(request); err != nil {
		fmt.Print("error in binding..")
		return err
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	handler := NewEventPayloadHandler[Order]()
	order, err := handler.Extract(request.Data)
	if err != nil {
		return err
	}
	fmt.Printf("**request** type %s => received %s \n", request.DataContentType, order.Id)
	return nil
}
