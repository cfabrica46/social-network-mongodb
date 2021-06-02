package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/cfabrica46/social-network-mongodb/server/token"
	"github.com/gin-gonic/gin"
)

func GetUserFromBody(c *gin.Context) {

	var user database.User

	d, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		if err != io.EOF {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}
	}

	fmt.Printf("%s\n", d)

	err = json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		if err != io.EOF {

			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}
	}

	fmt.Println("Usuario", user)

	c.Set("user-data", &user)
	c.Next()

}

func GetUserFromToken(c *gin.Context) {

	var tokenValue database.Token

	if err := c.ShouldBindHeader(&tokenValue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	check := database.CheckIfTokenIsInBlackList(tokenValue.Content)
	if check {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "El Token no es válido",
		})
		return
	}

	user, err := token.ExtractUserFromClaims(tokenValue.Content)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	user.Token = tokenValue.Content

	deadline, err := time.Parse(time.ANSIC, user.Deadline)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	checkTime := time.Now().Local().After(deadline)

	if !checkTime {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "El Token no es válido",
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
