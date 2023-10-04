package errors

import "errors"

var (
	ErrFirebaseCredentialsMissing = errors.New("firebase credentials have not been found")
)
