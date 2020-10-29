package Reminder

import (
	"GoReminder/Db"
	"GoReminder/EmailSender"
	"GoReminder/models"
)

func Reminder(){
	//channelSize := 10
	//ch的内容由Scanner获取
	//ch :=make(map[string]models.NovelInfo,channelSize)
	var chapter models.NovelInfo
	for{
		//Block for get
		select{

		}
		db_chapter,_:=Db.GetLastChapterAndIsInit(chapter.Url)
		if db_chapter>chapter.LastChapter{
			EmailSender.SendEmail("1","2","3")
		}
	}
}
