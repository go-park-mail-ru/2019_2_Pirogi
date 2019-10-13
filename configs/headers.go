package configs

var Headers = map[string]string{
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Methods": "GET, POST, OPTIONS, PUT, DELETE",
	"Cache-Control":                "no-store, no-cache, must-revalidate, post-check=0, pre-check=0",
	"Content-Type":                 "application/json; charset=UTF-8",
	"Vary":                         "Accept-Encoding",
}
