package render

import (
	"front-exercise/config"
	"log"
	"strconv"

	"github.com/CloudyKit/jet/v6"
	"github.com/leekchan/accounting"
)

// テンプレートを保持する構造体
type Template struct {
	Jet    *jet.Set       // Jetテンプレートエンジン
	VarMap *jet.VarMap    // 値マップ
	Templ  *jet.Template  // テンプレート
	Config *config.Config // アプリケーション環境情報
}

// 指定されたテンプレートを取得する
func (t *Template) GetTemplate(name string) *jet.Template {
	// テンプレート名を取得する
	viewName, err := t.Config.GetViewName(name)
	if err != nil {
		log.Fatalf("GetTemplate: %v", err)
	}
	// テンプレートを取得する
	tmpl, err := t.Jet.GetTemplate(viewName)
	if err != nil {
		log.Fatalf("GetTemplate: %v", err)
	}
	return tmpl
}

// コンストラクタ
func NewTemplate(config *config.Config, title string, viewName string) *Template {

	jetSet := jet.NewSet(
		jet.NewOSFileSystemLoader(config.App.Views.Base), // テンプレートのディレクトリを指定
		jet.InDevelopmentMode(),                          // 開発モード（テンプレートの自動リロードを有効化）
	)

	// 通貨フォーマット変換機能を生成する
	ac := accounting.Accounting{Symbol: "￥", Precision: 0, Thousand: ",", Decimal: "."}
	// 通貨フォーマット変換機能を登録する
	jetSet.AddGlobal("formatCurrency", func(priceStr string) string {
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			log.Fatalf("通貨フォーマット変換に失敗しました。 値: %s", priceStr)
		}
		return ac.FormatMoney(price)
	})

	// Templateを生成する
	template := Template{Jet: jetSet, VarMap: nil, Templ: nil, Config: config}
	// VarMapを生成し、タイトルを登録する
	v := make(jet.VarMap)
	v.Set("title", title)
	template.VarMap = &v
	// jet.Templateを生成して、Templフィールドに設定する
	template.Templ = template.GetTemplate(viewName)

	return &template
}
