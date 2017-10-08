package ping_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestPing(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Ping Suite")
}

var _ = Describe("Ping", func() {
	Describe("a true truth", func() {
		It("should always be true", func() {
			Expect(true).To(Equal(true))
		})
	})
})
