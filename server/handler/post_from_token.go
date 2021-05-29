package handler

import (
	"net/http"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
)

func ShowPostsFromUser(c *gin.Context) {

	var err error

	user := c.MustGet("user-data").(*database.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	user.Posts, err = database.GetPostsFromUser(user.ID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, *user)
}
