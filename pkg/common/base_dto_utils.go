package common

import (
	"fmt"

	"github.com/hainguyen27798/gin-boilerplate/global"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// ToBson converts the given interface{} value to a *bson.D document. It first
// marshals the value to BSON data, and then unmarshals the data into a *bson.D
// document.
func ToBson(dto interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(dto)
	if err != nil {
		return nil, err
	}

	if err := bson.Unmarshal(data, &doc); err != nil {
		return nil, err
	}

	return doc, err
}

// ValidateStruct validates the given struct against the global validator. If the
// global validator is not initialized, it returns an error.
func ValidateStruct(s interface{}) error {
	if global.Validator == nil {
		return fmt.Errorf("global.Validator is not initialized")
	}
	return global.Validator.Struct(s)
}
