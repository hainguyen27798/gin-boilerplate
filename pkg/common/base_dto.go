package common

import "go.mongodb.org/mongo-driver/v2/bson"

// BaseDto is a base data transfer object struct that can be embedded in other
// DTO structs to provide common fields and functionality.
type BaseDto struct{}

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
