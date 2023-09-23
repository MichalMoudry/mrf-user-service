package model

import (
	"user-service/internal/service"
	"user-service/internal/transport/model/ioc"

	"firebase.google.com/go/v4/auth"
	dapr "github.com/dapr/go-sdk/client"
)

// A structure representing a collection of public services in the service layer.
type ServiceCollection struct {
	UserService ioc.IUserService
}

// A constructor function for ServiceCollection structure.
func NewServiceCollection(dapr dapr.Client, fbAuth *auth.Client) *ServiceCollection {
	return &ServiceCollection{
		UserService: service.NewUserService(dapr, fbAuth),
	}
}
