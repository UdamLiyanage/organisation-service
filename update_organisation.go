package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateOrganisation(c *gin.Context) {
	var organisation Organisation
	err := json.NewDecoder(c.Request.Body).Decode(&organisation)
	if err != nil {
		panic(err)
	}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"name": organisation.Name,
		},
	}
	_, err = db.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	c.Status(200)
}
