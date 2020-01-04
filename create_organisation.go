package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createOrganisation(c *gin.Context) {
	var organisation Organisation
	err := json.NewDecoder(c.Request.Body).Decode(&organisation)
	if err != nil {
		panic(err)
	}
	insertResult, err := db.Collection.InsertOne(context.TODO(), organisation)
	if err != nil {
		panic(err)
	}
	organisation.ID = insertResult.InsertedID.(primitive.ObjectID)
	c.JSON(201, organisation)
}
