package user

import "errors"

var (
	ErrNoUser     = errors.New("user doesn't exists")
	ErrUserExists = errors.New("user is already exists")
)
