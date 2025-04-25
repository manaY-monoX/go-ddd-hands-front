package handler

import (
	"fmt"
	"front-exercise/presentation/echo/render"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type baseHandler struct {
	*render.Template // Templateをエンベデッドする
}

// エラーログを出力し、/errorにリダイレクトする
func (b *baseHandler) ErrorHandler(err error, context echo.Context) {
	msg := fmt.Sprintf("エラーログ:%s", err.Error())
	log.Println(msg)
	url, _ := b.Config.BuildAppURL("error")
	context.Redirect(http.StatusTemporaryRedirect, url)
}

// コンストラクタ
func NewbaseHandler(template *render.Template) *baseHandler {
	return &baseHandler{
		Template: template,
	}
}
