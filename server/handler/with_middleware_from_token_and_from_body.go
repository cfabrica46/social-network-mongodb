package handler

import (
	"net/http"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/cfabrica46/social-network-mongodb/server/token"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	userWithNewData := c.MustGet("old-and-new-user-data").(*struct {
		User                     database.User
		NewUsername, NewPassword string
	})
	if userWithNewData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	check, err := database.CheckIfUserAlreadyExist(userWithNewData.NewUsername)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}
	if check {
		c.JSON(http.StatusForbidden, gin.H{
			"ErrMessage": "El nombre del usuario ya esta en uso",
		})
		return
	}

	user := userWithNewData.User

	err = database.UpdateUser(&user, userWithNewData.NewUsername, userWithNewData.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	err = database.InsertTokenIntoBlackList(user.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	user.Token, err = token.GenerateToken(user.ID.Hex(), user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	c.JSON(http.StatusOK, user)

}
