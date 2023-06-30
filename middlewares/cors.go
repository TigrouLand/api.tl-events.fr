package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func Cors() gin.HandlerFunc {
	origins := []string{os.Getenv("FRONTEND_URL")}
	if os.Getenv("DEV_MODE") == "true" {
		origins = append(origins, "http://127.0.0.1:3000")
	}

	return cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
	})
}
