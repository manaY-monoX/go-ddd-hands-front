package preparation

import (
	"front-exercise/config"
	"front-exercise/presentation/echo/handler"
	"front-exercise/presentation/session"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// ルータ構造体
type Router struct {
	Echo *echo.Echo // Echoフレームワーク
}

func NewRouter(handlers handler.Handlers, config *config.Config, session *session.SessionManager) *Router {
	e := echo.New()             // Echoの生成
	e.Debug = false             // リリースモード
	e.Logger.SetLevel(log.INFO) // ログレベルInformetion

	// ロガーとリカバリーのミドルウェアを追加
	e.Use(middleware.Logger())       // ログ出力のミドルウェアを追加
	e.Use(middleware.Recover())      // パニックからのリカバリーミドルウェアを追加
	e.Use(session.SessionMiddleware) // セッション管理ミドルウェアを追加
	// リダイレクトミドルウェアをグローバルミドルウェアとして追加
	e.Use(handlers.PathChecker.RedirectMiddleware)

	// コンテキスト名「/exercise」を持つルートグループを作成
	group := e.Group(config.App.Group)

	// リクエストハンドラを設定する
	group.GET(config.App.Paths.Top, handlers.Top.Execute)                     // トップページを表示する
	group.GET(config.App.Paths.Search, handlers.Search.Enter)                 // 商品キーワード検索ページを表示する
	group.GET(config.App.Paths.Register, handlers.RegisterEntry.Enter)        // 商品を登録する
	group.POST(config.App.Paths.Search, handlers.Search.Search)               // 商品キーワード検索する
	group.POST(config.App.Paths.Register, handlers.RegisterComplete.Complete) // 商品を登録する
	// エラーページへの遷移
	group.GET(config.App.Paths.Error, handlers.Err.Execute)
	group.POST(config.App.Paths.Error, handlers.Err.Execute)

	return &Router{
		Echo: e,
	}
}
