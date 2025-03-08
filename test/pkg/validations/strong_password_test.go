package validations

import (
	"github.com/go-playground/validator/v10"
	validations2 "github.com/hainguyen27798/gin-boilerplate/pkg/validations"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestStrongPassword(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Strong Password Suite")
}

type TestStruct struct {
	Password string `validate:"strong_password"`
}

var _ = Describe("StrongPassword", func() {
	var validate *validator.Validate

	BeforeEach(func() {
		validate = validator.New()
		err := validate.RegisterValidation("strong_password", validations2.StrongPassword)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("with valid passwords", func() {
		validPasswords := []string{
			"Ab@12345",
			"SecurePass123!",
			"ValidP@ssw0rd",
		}

		for _, password := range validPasswords {
			// Use a local variable to avoid closure issues
			pwd := password
			It("should accept "+pwd, func() {
				test := TestStruct{Password: pwd}
				err := validate.Struct(test)
				Expect(err).NotTo(HaveOccurred(), "Password should be valid: "+pwd)
			})
		}
	})

	Context("with invalid passwords", func() {
		invalidPasswords := []string{
			"",
			"short",
			"1234567",
			"12345678",
			"!@#$%^&*",
			"abcdefg",
			"       ",
			"\t\n\r",
		}

		for _, password := range invalidPasswords {
			// Use a local variable to avoid closure issues
			pwd := password
			It("should reject "+pwd, func() {
				test := TestStruct{Password: pwd}
				err := validate.Struct(test)
				Expect(err).To(HaveOccurred(), "Password should be invalid: "+pwd)
			})
		}
	})

	Context("with edge cases", func() {
		edgeCases := []struct {
			password string
			valid    bool
		}{
			{"12345678", false},                         // Numbers only - invalid
			{"123456789012345678901234567890", false},   // Numbers only - invalid
			{"        ", false},                         // Spaces only - invalid
			{"!@#$%^&*", false},                         // Special chars only - invalid
			{"ñÑáéíóúÁÉÍÓÚ12345678", false},             // Unicode without required char mix
			{"\x00\x00\x00\x00\x00\x00\x00\x00", false}, // Null bytes
			{"Str0ng!Pass", true},                       // Valid mix of characters
			{"Complex1@Pwd", true},                      // Valid mix
		}

		for _, tc := range edgeCases {
			// Use local variables to avoid closure issues
			pwd := tc.password
			valid := tc.valid
			It("should correctly handle "+pwd, func() {
				test := TestStruct{Password: pwd}
				err := validate.Struct(test)
				if valid {
					Expect(err).NotTo(HaveOccurred(), "Password should be valid: "+pwd)
				} else {
					Expect(err).To(HaveOccurred(), "Password should be invalid: "+pwd)
				}
			})
		}
	})
})
