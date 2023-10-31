package service

import (
	"context"
	"user-service/internal/service/model/ioc"

	"firebase.google.com/go/v4/auth"
	dapr "github.com/dapr/go-sdk/client"
)

// A structure representing a service for dealing with users.
type UserService struct {
	Dapr         ioc.IDaprService
	FirebaseAuth *auth.Client
}

// A constructor for the UserService structure.
func NewUserService(dapr dapr.Client, fbAuth *auth.Client) UserService {
	return UserService{
		Dapr:         NewDapr(dapr),
		FirebaseAuth: fbAuth,
	}
}

// Method for deleting a user in the system.
func (srvc UserService) DeleteUser(ctx context.Context, userId string) error {
	err := srvc.FirebaseAuth.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}

	if err = srvc.Dapr.PublishEvent(ctx, "user-delete", userId); err != nil {
		return err
	}
	return nil
}
