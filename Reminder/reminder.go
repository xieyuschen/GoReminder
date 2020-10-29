package Reminder

import (
	"GoReminder/Db"
	"GoReminder/EmailSender"
	"GoReminder/WebScanner"
	"GoReminder/models"
	"fmt"
)

func Reminder(url string){
	host:="http://www.biquge.se"
	channelSize := 10
	//ch的内容由Scanner获取
	ch :=make(chan  map[int]models.Article,channelSize)
	lists := make(map[int]models.Article)
	for j:=0;j<1;j++{
		go WebScanner.ArticleUrlAndSubject(url,ch)
		//Block for get
		select{
			case lists = <-ch:
		}
		db_chapter,_:=Db.GetLastChapterAndIsInit(url)
		_,lastchapter:=WebScanner.GetNewestChapter(url,lists)

		//Just for testing:)
		db_chapter=901

		if db_chapter<lastchapter{
			for i:=db_chapter+1;i<=lastchapter;i++{
				println(host+lists[i].Url)
				Content:=WebScanner.GetContentAndSubject(host+lists[i].Url)
				EmailSender.SendEmail("1743432766@qq.com","test",Content)
			}
		}else {
			fmt.Println("Helloworld")
		}
	}
}
