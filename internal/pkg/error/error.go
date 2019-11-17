package error

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"strings"
)

func New(status int, details ...string) *domains.Error {
	return &domains.Error{
		Status: status,
		Error:  strings.Join(details, "; "),
	}
}
