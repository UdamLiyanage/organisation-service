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
	checkError(err, c)
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(err, c)
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"name": organisation.Name,
		},
	}
	_, err = DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(err, c)
	c.Status(200)
}
