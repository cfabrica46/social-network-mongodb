package handler

import (
	"net/http"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	user := c.MustGet("user-data").(database.User)
	if user.ID.Hex() == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func LogOut(c *gin.Context) {

	user := c.MustGet("user-data").(database.User)
	if user.ID.Hex() == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	err := database.InsertTokenIntoBlackList(user.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Sesión Cerrada",
	})
}

func DeleteUser(c *gin.Context) {

	user := c.MustGet("user-data").(database.User)
	if user.ID.Hex() == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	err := database.DeleteUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Tu cuenta ah sido eliminada",
	})

}
