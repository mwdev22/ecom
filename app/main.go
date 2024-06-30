package main

import (
	"fmt"
	"log"

	"github.com/mwdev22/ecom/app/api"
	env "github.com/mwdev22/ecom/app/config"
	dbConn "github.com/mwdev22/ecom/app/db"
)

func main() {

	dbCfg := env.GetDbCfg()

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=5432 sslmode=disable",
		dbCfg.User, dbCfg.Name, dbCfg.Pass, dbCfg.Host)

	db, err := dbConn.DbOpen(connStr, "postgres")
	if err != nil {
		fmt.Printf("db open failed: %v", err)
	}
	dbConn.InitConn(db)

	server := api.NewServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
