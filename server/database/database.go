package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func CheckIfTokenIsInBlackList(token string) (check bool) {

	var tokenAux Token

	err := BlackListCollection.FindOne(context.TODO(), bson.M{"content": token}).Decode(&tokenAux)
	if err != nil {
		return
	}

	check = true

	return
}

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

func InsertIntoBlackList(token string) (err error) {

	tokenAux := Token{Content: token}

	_, err = BlackListCollection.InsertOne(context.TODO(), tokenAux)
	if err != nil {
		return
	}

	return
}

func DeleteUser(id string) (err error) {

	_, err = UsersCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return
	}

	return
}

func CleanBlackList() {

	for {

		err := BlackListCollection.Drop(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

	}

}
