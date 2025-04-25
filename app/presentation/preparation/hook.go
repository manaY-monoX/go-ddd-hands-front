package preparation

import (
	"context"
	"fmt"
	"front-exercise/config"
	"net/http"

	"go.uber.org/fx"
)

// fxコンテナのライフサイクル
func RegisterHooks(lifecycle fx.Lifecycle, router *Router, config *config.Config) {
	lifecycle.Append(
		fx.Hook{
			// fxコンテナ起動時の処理
			OnStart: func(ctx context.Context) error {
				go func() {
					port := fmt.Sprintf(":%d", config.App.Port)
					fmt.Printf("サーバをポート%sで起動!!\n", port)

					if err := router.Echo.Start(port); err != nil && err != http.ErrServerClosed {
						fmt.Printf("サーバ起動エラー: %v\n", err)
					}
				}()
				return nil
			},
			// fxコンテナ停止時の処理
			OnStop: func(ctx context.Context) error {
				fmt.Println("サーバを停止!!")
				return router.Echo.Shutdown(ctx)
			},
		},
	)
}
