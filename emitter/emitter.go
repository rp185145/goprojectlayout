// Package emitter is used to emit strings to stdout.
package emitter

import (
	"fmt"
	"io"
	"os"

	"github.com/rp185145/goprojectlayout/idutils"
)

// output is used to demonstrate a private function
// and is used to write to a Writer
func output(writer io.Writer, s string) {
	fmt.Println("emitter.output()")
	fmt.Fprintln(writer, "Emitter:", s)
	fmt.Fprintln(writer, "Your ID is", idutils.GetUUID())
}

// Emitter emits a string to stdout
//
// Normally APIs should not 1) exit for any reason
// (calling os.Exit() is a mistake) 2) handle all
// errors and eventually return to the caller
// gracefully 3) generally not output to the user
// interface unless the API is specifically designed
// to do so (such as the Cobra package).
func Emitter(s string) {
	fmt.Println("emitter.Emitter()")
	output(os.Stdout, s)
}
