package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func attachUser(c *gin.Context) {
	var attach AttachOwner
	err := json.NewDecoder(c.Request.Body).Decode(&attach)
	checkError(err, c)
	filter := bson.M{"_id": attach.OrganisationID}
	update := bson.M{
		"$addToSet": bson.M{
			"users": bson.M{
				"user_id":   attach.UserID,
				"user_name": attach.UserName,
			},
		},
	}
	res, err := db.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(err, c)
	c.JSON(201, res)
}

func getAttachedUsers(c *gin.Context) {
	var res []map[string]interface{}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}
	cursor, err := db.Collection.Find(
		context.Background(),
		bson.M{"_id": objID},
		options.Find().SetProjection(bson.D{
			{"users", 1},
		}),
	)
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var user map[string]interface{}
		err := cursor.Decode(&user)
		if err != nil {
			panic(err)
		}
		res = append(res, user)
	}
	c.JSON(200, res)
}
