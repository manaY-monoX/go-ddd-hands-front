package handler

import (
	"fmt"
	"front-exercise/config"
	"front-exercise/presentation/echo/render"
	"front-exercise/presentation/errortype"
	"front-exercise/service/dto"
	"front-exercise/service/infcs"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

// 商品登録リクエストハンドラ
type registerEnterPageHandler struct {
	service      infcs.CategoryList // 商品カテゴリ取得サービス
	*baseHandler                    // 共通ハンドラをエンベデッドする
}

// 商品入力ページリクエストへの応答
func (r *registerEnterPageHandler) Enter(context echo.Context) error {
	// コンテキストからセッションを取得する
	session := context.Get("session").(*sessions.Session)
	// 商品カテゴリの有無を確認する
	if val, ok := session.Values["category_list"]; ok {
		// 空インターフェイスで返された値をdto.CategoryDTOに変換する
		if category_list, ok := val.([]dto.CategoryDTO); !ok {
			err := errortype.NewPresentationError("商品カテゴリリストのタイプは *dto.CategoryDTO ではありません。")
			r.baseHandler.ErrorHandler(err, context)
			return nil
		} else {
			r.VarMap.Set("category_list", category_list) // VarMapにcategory_listを登録する
		}
	} else {
		response, err := r.service.Execute() // サーバーから商品カテゴリを取得する
		if err != nil {
			r.baseHandler.ErrorHandler(err, context)
			return nil
		}
		// VarMapに商品カテゴリを登録する
		r.VarMap.Set("category_list", response.Categories)
		// Sessionに商品カテゴリを登録する
		session.Values["category_list"] = response.Categories
		// セッションの変更を保存する
		if err := session.Save(context.Request(), context.Response().Writer); err != nil {
			err := errortype.NewPresentationError(fmt.Sprintf("セッションの保存に失敗しました:%v", err))
			r.baseHandler.ErrorHandler(err, context)
			return nil
		}
	}
	// テンプレートの実行と送信
	return r.baseHandler.Templ.Execute(context.Response().Writer, *r.baseHandler.VarMap, nil)
}

// コンストラクタ
func NewregisterEnterPageHandler(config *config.Config, service infcs.CategoryList) *registerEnterPageHandler {
	template := render.NewTemplate(config, "商品登録", "register_enter")
	return &registerEnterPageHandler{
		service:     service,
		baseHandler: NewbaseHandler(template),
	}
}
