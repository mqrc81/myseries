package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/mqrc81/myseries/myseries"
)

type TokenStore struct {
	*sqlx.DB
}

func (store *TokenStore) GetToken(tokenID string) (myseries.Token, error) {
	var token myseries.Token

	query := `
		SELECT * 
		FROM tokens 
		WHERE token_id = $1
		`
	if err := store.Get(token, query, tokenID); err != nil {
		return myseries.Token{}, fmt.Errorf("error getting token: %w", err)
	}

	return token, nil
}

func (store *TokenStore) CreateToken(token myseries.Token) error {

	query := `
		INSERT INTO tokens(token_id, user_id, expiry) 
		VALUES ($1, $2, $3)
		`
	if _, err := store.Exec(query,
		token.TokenID,
		token.UserID,
		token.Expiry,
	); err != nil {
		return fmt.Errorf("error creating token: %w", err)
	}

	return nil
}

func (store *TokenStore) DeleteTokensByUser(userID int) error {

	query := `
		DELETE FROM tokens
		WHERE user_id = $1
		`
	if _, err := store.Exec(query, userID); err != nil {
		return fmt.Errorf("error deleting tokens by user: %w", err)
	}

	return nil
}
