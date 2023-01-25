package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

var dbCtx = context.Background()

func InitDBSession() error {
	var err error
	Conn, err = pgx.Connect(dbCtx, Config.DatabaseURL)

	err = Conn.Ping(context.Background())
	fmt.Println("pg connected and pinged!")
	return err
}
