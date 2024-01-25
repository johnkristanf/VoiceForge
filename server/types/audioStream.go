package types


type StreamBody struct {
	Text          string  `json:"text"`
	Voice         string  `json:"voice"`
	Output_format string  `json:"output_format"`
	Speed         float32 `json:"speed"`
}

type AudioStruct struct {
	ID int64  `json:"audio_id"`
	AudioStream []byte `json:"audioStream"`
	AudioText   string `json:"audioText"`
}