package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	UserId       primitive.ObjectID `json:"user_id,omitempty"`
	Street       string             `json:"street,omitempty"`
	Number       string             `json:"number,omitempty"`
	Complement   string             `json:"complement,omitempty"`
	Neighborhood string             `json:"neighborhood,omitempty"`
	City         string             `json:"city,omitempty"`
	State        string             `json:"state,omitempty"`
	ZipCode      string             `json:"zipcode,omitempty"`
}
