package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
)

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
