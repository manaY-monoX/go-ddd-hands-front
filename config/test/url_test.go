package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("URL生成の品質検証\n", Ordered, func() {
	var env *ConfigTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})

	Context("サーバURL生成を検証する\n", func() {
		When("部署一覧取得URL生成を検証する\n", func() {
			It("正しいpathkeyを指定するとエラーなくURLが返される\n", func() {
				url, err := env.conf.BuildServerURL("categoryList")
				Expect(err).To(BeNil())
				Expect(url).To(Equal("http://service_exercise:8085/category/list"))
			})
			It("正しくないpathkeyを指定するとエラーが返される\n", func() {
				url, err := env.conf.BuildServerURL("abd")
				Expect(url).To(Equal(""))
				Expect(err.Error()).To(Equal("不正なパスキーが指定されました。"))
			})
		})
		When("商品キーワード検索URL生成を検証する\n", func() {
			It("正しいpathkeyを指定するとエラーなくURLが返される\n", func() {
				url, err := env.conf.BuildServerURL("productKeyword")
				Expect(err).To(BeNil())
				Expect(url).To(Equal("http://service_exercise:8085/product/keyword/%s"))
			})
			It("正しくないpathkeyを指定するとエラーが返される\n", func() {
				url, err := env.conf.BuildServerURL("abd")
				Expect(url).To(Equal(""))
				Expect(err.Error()).To(Equal("不正なパスキーが指定されました。"))
			})
		})
		When("商品登録URL生成を検証する\n", func() {
			It("正しいpathkeyを指定するとエラーなくURLが返される\n", func() {
				url, err := env.conf.BuildServerURL("productRegister")
				Expect(err).To(BeNil())
				Expect(url).To(Equal("http://service_exercise:8085/product/register"))
			})
			It("正しくないpathkeyを指定するとエラーが返される\n", func() {
				url, err := env.conf.BuildServerURL("abd")
				Expect(url).To(Equal(""))
				Expect(err.Error()).To(Equal("不正なパスキーが指定されました。"))
			})
		})
	})

	Context("アプリケーションURL生成を検証する\n", func() {
		When("トップページURL生成を検証する\n", func() {
			It("正しいpathkeyを指定するとエラーなくURLが返される\n", func() {
				url, err := env.conf.BuildAppURL("top")
				Expect(err).To(BeNil())
				Expect(url).To(Equal("http://localhost:8081/exercise/top"))
			})
			It("正しくないpathkeyを指定するとエラーが返される\n", func() {
				url, err := env.conf.BuildAppURL("abd")
				Expect(url).To(Equal(""))
				Expect(err.Error()).To(Equal("不正なパスキーが指定されました。"))
			})
		})
		When("検索ページURL生成を検証する\n", func() {
			It("正しいpathkeyを指定するとエラーなくURLが返される\n", func() {
				url, err := env.conf.BuildAppURL("search")
				Expect(err).To(BeNil())
				Expect(url).To(Equal("http://localhost:8081/exercise/search"))
			})
			It("正しくないpathkeyを指定するとエラーが返される\n", func() {
				url, err := env.conf.BuildAppURL("abd")
				Expect(url).To(Equal(""))
				Expect(err.Error()).To(Equal("不正なパスキーが指定されました。"))
			})
		})
		When("登録ページURL生成を検証する\n", func() {
			It("正しいpathkeyを指定するとエラーなくURLが返される\n", func() {
				url, err := env.conf.BuildAppURL("register")
				Expect(err).To(BeNil())
				Expect(url).To(Equal("http://localhost:8081/exercise/register"))
			})
			It("正しくないpathkeyを指定するとエラーが返される\n", func() {
				url, err := env.conf.BuildAppURL("abd")
				Expect(url).To(Equal(""))
				Expect(err.Error()).To(Equal("不正なパスキーが指定されました。"))
			})
		})
	})
})
