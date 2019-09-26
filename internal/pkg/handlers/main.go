package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello, world!")
	if err != nil {
		log.Fatal(err)
	}
}

func HandleUser(w http.ResponseWriter, r *http.Request) {

}

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL)
}
