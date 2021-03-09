package web

import (
	"database/sql"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"

	"github.com/mqrc81/myseries/backend/myseries"
)

func NewSessionManager(dataSourceName string) (*scs.SessionManager, error) {

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(25)

	sessions := scs.New()
	sessions.Store = postgresstore.New(db)
	return sessions, nil
}

type SessionData struct {
	FlashMessageSuccess string
	FlashMessageInfo    string
	FlashMessageError   string
	Form                interface{}
	User                myseries.User
	LoggedIn            bool
}
