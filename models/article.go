package models

type Article struct {
	Content string
	Name string
	Chapter int
	Url string
}

type ArticleDetail struct{
	LastestChapter string
	IsInit bool
}