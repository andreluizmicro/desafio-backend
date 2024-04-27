package exception

import "errors"

var (
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrAccountNotFound      = errors.New("account not found")
	ErrUserNotFound         = errors.New("user not found")
)
