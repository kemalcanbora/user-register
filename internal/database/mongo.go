package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"rollic/internal/models"
	u "rollic/pkg/utils"
	"time"
)

type MongoClient struct {
	Client  *mongo.Client
	Context context.Context
	Cancel  func()
}

func MongoConnection() MongoClient {
	mongoClient := MongoClient{}
	var err error
	mongoClient.Client, err = mongo.NewClient(options.Client().ApplyURI(u.GetEnvVariable("MONGO_HOST")))
	mongoClient.Context, mongoClient.Cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer mongoClient.Cancel()

	err_ := mongoClient.Client.Connect(mongoClient.Context)
	if err_ != nil {
		log.Fatal(err)
	}
	return mongoClient
}

func (m *MongoClient) AddUser(data interface{}, collectionName string) (*mongo.InsertOneResult, error) {
	collection := m.Client.Database(u.GetEnvVariable("MONGO_DB")).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (m *MongoClient) DeleteUser(data interface{}, collectionName string) (*mongo.DeleteResult, error) {
	collection := m.Client.Database(u.GetEnvVariable("MONGO_DB")).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var r bson.M
	err := collection.FindOne(ctx, data).Decode(&r)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = fmt.Errorf("user with that id does not exist")
			return nil, err
		}
		log.Fatal(err)
	}
	result, err := collection.DeleteOne(ctx, data)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (m *MongoClient) UpdateUser(filter, update interface{}, collectionName string) (*mongo.UpdateResult, error) {
	collection := m.Client.Database(u.GetEnvVariable("MONGO_DB")).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (m *MongoClient) FindUser(data interface{}, collectionName string) (bson.M, error) {
	collection := m.Client.Database(u.GetEnvVariable("MONGO_DB")).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var r bson.M
	err := collection.FindOne(ctx, data).Decode(&r)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = fmt.Errorf("user with that id does not exist")
			return nil, err
		}
		log.Fatal(err)
	}
	fmt.Println(r)

	return r, err
}

func (m *MongoClient) GetAllUser(collectionName string) []models.UserResponse {
	var users []models.UserResponse
	collection := m.Client.Database(u.GetEnvVariable("MONGO_DB")).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.UserResponse
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}

	return users
}

func (m *MongoClient) FindUserWithEmail(user models.User) models.User {
	var dbUser models.User
	collection := m.Client.Database(u.GetEnvVariable("MONGO_DB")).Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		log.Println(err)
	}
	return dbUser
}
