package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createOrganisation(c *gin.Context) {
	var organisation Organisation
	err := json.NewDecoder(c.Request.Body).Decode(&organisation)
	checkError(err, c)
	insertResult, err := DB.Collection.InsertOne(context.TODO(), organisation)
	checkError(err, c)
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		c.JSON(201, map[string]interface{}{
			"id":      oid.String(),
			"name":    organisation.Name,
			"devices": organisation.Devices,
			"users":   organisation.Users,
		})
	} else {
		c.AbortWithStatus(500)
	}
	c.JSON(201, organisation)
}
