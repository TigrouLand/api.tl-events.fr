package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

func Sessions() gin.HandlerFunc {
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	return sessions.Sessions("tl_session", store)
}
