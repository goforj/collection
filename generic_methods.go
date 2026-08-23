package collection

// MapTo maps this collection to a collection with a different element type.
// @group Transformation
// @behavior immutable
// @chainable true
// @terminal false
func (c *Collection[T]) MapTo[R any](fn func(T) R) *Collection[R] {
	items := c.Items()
	out := make([]R, len(items))
	for i, value := range items {
		out[i] = fn(value)
	}
	return New(out)
}

// MinBy returns the item whose extracted key is the smallest.
// @group Aggregation
// @behavior readonly
// @chainable false
// @terminal true
func (c *Collection[T]) MinBy[K Number | ~string](keyFn func(T) K) (T, bool) {
	var zero T
	if len(c.items) == 0 {
		return zero, false
	}

	minItem := c.items[0]
	minKey := keyFn(minItem)
	for _, item := range c.items[1:] {
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
func (c *Collection[T]) MaxBy[K Number | ~string](keyFn func(T) K) (T, bool) {
	var zero T
	if len(c.items) == 0 {
		return zero, false
	}

	maxItem := c.items[0]
	maxKey := keyFn(maxItem)
	for _, item := range c.items[1:] {
		key := keyFn(item)
		if key > maxKey {
			maxKey = key
			maxItem = item
		}
	}
	return maxItem, true
}

// ToMap reduces this collection into a map using the provided key and value functions.
// @group Maps
// @behavior readonly
// @chainable false
// @terminal true
func (c *Collection[T]) ToMap[K comparable, V any](keyFn func(T) K, valueFn func(T) V) map[K]V {
	out := make(map[K]V, len(c.items))
	for _, item := range c.items {
		out[keyFn(item)] = valueFn(item)
	}
	return out
}

// GroupBy partitions this collection into collections keyed by the extracted value.
// @group Grouping
// @behavior readonly
// @chainable false
// @terminal true
func (c *Collection[T]) GroupBy[K comparable](keyFn func(T) K) map[K]*Collection[T] {
	out := make(map[K]*Collection[T], len(c.items))
	for _, item := range c.items {
		key := keyFn(item)
		group := out[key]
		if group == nil {
			out[key] = &Collection[T]{items: []T{item}}
			continue
		}
		group.items = append(group.items, item)
	}
	return out
}

// GroupBySlice partitions this collection into slices keyed by the extracted value.
// @group Grouping
// @behavior readonly
// @chainable false
// @terminal true
func (c *Collection[T]) GroupBySlice[K comparable](keyFn func(T) K) map[K][]T {
	out := make(map[K][]T)
	for _, item := range c.items {
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
func (c *Collection[T]) CountBy[K comparable](keyFn func(T) K) map[K]int {
	result := make(map[K]int)
	for _, item := range c.Items() {
		result[keyFn(item)]++
	}
	return result
}

// UniqueBy returns a collection containing the first item for each extracted key.
// @group Set Operations
// @behavior immutable
// @chainable true
// @terminal false
func (c *Collection[T]) UniqueBy[K comparable](keyFn func(T) K) *Collection[T] {
	items := c.items
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

// Pipe passes this collection to fn and returns fn's result.
// @group Transformation
// @behavior readonly
// @chainable false
// @terminal true
func (c *Collection[T]) Pipe[R any](fn func(*Collection[T]) R) R {
	return fn(c)
}

// ZipWith combines this collection with another collection using fn up to the shorter length.
// @group Transformation
// @behavior immutable
// @chainable true
// @terminal false
func (c *Collection[T]) ZipWith[U, R any](other *Collection[U], fn func(T, U) R) *Collection[R] {
	length := len(c.items)
	if len(other.items) < length {
		length = len(other.items)
	}

	out := make([]R, length)
	for i := 0; i < length; i++ {
		out[i] = fn(c.items[i], other.items[i])
	}
	return New(out)
}
