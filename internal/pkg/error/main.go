package Error

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"log"
	"net/http"
	"strings"
)

func New(status int, details ...string) *models.Error {
	newError := models.Error{
		Status: status,
		Error:  strings.Join(details, "; "),
	}
	return &newError
}

func Render(w http.ResponseWriter, error *models.Error) {
	w.WriteHeader(error.Status)
	log.Print(error.Status, " | ", error.Error)
	jsonError, _ := error.MarshalJSON()
	_, _ = fmt.Fprint(w, string(jsonError))
}
