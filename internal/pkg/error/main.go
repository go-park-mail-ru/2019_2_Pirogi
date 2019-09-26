package Error

type JSONError struct {
	Error string `json:"error"`
}

func (e *JSONError) String() string {
	return e.Error
}

func New(details string) string {
	return "{\"error\":\"" + details + "\"}\n"
}

func Wrap(text string, err error) string {
	return New(text + ": " + err.Error())
}

func InvalidMethod(method string) string {
	return New("invalid method: " + method)
}

func InvalidQueryArgument(key string) string {
	return New("invalid method: " + key)
}

func NotFound() string {
	return New("not found")
}

func InvalidJSON() string {
	return New("invalid json")
}
