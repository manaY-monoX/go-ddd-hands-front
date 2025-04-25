package tests

import (
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("商品カテゴリ取得サービスの品質検証\n", Ordered, func() {
	var env *ServiceTestEnvironment
	BeforeEach(func() {
		env = SetupTestEnvironment()
	})
	AfterEach(func() {
		TeardownTestEnvironment(env)
	})

	It("商品カテゴリを取得する\n", func() {
		results, err := env.list.Execute()
		Expect(err).To(BeNil())
		Expect(results.Product).To(BeNil())
		Expect(results.Products).To(BeNil())
		Expect(results.Error).To(BeNil())
		for _, category := range *results.Categories {
			log.Println(category)
		}
	})
})
