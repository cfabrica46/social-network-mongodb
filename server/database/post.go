package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPostsFromUser(id primitive.ObjectID) (posts []Post, err error) {

	postCur, err := PostsCollection.Find(context.TODO(), bson.M{"user_id": id})
	if err != nil {
		return
	}
	defer postCur.Close(context.TODO())

	for postCur.Next(context.TODO()) {
		var postAux Post
		err = postCur.Decode(&postAux)
		if err != nil {
			return
		}
		posts = append(posts, postAux)
	}

	return
}
