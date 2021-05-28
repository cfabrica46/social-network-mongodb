package middleware

import (
	"encoding/json"
	"time"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/cfabrica46/social-network-mongodb/server/token"
	"github.com/gin-gonic/gin"
)

func GetUserFromBody() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user database.User

		err := json.NewDecoder(c.Request.Body).Decode(&user)

		if err != nil {
			c.JSON(500, gin.H{
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

		var tokenValue database.Token

		if err := c.ShouldBindHeader(&tokenValue); err != nil {
			c.JSON(500, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}

		check := database.CheckIfTokenIsInBlackList(tokenValue.Content)
		if check {
			c.JSON(500, gin.H{
				"ErrMessage": "El Token no es válido",
			})
			return
		}

		user, err := token.ExtractUserFromClaims(tokenValue.Content)

		if err != nil {
			c.JSON(500, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}

		user.Token = tokenValue.Content

		deadline, err := time.Parse(time.ANSIC, user.Deadline)

		if err != nil {
			c.JSON(500, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}

		checkTime := time.Now().Local().After(deadline)

		if !checkTime {
			c.JSON(500, gin.H{
				"ErrMessage": "El Token no es válido",
			})
			return
		}

		c.Set("user-data", user)

		c.Next()

	}
}
