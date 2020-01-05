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
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
	c.JSON(201, res)
}
