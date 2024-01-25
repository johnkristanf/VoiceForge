package database

import (
	"bytes"

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


func (sql *SQL_DB) InsertAudioStream(text string, audiostreambase64 *bytes.Buffer) error {

	query := "INSERT INTO audio (text, audioStreamBase64) VALUES ($1, $2)"
    
	_, err := sql.database.Exec(query, text, audiostreambase64.Bytes())
	if err != nil{
		return err
	}

	return nil
}


func (sql *SQL_DB) FetchAudioStream() ([]*types.AudioStruct, error) {

	query := "SELECT * FROM audio;"

	rows, err := sql.database.Query(query)
	if err != nil{
		return nil, err
	}

	audioStreamData := []*types.AudioStruct{}

	for rows.Next(){
		
		fields := &types.AudioStruct{}

		if err := rows.Scan(&fields.ID, &fields.AudioText, &fields.AudioStream); err != nil{
			return nil, err
		}

		audioStreamData = append(audioStreamData, fields)
	}

	return audioStreamData, nil
}


func (sql *SQL_DB) DeleteAudioData(audio_id int64) (int64, error) {
	query := "DELETE FROM audio WHERE id = $1 RETURNING id"

	rows, err := sql.database.Query(query, audio_id)
	if err != nil {
		return 0, err
	}

	
    var lastDeletedID int64
	for rows.Next() {
		if err := rows.Scan(&lastDeletedID); err != nil {
			return 0, err
		}
	}

	return lastDeletedID, nil
}