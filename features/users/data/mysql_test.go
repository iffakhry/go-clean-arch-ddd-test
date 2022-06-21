package data

import (
	"be9/restclean/config"
	"be9/restclean/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	repo := NewUserRepository(db)

	t.Run("Test Create User", func(t *testing.T) {
		mockUser := users.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty"}
		row, err := repo.InsertData(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
}
