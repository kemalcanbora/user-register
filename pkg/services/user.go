package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"rollic/internal/models"
)

type UserService interface {
	AddUser(data interface{}, collectionName string) (*mongo.InsertOneResult, error)
	DeleteUser(data interface{}, collectionName string) (*mongo.DeleteResult, error)
	UpdateUser(filter, update interface{}, collectionName string) (*mongo.UpdateResult, error)
	FindUser(data interface{}, collectionName string) (bson.M, error)
	GetAllUser(collectionName string) []models.UserResponse
	FindUserWithEmail(user models.User) models.User
}
