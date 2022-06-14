package repository

import (
	  "golang-blog/model"
    "time"
)

// 記事全件取得
func ArticleList() ([]*model.Article) {
	var article_list []*model.Article
    db.Find(&article_list)

	return article_list
}

// 記事作成
func ArticleCreate(article *model.Article) (error){
	now := time.Now()

	article.Created = now
	article.Updated = now

	if result := db.Create(&article); result.Error != nil {
		return result.Error
	}

	return nil
}