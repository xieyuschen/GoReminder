package WebScanner

import (
	"GoReminder/models"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetPageNode(url string) (node *html.Node){

	resp,err:=http.Get(url)
	if err!=nil{
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer  resp.Body.Close()
	doc,err := html.Parse(strings.NewReader(string(body)))
	if err!=nil{
		log.Panic(err)
	}
	return doc
}
func ArticleUrlAndSubject(str string) (lists map[string]models.Article) {
	lists =make(map[string]models.Article)
	doc,err := html.Parse(strings.NewReader(str))
	if err!=nil{
		log.Panic(err)
	}
	r1:=getElementById(doc,"list")
	//fmt.Println(r1)

	var f func(*html.Node)
	f = func(n *html.Node) {
		var key string
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key=="href" {
					key = a.Val
					break
				}
			}
			//With the help of `https://stackoverflow.com/questions/18274501/how-can-i-get-the-content-of-an-html-node`
			//Powerful Internet Explorer!
			text := &bytes.Buffer{}
			collectText(n, text)

			chapter,name:=splitNameAndChapter(fmt.Sprintf("%s",text))
			lists[key]=models.Article{Chapter: chapter,Name: name}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(r1)

	return lists
}
func GetNewestChapter(url string, lists map[string]models.Article) (u string,lastchapter int){
	for _,value:=range lists{
		if value.Chapter>lastchapter{
			lastchapter=value.Chapter
		}
	}
	return url,lastchapter
}
func GetContentAndSubject(url string) (content string){
	node := GetPageNode(url)
	t:=getElementById(node,"content")
	text := &bytes.Buffer{}
	collectText(t, text)

	return fmt.Sprintf("%s",text)
}
