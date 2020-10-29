package Db

import (
	"GoReminder/EmailSender"
	"GoReminder/models"
)

var DbSetting models.DbSettings
func init(){
	conf := EmailSender.ReadSettingsFromFile("Config.json")
	DbSetting = conf.DbSettings
}
