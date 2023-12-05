package models

import (
	"github.com/Naveenchand06/go-projects/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)
var db *gorm.DB

type Book struct {

	gorm.Model
	Name string `gorm:"json":"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`

}


func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(Book{})
}


func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}


func (b *Book) getBookyById(Id int64) (*Book, *gorm.DB) {
	
}