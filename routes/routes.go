package routes

import (
	"be9/restclean/factory"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	// presenter := factory.InitFactory()
	e := echo.New()

	// e.GET("/users", presenter.UserPresenter.GetAll, middlewares.JWTMiddleware())
	e.GET("/users", presenter.UserPresenter.GetAll)
	// e.GET("/products", presenter.ProductPresenter.GetAll)
	// e.GET("/users/:id", _controllers.GetUserByIdController, middlewares.JWTMiddleware())
	e.POST("/users", presenter.UserPresenter.AddUser)
	// e.POST("/auth", _controllers.LoginUserController)
	e.POST("/books", presenter.BookPresenter.AddBook)
	e.GET("/books", presenter.BookPresenter.GetAll)
	return e
}
