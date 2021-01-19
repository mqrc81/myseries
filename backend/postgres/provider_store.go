package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/mqrc81/myseries/myseries"
)

type ProviderStore struct {
	*sqlx.DB
}

func (store *ProviderStore) GetProvider(providerID int) (myseries.Provider, error) {
	var provider myseries.Provider

	query := `
		SELECT * 
		FROM providers 
		WHERE provider_id = $1
		`

	if err := store.Get(&provider, query, providerID); err != nil {
		return myseries.Provider{}, fmt.Errorf("error getting provider by provider_id: %w", err)
	}

	return provider, nil
}

func (store *ProviderStore) GetProviders() ([]myseries.Provider, error) {
	var providers []myseries.Provider

	query := `
		SELECT * 
		FROM providers
		`

	if err := store.Select(&providers, query); err != nil {
		return []myseries.Provider{}, fmt.Errorf("error getting all providers: %w", err)
	}

	return providers, nil
}

func (store *ProviderStore) CreateProvider(provider myseries.Provider) error {

	query := `
		INSERT INTO providers(name, url) 
		VALUES ($1, $2)
		`

	if _, err := store.Exec(query,
		provider.Name,
		provider.URL); err != nil {
		return fmt.Errorf("error creating new provider: %w", err)
	}

	return nil
}

func (store *ProviderStore) UpdateProvider(provider myseries.Provider) error {

	query := `
		UPDATE providers 
		SET name = $1, 
		    url = $2 
		WHERE provider_id = $3
		`

	if _, err := store.Exec(query,
		provider.Name,
		provider.URL); err != nil {
		return fmt.Errorf("error updating provider: %w", err)
	}

	return nil
}

func (store *ProviderStore) DeleteProvider(providerID int) error {

	query := `
		DELETE FROM providers 
		WHERE provider_id = $1
		`

	if _, err := store.Exec(query, providerID); err != nil {
		return fmt.Errorf("error deleting provider: %w", err)
	}

	return nil
}
