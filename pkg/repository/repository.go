package word

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// repo skeleton
type Repository struct {
	Conn *pgxpool.Pool
}
