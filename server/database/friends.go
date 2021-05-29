package database

import (
	"context"

	"github.com/cfabrica46/leveldb-login/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFriendsFromUser(ids []primitive.ObjectID) (friends []database.User, err error) {

	for i := range ids {
		var friend database.User

		err = UsersCollection.FindOne(context.TODO(), bson.M{"_id": ids[i]}).Decode(&friend)
		if err != nil {
			return
		}

		friends = append(friends, friend)
	}

	return
}
