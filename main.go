package main

import (
	"DB/DB"
	"DB/api"
	"DB/util"
	"database/sql"
	"log"
)


func main() {
	config, err :=  util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config" , err)
	}

	conn, err := sql.Open(config.DBDriver,config.DBSource)

	if err != nil {
		log.Fatal("cannot load db" , err)
	}

	store := DB.NewStore(conn)
	server,err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server", err)
	}


	err = server.Start(config.SourceAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}



}