package repositories

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RomainC75/todo2/data/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var email = "myemail@email.com"
var password = "cryptedPassword"

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}

func TestAddUser_shouldSuccess(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	userRepo := &UserRepository{
		DB: db,
	}

	addRow := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow(1, email, password)

	expectedSQL := `INSERT INTO \"users\" (.+) VALUES (.+) RETURNING "id"`

	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
	mock.ExpectCommit()

	var reqUser models.User
	reqUser.Email = email
	reqUser.Password = password

	newUser, _ := userRepo.CreateUser(reqUser)
	fmt.Println("new user : ", newUser)
	assert.Nil(t, mock.ExpectationsWereMet())
}
