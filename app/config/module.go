package config

import "go.uber.org/fx"

// アプリケーション環境の依存定義
var ConfigModule = fx.Options(
	fx.Provide(
		NewConfig, // アプリケーション環境情報の生成
	),
)
