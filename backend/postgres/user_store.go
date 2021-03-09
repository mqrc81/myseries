package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/mqrc81/myseries/backend/myseries"
)

type UserStore struct {
	*sqlx.DB
}

func (store *UserStore) GetUser(userID int) (myseries.User, error) {
	var user myseries.User

	query := `
		SELECT *
		FROM users
		WHERE user_id = $1
		`

	if err := store.Get(&user, query, userID); err != nil {
		return myseries.User{}, fmt.Errorf("error getting user by user_id: %w", err)
	}

	return user, nil
}

func (store *UserStore) GetUserByUsername(username string) (myseries.User, error) {
	var user myseries.User

	query := `
		SELECT * 
		FROM users 
		WHERE username = $1
		`

	if err := store.Get(&user, query, username); err != nil {
		return myseries.User{}, fmt.Errorf("error getting user by username: %w", err)
	}

	return user, nil
}

func (store *UserStore) GetUsers() ([]myseries.User, error) {
	var users []myseries.User

	query := `
		SELECT * 
		FROM users
		`

	if err := store.Select(&users, query); err != nil {
		return []myseries.User{}, fmt.Errorf("error getting all users: %w", err)
	}

	return users, nil
}

func (store *UserStore) GetUsersByShow(showID int) ([]myseries.User, error) {
	var users []myseries.User

	query := `
		SELECT * 
		FROM users 
		    LEFT JOIN users_shows us ON users.user_id = us.user_id
		WHERE us.show_id = $1
		`

	if err := store.Select(&users, query, showID); err != nil {
		return []myseries.User{}, fmt.Errorf("error getting users by show_id: %w", err)
	}

	return users, nil
}

func (store *UserStore) CreateUser(user myseries.User) error {

	query := `
		INSERT INTO users(username, password, email) 
		VALUES ($1, $2, $3)
		`

	if _, err := store.Exec(query, user.Username, user.Password, user.Email); err != nil {
		return fmt.Errorf("error creating new user: %w", err)
	}

	return nil
}

func (store *UserStore) UpdateUser(user myseries.User) error {

	query := `
		UPDATE users 
		SET username = $1, 
		    password = $2, 
		    email = $3 
		WHERE user_id = $4
		`

	if _, err := store.Exec(query, user.Username, user.Password, user.Email, user.UserID); err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (store *UserStore) DeleteUser(userID myseries.User) error {

	query := `
		DELETE FROM users
		WHERE user_id = $1
		`

	if _, err := store.Exec(query, userID); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}
