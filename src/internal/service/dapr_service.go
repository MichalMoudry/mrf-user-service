package service

import (
	"context"
	"user-service/internal/service/errors"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	PUBSUB = "mrf-pub-sub"
)

// A wrapper service for a Dapr client.
type DaprService struct {
	DaprClient dapr.Client
}

// A constructor function for DaprService structure.
func NewDapr(client dapr.Client) *DaprService {
	return &DaprService{
		DaprClient: client,
	}
}

// Method for publishing an event in the system.
func (srvc DaprService) PublishEvent(ctx context.Context, topic string, data interface{}) error {
	if srvc.DaprClient == nil {
		return errors.ErrDaprIsNotInitialized
	}
	return srvc.DaprClient.PublishEvent(ctx, PUBSUB, topic, data)
}
