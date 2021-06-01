package handler

import (
	"fmt"
	"net/http"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFriendsFromUser(c *gin.Context) {
	var err error

	user := c.MustGet("user-data").(*database.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	friends, err := database.GetFriendsFromUser(user.Friends)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, friends)

}

func GetPostsOfFriend(c *gin.Context) {

	var check bool

	user := c.MustGet("user-data").(*database.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	usernameFriend := c.Param("username")

	friend, err := database.GetUserFromUsername(usernameFriend)
	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	for i := range user.Friends {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}
		if friend.ID == user.Friends[i] {
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

	err = database.GetUserFromID(&friend)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	posts, err := database.GetPostsFromUser(friend.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	friendAndPosts := struct {
		Friend database.User
		Posts  []database.Post
	}{
		friend,
		posts,
	}

	c.JSON(http.StatusOK, friendAndPosts)

}

func GetPostsFromFriends(c *gin.Context) {
	var ids []primitive.ObjectID

	user := c.MustGet("user-data").(*database.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	friends, err := database.GetFriendsFromUser(user.Friends)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	for i := range friends {
		ids = append(ids, friends[i].ID)
	}

	friendsPosts, err := database.GetPostsFromIDsFriends(ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	c.JSON(http.StatusOK, friendsPosts)

}
