package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	c.String(200, "Get Attached Devices")
}
