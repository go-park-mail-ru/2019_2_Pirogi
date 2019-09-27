package Error

import (
	"fmt"
	"net/http"
)

type JSONError struct {
	Error string `json:"error"`
}

func (e *JSONError) String() string {
	return e.Error
}

func New(details ...string) string {
	var s string
	for _, v := range details {
		s += v + " "
	}
	return "{\"error\":\"" + s + "\"}\n"
}

func Render(w http.ResponseWriter, statusCode int, details ...string) {
	w.WriteHeader(statusCode)
	_, _ = fmt.Fprint(w, New(details...))
}

func Wrap(text string, err error) string {
	return New(text + ": " + err.Error())
}

func InvalidQueryArgument(key string) string {
	return New("invalid method: " + key)
}

func NotFound() string {
	return New("not found")
}
