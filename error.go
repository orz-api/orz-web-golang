package orzweb

import (
	"fmt"
)

type ApiError struct {
	Code string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("code(`%s`)", e.Code)
}

func NewApiError(code string) error {
	return &ApiError{
		Code: code,
	}
}
