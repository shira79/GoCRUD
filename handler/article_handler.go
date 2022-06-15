package handler

import (
	"golang-blog/model"
	"golang-blog/repository"

	"net/http"

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

// 保存処理のレスポンスデータの構造体
type ArticleCreateOutput struct {
	Article *model.Article
	Message string
	ValidationErrors []string
}

// 保存処理
func ArticleCreate(c echo.Context) error {
	var article model.Article
	var out ArticleCreateOutput

	// フォームの内容を構造体に埋め込みます。
	if err := c.Bind(&article); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusBadRequest, out)
	}

	// バリデーション
	if err := c.Validate(&article); err != nil {
		c.Logger().Error(err)

		// エラーセットをレスポンスにセットして、返却
		out.ValidationErrors = article.ValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, out)
	}

	// 保存処理を実行
	if err := repository.ArticleCreate(&article); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, out)
	}

	// レスポンスに保存した記事のデータし返却
	out.Article = &article
	return c.JSON(http.StatusOK, out)
}