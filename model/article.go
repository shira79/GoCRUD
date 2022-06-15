package model

import (
    "time"

	"github.com/go-playground/validator"
)

type Article struct {
	ID      int       `db:"id" form:"id"`
	Title   string    `db:"title" form:"title" validate:"required,max=50"`
	Body    string    `db:"body" form:"body" validate:"required"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
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