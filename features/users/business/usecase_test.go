package business

import (
	"be9/restclean/features/users"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockUserDataSuccess struct{}

func (mock mockUserDataSuccess) SelectData(param string) (data []users.Core, err error) {
	return []users.Core{
		{ID: 1, Name: "alta", Email: "alta@mail.com", Password: "qwerty"},
	}, nil
}

func (mock mockUserDataSuccess) InsertData(data users.Core) (row int, err error) {
	return 1, nil
}

type mockUserDataFailed struct{}

func (mock mockUserDataFailed) SelectData(param string) (data []users.Core, err error) {
	return nil, fmt.Errorf("error")
}

func (mock mockUserDataFailed) InsertData(data users.Core) (row int, err error) {
	return 0, fmt.Errorf("error insert data")
}

func TestGetAllData(t *testing.T) {
	t.Run("Test Get All Success", func(t *testing.T) {
		userUsecase := NewUserBusiness(mockUserDataSuccess{})
		result, err := userUsecase.GetAllData(0, 0)
		assert.Nil(t, err)
		assert.Equal(t, "alta", result[0].Name)
	})

	t.Run("Test Get All Failed", func(t *testing.T) {
		userUsecase := NewUserBusiness(mockUserDataFailed{})
		result, err := userUsecase.GetAllData(0, 0)
		assert.NotNil(t, err)
		assert.Equal(t, nil, result)
	})
}

func TestInsertData(t *testing.T) {
	t.Run("Test Insert Data Success", func(t *testing.T) {
		userUsecase := NewUserBusiness(mockUserDataSuccess{})
		newUser := users.Core{
			ID:       1,
			Name:     "alta",
			Email:    "alta@mail.com",
			Password: "qwerty",
		}
		result, err := userUsecase.InsertData(newUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Test Insert Data Incompleted Failed", func(t *testing.T) {
		userUsecase := NewUserBusiness(mockUserDataSuccess{})
		newUser := users.Core{
			ID:    1,
			Name:  "alta",
			Email: "alta@mail.com",
		}
		result, err := userUsecase.InsertData(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, -1, result)
	})

	t.Run("Test Insert Data Incompleted Failed", func(t *testing.T) {
		userUsecase := NewUserBusiness(mockUserDataSuccess{})
		newUser := users.Core{
			ID:       1,
			Name:     "alta",
			Email:    "alta@mail.com",
			Password: "1234",
		}
		result, err := userUsecase.InsertData(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}
