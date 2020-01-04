package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func deleteOrganisation(c *gin.Context) {
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.D{{"_id", objID}}
	_, err = db.Collection.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	c.JSON(404, nil)
}
