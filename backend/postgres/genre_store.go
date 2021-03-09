package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/mqrc81/myseries/backend/myseries"
)

type GenreStore struct {
	*sqlx.DB
}

func (store *GenreStore) GetGenre(genreID int) (myseries.Genre, error) {
	var genre myseries.Genre

	query := `
		SELECT * 
		FROM genres 
		WHERE genre_id = $1
		`

	if err := store.Get(&genre, query, genreID); err != nil {
		return myseries.Genre{}, fmt.Errorf("error getting genre by genre_id: %w", err)
	}

	return genre, nil
}

func (store *GenreStore) GetGenres() ([]myseries.Genre, error) {
	var genres []myseries.Genre

	query := `
		SELECT * 
		FROM genres
		`

	if err := store.Select(&genres, query); err != nil {
		return []myseries.Genre{}, fmt.Errorf("error getting all genres: %w", err)
	}

	return genres, nil
}

func (store *GenreStore) GetGenresByShow(showID int) ([]myseries.Genre, error) {
	var genres []myseries.Genre

	query := `
		SELECT * 
		FROM genres 
		    LEFT JOIN shows_genres sg ON genres.genre_id = sg.genre_id 
		WHERE sg.show_id = $1
		`

	if err := store.Select(&genres, query, showID); err != nil {
		return []myseries.Genre{}, fmt.Errorf("error getting genres by show_id: %w", err)
	}

	return genres, nil
}

func (store *GenreStore) CreateGenre(genre myseries.Genre) error {

	query := `
		INSERT INTO genres(name) 
		VALUES ($1)
		`

	if _, err := store.Exec(query, genre.Name); err != nil {
		return fmt.Errorf("error creating new genre: %w", err)
	}

	return nil
}

func (store *GenreStore) DeleteGenre(genreID int) error {

	query := `
		DELETE FROM genres 
		WHERE genre_id = $1
		`

	if _, err := store.Exec(query, genreID); err != nil {
		return fmt.Errorf("error deleting genre: %w", err)
	}

	return nil
}

func (store *GenreStore) AddGenreToShow(showID int, genreID int) error {

	query := `
		INSERT INTO shows_genres(show_id, genre_id) 
		VALUES ($1, $2)
		`

	if _, err := store.Exec(query, showID, genreID); err != nil {
		return fmt.Errorf("error adding genre to show: %w", err)
	}

	return nil
}
