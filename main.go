package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()

    // テンプレートファイルのロード
    router.LoadHTMLGlob("templates/*")

    // ルートハンドラ
    router.GET("/", func(c *gin.Context) {
		
		imageURL, err := GenerateImage("クマの絵をお願いします", "1024x1024")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
        c.HTML(http.StatusOK, "index.html", gin.H{
            "Title":   "Hello, Gin!",
            "Message": "Welcome to the Gin framework!",
			"ImageURL": imageURL,
        })
    })
	
    router.Run(":8080")
}
