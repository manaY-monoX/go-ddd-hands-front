package handler

import (
	"front-exercise/config"
	"front-exercise/presentation/echo/render"

	"github.com/labstack/echo/v4"
)

type errorHandler struct {
	*baseHandler // basehandlerをエンベデッドする
}

// エラー画面を返す
func (t *errorHandler) Execute(context echo.Context) error {
	return t.Templ.Execute(context.Response().Writer, *t.baseHandler.VarMap, nil)
}

// コンストラクタ
func NewerrorHandler(config *config.Config) *errorHandler {
	template := render.NewTemplate(config, "エラー", "error")
	return &errorHandler{
		baseHandler: NewbaseHandler(template),
	}
}
