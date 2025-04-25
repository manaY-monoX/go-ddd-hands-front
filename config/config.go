package config

import (
	"embed"
	"fmt"
	"io/fs"
	"log"

	"gopkg.in/yaml.v3"
)

//go:embed config.yml
var fsys embed.FS

// 　アプリケーション環境情報
type Config struct {
	// 接続サーバ環境情報
	Server struct {
		Hostname string `yaml:"hostname"`
		Port     int    `yaml:"port"`
		Paths    struct {
			CategoryList    string `yaml:"category_list"`
			ProductKeyword  string `yaml:"product_keyword"`
			ProductRegister string `yaml:"product_register"`
		} `yaml:"paths"`
	} `yaml:"server"`
	// アプリケーション環境情報
	App struct {
		Hostname string `yaml:"hostname"`
		Port     int    `yaml:"port"`
		Group    string `yaml:"group"`
		// URLパス情報
		Paths struct {
			Top      string `yaml:"top"`
			Search   string `yaml:"search"`
			Register string `yaml:"register"`
			Error    string `yaml:"error"`
		} `yaml:"paths"`
		// View(HTML)情報
		Views struct {
			Base             string `yaml:"base"`
			Top              string `yaml:"top"`
			Search           string `yaml:"search"`
			RegisterEnter    string `yaml:"register_enter"`
			RegisterComplete string `yaml:"register_complete"`
			Error            string `yaml:"error"`
		} `yaml:"views"`
	} `yaml:"app"`
}

// サーバURLを生成する
func (c *Config) BuildServerURL(pathKey string) (string, error) {
	path, ok := map[string]string{
		"categoryList":    c.Server.Paths.CategoryList,
		"productKeyword":  c.Server.Paths.ProductKeyword,
		"productRegister": c.Server.Paths.ProductRegister,
	}[pathKey]
	if !ok {
		return "", NewConfigError("不正なパスキーが指定されました。")
	}
	return fmt.Sprintf("http://%s:%d%s", c.Server.Hostname, c.Server.Port, path), nil
}

// アプリケーションURLを生成する
func (c *Config) BuildAppURL(pathKey string) (string, error) {
	path, ok := map[string]string{
		"top":      c.App.Paths.Top,
		"search":   c.App.Paths.Search,
		"register": c.App.Paths.Register,
	}[pathKey]
	if !ok {
		return "", NewConfigError("不正なパスキーが指定されました。")
	}
	return fmt.Sprintf("http://%s:%d%s%s", c.App.Hostname, c.App.Port,
		c.App.Group, path), nil
}

// 利用するView(テンプレート)を名を取得する
func (c *Config) GetViewName(nameKey string) (string, error) {
	viewName, ok := map[string]string{
		"top":               c.App.Views.Top,
		"search":            c.App.Views.Search,
		"register_enter":    c.App.Views.RegisterEnter,
		"register_complete": c.App.Views.RegisterComplete,
		"error":             c.App.Views.Error,
	}[nameKey]
	if !ok {
		return "", NewConfigError("不正なパスキーが指定されました。")
	}
	return viewName, nil
}

// コンストラクタ
func NewConfig() *Config {
	var config Config
	// ファイルの読み込み
	data, err := fs.ReadFile(fsys, "config.yml")
	if err != nil {
		log.Fatalf("設定ファイルの読み取りエラー: %v", err)
	}

	// YAMLをパースしてConfig構造体に格納
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("YAML 解析エラー: %v", err)
	}
	return &config
}
