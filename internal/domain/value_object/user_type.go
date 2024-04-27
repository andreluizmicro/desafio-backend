package value_object

import "errors"

const (
	defaultUser = iota + 1
	shopkeeperUser
)

var ErrInvalidUserType = errors.New("invalid user type")

type UserType struct {
	Value int
}

func NewUserType(value int) *UserType {
	userType := &UserType{}
	switch value {
	case defaultUser:
		userType.Value = value
	case shopkeeperUser:
		userType.Value = value
	default:
		userType.Value = defaultUser
	}

	return userType
}
