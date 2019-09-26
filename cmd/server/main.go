package main

import (
	"../../internal/app/server"
	"flag"
	"log"
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
