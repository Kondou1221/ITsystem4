package main

import (
	"back-go/controllers"
	"back-go/dbconect"

	"github.com/gin-gonic/gin"
)

func main() {

	dbconect.Dbinit()

	r := gin.Default()

	// テスト用1
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hogehoge",
		})
	})

	// テスト用2
	r.GET("/test", controllers.Test)

	// ログイン
	r.POST("/login", controllers.Login)
	// 新規会員登録
	r.POST("/newlogin", controllers.New_login)
	// 依頼一覧取得
	r.GET("/request_get", controllers.Request_get)
	// 依頼新規作成
	r.POST("/new_request", controllers.New_request)
	// プロフィール取得
	r.GET("/profile_get/:id", controllers.Profile_get)

	// 404の時の処理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
	})

	// 500の時の処理
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
	})

	r.Run(":8080")

	dbconect.Close()

}
