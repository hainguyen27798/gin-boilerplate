package setting

import (
	setting2 "github.com/hainguyen27798/gin-boilerplate/pkg/setting"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestConfig_Structure", func() {
	It("should have all required fields", func() {
		config := setting2.Config{}
		Expect(config.Server).NotTo(BeNil())
		Expect(config.Logger).NotTo(BeNil())
		Expect(config.MongoDB).NotTo(BeNil())
	})
})

var _ = Describe("TestServerSettings_Validation", func() {
	It("should handle empty port", func() {
		settings := setting2.ServerSettings{}
		Expect(settings.Port).To(BeEmpty())
	})

	It("should store port value", func() {
		settings := setting2.ServerSettings{Port: "8080"}
		Expect(settings.Port).To(Equal("8080"))
	})
})

var _ = Describe("TestLoggerSettings_Validation", func() {
	It("should initialize with zero values", func() {
		settings := setting2.LoggerSettings{}
		Expect(settings.FileName).To(BeEmpty())
		Expect(settings.Level).To(BeEmpty())
		Expect(settings.MaxSize).To(BeZero())
		Expect(settings.MaxBackups).To(BeZero())
		Expect(settings.MaxAge).To(BeZero())
		Expect(settings.Compress).To(BeFalse())
	})

	It("should store all field values", func() {
		settings := setting2.LoggerSettings{
			FileName:   "app.log",
			Level:      "info",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     7,
			Compress:   true,
		}
		Expect(settings.FileName).To(Equal("app.log"))
		Expect(settings.Level).To(Equal("info"))
		Expect(settings.MaxSize).To(Equal(100))
		Expect(settings.MaxBackups).To(Equal(3))
		Expect(settings.MaxAge).To(Equal(7))
		Expect(settings.Compress).To(BeTrue())
	})
})

var _ = Describe("TestMongoDBSettings_Validation", func() {
	It("should initialize with zero values", func() {
		settings := setting2.MongoDBSettings{}
		Expect(settings.Host).To(BeEmpty())
		Expect(settings.Port).To(BeEmpty())
		Expect(settings.Username).To(BeEmpty())
		Expect(settings.Password).To(BeEmpty())
		Expect(settings.Database).To(BeEmpty())
		Expect(settings.MaxPoolSize).To(BeZero())
		Expect(settings.EnableLog).To(BeFalse())
		Expect(settings.DirectConnection).To(BeFalse())
	})

	It("should store all field values", func() {
		settings := setting2.MongoDBSettings{
			Host:             "localhost",
			Port:             "27017",
			Username:         "admin",
			Password:         "password",
			Database:         "testdb",
			MaxPoolSize:      100,
			EnableLog:        true,
			DirectConnection: true,
		}
		Expect(settings.Host).To(Equal("localhost"))
		Expect(settings.Port).To(Equal("27017"))
		Expect(settings.Username).To(Equal("admin"))
		Expect(settings.Password).To(Equal("password"))
		Expect(settings.Database).To(Equal("testdb"))
		Expect(settings.MaxPoolSize).To(Equal(uint64(100)))
		Expect(settings.EnableLog).To(BeTrue())
		Expect(settings.DirectConnection).To(BeTrue())
	})
})
