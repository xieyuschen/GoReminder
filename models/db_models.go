package models

type DbSettings struct{
	Username string `json:"Username"`
	Password string	`json:"Password"`
	Hostname string `json:"Hostname"`
	Dbname   string `json:"Dbname"`
}
type NovelInfo struct {
	Url string  `gorm:"primary_key"`
	LastChapter int
	IsInit bool
}