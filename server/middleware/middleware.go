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

		c.Set("user-data", user)

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
		}

		c.Set("user-data", user)

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
		}

		err = json.NewDecoder(c.Request.Body).Decode(&newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}

		oldUserAndNewUser := struct {
			OldUser database.User
			NewUser database.User
		}{
			user,
			newUser,
		}

		c.Set("old-and-new-user-data", oldUserAndNewUser)
		c.Next()
	}
}
