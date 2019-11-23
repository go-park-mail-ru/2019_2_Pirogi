package main

import (
	"encoding/json"
	"flag"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	json2 "github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
	"log"
	"net/http"
)

func main() {
	configsPath := flag.String("config", "configs/", "directory with configs")
	flag.Parse()

	err := configuration.UnmarshalConfigs(*configsPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	conn, err := database.InitMongo(configs.Default.MongoHost)
	if err != nil {
		log.Fatal(err)
	}
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
