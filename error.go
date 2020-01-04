package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkError(err error, c *gin.Context) {
	if err == nil {
		return
	}
	if err == mongo.ErrNoDocuments {
		notFoundResponse(c)
	}
}

func notFoundResponse(c *gin.Context) {
	_ = c.AbortWithError(404, nil)
}
