package handler

import (
	"golang-blog/model"
	"golang-blog/repository"

	"net/http"
	"fmt"
	"strconv"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/jinzhu/gorm"
)

// 一覧表示
func ArticleIndex(c echo.Context) error {

	// "/articles" のパスでリクエストがあったら "/" にリダイレクト
    if c.Request().URL.Path == "/articles" {
    	c.Redirect(http.StatusPermanentRedirect, "/")
    }

	var articles []model.Article
	initialCursor := 0
	// 保存処理を実行
	if err := repository.ArticleListByCursor(&articles, initialCursor); err != nil {
		c.Logger().Error(err.Error())
		// 500を返却
		return render(c, "error/500.html", map[string]interface{}{})
	}

	// 取得した最後の記事のIDをカーソルにする
    var cursor int
    if len(articles) != 0 {
    	cursor = articles[len(articles)-1].ID
    }

	data := map[string]interface{}{
		"Articles": articles,
		"Cursor":   cursor,
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

	var article model.Article
	id, _ := strconv.Atoi(c.Param("articleID"))

	if err := repository.ArticleByID(&article, id); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound){
			c.Logger().Error(err.Error())
			return render(c, "error/404.html", map[string]interface{}{})

		} else {
			c.Logger().Error(err.Error())
			return render(c, "error/500.html", map[string]interface{}{})
		}
	}

	data := map[string]interface{}{
		"Article": article,
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

// 削除処理
func ArticleDelete(c echo.Context) error {
	// パラメータからID を取得し、数値型にキャスト
	id, _ := strconv.Atoi(c.Param("articleID"))

	// 削除処理
	if err := repository.ArticleDelete(id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Article %d is deleted.", id))
}

// 記事データをjsonで取得
func ArticleList(c echo.Context) error {
	// cursorパラメーターをintにキャスト
	cursor, _ := strconv.Atoi(c.QueryParam("cursor"))

	var articles []model.Article
	// 取得実行
	if err := repository.ArticleListByCursor(&articles, cursor); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articles)
}