package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ✅ static 파일 제공 (/static 폴더 안의 파일)
	r.Static("/static", "./static")

	// ✅ 루트("/") 요청 시 index.html 반환
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	r.Run(":8080")
}
