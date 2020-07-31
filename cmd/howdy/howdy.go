package main

import (
	"fmt"

	"github.com/rp185145/goprojectlayout/emitter"
)

func main() {
	fmt.Println("howdy CLI")
	emitter.Emitter("Howdy Pardner!")
}
