package Reminder

import (
	"GoReminder/Db"
	"GoReminder/EmailSender"
	"GoReminder/WebScanner"
	"GoReminder/models"
	"fmt"
	"time"
)
var maps map[string]models.ArticleDetail

func Reminder(url string){

	maps[url] = models.ArticleDetail{IsInit: true}
	host:="http://www.biquge.se"
	channelSize := 10
	//ch的内容由Scanner获取
	ch :=make(chan  map[int]models.Article,channelSize)
	lists := make(map[int]models.Article)

	WebScanner.ArticleUrlAndSubject(url, ch)

	go WebScanner.ArticleUrlAndSubject(url, ch)
	//Block for get
	select {
		case lists = <-ch:
	}
	_, lastchapter := WebScanner.GetNewestChapter(url, lists)

	fmt.Println(lastchapter)
	info := models.NovelInfo{Url: url, LastChapter: lastchapter - 1, IsInit: true}
	Db.InsertArticle(info)
	fmt.Println("Restart service and Send the newest chapter to you:)")

	for i:=0;;i++{
		go WebScanner.ArticleUrlAndSubject(url, ch)
		//Block for get
		select {
			case lists = <-ch:
		}
		_, lastchapter := WebScanner.GetNewestChapter(url, lists)
		println("This is ",i,"turn at ",time.Now().String())
		if db_chapter, _ := Db.GetLastChapterAndIsInit(url); db_chapter < lastchapter {
			for i := db_chapter + 1; i <= lastchapter; i++ {
				Db.UpdateLastestChapter(url, i)
				Content := WebScanner.GetContentAndSubject(host + lists[i].Url)
				EmailSender.SendEmail("1743432766@qq.com", lists[i].Name, Content)

			}
		} else {
			time.Sleep(20 * time.Second)
		}
	}

}
