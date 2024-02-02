package models

type Login_user struct {
	User_email  string `json:"user_email" binding:"required"`
	User_passwd string `json:"user_passwd" binding:"required"`
}

type Request_create struct {
	Request_title       string `json:"request_title" binding:"required"`
	Request_description string `json:"request_description" binding:"required"`
	Request_user_id     int    `json:"request_user_id" binding:"required"`
	Request_photo       string `json:"request_photo"`
}
