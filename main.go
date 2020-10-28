package main

import (
	"GoReminder/WebScanner"
)

func main(){
	str := "http://www.biquge.se/23609/"
	body:=WebScanner.GetPage(str)
	WebScanner.HtmlParse(body)
}