package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/mqrc81/myseries/myseries"
)

type ShowStore struct {
	*sqlx.DB
}

func (store *ShowStore) GetShow(showID int) (myseries.Show, error) {
	var show myseries.Show

	query := `
		SELECT *
		FROM shows
		WHERE show_id = $1
		`

	if err := store.Get(&show, query, showID); err != nil {
		return myseries.Show{}, fmt.Errorf("error getting show by show_id: %w", err)
	}

	return show, nil
}

func (store *ShowStore) GetShowsByUser(userID int) ([]myseries.Show, error) {
	var shows []myseries.Show

	query := `
		SELECT * 
		FROM shows 
		    LEFT JOIN users_shows us ON shows.show_id = us.show_id
		WHERE us.user_id = $1
		`

	if err := store.Select(&shows, query, userID); err != nil {
		return []myseries.Show{}, fmt.Errorf("error getting shows by user_id: %w", err)
	}

	return shows, nil
}

func (store *ShowStore) GetShowsByGenres(genreID int) ([]myseries.Show, error) {
	var shows []myseries.Show

	query := `
		SELECT * 
		FROM shows 
		    LEFT JOIN shows_genres sg ON shows.show_id = sg.show_id 
		WHERE sg.genre_id = $1
		`

	if err := store.Select(&shows, query, genreID); err != nil {
		return []myseries.Show{}, fmt.Errorf("error getting shows by genre_id: %w", err)
	}

	return shows, nil
}

func (store *ShowStore) CreateShow(show myseries.Show) error {

	query := `
		INSERT INTO shows(provider_id, name, year, description, poster, seasons, imdb_rating, ended) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`

	if _, err := store.Exec(query,
		show.ProviderID,
		show.Name,
		show.Year,
		show.Description,
		show.Poster,
		show.Seasons,
		show.IMDBRating,
		show.SeasonEnded); err != nil {
		return fmt.Errorf("error creating new show: %w", err)
	}

	return nil
}

func (store *ShowStore) UpdateShow(show myseries.Show) error {

	query := `
		UPDATE shows 
		SET provider_id = $1, 
		    name = $2, 
		    year = $3, 
		    description = $4, 
		    poster = $5, 
		    seasons = $6, 
		    imdb_rating = $7, 
		    ended = $8 
		WHERE show_id = $9
		`

	if _, err := store.Exec(query,
		show.ProviderID,
		show.Name,
		show.Year,
		show.Description,
		show.Poster,
		show.Seasons,
		show.IMDBRating,
		show.SeasonEnded,
		show.ShowID); err != nil {
		return fmt.Errorf("error creating show: %w", err)
	}

	return nil
}

func (store *ShowStore) DeleteShow(showID myseries.Show) error {

	query := `
		DELETE FROM shows 
		WHERE show_id = $1
		`

	if _, err := store.Exec(query, showID); err != nil {
		return fmt.Errorf("error deleting show: %w", err)
	}

	return nil
}

func (store *ShowStore) AddShowToUser(userID int, showID int) error {

	query := `
		INSERT INTO users_shows(user_id, show_id) 
		VALUES ($1, $2)
		`

	if _, err := store.Exec(query, userID, showID); err != nil {
		return fmt.Errorf("error adding show to user: %w", err)
	}

	return nil
}
