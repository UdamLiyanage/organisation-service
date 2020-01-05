package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	c.JSON(200, "Get attached users function")
}
