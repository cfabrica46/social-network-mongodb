package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(user *User) (err error) {

	var userAux User

	err = UsersCollection.FindOne(context.TODO(), bson.M{"username": user.Username, "password": user.Password}).Decode(&userAux)
	if err != nil {
		return
	}

	*user = userAux

	return
}

func AddUser(user User) (err error) {

	user.Role = "member"

	_, err = UsersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return
	}

	return
}

func DeleteUser(id primitive.ObjectID) (err error) {

	_, err = UsersCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return
	}

	return
}

func GetUsers() (users []User, err error) {

	userCur, err := UsersCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return
	}
	defer userCur.Close(context.TODO())
	for userCur.Next(context.TODO()) {
		var userAux User
		err = userCur.Decode(&userAux)
		if err != nil {
			return
		}
		users = append(users, userAux)
	}

	posts := []Post{}

	postCur, err := PostsCollection.Find(context.TODO(), bson.D{{}})
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

	for i := range posts {
		for indx := range users {
			if posts[i].UserID == users[indx].ID {
				users[indx].Posts = append(users[indx].Posts, posts[i])
				break
			}
		}
	}
	return

}
