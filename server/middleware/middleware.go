package middleware

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Friend struct {
	ID string `json:"id"`
}

func GetUserFromBody() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user database.User

		err := json.NewDecoder(c.Request.Body).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}

		c.Set("user-data", &user)

		c.Next()
	}
}

func GetUserFromToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := getUser(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": err.Error(),
			})
			return
		}

		err = database.GetUserFromID(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": err.Error(),
			})
			return
		}

		c.Set("user-data", &user)
		c.Next()

	}
}

func GetUserFromTokenAndNewUserDataFromBody() gin.HandlerFunc {
	return func(c *gin.Context) {

		var newUser database.User

		user, err := getUser(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": err.Error(),
			})
			return
		}

		err = database.GetUserFromID(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": err.Error(),
			})
			return
		}

		err = json.NewDecoder(c.Request.Body).Decode(&newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}

		userWithNewData := struct {
			User                     database.User
			NewUsername, NewPassword string
		}{
			user,
			newUser.Username,
			newUser.Password,
		}

		c.Set("old-and-new-user-data", &userWithNewData)
		c.Next()
	}
}

func GetUserFromTokenAndIDFriend() gin.HandlerFunc {
	return func(c *gin.Context) {

		var check bool

		user, err := getUser(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": err.Error(),
			})
			return
		}

		err = database.GetUserFromID(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": err.Error(),
			})
			return
		}

		friendID := Friend{}

		err = json.NewDecoder(c.Request.Body).Decode(&friendID)
		if err != nil {
			if err != io.EOF {
				c.JSON(http.StatusInternalServerError, gin.H{
					"ErrMessage": "Internal Error",
				})
				return
			}
		}

		id, err := primitive.ObjectIDFromHex(friendID.ID)

		for i := range user.Friends {
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"ErrMessage": "Internal Error",
				})
				return
			}
			if id == user.Friends[i] {
				check = true
				break
			}
		}

		if !check {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrMessage": "The selected id is not from a friend",
			})
			return
		}

		userWithFriendID := struct {
			User     database.User
			FriendID primitive.ObjectID
		}{user, id}

		c.Set("user-data-friend-id", &userWithFriendID)
		c.Next()
	}
}
