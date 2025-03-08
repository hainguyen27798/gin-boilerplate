package common

import (
	"github.com/hainguyen27798/gin-boilerplate/pkg/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/v2/bson"
	"testing"
	"time"
)

func TestBaseModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Base Model Suite")
}

var _ = Describe("BaseModel", func() {
	var existingID bson.ObjectID

	BeforeEach(func() {
		existingID = bson.NewObjectID()
	})

	Describe("BeforeCreate", func() {
		Context("when fields are empty", func() {
			It("should set ID and timestamps", func() {
				model := &common.BaseModel{}
				model.BeforeCreate()

				Expect(model.ID).NotTo(BeEmpty())
				Expect(model.CreatedAt.IsZero()).To(BeFalse())
				Expect(model.UpdatedAt.IsZero()).To(BeFalse())
				Expect(model.CreatedAt.UTC().Location()).To(Equal(time.UTC))
				Expect(model.UpdatedAt.UTC().Location()).To(Equal(time.UTC))
			})
		})

		Context("when ID and CreatedAt already exist", func() {
			It("should preserve existing ID and CreatedAt", func() {
				existingTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

				model := &common.BaseModel{
					ID:        existingID,
					CreatedAt: existingTime,
				}
				model.BeforeCreate()

				Expect(model.ID).To(Equal(existingID))
				Expect(model.CreatedAt).To(Equal(existingTime))
				Expect(model.UpdatedAt.IsZero()).To(BeFalse())
				Expect(model.UpdatedAt.Equal(existingTime)).To(BeTrue())
			})
		})
	})

	Describe("BeforeUpdate", func() {
		Context("when updating a model", func() {
			It("should update the UpdatedAt timestamp", func() {
				model := &common.BaseModel{
					ID:        bson.NewObjectID(),
					CreatedAt: time.Now().Add(-24 * time.Hour),
					UpdatedAt: time.Now().Add(-24 * time.Hour),
				}

				oldUpdatedAt := model.UpdatedAt
				time.Sleep(time.Millisecond)
				model.BeforeUpdate()

				Expect(model.UpdatedAt).To(BeTemporally(">", oldUpdatedAt))
				Expect(model.UpdatedAt.Location()).To(Equal(time.UTC))
			})
		})

		Context("when preserving other fields", func() {
			It("should keep the original ID and CreatedAt", func() {
				originalID := bson.NewObjectID()
				originalCreatedAt := time.Now().Add(-24 * time.Hour)

				model := &common.BaseModel{
					ID:        originalID,
					CreatedAt: originalCreatedAt,
				}

				model.BeforeUpdate()

				Expect(model.ID.Hex()).To(Equal(originalID.Hex()))
				Expect(model.CreatedAt).To(Equal(originalCreatedAt))
			})
		})
	})
})
