package Db

import (
	"GoReminder/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var db *gorm.DB

func GetLastChapterAndIsInit(url string)(LastChapter int,IsInit bool){
	var info models.NovelInfo
	db.Where("url=?",url).Find(&info)
	return info.LastChapter,info.IsInit
}
func InsertArticle(info models.NovelInfo){
	var temp models.NovelInfo
	db.Where("url = ?",info.Url).Find(&temp)
	if temp.Url==""{
		db.Create(&info)
	}else {
		temp.LastChapter = info.LastChapter
		db.Save(&temp)
	}
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
