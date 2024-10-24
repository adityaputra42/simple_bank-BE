package main

import (
	"log"
	"simple_bank_solid/config"
	"simple_bank_solid/db"
	"simple_bank_solid/routes"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
		panic(err)
	}
	db.InitDB(conf.DbSource)
	routes.InitServer(conf)

}
