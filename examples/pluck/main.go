package main

import "github.com/goforj/collection"

func main() {
	  names := users.Pluck(func(u User) string { return u.Name })
	  // names is a Collection[string] of user names
}
