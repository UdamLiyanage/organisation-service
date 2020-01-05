package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Organisation struct {
	ID      primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string              `json:"name" bson:"name"`
	Devices []map[string]string `json:"devices,omitempty" bson:"devices,omitempty"`
	Users   []map[string]string `json:"users,omitempty" bson:"users,omitempty"`
}

type AttachOwner struct {
	UserID         string             `json:"user_id" bson:"user_id"`
	UserName       string             `json:"user_name" bson:"user_name"`
	OrganisationID primitive.ObjectID `json:"organisation_id" bson:"organisation_id"`
}
