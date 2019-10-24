package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/server"
	"github.com/labstack/gommon/log"
)

func UnmarshalConfigs() error {
	file, err := ioutil.ReadFile("../../configs/default.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &configs.Default)
	if err != nil {
		return err
	}

	file, err = ioutil.ReadFile("../../configs/headers.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &configs.Headers)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := UnmarshalConfigs()
	if err != nil {
		log.Fatalf(err.Error())
		return
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
