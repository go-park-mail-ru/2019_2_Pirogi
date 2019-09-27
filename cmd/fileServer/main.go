package main

import (
	"flag"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "3000", "port to start file server")
	router := mux.NewRouter()
	log.Println("Starting new file server...")
	router.HandleFunc("/api/images/", handlers.UploadImageHandler).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
