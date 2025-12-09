package main

import "github.com/goforj/collection"

func main() {
	  sorted := users.Sort(func(a, b User) bool { return a.Age < b.Age })
	 // sorted by Age ascending
}
