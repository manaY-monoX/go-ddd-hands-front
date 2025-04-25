package tests

import (
	"front-exercise/config"
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConfigSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "環境パラメータの取得")
}

// サービス層のテストに必要な環境
type ConfigTestEnvironment struct {
	app  *fxtest.App
	conf *config.Config
}

// テストに必要な環境を準備する
func SetupTestEnvironment() *ConfigTestEnvironment {
	var env ConfigTestEnvironment
	app := fxtest.New(
		GinkgoT(),
		config.ConfigModule,
		fx.Populate(&env.conf),
	)
	env.app = app
	env.app.RequireStart()
	return &env
}

// fxコンテナを停止する
func TeardownTestEnvironment(env *ConfigTestEnvironment) {
	env.app.RequireStop()
}
