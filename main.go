package main

import (
	"log"
)

func main(){
	//url :="http://www.biquge.se/23609/"//77693356.html"
	//Reminder.Reminder(url)
	toy()
}
func toy(){
	Mock()
	log.Println("AfterMath")
}
func Mock(){
	defer func() {
		if err:=recover();err==nil{
			log.Panic("Panic Again",err)
		}
		log.Println("Recover Succesfully")
	}()
	log.Panic("Panic!")

}