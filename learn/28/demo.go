package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// 绑定 JSON ({"user": "manu", "password": "123"})
	r.POST("/loginJson", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
