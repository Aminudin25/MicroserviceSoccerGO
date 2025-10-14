package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUsernameExists       = errors.New("user already exists")
	ErrEmailExist			= errors.New("email already exists")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
	ErrInvalidEmail         = errors.New("invalid email format")
	ErrUserDisabled         = errors.New("user account is disabled")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrUsernameExists,
	ErrEmailExist,
	ErrInvalidPassword,
	ErrPasswordDoesNotMatch,
	ErrInvalidEmail,
	ErrUserDisabled,
}
