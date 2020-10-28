package main

import (
	"GoReminder/WebScanner"
)

func main(){
	str := "http://www.biquge.se/23609/"
	body:=WebScanner.GetPage(str)
	m:=WebScanner.HtmlParse(body)

	for key,val:=range m{
		print(key)
		println(val.Name,val.Chapter)
	}
}