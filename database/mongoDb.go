package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	models "GO/models"
)

func HandleDatabaseInsert(DBname string, CollectionName string, email string, phone int, password string, fname string, lname string, uid string, created time.Time, updated time.Time) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database(DBname).Collection(CollectionName)

	_, errInsert := collection.InsertOne(ctx, bson.M{
		"email":      email,
		"phone":      phone,
		"password":   password,
		"first_name": fname,
		"last_name":  lname,
		"user_id":    uid,
		"created_at": created,
		"updated_at": updated,
	})

	if errInsert != nil {
		return false
	}

	return true
}

func HandleAuthentication(email string, password string, DBname string, CollectionName string) (bool, string, string, string, string) {

	var user models.AuthenticationModel

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database(DBname).Collection(CollectionName)

	errFind := collection.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user)

	if errFind != nil {
		return false, "", "", "", ""

	}

	return true, user.Email, user.First_name, user.Last_name, user.User_id
}
