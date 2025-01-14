package tododatastore

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoData struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Task        string             `bson:"task"`
	IsCompleted bool               `bson:"isCompleted"`
	UserId      primitive.ObjectID `bson:"userId"`
	IsDeleted   bool               `bson:"isDeleted"`
}
