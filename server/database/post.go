package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetPostsFromIDsFriends(ids []primitive.ObjectID) (friendsPosts []struct {
	Author string
	Post   Post
}, err error) {

	for i := range ids {
		var postAux Post

		var cur *mongo.Cursor
		cur, err = PostsCollection.Find(context.TODO(), bson.M{"user_id": ids[i]})
		if err != nil {
			return
		}

		for cur.Next(context.TODO()) {
			var userAux User

			err = cur.Decode(&postAux)
			if err != nil {
				return
			}

			err = UsersCollection.FindOne(context.TODO(), bson.M{"_id": ids[i]}).Decode(&userAux)
			if err != nil {
				return
			}

			friendPost := struct {
				Author string
				Post   Post
			}{userAux.Username, postAux}

			friendsPosts = append(friendsPosts, friendPost)

		}
	}

	err = ordenarStructFriendsPosts(friendsPosts)
	if err != nil {
		return
	}

	return
}

func ordenarStructFriendsPosts(friendsPosts []struct {
	Author string
	Post   Post
}) (err error) {

	for indx := 0; indx < len(friendsPosts)-1; indx++ {

		for i := 0; i < len(friendsPosts)-1; i++ {

			var t1, t2 time.Time

			t1, err = time.Parse(time.Stamp, (friendsPosts)[i].Post.Date)
			if err != nil {
				return
			}

			t2, err = time.Parse(time.Stamp, (friendsPosts)[i+1].Post.Date)
			if err != nil {
				return
			}

			if t1.Before(t2) {

				aux := (friendsPosts)[i]
				(friendsPosts)[i] = (friendsPosts)[i+1]
				(friendsPosts)[i+1] = aux

			}

		}

	}
	return
}
