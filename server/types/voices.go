package types

type VoiceStruct struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Sample       string `json:"sample" `
	Accent       string `json:"accent"`
	Age          string `json:"age" `
	Gender       string `json:"gender"`
	Language     string `json:"language" `
	LanguageCode string `json:"language_code" `
	Loudness     string `json:"loudness" `
	Style        string `json:"style"`
	Tempo        string `json:"tempo" `
	Texture      string `json:"texture" `
	IsCloned     bool   `json:"is_cloned" `
	VoiceEngine  string `json:"voice_engine"`
}

type FetchVoiceTypes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Sample      string `json:"sample" `
	Gender      string `json:"gender" `
	Accent      string `json:"accent"`
	Language    string `json:"language" `
	Style       string `json:"style"`
	VoiceEngine string `json:"voice_engine"`
}


type VoiceCloneType struct{
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Type string  `json:"type"`
	Voice_Engine string  `json:"voice_engine"`
}