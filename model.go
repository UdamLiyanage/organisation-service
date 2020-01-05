package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Organisation struct {
	ID      primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string              `json:"name" bson:"name"`
	Devices []map[string]string `json:"devices" bson:"devices"`
	Users   []map[string]string `json:"users" bson:"users"`
}
