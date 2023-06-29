package main

import (
	"aigr20/prom/database"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	dbUser     = flag.String("user", "", "Username for the database")
	dbPassword = flag.String("password", "", "Password for the database")
)

func init() {
	flag.Parse()
}

func main() {
	router := gin.Default()
	db, err := database.CreateConnection("prom", *dbUser, *dbPassword)
	if err != nil {
		log.Fatalf("Database connection failed: %s", err.Error())
	}
	defer db.Close()

	router.Run()
}
