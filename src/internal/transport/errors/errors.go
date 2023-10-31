package errors

import "errors"

var (
	ErrExpiredJwt = errors.New("token has already expired")
	ErrUidError   = errors.New("parsing of user id has failed")
)
