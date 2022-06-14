package handler

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

const tmplPath = "src/template/"

// テンプレートとデータからHTMLをバイトデータとして生成
func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

// 描画
func render(c echo.Context, file string, data map[string]interface{}) error {
	// CSRF対策用のトークンを
	data["CSRF"] = c.Get("csrf").(string)

	b, err := htmlBlob(file, data)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.HTMLBlob(http.StatusOK, b)
}