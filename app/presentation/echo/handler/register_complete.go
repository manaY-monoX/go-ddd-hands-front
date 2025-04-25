package handler

import (
	"front-exercise/config"
	"front-exercise/presentation/echo/render"
	"front-exercise/service/dto"
	"front-exercise/service/infcs"

	"github.com/labstack/echo/v4"
)

// 商品登録リクエストハンドラ
type registerCompletePageHandler struct {
	service      infcs.ProductRegister // 商品登録サービス
	*baseHandler                       //共通ハンドラをエンベデッドする
}

// 商品登録リクエストへの応答
func (r *registerCompletePageHandler) Complete(context echo.Context) error {
	// フォームデータの取得
	productName := context.FormValue("productName")
	productPrice := context.FormValue("productPrice")
	category := context.FormValue("category")
	// 商品登録用のDTOインスタンスを生成する
	newProduct := dto.ProductDTO{
		Id:    "",
		Name:  productName,
		Price: productPrice,
		Category: dto.CategoryDTO{
			Id:   category,
			Name: "dummy",
		},
	}
	// 登録サービスを実行する
	response, err := r.service.Execute(&newProduct)
	if err != nil {
		r.baseHandler.ErrorHandler(err, context)
	}
	r.VarMap.Set("response", response) // 実行結果をVarMapに登録する
	// テンプレートの実行と送信
	return r.Templ.Execute(context.Response().Writer, *r.VarMap, nil)
}

// コンストラクタ
func NewregisterCompletePageHandler(config *config.Config,
	service infcs.ProductRegister) *registerCompletePageHandler {
	template := render.NewTemplate(config, "商品登録", "register_complete")
	return &registerCompletePageHandler{
		service:     service,
		baseHandler: NewbaseHandler(template),
	}
}
