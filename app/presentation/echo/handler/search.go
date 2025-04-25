package handler

import (
	"front-exercise/config"
	"front-exercise/presentation/echo/render"
	"front-exercise/service/infcs"

	"github.com/labstack/echo/v4"
)

// 従業員キーワード検索リクエストハンドラ
type searchPageHandler struct {
	service      infcs.ProductKeyword // 商品キーワード検索サービス
	*baseHandler                      // 共通ハンドラをエンベデッドする
}

// 検索ページリクエストへの応答
func (s *searchPageHandler) Enter(context echo.Context) error {
	// テンプレートの実行と送信
	return s.baseHandler.Templ.Execute(context.Response().Writer, *s.baseHandler.VarMap, nil)
}

// 検索リクエストへの応答
func (s *searchPageHandler) Search(context echo.Context) error {
	keyword := context.FormValue("keyword") // Formからキーワードを取得
	resp, err := s.service.Execute(keyword) // 検索サービスを実行する
	if err != nil {
		s.baseHandler.ErrorHandler(err, context)
		return nil
	}
	s.VarMap.Set("response", resp) // サービス実行結果をVarMapに格納する
	// テンプレートの実行と送信
	return s.baseHandler.Templ.Execute(context.Response().Writer, *s.baseHandler.VarMap, nil)
}

// コンストラクタ
func NewsearchPageHandler(config *config.Config, service infcs.ProductKeyword) *searchPageHandler {
	template := render.NewTemplate(config, "商品検索", "search")

	return &searchPageHandler{
		service:     service,
		baseHandler: NewbaseHandler(template),
	}
}
