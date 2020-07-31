// Package idutils is a utility package used for working
// with IDs.
package idutils

import (
	"fmt"

	"github.com/google/uuid"
)

// GetUUID is used to abstract away the generation
// of a UUID in the event we wish to use a different
// API

// GetUUID generates and returns a new UUID
func GetUUID() uuid.UUID {
	fmt.Println("idutils.GetUUID()")
	return uuid.New()
}
