package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.LoadHTMLFiles("index.html")

	router.StaticFile("favicon.ico", "favicon2.ico")
	// 网页图标的静态文件

	router.StaticFile("WQYZenHeiMono.ttf", "WQYZenHeiMono.ttf")
	// 网页里的style对应的字体的静态文件

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"a": "b"})
	})

	_ = router.Run(":8080")
}
