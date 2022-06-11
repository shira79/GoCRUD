package handler

import (
	"golang-blog/repository"
	"github.com/labstack/echo/v4"
)

// 一覧表示
func ArticleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"article_list": repository.ArticleList(),
	}
	return render(c, "article/index.html", data)
}

// 新規作成
func ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
	}
	return render(c, "article/new.html", data)
}

// 詳細
func ArticleShow(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article Show",
	}
	return render(c, "article/show.html", data)
}

// 編集
func ArticleEdit(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article Edit",
	}
	return render(c, "article/edit.html", data)
}