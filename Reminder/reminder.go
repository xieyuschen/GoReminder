package Reminder

import (
	"GoReminder/Db"
	"GoReminder/EmailSender"
	"GoReminder/models"
)

func Reminder(){
	var ch chan models.NovelInfo
	for{
		chapter:=<-ch
		db_chapter,_:=Db.GetLastChapterAndIsInit(chapter.Url)
		if db_chapter>chapter.LastChapter{
			EmailSender.SendEmail("1","2","3")
		}
	}
}
