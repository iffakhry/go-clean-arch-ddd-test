package factory

import (
	_bookBusiness "be9/restclean/features/books/business"
	_bookData "be9/restclean/features/books/data"
	_bookPresentation "be9/restclean/features/books/presentation"
	_userBusiness "be9/restclean/features/users/business"
	_userData "be9/restclean/features/users/data"
	_userPresentation "be9/restclean/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter *_userPresentation.UserHandler
	BookPresenter *_bookPresentation.BookHandler
	// ProductPresenter *_productPresentation.ProductHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// dbConn := config.InitDB()

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	bookData := _bookData.NewBookRepository(dbConn)
	bookBusiness := _bookBusiness.NewBookBusiness(bookData)
	bookPresentation := _bookPresentation.NewBookHandler(bookBusiness)

	// productData := _productData.NewProductRepository(dbConn)
	// productBusiness := _productBusiness.NewProductBusiness(productData)
	// productPresentation := _productPresentation.NewProductHandler(productBusiness)

	return Presenter{
		UserPresenter: userPresentation,
		BookPresenter: bookPresentation,
		// ProductPresenter: productPresentation,
	}
}
