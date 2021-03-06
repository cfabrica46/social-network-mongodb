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

	/* err = migrate()
	if err != nil {
		log.Fatal(err)
	} */
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

	var users []interface{}

	ids := []primitive.ObjectID{getIDForMigration("10b1c4924ab293de961da0e7"), getIDForMigration("10b1c4924ab293de961da0e8"), getIDForMigration("10b1c4924ab293de961da0e9"), getIDForMigration("60b1c4924ab293de961da0ea")}

	if len(ids) == 4 {
		cfabrica46 := User{
			ID:       ids[0],
			Username: "cfabrica46",
			Password: "01234",
			Role:     "admin",
			Deadline: "",
			Token:    "",
			Friends:  []primitive.ObjectID{ids[1], ids[2]},
		}
		arthuronavah := User{
			ID:       ids[1],
			Username: "arthuronavah",
			Password: "12345",
			Role:     "admin",
			Deadline: "",
			Token:    "",
			Friends:  []primitive.ObjectID{ids[0]},
		}
		luis := User{
			ID:       ids[2],
			Username: "luis",
			Password: "lolsito123",
			Role:     "member",
			Deadline: "",
			Token:    "",
			Friends:  []primitive.ObjectID{ids[0]},
		}
		carlos := User{
			ID:       ids[3],
			Username: "carlos",
			Password: "789",
			Role:     "member",
			Deadline: "",
			Token:    "",
			Friends:  []primitive.ObjectID{},
		}

		users = []interface{}{
			cfabrica46,
			arthuronavah,
			luis,
			carlos,
		}
	} else {
		log.Fatal(err)
	}

	_, err = UsersCollection.InsertMany((context.TODO()), users)
	if err != nil {
		return
	}

	posts := []interface{}{}

	for indx := range ids {
		for i := 0; i < 4; i++ {
			postAux := Post{
				UserID:  ids[indx],
				Content: "Message N-" + strconv.Itoa(i+1),
				Date:    time.Now().Format(time.Stamp),
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
