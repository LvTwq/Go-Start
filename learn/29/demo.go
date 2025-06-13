package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := "./" + file.Filename
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.Run(":8080")
}
