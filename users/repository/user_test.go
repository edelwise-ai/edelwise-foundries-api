package repository

import (
	"Foundries/domain"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type TestSuite struct {
	db       *gorm.DB
	mock     sqlmock.Sqlmock
	userRepo domain.UserRepository
}

func (s *TestSuite) SetupSuite(t *testing.T) {
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
}

func TestFetch(t *testing.T) {
	s := &TestSuite{}

	s.SetupSuite(t)

	mockUser := []domain.User{
		{
			ID:       "IT_LbTlq16",
			Email:    "nevindra@nodeflux.io",
			Password: "$2a$10$P8Gv710T0BgzQkIFPIFYiuR1z7knjGdEe/9QzrtqNwDwGByC3L0Uq",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow(mockUser[0].ID, mockUser[0].Email, mockUser[0].Password)

	query := "SELECT * FROM \"users\""

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
	a := NewUserRepository(s.db)
	_, err := a.Fetch()
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}

	assert.NotEmpty(t, mockUser)
	assert.Len(t, mockUser, 1)
}

//func TestGetByID(t *testing.T) {
//	s := &TestSuite{}
//
//	s.SetupSuite(t)
//
//	mockUser := []domain.User{
//		{
//			ID:       "IT_LbTlq16",
//			Email:    "nevindra@nodeflux.io",
//			Password: "$2a$10$P8Gv710T0BgzQkIFPIFYiuR1z7knjGdEe/9QzrtqNwDwGByC3L0Uq",
//		},
//	}
//
//	rows := sqlmock.NewRows([]string{"id", "email", "password"}).
//		AddRow(mockUser[0].ID, mockUser[0].Email, mockUser[0].Password)
//
//	query := "SELECT * FROM \"users\" WHERE \"id\" = ? ORDER BY \"users\".\"id\" LIMIT 1"
//
//	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
//	a := NewUserRepository(s.db)
//	_, err := a.GetByID("IT_LbTlq16")
//	if err != nil {
//		t.Errorf("Error should be nil, got %v", err)
//	}
//
//	assert.NotEmpty(t, mockUser)
//	assert.Len(t, mockUser, 1)
//
//	// check if the mock was called as expected
//	if err := s.mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}
