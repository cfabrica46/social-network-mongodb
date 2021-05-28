package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

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

	//var usersLoaded []bson.M
	//lookupStage := bson.D{{"$lookup", bson.D{{"from", "posts"}, {"localField", "_id"}, {"foreignField", "userID"}, {"as", "user_posts"}}}}
	//
	//usersLoadedCursor, err := UsersCollection.Aggregate(context.TODO(), mongo.Pipeline{lookupStage})
	//if err != nil {
	//	return
	//}
	//
	//if usersLoadedCursor.All(context.TODO(), &usersLoaded); err != nil {
	//	return
	//}
	//
	//for _, v := range usersLoaded {
	//	var userAux User
	//	var userPost []byte
	//
	//	userPost, err = bson.Marshal(v)
	//	if err != nil {
	//		return
	//	}
	//
	//	err = bson.Unmarshal(userPost, &userAux)
	//	if err != nil {
	//		return
	//	}
	//
	//	users = append(users, userAux)
	//}

	return

}
