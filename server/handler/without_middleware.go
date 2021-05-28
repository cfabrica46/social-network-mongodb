package handler

import (
	"fmt"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
)

func ShowUsers(c *gin.Context) {
	users, err := database.GetUsers()
	if err != nil {
		c.JSON(403, gin.H{
			"ErrMessage": "El nombre del usuario ya esta en uso",
		})
		return
	}

	for i := range users {
		fmt.Println()
		fmt.Printf("ID: %s | Username: %s | Role: %s | Token: %s\n ", users[i].ID.Hex(), users[i].Username, users[i].Role, users[i].Token)
		fmt.Println()
		fmt.Println("Posts:")
		for indx := range users[i].Posts {
			fmt.Printf("\t Content: %s\n", users[i].Posts[indx].Content)
		}
	}

	fmt.Println()
	c.JSON(200, users)

}
