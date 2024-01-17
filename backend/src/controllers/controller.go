package controllers

import (
	"back-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type book struct {
// 	ID       string `json:"id"`
// 	Title    string `json:"title"`
// 	Author   string `json:"author"`
// 	Quantity int    `json:"quantity"`
// }

// var books = []book{
// 	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
// 	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
// 	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
// }

// テスト
func Test(c *gin.Context) {

	// c.IndentedJSON(http.StatusOK, books)

	// var newBook book

	// if err := c.BindJSON(&newBook); err != nil {
	// 	return
	// }

	// books = append(books, newBook)
	// c.IndentedJSON(http.StatusCreated, newBook)

	// id := c.Param("id")
	// book, err := getBookById(id)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
	// 	return
	// }
	// c.IndentedJSON(http.StatusOK, book)

}

// func getBookById(id string) (*book, error) {
// 	for i, b := range books {
// 		if b.ID == id {
// 			return &books[i], nil
// 		}
// 	}

// 	return nil, errors.New("book not found")
// }

// ログイン
func Login(c *gin.Context) {

}

// 新規会員登録
func New_login(c *gin.Context) {

	p, err := repository.User_create(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, p)
	}

}

// プロフィール取得
func Profile_get(c *gin.Context) {

}

// 依頼一覧取得
func Request_get(c *gin.Context) {
	p, err := repository.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

// 依頼新規作成
func New_request(c *gin.Context) {
}
