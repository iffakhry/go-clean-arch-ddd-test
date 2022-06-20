package migration

import (
	_mBooks "be9/restclean/features/books/data"
	_mUsers "be9/restclean/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mBooks.Book{})
}
