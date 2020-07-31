package main

import (
	"fmt"

	"github.com/rp185145/goprojectlayout/emitter"
)

func main() {
	fmt.Println("hello CLI")
	emitter.Emitter("Hello World!")
}
