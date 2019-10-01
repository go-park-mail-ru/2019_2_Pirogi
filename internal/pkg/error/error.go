package Error

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
	newError := models.Error{
		Status: status,
		Error:  strings.Join(details, "; "),
	}
	return &newError
}

func Render(w http.ResponseWriter, error *models.Error) {
	w.WriteHeader(error.Status)
	if f, err := os.OpenFile(configs.AccessLogPath+"error.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644); err != nil {
		log.Print("Can not open or create file to log: ", err.Error())
	} else {
		_, _ = fmt.Fprintf(f, "%s %d %s \n", time.Now().Format("02/01 15:04:05"), error.Status, error.Error)
		_ = f.Close()
	}
	jsonError, _ := error.MarshalJSON()
	_, _ = fmt.Fprint(w, string(jsonError))
}
