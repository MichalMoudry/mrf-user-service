package util

import "context"

type uId string

// Function for adding a user id to a context.
func WithUserId(ctx context.Context, userId string) context.Context {
	var key uId = "user_id"
	return context.WithValue(ctx, key, userId)
}
