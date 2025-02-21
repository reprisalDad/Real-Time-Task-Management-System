package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    FirstName    string             `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
    LastName     string             `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
    Email        string             `bson:"email" json:"email" validate:"required,email"`
    Password     string             `bson:"password" json:"password" validate:"required,min=6"`
    Phone        string             `bson:"phone" json:"phone" validate:"required"`
    UserType     string             `bson:"user_type" json:"user_type" validate:"required,oneof=ADMIN USER"`
    Token        string             `bson:"token,omitempty" json:"token,omitempty"`
    RefreshToken string             `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
    CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}
