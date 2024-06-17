package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type request struct {
	Prompt string `json:"prompt"`
}

func main() {
	router := gin.Default()

	// 静的ファイルの提供
	router.Static("/static", "./templates")
	// テンプレートファイルのロード
	router.LoadHTMLGlob("templates/*")

	// ルートハンドラ
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ImageURL": "/static/image.png",
		})
	})

	// ルートハンドラ
	router.POST("/generate", func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Prompt == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "prompt is required"})
			return
		}
		// ベースのプロンプトを環境変数から取得
		err := godotenv.Load(".env")
		if err != nil {
			panic("envFile not found")
		}
		basePrompt := os.Getenv("BASE_PROMPT")
		if basePrompt == "" {
			panic("BASE_PROMPT is empty")
		}

		// プロンプトを結合
		prompt := req.Prompt + basePrompt

		imageURL, err := GenerateImage(prompt, "1024x1024")

		// テンプレートファイルのロード
		router.LoadHTMLGlob("templates/*")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"imageUrl": imageURL})
	})

	router.Run(":8080")
}
