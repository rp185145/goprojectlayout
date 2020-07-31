// Package idutils is a utility package used for working
// with IDs.
package idutils

import (
	"fmt"

	"github.com/google/uuid"
)

// GetUUID generates and returns a new UUID
//
// It is used to astract away the generation of
// a UUID in the event we wish to use a different
// API
func GetUUID() uuid.UUID {
	fmt.Println("idutils.GetUUID()")
	return uuid.New()
}
