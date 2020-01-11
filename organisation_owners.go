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
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(err, c)
	c.JSON(201, res)
}

func removeAttachedUser(c *gin.Context) {
	var remove map[string]string
	err := json.NewDecoder(c.Request.Body).Decode(&remove)
	checkError(err, c)
	objectID, err := primitive.ObjectIDFromHex(remove["organisation_id"])
	checkError(err, c)
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$pull": bson.M{
			"users": bson.M{
				"user_id": remove["user_id"],
			},
		},
	}
	res, err := DB.Collection.UpdateOne(context.TODO(), filter, update)
	checkError(err, c)
	c.JSON(404, res)
}

func getAttachedUsers(c *gin.Context) {
	var res []map[string]interface{}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(err, c)
	cursor, err := DB.Collection.Find(
		context.Background(),
		bson.M{"_id": objID},
		options.Find().SetProjection(bson.D{
			bson.E{Key: "users", Value: 1},
		}),
	)
	checkError(err, c)
	for cursor.Next(context.TODO()) {
		var user map[string]interface{}
		err := cursor.Decode(&user)
		checkError(err, c)
		res = append(res, user)
	}
	c.JSON(200, res)
}
