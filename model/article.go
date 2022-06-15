package model

import (
    "time"

	"github.com/go-playground/validator"
)

type Article struct {
	ID      int       `db:"id" form:"id" json:"id"`
	Title   string    `db:"title" form:"title" validate:"required,max=50" json:"title"`
	Body    string    `db:"body" form:"body" validate:"required" json:"body"`
	Created time.Time `db:"created"  json:"created"`
	Updated time.Time `db:"updated"  json:"updated"`
}

// バリデーションのメッセージを整形する
func (a *Article) ValidationErrors(err error) []string {
	// メッセージを格納するスライス
	var errMessages []string

	// errをvalidator.ValidationErrorsにアサーション
	for _, err := range err.(validator.ValidationErrors) {

		var message string

		// エラーになったフィールドで分岐
		switch err.Field() {
			case "Title":
				// エラーになったルールで分岐
				switch err.Tag() {
					case "required":
						message = "タイトルは必須です。"
					case "max":
						message = "タイトルは最大50文字です。"
				}
			case "Body":
				switch err.Tag() {
					case "required":
						message = "本文は必須です。"
			}
		}

		// スライスにメッセージを格納
		if message != "" {
			errMessages = append(errMessages, message)
		}
	}

	return errMessages
}