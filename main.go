package main

import (
	"context"
	"github.com/scSZn/blog/internal/dao"
	"log"
)

func main() {
	Init()
	//err := dao.NewArticleDAO().CreateArticle(context.Background(), &model.Article{
	//	ArticleID: "2047483647",
	//})
	//log.Println(err)
	article, err := dao.NewArticleDAO().GetArticleByArticleID(context.Background(), "2047483647")
	log.Println(err)
	log.Println(article)
}
