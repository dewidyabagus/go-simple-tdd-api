package user

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-simple-api/business/user"
)

type Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *Repository
}

func TestUser(t *testing.T) {
	s := &Suite{}

	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		t.Errorf("Failed initiate mock, got error %v", err)
	} else if s.mock == nil {
		t.Errorf("mock value is null")
	} else if db == nil {
		t.Errorf("db value is null")
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed open database session, got error %v", err)
	}
	s.repo = NewRepository(s.db)

	user := &user.User{
		ID:        1,
		UniqueID:  "xxxxx-xxxxx-xxx",
		Username:  "joedo",
		Email:     "joedo@example.com",
		FirstName: "joe",
		LastName:  "do",
		Password:  "random_password",
		Verify:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("username","email","first_name","last_name","password","verify","created_at","updated_at","deleted_at") 
							VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "unique_id","id"`)).
		WithArgs(
			user.Username, user.Email, user.FirstName, user.LastName, user.Password, user.Verify, user.CreatedAt, user.UpdatedAt, nil,
		).
		WillReturnRows(sqlmock.NewRows([]string{"unique_id", "id"}).AddRow(user.UniqueID, user.ID))
	s.mock.ExpectCommit()

	if err := s.repo.InsertNew(user); err != nil {
		t.Errorf("Failed insert new user, got error %v", err)
	}

	if err = s.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Failed to meet expectations, got error %v", err)
	}
}
