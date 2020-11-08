package WebScanner

import (
	"GoReminder/models"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
)

func ArticleUrlAndSubject(str string,Articles chan map[int]models.Article) {

	doc,err := getPageNode(str)
	if err!=nil{
		Articles<- nil
	}else {
		r1:=getElementById(doc,"list")

		lists := make( map[int]models.Article)
		var f func(*html.Node,*map[int]models.Article)
		f = func(n *html.Node,list *map[int]models.Article) {
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
				lists[chapter]=models.Article{Chapter: chapter,Name: name,Url: key}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c,list)
			}
		}
		f(r1,&lists)
		Articles<-lists
	}
	return
}
func GetNewestChapter(Url string, lists map[int]models.Article) (url string,lastchapter int){
	for key,value:=range lists{
		if value.Chapter>lastchapter{
			lastchapter=key
		}
	}
	return Url,lastchapter
}
func GetContentAndSubject(url string) (content string){
	node,_ := getPageNode(url)
	t:=getElementById(node,"content")
	text := &bytes.Buffer{}
	collectText(t, text)

	return fmt.Sprintf("%s",text)
}
