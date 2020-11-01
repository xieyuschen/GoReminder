package Db

import (
	"GoReminder/EmailSender"
	"GoReminder/models"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var DbSetting models.DbSettings
func init(){
	conf := EmailSender.ReadSettingsFromFile("Config.json")
	fmt.Println(conf)
	DbSetting = conf.DbSettings
	//"root:@tcp(127.0.0.1:3306)/?parseTime=true&charset=utf8"
	connStr := dsn(DbSetting)
	msdb, err := sql.Open("mysql",connStr)
	if err!=nil{
		log.Panic(err)
	}
	msdb.Exec("create database if not exists "+DbSetting.Dbname +" character set utf8")
	msdb.Close()

	db, _ = gorm.Open("mysql",dsn(DbSetting))
	var novelInfo models.NovelInfo
	if !db.HasTable(&novelInfo){
		db.CreateTable(&novelInfo)
	}
}
