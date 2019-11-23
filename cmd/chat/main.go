package main

import (
	"flag"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/chat"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
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

	conn, err := database.InitMongo(configs.Default.MongoHost)
	if err != nil {
		log.Fatal(err)
	}
	//adminServer := http.NewServeMux()
	//adminServer.HandleFunc("/admin", getHandlerAdmin(conn))
	//go log.Fatal(http.ListenAndServe(":9000", adminServer))

	server := chat.NewServer("/", conn)
	go server.Listen()
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//
//func getHandlerAdmin(conn database.Database) func(res http.ResponseWriter, req *http.Request) {
//	return func(res http.ResponseWriter, req *http.Request) {
//		chatsInterfaces, err := conn.GetAll(configs.Default.ChatTargetName)
//		if err != nil {
//			http.Error(res, "no chats", 404)
//			return
//		}
//		var chats [][]byte
//		for _, chatInterface := range chatsInterfaces {
//			chat  = chatInterface.(model.Chat)
//			chats = append(chats,  json.Marshal(chat))
//
//		}
//		json.MakeJSONArray(chats)
//	}
//}
