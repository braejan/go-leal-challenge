package main

import (
	"log"

	"github.com/braejan/go-leal-challenge/cmd/web/branches"
	"github.com/braejan/go-leal-challenge/cmd/web/businesses"
	"github.com/braejan/go-leal-challenge/internal/db/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	_, err := postgres.OpenDefautDatabase()
	if err != nil {
		log.Fatalf("error starting database: %v", err)
	}
	businesses.RegisterRoutes(r)
	branches.RegisterRoutes(r)
	r.Run(":8010")
}
