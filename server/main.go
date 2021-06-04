package main

import (
	"context"
	"log"
	"net/http"
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

	r.StaticFS("/index", http.Dir("index"))

	s := r.Group("/api/v1")
	{
		s.GET("/users", handler.ShowUsers)

		sGetUserFromBody := s.Group("/")
		sGetUserFromBody.Use(middleware.GetUserFromBody)
		{
			sGetUserFromBody.POST("/signin", handler.SignIn)
			sGetUserFromBody.POST("/signup", handler.SignUp)
		}

		sGetUserFromToken := s.Group("/")
		sGetUserFromToken.Use(middleware.GetUserFromToken)
		{
			sGetUserFromToken.GET("/logout", handler.LogOut)
			sGetUserFromToken.GET("/user", handler.Profile)
			sGetUserFromToken.DELETE("/user", handler.DeleteUser)
			sGetUserFromToken.PUT("/user", handler.UpdateUser)
			sGetUserFromToken.GET("/user/posts", handler.GetPostsFromUser)
			sGetUserFromToken.GET("/user/friends", handler.GetFriendsFromUser)
			sGetUserFromToken.GET("/friend/:username/posts", handler.GetPostsOfFriend)
			sGetUserFromToken.GET("/friends/posts", handler.GetPostsFromFriends)
		}
	}

	r.Run(":8080")

}
