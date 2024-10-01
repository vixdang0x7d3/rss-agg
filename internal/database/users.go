package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DbUser struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      sql.NullString
}

const sqlCreateUser = `
	INSERT INTO rssagg.users (id, created_at, updated_at, name)
	VALUES (:id, :createdat, :updatedat, :name)
`

func CreateUser(db *sqlx.DB, user DbUser) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.NamedExec(sqlCreateUser, &user); err != nil {
		return err
	}

	return tx.Commit()
}
