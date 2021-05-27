package main

import (
	"log"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/cfabrica46/social-network-mongodb/server/handler"
	"github.com/cfabrica46/social-network-mongodb/server/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	log.SetFlags(log.Lshortfile)

	go database.CleanBlackList()

	r := gin.Default()

	rGetUserFromBody := r.Group("/")
	rGetUserFromBody.Use(middleware.GetUserFromBody())
	{
		rGetUserFromBody.POST("/login", handler.Login)
		rGetUserFromBody.POST("/register", handler.Register)
	}

	rGetUserFromToken := r.Group("/")
	rGetUserFromToken.Use(middleware.GetUserFromToken())
	{
		rGetUserFromToken.GET("/user", handler.Profile)
		rGetUserFromToken.GET("/logout", handler.LogOut)
		rGetUserFromToken.DELETE("/user", handler.DeleteUser)
	}

	r.Run(":8080")

}
