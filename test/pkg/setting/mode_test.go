package setting

import (
	setting2 "github.com/hainguyen27798/gin-boilerplate/pkg/setting"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestAppMode_String", func() {
	It("should return correct string for ProdMode", func() {
		mode := setting2.ProdMode
		Expect(string(mode), "prod")
	})

	It("should return correct string for DevMode", func() {
		mode := setting2.DevMode
		Expect(string(mode), "dev")
	})
})

var _ = Describe("TestAppMode_Constants", func() {
	It("should have correct constant values", func() {
		Expect(setting2.ProdMode).NotTo(Equal(setting2.DevMode))
		Expect(setting2.ProdMode).To(Equal(setting2.AppMode("prod")))
		Expect(setting2.DevMode).To(Equal(setting2.AppMode("dev")))
	})
})

var _ = Describe("TestAppMode_StringConversion", func() {
	It("should allow conversion from string", func() {
		mode := setting2.AppMode("prod")
		Expect(mode).To(Equal(setting2.ProdMode))

		mode = "dev"
		Expect(mode).To(Equal(setting2.DevMode))
	})

	It("should handle empty string conversion", func() {
		mode := setting2.AppMode("")
		Expect(mode).NotTo(Equal(setting2.ProdMode))
		Expect(mode).NotTo(Equal(setting2.DevMode))
		Expect(mode).To(Equal(setting2.AppMode("")))
	})
})
