package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dsn string) (*Store, error) {

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &Store{
		&UserStore{DB: db},
		&ShowStore{DB: db},
		&GenreStore{DB: db},
		&ReleaseStore{DB: db},
		&ProviderStore{DB: db},
		&TokenStore{DB: db},
	}, nil
}

type Store struct {
	*UserStore
	*ShowStore
	*GenreStore
	*ReleaseStore
	*ProviderStore
	*TokenStore
}
