package models

import (
	//"encoding/json"
	//"github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/conf"
	//"github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/database"
	//"github.com/dgrijalva/jwt-go"
	//"github.com/google/uuid"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"log"
	//"time"
)

func NewResponse(status string, message string, data map[string]interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

type Response struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

//-----Struct for user interaction-----//
type UserCredentials struct {
	Email       string `json:"email,omitempty" bson:"email,omitempty"`
	CountryCode string `json:"country_code,omitempty" bson:"country_code,omitempty"`
	Mobile      int    `json:"mobile,omitempty" bson:"mobile,omitempty"`
}
//-----Struct for DB interaction-----//
type User struct {
	Id                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName           string             `json:"first_name" bson:"first_name"`
	LastName            string             `json:"last_name" bson:"last_name"`
	Picture             string             `json:"picture" bson:"picture"`
	Email               string             `json:"email" bson:"email,omitempty"`
	CountryCode         string             `json:"country_code" bson:"country_code,omitempty"`
	Mobile              int                `json:"mobile" bson:"mobile,omitempty"`
	Address             string             `json:"address" bson:"address,omitempty"`
	UserClassification  string             `json:"user_classification" bson:"user_classification,omitempty"`
	SignUpThrough       string             `json:"sign_up_through" bson:"sign_up_through"`
	Newsletter          bool               `json:"newsletter" bson:"newsletter"`
	IsAdmin             bool               `json:"is_admin" bson:"is_admin"`
	IsActive            bool               `json:"is_active" bson:"is_active"`
	AccountCreationTime int64              `json:"account_creation_time" bson:"account_creation_time"`
	LastLoggedIn        int64              `json:"last_logged_in" bson:"last_logged_in"`
	LastUpdatedOn       int64              `json:"last_updated_on" bson:"last_updated_on"`
}

func NewUser(email string, countryCode string, mobile int) *User {
	return &User{
		Email:       email,
		CountryCode: countryCode,
		Mobile:      mobile,
		IsActive:    true,
	}
}

func NewUserByEmail(email string) *User {
	return NewUser(email, "", 0)
}

func NewUserByMobile(countryCode string, mobile int) *User {
	return NewUser("", countryCode, mobile)
}


