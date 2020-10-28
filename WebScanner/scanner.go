package WebScanner

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"golang.org/x/net/html"
	"strings"
)

func GetPage(url string) (pageContent string){

	resp,err:=http.Get(url)
	if err!=nil{
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer  resp.Body.Close()
	return string(body)
}
func HtmlParse(str string) {
	doc,err := html.Parse(strings.NewReader(str))
	if err!=nil{
		log.Panic(err)
	}
	r1:=getElementById(doc,"list")
	//fmt.Println(r1)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			fmt.Println()
			for _, a := range n.Attr {
				if a.Key=="href" {
					fmt.Println(a.Val)
					break
				}
			}
			//With the help of `https://stackoverflow.com/questions/18274501/how-can-i-get-the-content-of-an-html-node`
			//Powerful Internet Explorer!
			text := &bytes.Buffer{}
			collectText(n, text)
			fmt.Println(text)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(r1)
}
func GetNewestChapter(url string, lists chan map[string]int){

}

