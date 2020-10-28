package models
type Config struct {
	DbSettings DbSettings `json:"DbSettings"`
	EmailSenderSettings EmailSenderSettings `json:"EmailSenderSettings"`
}
