package repository

import (
	"back-go/dbconect"
	"back-go/models"
)

type Request models.Request

// 依頼一覧取得
func Request_getAll() ([]Request, error) {
	db := dbconect.GetDB()

	var r []Request
	if err := db.Table("requests").Find(&r).Error; err != nil {
		return r, err
	}

	return r, nil
}

// 依頼新規作成
func Request_create(r_title string, r_description string, r_user_id int, r_photo string) error {
	db := dbconect.GetDB()

	sql := "INSERT INTO REQUESTS(title, description, user_id, photo) VALUES(?, ?, ?, ?)"
	if err := db.Exec(sql, r_title, r_description, r_user_id, r_photo).Error; err != nil {
		return err
	}

	return nil
}
