package handler

import (
	"fmt"
	"net/http"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
)

func ShowUsers(c *gin.Context) {
	users, err := database.GetUsers()
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"ErrMessage": "El nombre del usuario ya esta en uso",
		})
		return
	}

	usersWithPosts := []struct {
		User  database.User
		Posts []database.Post
	}{}

	for i := range users {
		fmt.Println()
		fmt.Printf("Debbug: ID: %24s | Username: %15s | Role: %5s | Token: %32s\n", users[i].ID.Hex(), users[i].Username, users[i].Role, users[i].Token)
		fmt.Println()
		fmt.Println("Posts:")
		posts, err := database.GetPostsFromUser(users[i].ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ErrMessage": "Internal Error",
			})
			return
		}
		for indx := range posts {
			fmt.Printf("\t Content: %s\n", posts[indx].Content)
		}
		fmt.Println("Friends:")
		for indx := range users[i].Friends {
			fmt.Printf("\t Content: %s\n", users[i].Friends[indx].Hex())
		}

		userWithPosts := struct {
			User  database.User
			Posts []database.Post
		}{
			users[i],
			posts,
		}

		usersWithPosts = append(usersWithPosts, userWithPosts)
	}

	fmt.Println()
	c.JSON(http.StatusOK, usersWithPosts)

}
