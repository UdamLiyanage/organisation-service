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

	r.GET("/organisations/:id", readOrganisation)

	r.POST("/organisations", createOrganisation)
	r.POST("/organisations/attach/user", attachUser)

	r.PUT("/organisations/:id", updateOrganisation)

	r.DELETE("/organisations/:id", deleteOrganisation)
	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run())
}
