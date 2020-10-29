package Db

import (
	"GoReminder/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)
var db *gorm.DB
func init(){
	settings := models.DbSettings{Username: "root", Password: "root", Hostname: "127.0.0.1:3306", Dbname: "goreminder"}
	//"root:@tcp(127.0.0.1:3306)/?parseTime=true&charset=utf8"
	connStr := dsn(settings)
	msdb, err := sql.Open("mysql",connStr)
	if err!=nil{
		log.Panic(err)
	}
	msdb.Exec("create database if not exists "+settings.Dbname +" character set utf8")
	msdb.Close()

	db, _ = gorm.Open("mysql",dsn(settings))
	var novelInfo models.NovelInfo
	if !db.HasTable(&novelInfo){
		db.CreateTable(&novelInfo)
	}
}
func GetLastChapterAndIsInit(url string)(LastChapter int,IsInit bool){
	var info models.NovelInfo
	db.Where("url=?",url).Find(&info)
	return info.LastChapter,info.IsInit
}
func InsertArticle(info models.NovelInfo){
	db.Create(&info)
}
func dsn(settings models.DbSettings) string {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	// Add ?parseTime=true
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8", settings.Username,settings.Password, settings.Hostname,settings.Dbname)
}
func UpdateLastestChapter(url string,lastestChapter int){
	var info models.NovelInfo
	db.Where("url=?",url).Find(&info)
	info.LastChapter=lastestChapter
	db.Save(&info)
	return
}