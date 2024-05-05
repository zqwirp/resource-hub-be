package handler

import "github.com/gin-gonic/gin"

type PingHandler interface {
	Pong(c *gin.Context)
	Ping(c *gin.Context)
}

type Ping struct {
	Message string
}

func NewPingHandler() PingHandler {
	return &Ping{}
}

func (h *Ping) Ping(c *gin.Context) {
	arr := []int{1, 2, 3}

	c.JSON(200, gin.H{
		"message": "pong",
		"value":   arr[4],
	})
}

func (h *Ping) Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}
