package tests

import (
	"front-exercise/service/dto"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("商品登録サービスの品質検証\n", Ordered, func() {
	var env *ServiceTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})
	When("存在しない商品を登録する\n", func() {
		It("エラーなく登録結果が返される\n", func() {
			category := dto.CategoryDTO{
				Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c0",
				Name: "文房具",
			}
			product := dto.ProductDTO{
				Id:       "",
				Name:     "消しゴム",
				Price:    "150",
				Category: category,
			}
			result, err := env.register.Execute(&product)
			Expect(err).To(BeNil())
			Expect(result.Categories).To(BeNil())
			Expect(result.Products).To(BeNil())
			Expect(result.Error).To(BeNil())
			Expect(result.Product).NotTo(BeNil())
			log.Println(result.Product)
		})
	})
	When("存在する商品を登録する\n", func() {
		It("エラーが返される\n", func() {
			category := dto.CategoryDTO{
				Id:   "b1524011-b6af-417e-8bf2-f449dd58b5c0",
				Name: "文房具",
			}
			product := dto.ProductDTO{
				Id:       "",
				Name:     "消しゴム",
				Price:    "150",
				Category: category,
			}
			result, err := env.register.Execute(&product)
			Expect(err).To(BeNil())
			Expect(result.Categories).To(BeNil())
			Expect(result.Products).To(BeNil())
			Expect(result.Product).To(BeNil())
			Expect(result.Error).NotTo(BeNil())
			Expect(result.Error.Message).To(Equal("商品:消しゴムは、既に登録済です。"))
			log.Println()
		})
	})
})
