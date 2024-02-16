package database

import (
	"database/sql"

	"github.com/johnkristanf/VoiceForge/server/types"
)

func (sql *SQL_DB) CreateAudioTable() error {

	query := `CREATE TABLE IF NOT EXISTS audio (
		id SERIAL PRIMARY KEY,
		text TEXT NOT NULL,
		audioStreamBase64 BYTEA NOT NULL
	);`

	_, err := sql.database.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (sql *SQL_DB) CreateAudioIndex() error {

	query := "CREATE INDEX IF NOT EXISTS idx_audio_id ON audio (id);"

	_, err := sql.database.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (sql *SQL_DB) InsertAudioStream(text string, audiostreambase64 []byte) (int64, error) {

	var lastInsertedID int64

	query := "INSERT INTO audio (text, audioStreamBase64) VALUES ($1, $2) RETURNING id"

    err := sql.database.QueryRow(query, text, audiostreambase64).Scan(&lastInsertedID)
	if err != nil{
		return 0, err
	}

	return lastInsertedID, nil
}

func (sql *SQL_DB) FetchAudioStream() ([]*types.AudioStruct, error) {

	query := "SELECT * FROM audio;"

	errorChan := make(chan error, 1)
	audioStreamChan := make(chan []*types.AudioStruct, 1)

	rows, err := sql.database.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	go func ()  {
		defer close(errorChan)
		defer close(audioStreamChan)
		
		audioStreamData, err := ScantoRowsAudioData(rows)
		if err != nil{
			errorChan <- err
		}

		audioStreamChan <- audioStreamData
		
	}()

	select{

	   case err := <- errorChan:
		return nil, err

	   case audioStreamData := <- audioStreamChan:
		return audioStreamData, nil
		
	}
	
}


func (sql *SQL_DB) DeleteAudioData(audio_id int64) (int64, error) {

	var lastDeletedID int64

	query := "DELETE FROM audio WHERE id = $1 RETURNING id"

	err := sql.database.QueryRow(query, audio_id).Scan(&lastDeletedID)
	if err != nil{
		return 0, err
	}

	return lastDeletedID, nil
}




func ScantoRowsAudioData(rows *sql.Rows) ([]*types.AudioStruct, error)  {

	audioStreamData := []*types.AudioStruct{}

	for rows.Next() {

		fields := &types.AudioStruct{}

		if err := rows.Scan(&fields.ID, &fields.AudioText, &fields.AudioStream); err != nil {
			return nil, err
		}

		audioStreamData = append(audioStreamData, fields)
	}

	return audioStreamData, nil
}