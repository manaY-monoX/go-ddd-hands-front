package tests

import (
	"front-exercise/config"
	"front-exercise/service"
	"front-exercise/service/infcs"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestServiceSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "サービス層のテストスイート")
}

// サービス層のテストに必要な環境
type ServiceTestEnvironment struct {
	app      *fxtest.App
	conf     *config.Config
	list     infcs.CategoryList
	keyword  infcs.ProductKeyword
	register infcs.ProductRegister
}

// テストに必要な環境を準備する
func SetupTestEnvironment() *ServiceTestEnvironment {
	var env ServiceTestEnvironment
	app := fxtest.New(
		GinkgoT(),
		service.ServiceModeul,
		fx.Populate(&env.conf),
		fx.Populate(&env.list),
		fx.Populate(&env.keyword),
		fx.Populate(&env.register),
	)
	env.app = app
	env.app.RequireStart()
	return &env
}

// fxコンテナを停止する
func TeardownTestEnvironment(env *ServiceTestEnvironment) {
	env.app.RequireStop()
}
