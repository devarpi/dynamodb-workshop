package dao

import "errors"

var (
	ErrItemNotFound    = errors.New("item not found")
	ErrInvalidInput    = errors.New("invalid input")
	ErrDuplicateItem   = errors.New("item already exists")
	ErrConnectionError = errors.New("failed to connect to DynamoDB")
)
