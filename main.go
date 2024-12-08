package main

import (
	"ginSample/database"
	"ginSample/routes"
	"net/http"

	_ "ginSample/docs" // Swaggerのドキュメントをインポート

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           User Registration API
// @version         1.0
// @description     API for user registration
// @host            localhost:8080
// @BasePath        /
func main() {
	// Initialize database
	database.ConnectDatabase()

	// Create Gin router
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// テンプレートレンダリングの設定
	r.LoadHTMLFiles("templates/index.html", "templates/main.html")

	// 静的ファイルの提供
	r.Static("/static", "./static")
	// Routes
	routes.RegisterAuthRoutes(r)

	// トップページをレンダリング（index.html）
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) // index.htmlをレンダリング
	})

	// ログイン後にmain.htmlをレンダリングするルート
	r.GET("/main", func(c *gin.Context) {
		// トークンが存在するか確認（認証チェック）
		token := c.GetHeader("Authorization")
		if token == "" {
			// トークンが無い場合は、ログイン画面にリダイレクト
			c.Redirect(http.StatusFound, "/")
			return
		}

		// 認証済みなら、main.htmlをレンダリング
		c.HTML(http.StatusOK, "main.html", nil)
	})

	// Start server
	r.Run(":8080")
}
