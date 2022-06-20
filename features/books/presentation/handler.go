package presentation

import (
	"be9/restclean/features/books"
	_requestBook "be9/restclean/features/books/presentation/request"
	_responseBook "be9/restclean/features/books/presentation/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookBusiness books.Business
}

func NewBookHandler(business books.Business) *BookHandler {
	return &BookHandler{
		bookBusiness: business,
	}
}

func (h *BookHandler) GetAll(c echo.Context) error {
	//param, query param, binding
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.bookBusiness.GetAllData(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseBook.FromCoreList(result),
	})
}

func (h *BookHandler) AddBook(c echo.Context) error {
	var newBook _requestBook.Book
	errBind := c.Bind(&newBook)
	//baca token/extract token
	// newBook.UserId = idtoken
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}
	dataUser := _requestBook.ToCore(newBook)
	row, err := h.bookBusiness.InsertData(dataUser)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to insert data",
	})

}
