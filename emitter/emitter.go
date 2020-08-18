// Package emitter is used to emit strings to stdout.
package emitter

import (
	"fmt"
	"io"
	"os"

	// Because the Go tools use the root go.mod file
	// to help determine its actions, it will check
	// the go.mod file in this project, determine the
	// local filesystem contains the files needed
	// for github.com/rp185145/goprojectlayout,
	// including idutils/, then use the local copy to
	// resolve the package rather than using the
	// network location.
	"github.com/rp185145/goprojectlayout/idutils"
)

// output is used to demonstrate a private function
// and is used to send a given string to a Writer
func output(writer io.Writer, s string) {
	fmt.Println("emitter.output()")
	fmt.Fprintln(writer, "Emitter:", s)
	fmt.Fprintln(writer, "Your ID is", idutils.GetUUID())
}

// Normally APIs should not:
//   1. exit for any reason (calling os.Exit() is a
//      mistake)
//   2. handle all errors and eventually return to
//      the caller gracefully
//   3. generally not output to the user interface
//      unless the API is specifically designed to
//      do so (such as the Cobra package).

// Emitter emits a string to stdout
func Emitter(s string) {
	fmt.Println("emitter.Emitter()")
	output(os.Stdout, s)
}
