package main

import (
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const tmplPath = "src/template/"

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

// 一覧表示
func articleIndex(c echo.Context) error {
	// 空のインターフェースは、任意の型の値を保持できる。
	data := map[string]interface{}{
		"Message" : "Hello, World!",
		"Now"     : time.Now(),
	}
	  return render(c, "article/index.html", data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	// 生成された HTML をバイトデータとして受け取る
    b, err := htmlBlob(file, data)
    if err != nil {
      return c.NoContent(http.StatusInternalServerError)
    }
    return c.HTMLBlob(http.StatusOK, b)
}

// テンプレートとデータからHTMLをバイトデータとして生成
func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
    return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}