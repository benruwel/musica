package utils

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

var dbCtx = context.Background()

func InitDBSession() error {
	var err error
	Conn, _ = pgx.Connect(dbCtx, Config.DatabaseURL)

	err = Conn.Ping(dbCtx)
	return err
}
