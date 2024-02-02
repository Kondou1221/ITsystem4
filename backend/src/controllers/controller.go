package controllers

import (
	"back-go/models"
	"back-go/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginInput models.Login_user
type requestInput models.Request_create

// テスト
func Test(c *gin.Context) {
	u, err := repository.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, u)
	}

}

// ログイン
func Login(c *gin.Context) {

	var l LoginInput

	if err := c.ShouldBindJSON(&l); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストが不正です"})
	} else {
		err := repository.User_login(l.User_email, l.User_passwd)

		if err != nil {
			c.AbortWithStatus(404)
			c.JSON(http.StatusBadRequest, gin.H{"error": "ログイン失敗"})
		} else {
			c.JSON(200, "ログイン成功")
		}
	}
}

// 新規会員登録
func New_login(c *gin.Context) {

	p, err := repository.User_create(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, p.User_id)
	}

}

// プロフィール取得
func Profile_get(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, _ := strconv.Atoi(id)
	p, err := repository.GetID(idInt)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// 依頼一覧取得
func Request_get(c *gin.Context) {
	r, err := repository.Request_getAll()

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		log.Println(r)
		c.JSON(200, r)
	}

}

// 依頼新規作成
func New_request(c *gin.Context) {
	var r requestInput

	if err := c.ShouldBindJSON(&r); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストが不正です"})
	} else {
		err := repository.Request_create(r.Request_title, r.Request_description, r.Request_user_id, r.Request_photo)

		if err != nil {
			c.AbortWithStatus(400)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, "依頼作成しました")
		}
	}

}
