package handler

import (
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

	userWithFriends := struct {
		User    database.User
		Friends []database.User
	}{
		*user,
		friends,
	}

	c.JSON(http.StatusOK, userWithFriends)

}

func GetPostsOfFriends(c *gin.Context) {

	userWithFriendID := c.MustGet("user-data-friend-id").(*struct {
		User     database.User
		FriendID primitive.ObjectID
	})
	if userWithFriendID == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	friend := database.User{ID: userWithFriendID.FriendID}

	database.GetUserFromID(&friend)

	posts, err := database.GetPostsFromUser(friend.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	friendStruct := struct {
		User  database.User
		Posts []database.Post
	}{
		friend,
		posts,
	}

	userWithFriendAndPosts := struct {
		User         database.User
		FriendStruct struct {
			User  database.User
			Posts []database.Post
		}
	}{
		userWithFriendID.User,
		friendStruct,
	}

	c.JSON(http.StatusOK, userWithFriendAndPosts)

}

func GetPostsFromFriends(c *gin.Context) {

	user := c.MustGet("user-data").(*database.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	friendsPosts := []struct {
		Post, Date, Author string
	}{}

	friends, err := database.GetFriendsFromUser(user.Friends)
	if err != nil {
		return
	}

	for i := range friends {
		posts, err := database.GetPostsFromUser(friends[i].ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}

		for indx := range posts {
			friendPost := struct {
				Post, Date, Author string
			}{posts[indx].Content, posts[indx].Date, friends[i].Username}

			friendsPosts = append(friendsPosts, friendPost)
		}

	}

	err = ordenarStructFriendsPosts(&friendsPosts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	friendsStruct := struct {
		User        database.User
		FriendPosts []struct {
			Post, Date, Author string
		}
	}{
		*user,
		friendsPosts,
	}

	c.JSON(http.StatusOK, friendsStruct)

}
