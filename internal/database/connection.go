package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	Conn       *pgxpool.Pool
	connString string
)

func init() {
	connString = "postgresql://posts:p0stgr3s@localhost:5432/posts"
}

func NewConnection() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Conn, err = pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	return Conn, nil
}
