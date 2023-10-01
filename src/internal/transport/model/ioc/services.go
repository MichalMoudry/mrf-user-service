package ioc

import "context"

// An interface for a user service.
type IUserService interface {
	// Method for deleting a user in the system.
	DeleteUser(ctx context.Context, userId string) error
}
