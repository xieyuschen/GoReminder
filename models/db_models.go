package models
type DbSettings struct{
	Username string `json:"Username"`
	Password string	`json:"Password"`
	Hostname string `json:"Hostname"`
	Dbname   string `json:"Dbname"`
}
type NovelInfo struct {
	Url string
	LastChapter int
	IsInit bool
}