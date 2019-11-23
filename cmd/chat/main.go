package main

import (
	"flag"
	"github.com/go-park-mail-ru/2019_2_Pirogi/chat"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func CreateLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"stdout",
	}
	cfg.ErrorOutputPaths = []string{
		"stderr",
	}
	return cfg.Build()
}

func main() {
	configsPath := flag.String("config", "configs/", "directory with configs")
	flag.Parse()
	err := configuration.UnmarshalConfigs(*configsPath)
	if err != nil {
		log.Fatal(err.Error())
	}



	logger, err := CreateLogger()
	if err != nil {
		log.Fatal(err.Error())
	}
	zap.ReplaceGlobals(logger)


	server := chat.NewServer("/")
	go server.Listen()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
