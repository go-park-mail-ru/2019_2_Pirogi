package main

import (
	"os"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/server"
	"github.com/labstack/gommon/log"
)

func main() {
	mode := os.Getenv("mode")
	conn, err := database.InitMongo()
	if err != nil {
		log.Fatal(err)
	}
	// Do it one time
	conn.FakeFillDB()
	apiServer, err := server.CreateAPIServer(conn)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	if mode == "production" {
		log.Fatal(apiServer.Server.ListenAndServeTLS(configs.CertFile, configs.KeyFile))
	} else {
		log.Fatal(apiServer.Server.ListenAndServe())
	}
}
