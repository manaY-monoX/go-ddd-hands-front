package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("動作情報取得の品質検証\n", Ordered, func() {
	var env *ConfigTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})
	It("動作情報の内容を確認する", func() {
		Expect(env.conf).ToNot(BeNil())
		Expect(env.conf.Server).ToNot(BeNil())
		Expect(env.conf.App).ToNot(BeNil())
		Expect(env.conf.Server.Hostname).To(Equal("service_exercise"))
		Expect(env.conf.Server.Port).To(Equal(8085))
		Expect(env.conf.Server.Paths.DepartmentList).To(Equal("/category/list"))
		Expect(env.conf.Server.Paths.EmployeeKeyword).To(Equal("/product/keyword/%s"))
		Expect(env.conf.Server.Paths.EmployeeRegister).To(Equal("/product/register"))
		Expect(env.conf.App.Hostname).To(Equal("localhost"))
		Expect(env.conf.App.Port).To(Equal(8080))
		Expect(env.conf.App.Group).To(Equal("/exercise"))
		Expect(env.conf.App.Paths.Top).To(Equal("/top"))
		Expect(env.conf.App.Paths.Search).To(Equal("/search"))
		Expect(env.conf.App.Paths.Register).To(Equal("/register"))
		Expect(env.conf.App.Views.Base).To(Equal("./views"))
		Expect(env.conf.App.Views.Top).To(Equal("top/top.html"))
		Expect(env.conf.App.Views.Search).To(Equal("search/search.html"))
		Expect(env.conf.App.Views.RegisterEnter).To(Equal("register/enter.html"))
		Expect(env.conf.App.Views.RegisterComplete).To(Equal("register/complete.html"))
	})
})
