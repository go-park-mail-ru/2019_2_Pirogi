package error

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

func New(status int, details ...string) *models.Error {
	return &models.Error{
		Status: status,
		Error:  strings.Join(details, "; "),
	}
}

func Render(w http.ResponseWriter, err *models.Error) {
	w.WriteHeader(err.Status)
	if f, e := os.OpenFile(configs.ErrorLog, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644); e != nil {
		log.Print("Can not open or create file to log: ", e.Error())
	} else {
		_, _ = fmt.Fprintf(f, "%s %d %s \n", time.Now().Format("02/01 15:04:05"), err.Status, err.Error)
		_ = f.Close()
	}
	jsonError, _ := err.MarshalJSON()
	_, _ = fmt.Fprint(w, string(jsonError))
}
