package main

import (
	"fmt"
	"time"
	"net/http"
	"os"

	"github.com/flosch/pongo2"
	"github.com/jinzhu/gorm"
  	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/joho/godotenv"
)

const tmplPath = "src/template/"

// グローバル変数
var e = createMux()
var db *gorm.DB

type Article struct {
	Id int
    Title string
  }

func main() {

	// .env取得
	err := godotenv.Load()
	if err != nil {
		panic(".envが見つからないよ")
	}

	db = sqlConnect()
	defer db.Close()

  // "/"にGETメソッドでアクセスがあった場合にarticleIndex関数を実行
	e.GET("/"         , articleIndex)
	e.GET("/new"      , articleNew)
	e.GET("/:id"      , articleShow)
	e.GET("/:id/edit" , articleEdit)

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

// DBとのコネクション確立
func sqlConnect() (database *gorm.DB) {

	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := os.Getenv("DB_PORT")
	NAME := os.Getenv("DB_NAME")
	CONNECT := USER + ":" + PASS + "@tcp(db:" + PROTOCOL + ")/" + NAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
	  for {
		if err == nil {
		  fmt.Println("")
		  break
		}
		fmt.Print(".")
		time.Sleep(time.Second)
		count++
		if count > 1 {
		  fmt.Println("")
		  fmt.Println("DB接続失敗")
		  panic(err)
		}
		db, err = gorm.Open(DBMS, CONNECT)
	  }
	}
	fmt.Println("DB接続成功")

	db.LogMode(true)

	return db
  }

// 一覧表示
func articleIndex(c echo.Context) error {
	var article_list []Article
    db.Find(&article_list)

	// 空のインターフェースは、任意の型の値を保持できる。
	data := map[string]interface{}{
		"article_list" : article_list,
	}
	return render(c, "article/index.html", data)
}

// 新規作成
func articleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message" : "new",
	}
	return render(c, "article/new.html", data)
}

// 詳細
func articleShow(c echo.Context) error {
	data := map[string]interface{}{
		"Message" : "show",
	}
	return render(c, "article/show.html", data)
}

// 編集
func articleEdit(c echo.Context) error {
	data := map[string]interface{}{
		"Message" : "edit",
	}
	return render(c, "article/edit.html", data)
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