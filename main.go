package main

import (
	"GoReminder/Reminder"
	"log"
)

func main(){
	url :="http://www.biquge.se/23609/"//77693356.html"
	Reminder.Reminder(url)
}
func Toy(){
	mock()
	log.Println("AfterMath")
}
func mock(){
	defer func() {
		if err:=recover();err!=nil{
			log.Println("Recover Succesfully")
		}
	}()
	log.Panic("Panic!")

}