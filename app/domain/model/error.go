package model

import (
	"errors"
	"github.com/labstack/echo"
)

type Error struct {
	Status int    `json:"status" valid:"numeric"`
	Error  string `json:"error" valid:"optional"`
}

func NewError(status int, details ...string) *Error {
	return &Error{
		Status: status,
		Error:  details[0],
	}
}

func (e *Error) String() string {
	return string(e.Status) + " " + e.Error
}

func (e *Error) HTTP() *echo.HTTPError {
	return echo.NewHTTPError(e.Status, e.Error)
}

func (e *Error) Common() error {
	if e == nil {
		return nil
	}
	return errors.New(e.String())
}
