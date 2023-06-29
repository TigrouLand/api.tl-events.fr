package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tigrouland/api/middlewares"
	"github.com/tigrouland/api/mongo"
	"github.com/tigrouland/api/routes"
	"github.com/tigrouland/api/routes/auth"
	"log"
)

func main() {
	// connect the mongodb database
	mongo.Connect()

	// setup web server
	r := gin.Default()

	r.Use(middlewares.Sessions())
	r.Use(cors.Default())

	userFetch := middlewares.UserFetch()
	userRequired := middlewares.UserRequired()

	r.GET("/", routes.Main)
	r.GET("/members", routes.Members)
	r.GET("/modifiers", routes.Modifiers)
	r.GET("/games", routes.Games)
	r.GET("/stats", routes.Stats)

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/login", auth.Login)
		authGroup.GET("/callback", auth.Callback)
		authGroup.GET("/self", userFetch, userRequired, auth.Self)
	}

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
