package errors

import "errors"

var (
	ErrExpiredJwt = errors.New("token has already expired")
)
