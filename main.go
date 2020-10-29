package main

import (
	"GoReminder/WebScanner"
	"fmt"
)

func main(){
	str :="http://www.biquge.se/23609/77693356.html"
	Content:=WebScanner.GetContentAndSubject(str)
	fmt.Println(Content)
}