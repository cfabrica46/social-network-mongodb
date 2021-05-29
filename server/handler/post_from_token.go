package handler

import (
	"net/http"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
)

func GetPostsFromUser(c *gin.Context) {

	var err error

	user := c.MustGet("user-data").(*database.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	posts, err := database.GetPostsFromUser(user.ID)
	if err != nil {
		return
	}

	userWithPosts := struct {
		User  database.User
		Posts []database.Post
	}{
		*user,
		posts,
	}

	c.JSON(http.StatusOK, userWithPosts)
}