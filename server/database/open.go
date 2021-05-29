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
	ids := []primitive.ObjectID{getIDForMigration("60b1c4924ab293de961da0e7"), getIDForMigration("60b1c4924ab293de961da0e8"), getIDForMigration("60b1c4924ab293de961da0e9"), getIDForMigration("60b1c4924ab293de961da0ea")}

	users := []interface{}{
		User{
			ID:       ids[0],
			Username: "cfabrica46",
			Password: "01234",
			Role:     "admin",
			Deadline: "",
			Token:    "",
			Posts:    []Post{},
			Friends:  []primitive.ObjectID{ids[1], ids[2]},
		},
		User{
			ID:       ids[1],
			Username: "arthuronavah",
			Password: "12345",
			Role:     "admin",
			Deadline: "",
			Token:    "",
			Posts:    []Post{},
			Friends:  []primitive.ObjectID{ids[0]},
		},
		User{
			ID:       ids[2],
			Username: "luis",
			Password: "lolsito123",
			Role:     "member",
			Deadline: "",
			Token:    "",
			Posts:    []Post{},
			Friends:  []primitive.ObjectID{ids[0]},
		},
		User{
			ID:       ids[3],
			Username: "carlos",
			Password: "789",
			Role:     "member",
			Deadline: "",
			Token:    "",
			Posts:    []Post{},
			Friends:  []primitive.ObjectID{},
		},
	}

	_, err = UsersCollection.InsertMany((context.TODO()), users)
	if err != nil {
		return
	}

	posts := []interface{}{}

	for indx := range ids {
		fmt.Println(len(ids))
		for i := 0; i < 4; i++ {
			postAux := Post{
				UserID:  ids[indx],
				Content: "Message N-" + strconv.Itoa(i+1),
				Date:    time.Now().String(),
			}
			posts = append(posts, postAux)
		}

	}

	_, err = PostsCollection.InsertMany(context.TODO(), posts)
	if err != nil {
		return
	}
	return
}

func getIDForMigration(idString string) (id primitive.ObjectID) {
	id, _ = primitive.ObjectIDFromHex(idString)
	return
}
