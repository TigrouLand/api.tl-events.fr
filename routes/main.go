package routes

import "github.com/gin-gonic/gin"

func Main(c *gin.Context) {
	c.JSON(200, struct {
		Welcome   string `json:"welcome"`
		Members   string `json:"members"`
		Modifiers string `json:"modifiers"`
		Games     string `json:"games"`
		Stats     string `json:"stats"`
	}{
		Welcome:   "Welcome to the Events API, check out the available public endpoints:",
		Members:   "/members",
		Modifiers: "/modifiers",
		Games:     "/games",
		Stats:     "/stats",
	})
}
