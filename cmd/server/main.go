package main

import (
	"flag"
	"log"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/app/server"
)

func main() {
	port := flag.String("p", "8080", "port for server")
	flag.Parse()
	s := server.New(*port)
	s.Init()
	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
