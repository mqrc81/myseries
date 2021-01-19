package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/mqrc81/myseries/myseries"
)

type ReleaseStore struct {
	*sqlx.DB
}

func (store *ReleaseStore) GetRelease(showID int) (myseries.Release, error) {
	var release myseries.Release

	query := `
		SELECT * 
		FROM releases 
		WHERE show_id = $1
		ORDER BY date 
		LIMIT 1
		`

	if err := store.Get(&release, query, showID); err != nil {
		return myseries.Release{}, fmt.Errorf("error getting release: %w", err)
	}

	return release, nil
}

func (store *ReleaseStore) GetReleasesByTime(from time.Time, to time.Time) ([]myseries.Release, error) {
	var releases []myseries.Release

	query := `
		SELECT * 
		FROM releases 
		WHERE date >= $1 
		  AND date <= $2
		`

	if err := store.Select(&releases, query, from, to); err != nil {
		return []myseries.Release{}, fmt.Errorf("error getting releases by time: %w", err)
	}

	return releases, nil
}

func (store *ReleaseStore) CreateRelease(release myseries.Release) error {

	query := `
		INSERT INTO releases(show_id, season, part, date, episodes) 
		VALUES ($1, $2, $3, $4, $5)
		`

	if _, err := store.Exec(query,
		release.ShowID,
		release.Season,
		release.Part,
		release.Date,
		release.Episodes); err != nil {
		return fmt.Errorf("error creating new release: %w", err)
	}

	return nil
}

func (store *ReleaseStore) UpdateRelease(release myseries.Release) error {

	query := `
		UPDATE releases 
		SET season = $1, 
		    part = $2, 
		    date = $3, 
		    episodes = $4 
		WHERE show_id = $5
		`

	if _, err := store.Exec(query,
		release.Season,
		release.Part,
		release.Date,
		release.Episodes,
		release.ShowID); err != nil {
		return fmt.Errorf("error updating release: %w", err)
	}

	return nil
}

func (store *ReleaseStore) DeleteRelease(showID int, season int, part string) error {

	query := `
		DELETE FROM releases 
		WHERE show_id = $1 
		  AND season = $2 
		  AND part = $3
		`

	if _, err := store.Exec(query, showID, season, part); err != nil {
		return fmt.Errorf("error deleting release: %w", err)
	}

	return nil
}
