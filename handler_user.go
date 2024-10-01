package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/vixdang0x7d3/rss-agg/internal/database"
)

type AppUser struct {
	Name string `json:"name"`
}

func handlerCreateUser(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		u := AppUser{}

		if err = c.Bind(&u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to bind user object: %v", err))
		}

		if err = database.CreateUser(db, database.DbUser{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      sql.NullString{String: u.Name, Valid: true},
		}); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, u)
	}
}
