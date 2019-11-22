package error

import (
	"strings"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

func New(status int, details ...string) *models.Error {
	return &models.Error{
		Status: status,
		Error:  strings.Join(details, "; "),
	}
}
