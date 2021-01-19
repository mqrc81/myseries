package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dsn string) (*Store, error) {

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error opening database: %e", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("error pinging database: %e", err)
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
