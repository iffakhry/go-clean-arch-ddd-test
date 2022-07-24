package data

import (
	"be9/restclean/config"
	"be9/restclean/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

var DbConn = config.InitDBTest()

func TestInsertData(t *testing.T) {
	DbConn.Migrator().DropTable(&User{})
	DbConn.AutoMigrate(&User{})

	repo := NewUserRepository(DbConn)

	t.Run("Test Create User", func(t *testing.T) {
		mockUser := users.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty"}
		row, err := repo.InsertData(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
}

func TestInsertDataFailed(t *testing.T) {
	DbConn.Migrator().DropTable(&User{})
	// DbConn.AutoMigrate(&User{})

	repo := NewUserRepository(DbConn)

	t.Run("Test Create User", func(t *testing.T) {
		mockUser := users.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty"}
		row, err := repo.InsertData(mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestSelectData(t *testing.T) {
	DbConn.Migrator().DropTable(&User{})
	DbConn.AutoMigrate(&User{})
	mockUser := users.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty"}
	dataInput := fromCore(mockUser)
	DbConn.Save(&dataInput)
	repo := NewUserRepository(DbConn)

	t.Run("Test Select Data User", func(t *testing.T) {
		dataResult, err := repo.SelectData("data")
		assert.Nil(t, err)
		// assert.Equal(t, 1, row)
		assert.Equal(t, mockUser.Name, dataResult[0].Name)
	})
}

func TestSelectDataFailed(t *testing.T) {
	DbConn.Migrator().DropTable(&User{})
	// DbConn.AutoMigrate(&User{})

	repo := NewUserRepository(DbConn)

	t.Run("Test Select Data User failed, because table user not found", func(t *testing.T) {
		// mockUser := users.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty"}
		dataResult, err := repo.SelectData("data")
		assert.NotNil(t, err)
		assert.Equal(t, []users.Core{}, dataResult)
	})
}
