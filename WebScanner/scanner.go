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
func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}

//========================================
func GetAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func checkId(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		s, ok := GetAttribute(n, "id")
		if ok && s == id {
			return true
		}
	}
	return false
}

func traverse(n *html.Node, id string) *html.Node {
	if checkId(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := traverse(c, id)
		if result != nil {
			return result
		}
	}

	return nil
}

func getElementById(n *html.Node, id string) *html.Node {
	return traverse(n, id)
}
