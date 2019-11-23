package main

import (
	"encoding/json"
	"flag"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/chat"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	json2 "github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
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

	conn, err := database.InitMongo(configs.Default.MongoHost)
	if err != nil {
		log.Fatal(err)
	}

	server := chat.NewServer("/", conn)
	go server.Listen()
	go log.Fatal(http.ListenAndServe(":8080", nil))
	adminServer := http.NewServeMux()
	adminServer.HandleFunc("/admin", getHandlerAdmin(conn))
	log.Fatal(http.ListenAndServe(":9000", adminServer))

}

func getHandlerAdmin(conn database.Database) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		chatsInterfaces, err := conn.GetAll(configs.Default.ChatTargetName)
		if err != nil {
			http.Error(res, "no chats", 404)
			return
		}
		var chats [][]byte
		for _, chatInterface := range chatsInterfaces {
			chat := chatInterface.(model.Chat)
			body, err := json.Marshal(chat)
			if err != nil {
				continue
			}
			chats = append(chats, body)

		}
		res.Write(json2.MakeJSONArray(chats))
	}
}
