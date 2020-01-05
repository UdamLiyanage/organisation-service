package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func attachDevice(c *gin.Context) {
	var attach AttachDevice
	err := json.NewDecoder(c.Request.Body).Decode(&attach)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": attach.OrganisationID}
	update := bson.M{
		"$addToSet": bson.M{
			"devices": bson.M{
				"device_id":   attach.DeviceID,
				"device_name": attach.DeviceName,
			},
		},
	}
	res, err := db.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	c.JSON(201, res)
}

func getAttachedDevices(c *gin.Context) {
	var res []map[string]interface{}
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	checkError(err, c)
	cursor, err := db.Collection.Find(
		context.Background(),
		bson.M{"_id": objID},
		options.Find().SetProjection(bson.D{
			{"devices", 1},
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
