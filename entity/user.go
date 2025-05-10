package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                  primitive.ObjectID `json:"id" bson:"_id"`
	FirstName           string             `json:"first_name" bson:"first_name"`
	LastName            string             `json:"last_name" bson:"last_name"`
	FullName            string             `json:"full_name" bson:"full_name"`
	Email               string             `json:"email" bson:"email"`
	Phone               string             `json:"phone" bson:"phone"`
	Roles               []Role             `json:"roles" bson:"roles"`
	Username            string             `json:"username" bson:"username"`
	Password            string             `json:"password" bson:"password"`
	Avatar              string             `json:"avatar" bson:"avatar"`
	CreatedAt           primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt           primitive.DateTime `json:"updated_at" bson:"updated_at"`
	DeletedAt           primitive.DateTime `json:"deleted_at" bson:"deleted_at"`
	IsDeleted           bool               `json:"is_deleted" bson:"is_deleted"`
	IsActive            bool               `json:"is_active" bson:"is_active"`
	RememberToken       string             `json:"remember_token,omitempty" bson:"remember_token,omitempty"`
	ForgotPasswordToken string             `json:"forgot_password_token,omitempty" bson:"forgot_password_token,omitempty"`
}

type AccessToken struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	AccessToken string             `json:"access_token " bson:"access_token"`
	ExpiresIn   int                `json:"expires_in" bson:"expires_in"`
	TokenType   string             `json:"token_type" bson:"token_type"`
	Scope       string             `json:"scope" bson:"scope"`
}
