package business

import (
	"be9/restclean/features/books"
	"errors"
)

type bookUsecase struct {
	bookData books.Data
}

func NewBookBusiness(bkData books.Data) books.Business {
	return &bookUsecase{
		bookData: bkData,
	}
}

func (uc *bookUsecase) GetAllData(limit, offset int) (resp []books.Core, err error) {
	var param string
	resp, err = uc.bookData.SelectData(param)
	return resp, err
}

func (uc *bookUsecase) InsertData(input books.Core) (row int, err error) {
	if input.Title == "" || input.Author == "" || input.Publisher == "" || input.User.ID == 0 {
		return -1, errors.New("all input data must be filled")
	}
	row, err = uc.bookData.InsertData(input)
	return row, err
}
