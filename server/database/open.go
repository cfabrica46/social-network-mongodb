package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client                                                *mongo.Client
	DB                                                    *mongo.Database
	UsersCollection, PostsCollection, BlackListCollection *mongo.Collection
)

func init() {
	var err error
	Client, err = open()
	if err != nil {
		log.Fatal(err)
	}

	DB = Client.Database("social_network")

	UsersCollection = DB.Collection("users")
	PostsCollection = DB.Collection("posts")
	BlackListCollection = DB.Collection("black_list")

}

func open() (client *mongo.Client, err error) {
	host := "localhost"
	port := 27017
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	client, err = mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return
	}
	return
}
