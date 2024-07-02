package server

import (
	"CoinBot/internal/database"
	"CoinBot/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) tap(c *gin.Context) {
	rdp := database.NewCache()
	client := repository.NewCacheRepository(rdp)

	client.PUSH("command")

	c.JSON(200, gin.H{
		"response": "ok",
	})
}
