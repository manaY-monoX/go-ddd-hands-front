package tests

import (
	"front-exercise/service/impls"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("サーバ接続確認の品質検証\n", Ordered, func() {
	var env *ServiceTestEnvironment
	BeforeEach(func() {
		env = SetupTestEnvironment()
	})
	AfterEach(func() {
		TeardownTestEnvironment(env)
	})
	It("サーバーとの接続を確認する\n", func() {
		base := impls.NewBaseService(env.conf)
		err := base.Ping()
		Expect(err).To(BeNil())
	})
})
