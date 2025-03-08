package common

import (
	"github.com/hainguyen27798/gin-boilerplate/pkg/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBsonUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bson Utils Suite")
}

var _ = Describe("IsValidObjectID", func() {
	Context("with valid ObjectID hex strings", func() {
		It("should return true", func() {
			validID := "507f1f77bcf86cd799439011"
			Expect(common.IsValidObjectID(validID)).To(BeTrue())
		})
	})

	Context("with invalid hex strings", func() {
		It("should return false", func() {
			invalidID := "invalid-hex-string"
			Expect(common.IsValidObjectID(invalidID)).To(BeFalse())
		})
	})

	Context("with empty strings", func() {
		It("should return false", func() {
			Expect(common.IsValidObjectID("")).To(BeFalse())
		})
	})

	Context("with short hex strings", func() {
		It("should return false", func() {
			shortID := "507f1f"
			Expect(common.IsValidObjectID(shortID)).To(BeFalse())
		})
	})

	Context("with long hex strings", func() {
		It("should return false", func() {
			longID := "507f1f77bcf86cd799439011507f1f77bcf86cd799439011"
			Expect(common.IsValidObjectID(longID)).To(BeFalse())
		})
	})
})
