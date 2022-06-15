package repository

import (
	"time"
	"math"

	"github.com/pkg/errors"

	"golang-blog/model"
)

// 記事取得
func ArticleListByCursor(articles *[]model.Article, cursor int) error {

	// 0以下の場合、intの最大値
	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	const limit int = 10

	if err := db.Where("id < ?", cursor).Limit(limit).Order("id desc").Find(&articles).Error; err != nil {
		return err
	}
	return nil
}

// 記事作成
func ArticleCreate(article *model.Article) error {
	now := time.Now()

	article.Created = now
	article.Updated = now

	if result := db.Create(&article); result.Error != nil {
		return result.Error
	}

	return nil
}

// 記事削除
func ArticleDelete(id int) error {

	result := db.Delete(&model.Article{}, id)
	if  result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.Errorf("id=%w のTodoデータが存在しません。", id)
	}

	return nil
}