package helpers

import (
	helpers2 "github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestHashPassword(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hash Password Suite")
}

var _ = Describe("TestHashPassword", func() {
	It("should successfully hash password", func() {
		password := "mySecurePassword123"
		hash, err := helpers2.HashPassword(password)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(hash).NotTo(BeEmpty())
		Expect(hash).NotTo(Equal(password))
	})

	It("should generate different hashes for same password", func() {
		password := "mySecurePassword123"
		hash1, _ := helpers2.HashPassword(password)
		hash2, _ := helpers2.HashPassword(password)

		Expect(hash1).NotTo(Equal(hash2))
	})

	It("should handle empty password", func() {
		hash, err := helpers2.HashPassword("")

		Expect(err).ShouldNot(HaveOccurred())
		Expect(hash).NotTo(BeEmpty())
	})
})

var _ = Describe("TestCheckPasswordHash", func() {
	It("should verify correct password", func() {
		password := "mySecurePassword123"
		hash, _ := helpers2.HashPassword(password)
		Expect(helpers2.CheckPasswordHash(password, hash)).To(BeTrue())
	})

	It("should reject incorrect password", func() {
		password := "mySecurePassword123"
		hash, _ := helpers2.HashPassword(password)
		Expect(helpers2.CheckPasswordHash("wrongPassword", hash)).To(BeFalse())
	})

	It("should handle empty password and hash", func() {
		Expect(helpers2.CheckPasswordHash("", "")).To(BeFalse())
	})

	It("should reject invalid hash format", func() {
		Expect(helpers2.CheckPasswordHash("password", "invalidhash")).To(BeFalse())
	})
})
