package models

import "github.com/jinzhu/gorm"

type DbSettings struct{
	Username string `json:"Username"`
	Password string	`json:"Password"`
	Hostname string `json:"Hostname"`
	Dbname   string `json:"Dbname"`
}
type NovelInfo struct {
	gorm.Model
	Url string
	LastChapter int
	IsInit bool
}