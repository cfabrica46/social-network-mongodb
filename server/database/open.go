package database

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
	Deadline string             `bson:"deadline"`
	Token    string             `bson:"token"`
	Posts    []Post             `bson:"user_posts",omitempty`
}

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  primitive.ObjectID `bson:"user_id"`
	Content string             `bson:"content"`
	Date    string             `bson:"date"`
}

type Token struct {
	Content string `bson:"content"`
}

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

	err = migrate()
	if err != nil {
		log.Fatal(err)
	}
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

func migrate() (err error) {
	users := []interface{}{
		User{
			ID:       [12]byte{},
			Username: "cfabrica46",
			Password: "01234",
			Role:     "admin",
			Deadline: "",
			Token:    "",
			Posts:    []Post{},
		},
		User{
			ID:       [12]byte{},
			Username: "arthuronavah",
			Password: "12345",
			Role:     "admin",
			Deadline: "",
			Token:    "",
			Posts:    []Post{},
		},
		User{
			ID:       [12]byte{},
			Username: "luis",
			Password: "lolsito123",
			Role:     "member",
			Deadline: "",
			Token:    "",
			Posts:    []Post{},
		},
	}

	insertManyUsers, err := UsersCollection.InsertMany((context.TODO()), users)
	if err != nil {
		return
	}

	ids := []primitive.ObjectID{}

	for _, v := range insertManyUsers.InsertedIDs {
		if id, ok := v.(primitive.ObjectID); ok {
			ids = append(ids, id)
		}
	}

	posts := []interface{}{}

	for i := range ids {

		postAux := Post{
			UserID:  ids[i],
			Content: "Message " + strconv.Itoa(i+1),
			Date:    time.Now().String(),
		}
		posts = append(posts, postAux)
	}

	_, err = PostsCollection.InsertMany(context.TODO(), posts)
	if err != nil {
		return
	}
	return
}
