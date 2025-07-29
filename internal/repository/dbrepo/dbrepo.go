package dbrepo

import (
	"database/sql"

	"github.com/heavensfavorite/bookings/internal/config"
	"github.com/heavensfavorite/bookings/internal/models"
	"github.com/heavensfavorite/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// GetUserByID implements repository.DatabaseRepo.
func (m *postgresDBRepo) GetUserByID(id int) (models.User, error) {
	panic("unimplemented")
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}
