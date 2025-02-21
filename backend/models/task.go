package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title       string             `bson:"title" json:"title" validate:"required"`
    Description string             `bson:"description" json:"description"`
    AssignedTo  string             `bson:"assigned_to" json:"assigned_to"` // user ID
    Status      string             `bson:"status" json:"status" validate:"required,oneof=pending in-progress completed"`
    CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
