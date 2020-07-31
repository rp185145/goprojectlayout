package main

import (
	"fmt"

	"github.com/rp185145/goprojectlayout/idutils"
)

func main() {
	fmt.Println("uuid CLI")
	fmt.Println("UUID:", idutils.GetUUID())
}
