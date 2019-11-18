package domains

import "strings"

type Error struct {
	Status int    `json:"status" valid:"numeric"`
	Error  string `json:"error" valid:"optional"`
}

func NewError(status int, details ...string) *Error {
	return &Error{
		Status: status,
		Error:  strings.Join(details, "; "),
	}
}
