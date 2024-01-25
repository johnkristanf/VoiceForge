package database

import (
	"database/sql"

	"github.com/johnkristanf/VoiceForge/server/types"
)

func (sql *SQL_DB) CreateVoicesTable() error {

	query := `CREATE TABLE IF NOT EXISTS voices(
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		name VARCHAR(150) NOT NULL,
		sample VARCHAR(150) NOT NULL,
		accent VARCHAR(150) NOT NULL,
		age VARCHAR(150) NOT NULL,
		gender VARCHAR(150) NOT NULL,
		language VARCHAR(150) NOT NULL,
		language_code VARCHAR(150) NOT NULL,
		loudness VARCHAR(150) NOT NULL,
		style VARCHAR(150) NOT NULL,
		tempo VARCHAR(150) NOT NULL,
		texture VARCHAR(150) NOT NULL,
		is_cloned boolean,
		voice_engine VARCHAR(150) NOT NULL
	)`

	_, err := sql.database.Exec(query)
	if err != nil {
		return err
	}

	return nil

}

func (sql *SQL_DB) CreateVoiceIndex() error {

	query := "CREATE INDEX IF NOT EXISTS idx_voice_name ON voices (name);"

	_, err := sql.database.Exec(query)
	if err != nil{
		return err
	}

	return nil
}


func (sql *SQL_DB) InsertVoice(voice *types.VoiceStruct) error{
	query := `INSERT INTO voices(
		id, name, sample, accent, age, gender, language, language_code, loudness, style, tempo, texture, is_cloned, voice_engine)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`

	_, err := sql.database.Exec(query, 
	                        voice.ID,
							voice.Name,
							voice.Sample,
							voice.Accent,
							voice.Age,
							voice.Gender,
							voice.Language,
							voice.LanguageCode,
							voice.Loudness,
							voice.Style,
							voice.Tempo,
							voice.Texture,
							voice.IsCloned,
							voice.VoiceEngine,
						)

	if err != nil{
		return err
	}	

	return nil
}


func (sql *SQL_DB) Voices() ([]*types.FetchVoiceTypes, error) {
	query := `SELECT id, name, sample, gender, accent, style, language, voice_engine FROM voices  WHERE id LIKE 's3%' limit 40;`

	rows, err := sql.database.Query(query)
	if err != nil{
		return nil, err
	}

	var fetchVoices []*types.FetchVoiceTypes

	for rows.Next(){
		fields, err := voicesScantoRows(rows)
		if err != nil{
			return nil, err
		}
		
		fetchVoices = append(fetchVoices, fields)
	}

	return fetchVoices, nil
}


func voicesScantoRows(rows *sql.Rows) (*types.FetchVoiceTypes, error){
	fields := &types.FetchVoiceTypes{}
	
	err := rows.Scan(
		&fields.ID,
		&fields.Name,
		&fields.Sample,
		&fields.Gender,
		&fields.Accent,
		&fields.Style,
		&fields.Language,
		&fields.VoiceEngine,
	)

	if err != nil{
		return nil, err
	}

	return fields, nil

}