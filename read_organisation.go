package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readOrganisation(c *gin.Context) {
	var organisation Organisation
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(err, c)
	filter := bson.M{"_id": objID}
	err = db.Collection.FindOne(context.TODO(), filter).Decode(&organisation)
	checkError(err, c)
	c.JSON(200, organisation)
}
