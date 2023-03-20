package repository

import (
	"Foundries/domain"
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

type TestSuite struct {
	db       *gorm.DB
	mock     sqlmock.Sqlmock
	userRepo domain.UserRepository
}

func TestFetch(t *testing.T) {
	s := &TestSuite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	s.db, err = gorm.Open(dialector, &gorm.Config{})

	mockUser := []domain.User{
		{
			ID:       "IT_LbTlq16",
			Email:    "nevindra@nodeflux.io",
			Password: "$2a$10$P8Gv710T0BgzQkIFPIFYiuR1z7knjGdEe/9QzrtqNwDwGByC3L0Uq",
		},
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow(mockUser[0].ID, mockUser[0].Email, mockUser[0].Password)

	query := `SELECT * FROM "users"`

	s.mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewUserRepository(s.db)
	_, err = a.Fetch(context.Background())
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}

	assert.NotEmpty(t, mockUser)
	assert.Len(t, mockUser, 1)
}

func TestGetByID(t *testing.T) {
	
}
