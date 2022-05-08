package models

import "errors"

var (
	ErrorLoginExist             = errors.New("this login already exist")
	ErrorInvalidLoginOrPassword = errors.New("invalid login or password")
)
