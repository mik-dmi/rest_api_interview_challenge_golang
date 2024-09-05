package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mik-dmi/cmd/api"
	"github.com/mik-dmi/configs"
	"github.com/mik-dmi/db"
)

func main() {

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s  port=%s",
		configs.Envs.DBUSER, configs.Envs.DBPASSWORD, configs.Envs.DBHOST, configs.Envs.DBNAME, configs.Envs.DBPORT)

	db, err := db.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.PORT), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected!")
}
