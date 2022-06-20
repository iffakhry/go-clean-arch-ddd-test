package request

import "be9/restclean/features/books"

type Book struct {
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	UserId    int    `json:"user_id" form:"user_id"`
}

func ToCore(req Book) books.Core {
	return books.Core{
		Title:     req.Title,
		Author:    req.Author,
		Publisher: req.Publisher,
		User: books.User{
			ID: req.UserId,
		},
	}
}
