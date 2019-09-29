package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.String("p", "3000", "port to start file server")
	flag.Parse()
	router := mux.NewRouter()
	log.Printf("Starting new file server on port :%s...", *port)
	router.HandleFunc("/api/images/", handlers.UploadImageHandler).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
