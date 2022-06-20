package books

import "time"

type Core struct {
	ID        int
	Title     string
	Author    string
	Publisher string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}

type User struct {
	ID   int
	Name string
}

type Business interface {
	GetAllData(limit, offset int) (data []Core, err error)
	InsertData(data Core) (row int, err error)
}

type Data interface {
	SelectData(param string) (data []Core, err error)
	InsertData(data Core) (row int, err error)
}
