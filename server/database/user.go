package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetUserFromID(user *User) (err error) {
	var userAux User

	err = UsersCollection.FindOne(context.TODO(), bson.M{"_id": user.ID}).Decode(&userAux)
	if err != nil {
		return
	}

	*user = userAux
	return
}

func GetUserFromUsername(username string) (user User, err error) {

	err = UsersCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return
	}

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

func UpdateUser(user *User, newUsername, newPassword string) (err error) {

	_, err = UsersCollection.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.D{
		{"$set", bson.D{{"username", newUsername}, {"password", newPassword}}},
	})
	if err != nil {
		return
	}

	user.Username = newUsername
	user.Password = newPassword

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

	return

}

func CheckIfUserAlreadyExist(username string) (check bool, err error) {

	var userAux User

	err = UsersCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&userAux)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = nil
		}
		return
	}

	check = true
	return
}
