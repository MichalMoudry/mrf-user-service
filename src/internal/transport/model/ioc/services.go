package ioc

import "context"

type IUserService interface {
	// Method for deleting a user in the system.
	DeleteUser(ctx context.Context, userId string) error
}
