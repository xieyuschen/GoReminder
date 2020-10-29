package Reminder

import (
	"GoReminder/Db"
	"GoReminder/EmailSender"
	"GoReminder/WebScanner"
	"GoReminder/models"
)

func Reminder(url string){
	host:="http://www.biquge.se"
	channelSize := 10
	//ch的内容由Scanner获取
	ch :=make(chan  map[int]models.Article,channelSize)
	var lists  map[int]models.Article
	for j:=0;j<1;j++{
		go WebScanner.ArticleUrlAndSubject(url,ch)
		//Block for get
		select{
			case lists = <-ch:
		}
		db_chapter,_:=Db.GetLastChapterAndIsInit(url)
		if _,lastchapter:=WebScanner.GetNewestChapter(url,lists);db_chapter>lastchapter{
			for i:=db_chapter+1;i<=lastchapter;i++{
				println(host+lists[i].Url)
				Content:=WebScanner.GetContentAndSubject(host+lists[i].Url)
				EmailSender.SendEmail("2016231075@qq.com","test",Content)
			}
		}
	}
}
