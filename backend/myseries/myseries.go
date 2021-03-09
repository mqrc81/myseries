package myseries

import (
	"time"
)

type User struct {
	UserID             int    `db:"user_id"`
	Username           string `db:"username"`
	Password           string `db:"password"`
	Email              string `db:"email"`
	EmailConfirmed     bool   `db:"email_confirmed"`
	EmailNotifications bool   `db:"notifications"`
}

type Show struct {
	ShowID      int     `db:"show_id"`
	Name        string  `db:"name"`
	Year        int     `db:"year"`
	Description string  `db:"description"`
	Poster      string  `db:"poster"`
	Seasons     int     `db:"seasons"`
	Rating      float64 `db:"rating"`
	SeasonEnded bool    `db:"season_ended"`
	ProviderID  int     `db:"provider_id"`
}

type Genre struct {
	GenreID int    `db:"genre_id"`
	Name    string `db:"name"`
}

type Release struct {
	ShowID   int       `db:"show_id"`
	Season   string    `db:"season"`
	Part     int       `db:"part"`
	Date     time.Time `db:"date"`
	Episodes int       `db:"episodes"`
}

type Provider struct {
	ProviderID int    `db:"provider_id"`
	Name       string `db:"name"`
	URL        string `db:"url"`
}

type Token struct {
	TokenID string    `db:"token_id"`
	UserID  int       `db:"user_id"`
	Expiry  time.Time `db:"expiry"`
}

type UserStore interface {
	GetUser(userID int) (User, error)
	GetUserByUsername(username string) (User, error)
	GetUsers() ([]User, error)
	GetUsersByShow(showID int) ([]User, error)
	CreateUser(user User) error
	UpdateUser(user User) error
	DeleteUser(userID User) error
}

type ShowStore interface {
	GetShow(showID int) (Show, error)
	GetShows() ([]Show, error)
	GetShowsByUser(userID int) ([]Show, error)
	GetShowsByGenreAndDate(genreID int, from time.Time, to time.Time) ([]Show, error)
	CreateShow(show Show) error
	UpdateShow(show Show) error
	DeleteShow(showID Show) error
	AddShowToUser(userID int, showID int) error
}

type GenreStore interface {
	GetGenre(genreID int) (Genre, error)
	GetGenres() ([]Genre, error)
	GetGenresByShow(showID int) ([]Genre, error)
	CreateGenre(genre Genre) error
	DeleteGenre(genreID int) error
	AddGenreToShow(showID int, genreID int) error
}

type ReleaseStore interface {
	GetRelease(releaseID int) (Release, error)
	GetReleasesByTime(from time.Time, to time.Time) ([]Release, error)
	CreateRelease(release Release) error
	UpdateRelease(release Release) error
	DeleteRelease(showID int, season int, part string) error
}

type ProviderStore interface {
	GetProvider(providerID int) (Provider, error)
	GetProviders() ([]Provider, error)
	CreateProvider(provider Provider) error
	UpdateProvider(provider Provider) error
	DeleteProvider(providerID int) error
}

type TokenStore interface {
	GetToken(tokenID string) (Token, error)
	CreateToken(token Token) error
	DeleteTokensByUser(userID int) error
}

type Store interface {
	UserStore
	ShowStore
	GenreStore
	ReleaseStore
	ProviderStore
	TokenStore
}
