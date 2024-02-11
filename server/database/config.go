package database

import (
	"bytes"
	"database/sql"
	"os"

	"github.com/johnkristanf/VoiceForge/server/types"
	_ "github.com/lib/pq"
)

type Method interface {
	CheckVoicesValues() (int, error)
	InsertVoice(*types.VoiceStruct) error
	Voices(string) ([]*types.FetchVoiceTypes, error)

	InsertAudioStream(string, *bytes.Buffer) error
	FetchAudioStream() ([]*types.AudioStruct, error)
	DeleteAudioData(int64) (int64, error)

	SignUp(*types.SignupCredentials) error
	CheckEmailExist(string) (*types.UserEmailExist, error)

	VerifyUser(int64, string) error
}

type SQL_DB struct {
	database *sql.DB
}

func VoiceForgeDB() (*SQL_DB, error) {

	// var (
	// 	host = os.Getenv("DB_HOST")
	// 	port = os.Getenv("DB_PORT")

	// 	username = os.Getenv("DB_USERNAME")
	// 	password = os.Getenv("DB_PASSWORD")

	// 	dbname             = os.Getenv("DB_NAME")
	// 	sslmode            = os.Getenv("DB_SSLMODE")
	// 	connection_timeout = os.Getenv("DB_CONNECTION_TIMEOUT")
	// )

	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s",
	// 	host, port, username, password, dbname, sslmode, connection_timeout)

	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

	return &SQL_DB{database: db}, nil

}

func (sql *SQL_DB) DBInit() error {

	if err := sql.CreateVoicesTable(); err != nil {
		return err
	}

	if err := sql.CreateVoiceIndex(); err != nil {
		return err
	}

	if err := sql.CreateAudioTable(); err != nil {
		return err
	}

	if err := sql.CreateUserTable(); err != nil {
		return err
	}

	if err := sql.CreateUserCredentialsIndex(); err != nil {
		return err
	}

	return nil
}
