package database

import (
	"github.com/johnkristanf/VoiceForge/server/types"
	"golang.org/x/crypto/bcrypt"
)

func (sql *SQL_DB) CreateUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS "user"(
		user_id SERIAL PRIMARY KEY,
		email VARCHAR(150) NOT NULL,
		password VARCHAR(255) NOT NULL,
		verification_token VARCHAR(255) DEFAULT 'Unverified'
	);`

	_, err := sql.database.Exec(query)
	if err != nil {
		return err
	}

	return nil
}


func (sql *SQL_DB) CreateUserCredentialsIndex() error {
    query := `CREATE INDEX IF NOT EXISTS idx_userCredentials ON "user" (email, user_id);`

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


func (sql *SQL_DB) CheckEmailExist(email string) (*types.UserEmailExist, error) {
    query := `SELECT * FROM "user" WHERE email = $1`
    
    rows, err := sql.database.Query(query, email)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    cred := &types.UserEmailExist{}

    if rows.Next() {
        if err := rows.Scan(&cred.ID, &cred.Email, &cred.Password, &cred.Verification_Token); err != nil {
            return nil, err
        }
    }

    return cred, nil
}


func (sql *SQL_DB) VerifyUser(user_id int64, hashedCode string) error {

	query := `UPDATE "user" SET verification_token = $1 WHERE user_id = $2;`;

	_, err := sql.database.Exec(query, hashedCode, user_id)
	if err != nil{
		return err
	}

	return nil
}