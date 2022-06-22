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
	"github.com/go-playground/validator"
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

    // TOP
	e.GET("/", handler.ArticleIndex)

	// HTML返却
	e.GET("/articles", handler.ArticleIndex)
	e.GET("/articles/new", handler.ArticleNew)
	e.GET("/articles/:articleID", handler.ArticleShow)
	e.GET("/articles/:articleID/edit", handler.ArticleEdit)

	// JSON返却
	e.GET("/api/articles", handler.ArticleList)
	e.POST("/api/articles", handler.ArticleCreate)
	e.DELETE("/api/articles/:articleID", handler.ArticleDelete)
	e.PATCH("/api/articles/:articleID", handler.ArticleUpdate)

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
	e.Static("/css", "src/css")

	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}

// Validator用の構造体
type CustomValidator struct {
	validator *validator.Validate
}

// Validate関数を実装することで、Validatorインターフェイスを実装。
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
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
