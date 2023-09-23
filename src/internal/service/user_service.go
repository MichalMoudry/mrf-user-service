package service

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
)

// A structure representing a service for dealing with users.
type UserService struct {
	Dapr dapr.Client
}

// Method for deleting a user in the system.
func (srvc UserService) DeleteUser(ctx context.Context, userId string) error {
	err := srvc.Dapr.PublishEvent(ctx, "mrf-pub-sub", "user-delete", userId)
	if err != nil {
		return err
	}

	return nil
}
