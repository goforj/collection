package collection

// Map maps this Slice to a newly allocated Slice with a potentially different
// element type.
// @group Transformation
// @behavior immutable
// @chainable true
// @terminal false
//
// Example: map integers to labels
//
//	numbers := collection.New([]int{1, 2, 3, 4})
//	labels := numbers.Map(func(number int) string {
//		if number%2 == 0 {
//			return "even"
//		}
//		return "odd"
//	})
//	collection.Dump(labels)
//	// #[]string [
//	//   0 => "odd" #string
//	//   1 => "even" #string
//	//   2 => "odd" #string
//	//   3 => "even" #string
//	// ]
//	fmt.Println(numbers[0])
//	// 1
func (c Slice[T]) Map[R any](fn func(T) R) Slice[R] {
	out := make([]R, len(c))
	for i, value := range c {
		out[i] = fn(value)
	}
	return New(out)
}

// MinBy returns the item whose extracted key is the smallest.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: shortest string
//
//	words := collection.New([]string{"pear", "fig", "banana"})
//	shortest, ok := words.MinBy(func(word string) int {
//		return len(word)
//	})
//	collection.Dump(shortest, ok)
//	// "fig" #string
//	// true #bool
func (c Slice[T]) MinBy[K Number | ~string](keyFn func(T) K) (T, bool) {
	var zero T
	if len(c) == 0 {
		return zero, false
	}

	minItem := c[0]
	minKey := keyFn(minItem)
	for _, item := range c[1:] {
		key := keyFn(item)
		if key < minKey {
			minKey = key
			minItem = item
		}
	}
	return minItem, true
}

// MaxBy returns the item whose extracted key is the largest.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: longest string
//
//	words := collection.New([]string{"pear", "fig", "banana"})
//	longest, ok := words.MaxBy(func(word string) int {
//		return len(word)
//	})
//	collection.Dump(longest, ok)
//	// "banana" #string
//	// true #bool
func (c Slice[T]) MaxBy[K Number | ~string](keyFn func(T) K) (T, bool) {
	var zero T
	if len(c) == 0 {
		return zero, false
	}

	maxItem := c[0]
	maxKey := keyFn(maxItem)
	for _, item := range c[1:] {
		key := keyFn(item)
		if key > maxKey {
			maxKey = key
			maxItem = item
		}
	}
	return maxItem, true
}

// ToMap reduces this collection into a map using the provided key and value functions.
// If multiple items produce the same key, the value derived from the last item wins.
// @group Maps
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: index words by their value
//
//	words := collection.New([]string{"go", "forj"})
//	lengths := words.ToMap(
//		func(word string) string { return word },
//		func(word string) int { return len(word) },
//	)
//	collection.Dump(lengths)
//	// #map[string]int {
//	//   forj => 4 #int
//	//   go => 2 #int
//	// }
func (c Slice[T]) ToMap[K comparable, V any](keyFn func(T) K, valueFn func(T) V) map[K]V {
	out := make(map[K]V, len(c))
	for _, item := range c {
		out[keyFn(item)] = valueFn(item)
	}
	return out
}

// GroupBy partitions this Slice into independent built-in slices keyed by the
// extracted value.
// @group Grouping
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: group integers by parity
//
//	numbers := collection.New([]int{1, 2, 3, 4})
//	groups := numbers.GroupBy(func(number int) string {
//		if number%2 == 0 {
//			return "even"
//		}
//		return "odd"
//	})
//	collection.Dump(groups["even"], groups["odd"])
//	// #[]int [
//	//   0 => 2 #int
//	//   1 => 4 #int
//	// ]
//	// #[]int [
//	//   0 => 1 #int
//	//   1 => 3 #int
//	// ]
//	fmt.Println(len(groups["even"]))
//	// 2
//	fmt.Println(groups["odd"][0])
//	// 1
//	collection.Dump(groups["even"][:1])
//	// #[]int [
//	//   0 => 2 #int
//	// ]
func (c Slice[T]) GroupBy[K comparable](keyFn func(T) K) map[K][]T {
	out := make(map[K][]T)
	for _, item := range c {
		key := keyFn(item)
		out[key] = append(out[key], item)
	}
	return out
}

// CountBy returns occurrence counts keyed by the extracted value.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
//
// Example: count integers by parity
//
//	numbers := collection.New([]int{1, 2, 3, 5})
//	counts := numbers.CountBy(func(number int) string {
//		if number%2 == 0 {
//			return "even"
//		}
//		return "odd"
//	})
//	collection.Dump(counts)
//	// #map[string]int {
//	//   even => 1 #int
//	//   odd => 3 #int
//	// }
func (c Slice[T]) CountBy[K comparable](keyFn func(T) K) map[K]int {
	result := make(map[K]int)
	for _, item := range c {
		result[keyFn(item)]++
	}
	return result
}

// UniqueBy returns a collection containing the first item for each extracted key.
// @group Set Operations
// @behavior immutable
// @chainable true
// @terminal false
//
// Example: keep the first word of each length
//
//	words := collection.New([]string{"go", "up", "forj", "code"})
//	unique := words.UniqueBy(func(word string) int {
//		return len(word)
//	})
//	collection.Dump(unique)
//	// #[]string [
//	//   0 => "go" #string
//	//   1 => "forj" #string
//	// ]
func (c Slice[T]) UniqueBy[K comparable](keyFn func(T) K) Slice[T] {
	items := c
	if len(items) == 0 {
		return New([]T{})
	}

	seen := make(map[K]struct{}, len(items))
	out := make([]T, 0, len(items))
	for _, item := range items {
		key := keyFn(item)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, item)
	}
	return New(out)
}

// ZipWith combines this collection with a slice using fn up to the shorter length.
// @group Transformation
// @behavior immutable
// @chainable true
// @terminal false
//
// Example: add corresponding integers
//
//	left := collection.New([]int{1, 2, 3})
//	right := collection.New([]int{10, 20})
//	sums := left.ZipWith(right, func(a, b int) int {
//		return a + b
//	})
//	collection.Dump(sums)
//	// #[]int [
//	//   0 => 11 #int
//	//   1 => 22 #int
//	// ]
func (c Slice[T]) ZipWith[U, R any](other []U, fn func(T, U) R) Slice[R] {
	length := len(c)
	if len(other) < length {
		length = len(other)
	}

	out := make([]R, length)
	for i := 0; i < length; i++ {
		out[i] = fn(c[i], other[i])
	}
	return New(out)
}
