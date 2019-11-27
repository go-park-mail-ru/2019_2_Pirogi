package main

import (
	"flag"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/sessions"
	v1 "github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/sessions/protobuf"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/interfaces"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
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
	defer logger.Sync()

	conn, err := database.InitMongo(configs.Default.MongoHost)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := interfaces.NewUserRepository(conn)
	cookieRepo := interfaces.NewCookieRepository(conn)

	lis, err := net.Listen("tcp", configs.Default.SessionsMicroservicePort)
	if err != nil {
		log.Fatalln("can not listen on port: ", err)
	}
	server := grpc.NewServer()
	v1.RegisterAuthServiceServer(server, sessions.NewAuthManager(userRepo, cookieRepo))
	log.Fatal(server.Serve(lis))
}
