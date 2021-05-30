package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/cfabrica46/social-network-mongodb/server/handler"
	"github.com/cfabrica46/social-network-mongodb/server/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	//Solo para evitar acumulacion de datos en la Database.
	//No se utilizara para el producto final.
	{
		sigs := make(chan os.Signal)
		signal.Notify(sigs, syscall.SIGINT)
		go func() {
			<-sigs
			database.UsersCollection.Drop(context.TODO())
			database.PostsCollection.Drop(context.TODO())
			os.Exit(0)
		}()
	}

	log.SetFlags(log.Lshortfile)

	go func() {
		for {
			database.CleanBlackList()
			time.Sleep(time.Hour)
		}
	}()

	r := gin.Default()

	r.GET("/users", handler.ShowUsers)

	rGetUserFromBody := r.Group("/")
	rGetUserFromBody.Use(middleware.GetUserFromBody)
	{
		rGetUserFromBody.POST("/signin", handler.SignIn)
		rGetUserFromBody.POST("/signup", handler.SignUp)
	}

	rGetUserFromToken := r.Group("/")
	rGetUserFromToken.Use(middleware.GetUserFromToken)
	{
		rGetUserFromToken.GET("/logout", handler.LogOut)
		rGetUserFromToken.GET("/user", handler.Profile)
		rGetUserFromToken.DELETE("/user", handler.DeleteUser)
		rGetUserFromToken.PUT("/user", handler.UpdateUser)
		rGetUserFromToken.GET("/user/posts", handler.GetPostsFromUser)
		rGetUserFromToken.GET("/user/friends", handler.GetFriendsFromUser)
		rGetUserFromToken.GET("/user/friend/posts", handler.GetPostsOfFriend)
		rGetUserFromToken.GET("/user/friends/posts", handler.GetPostsFromFriends)
	}

	r.Run(":8080")

}
