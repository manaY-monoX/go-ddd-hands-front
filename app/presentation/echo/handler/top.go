package handler

import (
	"front-exercise/config"
	"front-exercise/presentation/echo/render"

	"github.com/labstack/echo/v4"
)

// トップ画面ハンドラ
type topPageHandler struct {
	*baseHandler // basehandlerをエンベデッドする
}

// トップ画面を返す
func (t *topPageHandler) Execute(context echo.Context) error {
	return t.baseHandler.Templ.Execute(context.Response().Writer, *t.baseHandler.VarMap, nil)
}

// コンストラクタ
func NewtopPageHandler(config *config.Config) *topPageHandler {
	template := render.NewTemplate(config, "トップページ", "top")
	return &topPageHandler{baseHandler: NewbaseHandler(template)}
}
