package WebScanner

import (
	"bytes"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

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

func splitNameAndChapter(combineStr string) (num int,name string){

	re := strings.Split(combineStr,"、")
	if len(re)!=2{
		num=-1
		name="no chapter info"
		return
	}else {
		num,err:=strconv.Atoi(re[0])
		if err!=nil{
			num=0
		}
		return num,re[1]
	}

}