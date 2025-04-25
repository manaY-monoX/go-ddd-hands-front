package tests

import (
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("商品キーワード検索サービスの品質検証\n", Ordered, func() {
	var env *ServiceTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})
	When("存在するキーワードで検索する\n", func() {
		It("検索結果が返される\n", func() {
			results, err := env.keyword.Execute("ボールペン")
			Expect(err).To(BeNil())
			Expect(results.Categories).To(BeNil())
			Expect(results.Product).To(BeNil())
			Expect(results.Error).To(BeNil())
			Expect(results.Products).NotTo(BeNil())
			for _, product := range *results.Products {
				log.Println(product)
			}
		})
	})
	When("存在しないキーワードで検索する\n", func() {
		It("エラーメッセージが返される\n", func() {
			result, err := env.keyword.Execute("川")
			Expect(err).To(BeNil())
			Expect(result.Categories).To(BeNil())
			Expect(result.Products).To(BeNil())
			Expect(result.Product).To(BeNil())
			Expect(result.Error).NotTo(BeNil())
			Expect(result.Error.Message).To(Equal("キーワード:'川'に該当する商品は見つかりませんでした。"))
			log.Println(result.Error.Message)
		})
	})
})
