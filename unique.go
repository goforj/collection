package collection

// Unique returns a new collection with duplicate items removed, based on the
// equality function `eq`. The first occurrence of each unique value is kept,
// and order is preserved.
// @group Set Operations
// @behavior immutable
// @fluent true
//
// The `eq` function should return true when two values are considered equal.
//
// Example: integers
//
//	c1 := collection.New([]int{1, 2, 2, 3, 4, 4, 5})
//	out1 := c1.Unique(func(a, b int) bool { return a == b })
//	collection.Dump(out1.Items())
//	// #[]int [
//	//	0 => 1 #int
//	//	1 => 2 #int
//	//	2 => 3 #int
//	//	3 => 4 #int
//	//	4 => 5 #int
//	// ]
//
// Example: strings (case-insensitive uniqueness)
//
//	c2 := collection.New([]string{"A", "a", "B", "b", "A"})
//	out2 := c2.Unique(func(a, b string) bool {
//		return strings.EqualFold(a, b)
//	})
//	collection.Dump(out2.Items())
//	// #[]string [
//	//	0 => "A" #string
//	//	1 => "B" #string
//	// ]
//
// Example: structs (unique by ID)
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
//	c3 := collection.New([]User{
//		{ID: 1, Name: "Alice"},
//		{ID: 2, Name: "Bob"},
//		{ID: 1, Name: "Alice Duplicate"},
//	})
//
//	out3 := c3.Unique(func(a, b User) bool {
//		return a.ID == b.ID
//	})
//
//	collection.Dump(out3.Items())
//	// #[]collection.User [
//	//	0 => {ID:1 Name:"Alice"} #collection.User
//	//	1 => {ID:2 Name:"Bob"}   #collection.User
//	// ]
func (c *Collection[T]) Unique(eq func(a, b T) bool) *Collection[T] {
	out := make([]T, 0, len(c.items))

	for _, v := range c.items {
		found := false
		for _, existing := range out {
			if eq(v, existing) {
				found = true
				break
			}
		}
		if !found {
			out = append(out, v)
		}
	}

	return Attach(out)
}
