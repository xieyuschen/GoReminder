package WebScanner

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetPage(url string) (pageContent string){
	resp,err:=http.Get(url)
	if err!=nil{
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}
func HtmlParse(){

}
func GetNewestChapter(url string, lists chan map[string]int){

}
