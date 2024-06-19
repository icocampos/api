package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

// NewUsers creates one repository ov users
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statements, err := repository.db.Prepare(
		"insert into users (name. nick, email, pass) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statements.Close()

	result, err := statements.Exec(user.Name, user.Nick, user.Email, user.Pass)
	if err != nil {
		return 0, err
	}
	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil

}
