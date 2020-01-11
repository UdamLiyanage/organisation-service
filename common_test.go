package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	DB.Collection = connect()
	// Run the other tests
	os.Exit(m.Run())
}

func newRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/organisations/:id", readOrganisation)
	r.POST("/organisations", createOrganisation)
	r.DELETE("/organisations/:id", deleteOrganisation)
	r.PUT("/organisations/:id", updateOrganisation)
	r.POST("/organisations/attach/device", attachDevice)
	r.POST("/organisations/attach/user", attachUser)
	r.POST("/organisations/remove/device", removeAttachedDevice)
	return r
}
