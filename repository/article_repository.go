package repository

import (
	"golang-blog/model"
)

// 記事全件取得
func ArticleList() ([]*model.Article) {
	var article_list []*model.Article
    db.Find(&article_list)

	return article_list
}