package main

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/labstack/gommon/log"
)

func main() {
	conn, err := database.InitMongo()
	if err != nil {
		log.Fatal(err)
	}

	conn.ClearDB()
	err = conn.InitCounters()
	if err != nil {
		log.Fatal(err)
	}
	conn.FakeFillDB()
}
