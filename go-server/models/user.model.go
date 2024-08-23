package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	UserId        *string            `bson:"user_id"`
	FirstName     *string            `bson:"first_name" validate:"required,min=3,max=32"`
	LastName      *string            `bson:"last_name" validate:"required,min=3,max=32"`
	Email         *string            `bson:"email" validate:"email,required"`
	Password      *string            `bson:"password" validate:"required,min=6,max=32"`
	Avatar        *string            `bson:"avater"`
	Token         *string            `bson:"token"`
	RefreshToken  *string            `bson:"refresh_token"`
	HasRestaurant *bool              `bson:"has_restaurant"`
	CreatedAt     *time.Time         `bson:"created_at"`
	UpdatedAt     *time.Time         `bson:"updated_at"`
}
