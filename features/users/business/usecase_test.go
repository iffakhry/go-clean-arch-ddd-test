package business

import (
	"be9/restclean/features/users"
	"be9/restclean/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllData(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := []users.Core{{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("SelectData", mock.Anything).Return(returnData, nil).Once()

		srv := NewUserBusiness(repo)

		res, err := srv.GetAllData(10, 0)
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		repo.On("SelectData", "").Return(nil, errors.New("data not found")).Once()

		srv := NewUserBusiness(repo)

		res, err := srv.GetAllData(10, 0)
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestInsertData(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := users.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}
	// returnData := users.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, nil).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.InsertData(insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(0, errors.New("there is some error")).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.InsertData(insertData)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert when incomplete data", func(t *testing.T) {
		/*
			dont need to write repo.On because this test case dont need to call data layer. just handle on business layer.
		*/
		// repo.On("InsertData", mock.Anything).Return(-1, errors.New("all input data must be filled")).Once()
		srv := NewUserBusiness(repo)

		_, err := srv.InsertData(users.Core{})
		assert.EqualError(t, err, "all input data must be filled")
		repo.AssertExpectations(t)
	})
}

// //mock data success case
// type mockUserData struct{}

// func (mock mockUserData) SelectData(param string) (data []users.Core, err error) {
// 	return []users.Core{
// 		{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"},
// 	}, nil
// }

// func (mock mockUserData) InsertData(data users.Core) (row int, err error) {
// 	return 1, nil
// }

// //mock data failed case
// type mockUserDataFailed struct{}

// func (mock mockUserDataFailed) SelectData(param string) (data []users.Core, err error) {
// 	return nil, fmt.Errorf("Failed to select data")
// }

// func (mock mockUserDataFailed) InsertData(data users.Core) (row int, err error) {
// 	return 0, fmt.Errorf("failed to insert data ")
// }

// func TestGetAllData(t *testing.T) {
// 	t.Run("Test Get All Data Success", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		result, err := userBusiness.GetAllData(0, 0)
// 		assert.Nil(t, err)
// 		assert.Equal(t, "alta", result[0].Name)
// 	})

// 	t.Run("Test Get All Data Failed", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		result, err := userBusiness.GetAllData(0, 0)
// 		assert.NotNil(t, err)
// 		assert.Nil(t, result)
// 	})
// }

// func TestInsertData(t *testing.T) {
// 	t.Run("Test Insert Data Success", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		newUser := users.Core{
// 			Name:     "alta",
// 			Email:    "alta@mail.id",
// 			Password: "qwerty",
// 		}
// 		result, err := userBusiness.InsertData(newUser)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, result)
// 	})

// 	t.Run("Test Insert Data Failed", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		newUser := users.Core{
// 			Name:     "alta",
// 			Email:    "alta@mail.id",
// 			Password: "qwerty",
// 		}
// 		result, err := userBusiness.InsertData(newUser)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, 0, result)
// 	})

// 	t.Run("Test Insert Data Failed When Email Empty", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		newUser := users.Core{
// 			Name:     "alta",
// 			Password: "qwerty",
// 		}
// 		result, err := userBusiness.InsertData(newUser)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, -1, result)
// 	})

// 	t.Run("Test Insert Data Failed When Password Empty", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		newUser := users.Core{
// 			Name:  "alta",
// 			Email: "alta@mail.id",
// 		}
// 		result, err := userBusiness.InsertData(newUser)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, -1, result)
// 	})
// }
