package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Database struct {
	Collection *mongo.Collection
}

var db = Database{Collection: connect()}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/organisations/:id")

	r.POST("/organisation")

	r.PUT("/organisations/:id")

	r.DELETE("/organisations/:id")
	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run())
}
