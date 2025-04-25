package service

import (
	"front-exercise/config"
	"front-exercise/service/impls"
	"log"

	"go.uber.org/fx"
)

// サービス層の依存性定義
var ServiceModeul = fx.Options(
	config.ConfigModule,
	fx.Provide(
		impls.NewcategoryListImpl,    // 商品カテゴリ取得サービスのインスタンス生成する
		impls.NewproductKeywordImpl,  // 商品キーワード検索サービスのインスタンス生成する
		impls.NewproductRegisterImpl, // 商品登録サービスのインスタンスを生成する
	),
	fx.Invoke(setupEnd),
)

func setupEnd() {
	log.Println("サービス層の構築が完了しました。")
}
