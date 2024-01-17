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
	if err := db.Table("members").Select("name, id").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// ユーザー作成
func User_create(c *gin.Context) (User, error) {
	db := dbconect.GetDB()
	var u User
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}
	// if err := db.Exec("INSERT INTO members (id, name, age) VALUES(?,?,?)", &u.Id, &u.Name, &u.Age).Error; err != nil {
	if err := db.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil

}
