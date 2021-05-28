package handler

import (
	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	user := c.MustGet("user-data").(database.User)

	c.JSON(200, user)
}

func LogOut(c *gin.Context) {

	user := c.MustGet("user-data").(database.User)

	err := database.InsertTokenIntoBlackList(user.Token)

	if err != nil {
		c.JSON(500, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Sesión Cerrada",
	})
}

func DeleteUser(c *gin.Context) {

	user := c.MustGet("user-data").(database.User)

	err := database.DeleteUser(user.ID)

	if err != nil {
		c.JSON(500, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"Message": "Tu cuenta ah sido eliminada",
	})

}
