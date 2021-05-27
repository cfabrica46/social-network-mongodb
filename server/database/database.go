package database

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	UserID  primitive.ObjectID `bson:"userID"`
	Content string             `bson:"content"`
	Date    string             `bson:"date"`
}
