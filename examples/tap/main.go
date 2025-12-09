//go:build ignore
// +build ignore

package main

import "github.com/goforj/collection"

func main() {

	  captured := []int{}
	  c := New([]int{3,1,2}).
	      Sort(func(a,b int) bool { return a < b }).  // → [1,2,3]
	      Tap(func(col *Collection[int]) {
	          captured = append([]int(nil), col.items...) // snapshot
	      }).
	      Filter(func(v int) bool { return v >= 2 })     // → [2,3]

	After Tap, 'captured' contains the sorted state: []int{1,2,3}
	and the chain continues unaffected.
}
