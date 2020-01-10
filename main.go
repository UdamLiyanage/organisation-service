package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Database struct {
	Collection *mongo.Collection
}

var DB = Database{}

func init() {
	DB.Collection = connect()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/organisations/:id", readOrganisation)
	r.GET("/organisations/:id/users", getAttachedUsers)
	r.GET("/organisations/:id/devices", getAttachedDevices)

	r.POST("/organisations", createOrganisation)
	r.POST("/organisations/attach/user", attachUser)
	r.POST("/organisations/attach/device", attachDevice)

	r.PUT("/organisations/:id", updateOrganisation)

	r.DELETE("/organisations/:id", deleteOrganisation)
	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run())
}
