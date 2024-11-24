package service

import "errors"

var (
	ErrCustomerNotFound error = errors.New("customer not found")
)
