package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleDatabaseInsert(DBname string, CollectionName string, user interface{}) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database(DBname).Collection(CollectionName)

	_, errInsert := collection.InsertOne(ctx, user)

	if errInsert != nil {
		return false
	}

	return true
}

func HandleAuthentication(email string, password string, DBname string, CollectionName string, user interface{}) bool {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database(DBname).Collection(CollectionName)

	errFind := collection.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user)

	if errFind != nil {
		return false

	}

	return true
}
