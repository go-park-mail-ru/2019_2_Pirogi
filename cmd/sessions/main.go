package main

import (
	"flag"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/interfaces"
	v1 "github.com/go-park-mail-ru/2019_2_Pirogi/cmd/sessions/protobuf"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	configsPath := flag.String("config", "configs/", "directory with configs")
	flag.Parse()

	err := configuration.UnmarshalConfigs(*configsPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	logger, err := network.CreateLogger()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	conn, err := database.InitMongo(configs.Default.MongoHost)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := interfaces.NewUserRepository(conn)
	cookieRepo := interfaces.NewCookieRepository(conn)

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("can not listen on port: ", err)
	}
	server := grpc.NewServer()
	v1.RegisterAuthServiceServer(server, NewAuthManager(userRepo, cookieRepo))
	log.Fatal(server.Serve(lis))
}
