package Db

import (
	"GoReminder/models"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
)
var db *gorm.DB

func databaseInit(){
	settings := models.DbSettings{Username: "root", Password: "root", Hostname: "127.0.0.1:3306", Dbname: "husthole"}
	//"root:@tcp(127.0.0.1:3306)/?parseTime=true&charset=utf8"
	connStr := dsn(settings)
	msdb, _ := sql.Open("mysql",connStr)

	msdb.Exec("create database if not exists "+settings.Dbname +" character set utf8")
	msdb.Close()

	db, _ = gorm.Open("mysql",dsn(settings))
	//db.DB().SetMaxIdleConns(0)

}
func dsn(settings models.DbSettings) string {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	// Add ?parseTime=true
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8", settings.Username,settings.Password, settings.Hostname,settings.Dbname)
}
