package response

import (
	"be9/restclean/features/books"
	"time"
)

type Book struct {
	ID        int       `json:"id" form:"id"`
	Title     string    `json:"title" form:"title"`
	Author    string    `json:"author" form:"author"`
	Publisher string    `json:"publisher" form:"publisher"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	User      User      `json:"user" form:"created_at"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCore(data books.Core) Book {
	return Book{
		ID:        data.ID,
		Title:     data.Title,
		Author:    data.Author,
		Publisher: data.Publisher,
		CreatedAt: data.CreatedAt,
		User: User{
			ID:   data.User.ID,
			Name: data.User.Name,
		},
	}
}

func FromCoreList(data []books.Core) []Book {
	result := []Book{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
