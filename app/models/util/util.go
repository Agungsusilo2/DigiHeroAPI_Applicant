package util

import "github.com/google/uuid"

var UniqueID string

func init() {
	UniqueID = generateUniqueID()
}

func generateUniqueID() string {
	return uuid.New().String()
}
