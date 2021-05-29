package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty"`
	Username string               `bson:"username"`
	Password string               `bson:"password"`
	Role     string               `bson:"role"`
	Deadline string               `bson:"deadline"`
	Token    string               `bson:"token"`
	Friends  []primitive.ObjectID `bson:"friends"`
}

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  primitive.ObjectID `bson:"user_id"`
	Content string             `bson:"content"`
	Date    string             `bson:"date"`
}

type Token struct {
	Content string `header:"Authorization-header"`
}
