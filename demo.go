// Package goprojectlayout is used as a base package in a
// project to demonstrate how to implement and use exported
// functions from the base package.
package goprojectlayout

import (
	"fmt"
	"runtime"
)

// Demo demonstraes implementation of an exported function in
// root package. It also emits the Go version used for the build
// along with the target OS and architecture to stdout.
func Demo() {
	fmt.Println("gcpprojectlayout.Demo() in the root package")
	fmt.Printf("Compiled with %s for %s/%s\n", runtime.Version(),
		runtime.GOOS, runtime.GOARCH)
}
