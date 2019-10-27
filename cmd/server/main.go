package main

import (
	"flag"
	"os"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/server"
	"github.com/labstack/gommon/log"
)

func main() {
	configsPath := flag.String("config", "../../configs/", "directory with configs")
	flag.Parse()

	err := common.UnmarshalConfigs(configsPath)
	if err != nil {
		log.Fatal(err.Error())
	}


	conn, err := database.InitMongo()
	if err != nil {
		log.Fatal(err)
	}

	apiServer, err := server.CreateAPIServer(conn)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	mode := os.Getenv("mode")
	if mode == "production" {
		log.Fatal(apiServer.Server.ListenAndServeTLS(configs.Default.CertFile, configs.Default.KeyFile))
	} else {
		log.Fatal(apiServer.Server.ListenAndServe())
	}
}
