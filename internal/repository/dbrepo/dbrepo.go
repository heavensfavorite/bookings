package dbrepo

import (
	"database/sql"

	"github.com/heavensfavorite/bookings/internal/config"
	"github.com/heavensfavorite/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// InsertRoomRestriction implements repository.DatabaseRepo.
/* func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	panic("unimplemented")
} */

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
