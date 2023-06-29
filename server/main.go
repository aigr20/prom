package main

import (
	"aigr20/prom/api"
	"aigr20/prom/database"
	"flag"
	"log"
)

var (
	dbUser     = flag.String("user", "", "Username for the database")
	dbPassword = flag.String("password", "", "Password for the database")
)

func init() {
	flag.Parse()
}

func main() {
	db, err := database.CreateConnection("prom", *dbUser, *dbPassword)
	if err != nil {
		log.Fatalf("Database connection failed: %s", err.Error())
	}
	defer db.Close()
	server := api.NewAPI(db)
	server.Serve()
}
