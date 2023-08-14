package models

import (
	"github.com/hardzal/go-example/go-books/pkg/config"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	Db = config.GetDB()
	Db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	Db.NewRecord(b)
	Db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	Db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := Db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	Db.Where("ID = ? ", ID).Delete(book)
	return book
}
