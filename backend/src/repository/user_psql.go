package repository

import (
	"back-go/dbconect"
	"back-go/models"

	"github.com/gin-gonic/gin"
)

type User models.User

// ユーザー全件表示
func GetAll() ([]User, error) {
	db := dbconect.GetDB()
	var u []User
	if err := db.Table("users").Find(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// ユーザー表示
func GetID(id int) (User, error) {
	db := dbconect.GetDB()
	var u User
	if err := db.Table("users").Where("user_id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// ユーザー作成
func User_create(c *gin.Context) (User, error) {
	db := dbconect.GetDB()
	var u User
	sql := "INSERT INTO USERS (user_name, user_email, user_passwd, user_gender) VALUES(?,?,?,?)"
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}
	if err := db.Exec(sql, &u.User_name, &u.User_email, &u.User_passwd, &u.User_gender).Error; err != nil {
		return u, err
	}
	return u, nil

}

// ユーザーログイン
func User_login(u_email string, u_passwd string) error {
	db := dbconect.GetDB()
	var u User

	if err := db.Table("users").Where("user_email = ? AND user_passwd = ?", u_email, u_passwd).First(&u).Error; err != nil {
		return err
	}

	return nil
}
