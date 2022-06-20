package request

import "be9/restclean/features/users"

type User struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(req User) users.Core {
	return users.Core{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
