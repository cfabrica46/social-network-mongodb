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

func InsertTokenIntoBlackList(token string) (err error) {

	tokenAux := Token{Content: token}

	_, err = BlackListCollection.InsertOne(context.TODO(), tokenAux)
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
