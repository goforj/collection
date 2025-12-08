package main

import "github.com/goforj/collection"

func main() {

	type User struct {
	    Name string
	    Role string
	}

	users := collection.New([]User{
	    {Name: "Alice", Role: "admin"},
	    {Name: "Bob", Role: "user"},
	    {Name: "Charlie", Role: "admin"},
	    {Name: "David", Role: "user"},
	    {Name: "Eve", Role: "admin"},
	    {Name: "Frank", Role: "user"},
	    {Name: "Grace", Role: "user"},
	    {Name: "Heidi", Role: "user"},
	})
	counts := CountBy(users, func(u User) string { return u.Role == "admin" })
	// map[string]int{"admin": 3, "user": 5}
}
