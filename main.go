package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// グローバル変数
var e = createMux()

func main() {
  // "/"にGETメソッドでアクセスがあった場合にarticleIndex関数を実行
	e.GET("/", articleIndex)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
  // アプリケーションインスタンスを生成
	e := echo.New()

  // アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}