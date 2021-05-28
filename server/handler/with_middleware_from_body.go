package handler

import (
	"fmt"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/cfabrica46/social-network-mongodb/server/token"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {

	user := c.MustGet("user-data").(database.User)

	err := database.GetUser(&user)

	if err != nil {
		c.JSON(403, gin.H{
			"ErrMessage": "Usuario no encontrado",
		})
		return
	}

	user.Token, err = token.GenerateToken(user.ID.Hex(), user.Username, user.Role)
	if err != nil {
		c.JSON(500, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	token := database.Token{Content: user.Token}

	c.JSON(200, token)

}

func SignUp(c *gin.Context) {

	user := c.MustGet("user-data").(database.User)

	err := database.AddUser(user)
	if err != nil {
		c.JSON(403, gin.H{
			"ErrMessage": "El nombre del usuario ya esta en uso",
		})
		return
	}

	err = database.GetUser(&user)
	if err != nil {
		c.JSON(403, gin.H{
			"ErrMessage": "El nombre del usuario ya esta en uso",
		})
		return
	}

	user.Token, err = token.GenerateToken(user.ID.Hex(), user.Username, user.Role)
	if err != nil {
		c.JSON(500, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	token := database.Token{Content: user.Token}

	c.JSON(200, token)

}

func ShowUsers(c *gin.Context) {
	users, err := database.GetUsers()
	if err != nil {
		c.JSON(403, gin.H{
			"ErrMessage": "El nombre del usuario ya esta en uso",
		})
		return
	}
	fmt.Println(users)
	c.JSON(200, users)

}
