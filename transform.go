package collection

// Transform applies fn to every item *in place*, mutating the collection.
//
// This mirrors Laravel's transform(), which modifies the underlying values
// instead of returning a new collection.
//
// Example: integers
//	c1 := collection.New([]int{1, 2, 3})
//	c1.Transform(func(v int) int { return v * 2 })
//	collection.Dump(c1.Items())
//	// #[]int [
//	//	0 => 2 #int
//	//	1 => 4 #int
//	//	2 => 6 #int
//	// ]
//
// Example: strings
//	c2 := collection.New([]string{"a", "b", "c"})
//	c2.Transform(func(s string) string { return strings.ToUpper(s) })
//	collection.Dump(c2.Items())
//	// #[]string [
//	//	0 => "A" #string
//	//	1 => "B" #string
//	//	2 => "C" #string
//	// ]
//
// Example: structs
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	c3 := collection.New([]User{
//		{ID: 1, Name: "alice"},
//		{ID: 2, Name: "bob"},
//	})
//
//	c3.Transform(func(u User) User {
//		u.Name = strings.ToUpper(u.Name)
//		return u
//	})
//
//	collection.Dump(c3.Items())
//	// #[]collection.User [
//	//	0 => {ID:1 Name:"ALICE"} #collection.User
//	//	1 => {ID:2 Name:"BOB"}   #collection.User
//	// ]
func (c *Collection[T]) Transform(fn func(T) T) {
	for i, v := range c.items {
		c.items[i] = fn(v)
	}
}
