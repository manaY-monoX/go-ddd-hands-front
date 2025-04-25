package presentation

import (
	"front-exercise/presentation/echo/handler"
	"front-exercise/presentation/echo/render"
	"front-exercise/presentation/preparation"
	"front-exercise/presentation/session"
	"front-exercise/service"
	"log"

	"go.uber.org/fx"
)

// プレゼンテーション層の依存定義
var PresentationModule = fx.Options(
	service.ServiceModeul,
	fx.Provide(
		session.NewSessionManager,              // セッション管理機能のインスタンスを生成する
		render.NewTemplate,                     // Jet Temaplete Engineのインスタンスを生成する
		handler.NewpathChecker,                 // URLパスチェックミドルウェアのインスタンスを生成する
		handler.NewtopPageHandler,              // トップページハンドラのインスタンスを生成する
		handler.NewsearchPageHandler,           // 商品キーワード検索ハンドラのインスタンスを生成する
		handler.NewregisterEnterPageHandler,    // 商品登録ハンドラのインスタンスを生成する
		handler.NewregisterCompletePageHandler, // 商品登録ハンドラのインスタンスを生成する
		handler.NewerrorHandler,                // エラーハンドラのインスタンス
		handler.ProvideHandlers,                // リクエストハンドラをグループ化する
		preparation.NewRouter,                  // Echoを生成し、リクエストハンドラを設定する
	),
	fx.Invoke(preparation.RegisterHooks),
	fx.Invoke(setupEnd),
)

func setupEnd() {
	log.Println("プレゼンテーション層の構築が完了しました。")
}
