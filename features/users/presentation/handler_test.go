package presentation

import (
	"be9/restclean/features/users"
	"be9/restclean/features/users/presentation/request"
	"be9/restclean/features/users/presentation/response"
	"be9/restclean/middlewares"
	"be9/restclean/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//with slice
type UsersResponseSuccess struct {
	Message string
	Data    []response.User
}

type SingleUserResponseSuccess struct {
	Message string
	Data    response.User
}

type ResponseGlobal struct {
	Message string
}

var (
	mock_data_user = request.User{
		Name:     "alta",
		Email:    "alta@gmail.com",
		Password: "qwerty",
	}
)

func TestGetAll(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UserBusiness)
	returnData := []users.Core{{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}}

	t.Run("Success Get All", func(t *testing.T) {
		usecase.On("GetAllData", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(returnData, nil).Once()
		srv := NewUserHandler(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")
		// if you want to add param, add this code below
		// context.SetParamNames("id")
		// context.SetParamValues("1")

		responseData := UsersResponseSuccess{}
		// result :=
		if assert.NoError(t, srv.GetAll(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData[0].Name, responseData.Data[0].Name)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		usecase.On("GetAllData", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, errors.New("failed to get all data")).Once()

		srv := NewUserHandler(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		srv.GetAll(echoContext)
		responseBody := rec.Body.String()
		var responseData ResponseGlobal
		err := json.Unmarshal([]byte(responseBody), &responseData)
		fmt.Println("res", responseData)
		if err != nil {
			assert.Error(t, err, "error")
		}
		// assert.Error(t, result)
		assert.Equal(t, "failed to get all data", responseData.Message)
		// assert.Nil(t, res)
		usecase.AssertExpectations(t)
	})
}

/*
contoh testing func yang membutuhkan jwt.
terdapat perbedaan saat pemanggilan fungsi.
Panggil func di handler dengan cara -->

middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllWithJWT))(echoContext)

*/
func TestGetAllWithJWT(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UserBusiness)
	returnData := []users.Core{{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}}

	t.Run("Success Get All With JWT", func(t *testing.T) {
		usecase.On("GetAllData", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(returnData, nil).Once()
		token, errToken := middlewares.CreateToken(returnData[0].ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}
		srv := NewUserHandler(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")
		// if you want to add param, add this code below
		// context.SetParamNames("id")
		// context.SetParamValues("1")

		responseData := UsersResponseSuccess{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllWithJWT))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData[0].Name, responseData.Data[0].Name)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Error Get All when failed to get data", func(t *testing.T) {
		usecase.On("GetAllData", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, errors.New("failed to get all data")).Once()
		token, errToken := middlewares.CreateToken(returnData[0].ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}
		srv := NewUserHandler(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllWithJWT))(echoContext)
		responseBody := rec.Body.String()
		var responseData ResponseGlobal
		err := json.Unmarshal([]byte(responseBody), &responseData)
		fmt.Println("res", responseData)
		if err != nil {
			assert.Error(t, err, "error")
		}
		// assert.Error(t, result)
		assert.Equal(t, "failed to get all data", responseData.Message)
		// assert.Nil(t, res)
		usecase.AssertExpectations(t)
	})

	t.Run("Error Get All when token false (id user < 1)", func(t *testing.T) {
		// usecase.On("GetAllData", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, errors.New("failed to get all data")).Once()
		token, errToken := middlewares.CreateToken(0)
		if errToken != nil {
			assert.Error(t, errToken)
		}
		srv := NewUserHandler(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllWithJWT))(echoContext)
		responseBody := rec.Body.String()
		var responseData ResponseGlobal
		err := json.Unmarshal([]byte(responseBody), &responseData)
		fmt.Println("res", responseData)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "unauthorized", responseData.Message)
		usecase.AssertExpectations(t)
	})
}

func TestAddUser(t *testing.T) {
	reqBody, err := json.Marshal(mock_data_user)
	if err != nil {
		t.Error(t, err, "error")
	}

	e := echo.New()
	usecase := new(mocks.UserBusiness)
	// returnData := []users.Core{{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}}

	t.Run("Success Add User", func(t *testing.T) {
		usecase.On("InsertData", mock.Anything).Return(1, nil).Once()

		srv := NewUserHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.AddUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "success to insert data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Add User when bind error", func(t *testing.T) {

		var dataFail = map[string]int{
			"name": 1,
		}
		reqBodyFail, _ := json.Marshal(dataFail)
		srv := NewUserHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.AddUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed to bind data, check your input", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Add User when insert data failed", func(t *testing.T) {
		usecase.On("InsertData", mock.Anything).Return(-1, errors.New("failed to insert data")).Once()
		// var dataFail = map[string]int{
		// 	"name": 1,
		// }
		// reqBodyFail, _ := json.Marshal(dataFail)
		srv := NewUserHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.AddUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed to insert data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Add User when insert data failed", func(t *testing.T) {
		usecase.On("InsertData", mock.Anything).Return(0, errors.New("failed to insert data")).Once()

		srv := NewUserHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.AddUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "failed to insert data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}
