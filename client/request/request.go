package request

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
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

type ErrMessage struct {
	ErrMessage string
}

func Login(username, password, url string) (tokenString string, err error) {
	var user User
	var token Token
	var errMessage ErrMessage

	user.Username = username
	user.Password = password

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")

	dataJSON, err := json.Marshal(user)

	if err != nil {
		return
	}

	req.SetBody(dataJSON)

	resp := fasthttp.AcquireResponse()

	client := &fasthttp.Client{}
	client.Do(req, resp)

	err = json.Unmarshal(resp.Body(), &token)

	if err != nil {
		return
	}

	if token.Content == "" {

		json.Unmarshal(resp.Body(), &errMessage)

		if err != nil {
			return
		}

		fmt.Println(errMessage.ErrMessage)

		return
	}

	tokenString = token.Content
	return
}
