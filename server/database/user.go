package database

import (
	"github.com/johnkristanf/VoiceForge/server/types"
	"golang.org/x/crypto/bcrypt"
)

func (sql *SQL_DB) CreateUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS "user"(
		user_id SERIAL PRIMARY KEY,
		email VARCHAR(150) NOT NULL,
		password VARCHAR(255) NOT NULL
	);`

	_, err := sql.database.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (sql *SQL_DB) SignUp(signUpCredentials *types.SignupCredentials) error {
	query := `INSERT INTO "user"(email, password) VALUES ($1, $2);`

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(signUpCredentials.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return hashErr
	}

	_, err := sql.database.Exec(query, signUpCredentials.Email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
