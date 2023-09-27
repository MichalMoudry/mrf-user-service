package service

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	dapr "github.com/dapr/go-sdk/client"
)

// A structure representing a service for dealing with users.
type UserService struct {
	Dapr         dapr.Client
	FirebaseAuth *auth.Client
}

// A constructor for the UserService structure.
func NewUserService(dapr dapr.Client, fbAuth *auth.Client) UserService {
	return UserService{
		Dapr:         dapr,
		FirebaseAuth: fbAuth,
	}
}

// Method for deleting a user in the system.
func (srvc UserService) DeleteUser(ctx context.Context, userId string) error {
	/*err := srvc.FirebaseAuth.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}*/

	err := srvc.Dapr.PublishEvent(ctx, "mrf-pub-sub", "user-delete", userId)
	if err != nil {
		return err
	}
	log.Printf("Published user delete event for: %s\n", userId)

	return nil
}
