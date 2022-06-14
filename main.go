package main

import (
	"fmt"
	"time"
	"os"

	"golang-blog/handler"
	"golang-blog/repository"

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


func main() {

	// .env取得
	err := godotenv.Load()
	if err != nil {
		panic(".envが見つからないよ")
	}

	// DB接続
	db = sqlConnect()
	repository.SetDB(db)
	defer db.Close()

    // "/"にGETメソッドでアクセスがあった場合にarticleIndex関数を実行
	e.GET("/"         , handler.ArticleIndex)
	e.GET("/new"      , handler.ArticleNew)
	e.GET("/:id"      , handler.ArticleShow)
	e.GET("/:id/edit" , handler.ArticleEdit)
	e.POST("/", handler.ArticleCreate)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
  // アプリケーションインスタンスを生成
	e := echo.New()

  // アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())

	e.Static("/js", "src/js")

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
