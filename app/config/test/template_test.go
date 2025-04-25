package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("テンプレート取得の品質検証\n", Ordered, func() {
	var env *ConfigTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})

	When("トップページテンプレート\n", func() {
		It("トップページを取得する", func() {
			templ, err := env.conf.GetViewName("top")
			Expect(err).To(BeNil())
			Expect(templ).To(Equal("top/top.html"))
		})
	})
	When("検索ページテンプレート\n", func() {
		It("検索ページを取得する", func() {
			templ, err := env.conf.GetViewName("search")
			Expect(err).To(BeNil())
			Expect(templ).To(Equal("search/search.html"))
		})
	})
	When("登録ページテンプレート\n", func() {
		It("入力ページを取得する", func() {
			templ, err := env.conf.GetViewName("register_enter")
			Expect(err).To(BeNil())
			Expect(templ).To(Equal("register/enter.html"))
		})
		It("登録結果ページを取得する", func() {
			templ, err := env.conf.GetViewName("register_complete")
			Expect(err).To(BeNil())
			Expect(templ).To(Equal("register/complete.html"))
		})
	})
	When("エラーページテンプレート\n", func() {
		It("エラーページを取得する", func() {
			templ, err := env.conf.GetViewName("error")
			Expect(err).To(BeNil())
			Expect(templ).To(Equal("error/error.html"))
		})
	})
})
