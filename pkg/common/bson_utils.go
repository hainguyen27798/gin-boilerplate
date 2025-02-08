package common

import "go.mongodb.org/mongo-driver/v2/bson"

// IsValidObjectID checks if the given string is a valid BSON ObjectID.
func IsValidObjectID(id string) bool {
	_, err := bson.ObjectIDFromHex(id)
	return err == nil
}
