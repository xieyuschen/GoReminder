package main

import (
	"fmt"
	"log"
)

func DoPanic(){
	log.Panic("Just For try")
}
func Recover(){
	fmt.Println("Recover plays its role")
	recover()
}
func try()  {
	for{
		defer Recover()
		DoPanic()

	}
}
func main(){
	//url :="http://www.biquge.se/23609/"//77693356.html"
	//Reminder.Reminder(url)
	try()
	fmt.Println("Helloworld")
	try()
}