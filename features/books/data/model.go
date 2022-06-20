package data

import (
	"be9/restclean/features/books"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	UserID    uint   `json:"user_id" form:"user_id"`
	User      User
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Books    []Book
}

// DTO

func (data *Book) toCore() books.Core {
	return books.Core{
		ID:        int(data.ID),
		Title:     data.Title,
		Author:    data.Author,
		Publisher: data.Publisher,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		User: books.User{
			ID:   int(data.User.ID),
			Name: data.User.Name,
		},
	}
}

func toCoreList(data []Book) []books.Core {
	result := []books.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core books.Core) Book {
	return Book{
		Title:     core.Title,
		Author:    core.Author,
		Publisher: core.Publisher,
		UserID:    uint(core.User.ID),
	}
}
