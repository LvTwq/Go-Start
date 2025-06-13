package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Multipart/Urlencoded 绑定
*/
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("./templates/*")
	r.GET("/demo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "admin",
			"pwd":  "123456",
		})
	})
	r.GET("/student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "查询学生信息成功",
		})
	})

	// curl -v --form user=user --form password=123456 http://localhost:8080/login
	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "123456" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": "user or password is wrong"})
			}
		}
	})

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	r.Run(":8080")
}
