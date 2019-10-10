package main

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/server"
	"github.com/labstack/gommon/log"
)

func main() {
	conn, err := database.InitInmemory()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	conn.FakeFillDB()
	apiServer, err := server.CreateAPIServer(conn)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Fatal(apiServer.Start(configs.APIPort))
}
