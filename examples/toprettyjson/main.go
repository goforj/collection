//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/goforj/collection"
)

func main() {
	// strings → pretty JSON
	pj1 := collection.New([]string{"a", "b"})
	out1, _ := pj1.ToPrettyJSON()
	fmt.Println(out1)
	// [
	//  "a",
	//  "b"
	// ]
}
