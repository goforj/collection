# Migrating to v3

Version 3 makes collections ordinary, named slices. The release also removes helpers that duplicate Go built-ins or the standard library, and gives mutating operations explicit names.

## Module path

Version 3 uses the semantic import version suffix `/v3`. Update both `go.mod` and imports.

### Before

```go
require github.com/goforj/collection v2.0.0+incompatible
```

```go
import "github.com/goforj/collection"
```

### After

```go
require github.com/goforj/collection/v3 v3.0.0
```

```go
import "github.com/goforj/collection/v3"
```

## Slice-backed collections

### Collection becomes Slice

`Collection[T]` and `*Collection[T]` are replaced by `Slice[T]`. A `Slice` supports `len`, indexing, slicing, and `range` directly.

### Before

```go
var values *collection.Collection[int] = collection.New([]int{10, 20, 30})
first := values.Items()[0]
```

### After

```go
var values collection.Slice[int] = collection.New([]int{10, 20, 30})
first := values[0]
```

### New is zero-copy

`New` converts the supplied slice to `Slice` without copying it. Mutations through either value are visible through the shared backing array. Call `Clone` when ownership must be independent.

### Before

```go
source := []int{1, 2, 3}
values := collection.New(source)
copy := values.ItemsCopy()
```

### After

```go
source := []int{1, 2, 3}
values := collection.New(source)
copy := values.Clone()
```

### Items is removed

Use the `Slice` value directly wherever a `[]T` is accepted. An explicit conversion is available when a concrete `[]T` value is required.

### Before

```go
values := collection.New([]int{1, 2, 3})
consume(values.Items())
```

### After

```go
values := collection.New([]int{1, 2, 3})
consume(values)
```

### ItemsCopy is removed

Use `Clone` to get an independent `Slice`, or `slices.Clone` when the result should be a built-in slice.

### Before

```go
copy := values.ItemsCopy()
```

### After

```go
copy := slices.Clone(values)
```

### NumericCollection is removed

Numeric operations now accept any slice-like value directly. `NewNumeric` and the `NumericCollection` wrapper are no longer needed.

### Before

```go
numbers := collection.NewNumeric([]int{1, 2, 3})
total := numbers.Sum()
average := numbers.Avg()
```

### After

```go
numbers := []int{1, 2, 3}
total := collection.Sum(numbers)
average := collection.Avg(numbers)
```

## Transformations

### Map is pure and generic

`Map` no longer mutates its receiver and is no longer restricted to `T` to `T`. It allocates a result and may change the element type.

### Before

```go
values := collection.New([]int{1, 2, 3})
mapped := values.Map(func(value int) int { return value * 2 })
// values is now [2 4 6]
```

### After

```go
values := collection.New([]int{1, 2, 3})
mapped := values.Map(func(value int) string { return strconv.Itoa(value) })
// values is still [1 2 3]
// mapped is ["1" "2" "3"]
```

### Transform is the in-place map

Use `Transform` for a same-type mutation. It returns the receiver so it remains chainable.

### Before

```go
values := collection.New([]int{1, 2, 3})
values.Map(func(value int) int { return value * 2 })
```

### After

```go
values := collection.New([]int{1, 2, 3})
values.Transform(func(value int) int { return value * 2 })
```

### MapTo is removed

Generic `Map` replaces both the `MapTo` method and free function.

### Before

```go
labels := collection.MapTo(values, func(value int) string {
	return strconv.Itoa(value)
})
```

### After

```go
labels := values.Map(func(value int) string {
	return strconv.Itoa(value)
})
```

### Filter is pure

`Filter` now allocates a result and leaves the receiver and its backing array unchanged.

### Before

```go
values := collection.New([]int{1, 2, 3, 4})
values.Filter(func(value int) bool { return value%2 == 0 })
// values is [2 4]
```

### After

```go
values := collection.New([]int{1, 2, 3, 4})
evens := values.Filter(func(value int) bool { return value%2 == 0 })
// values is still [1 2 3 4]
// evens is [2 4]
```

### Retain is the in-place filter

Use `Retain` to compact the existing backing array. Capture its return value because a slice value cannot update the caller's slice header.

### Before

```go
values := collection.New([]int{1, 2, 3, 4})
values.Filter(func(value int) bool { return value%2 == 0 })
```

### After

```go
values := collection.New([]int{1, 2, 3, 4})
values = values.Retain(func(value int) bool { return value%2 == 0 })
```

### Reduce has a generic accumulator

The accumulator may now differ from the element type.

### Before

```go
total := collection.New([]int{1, 2, 3}).Reduce(0, func(total, value int) int {
	return total + value
})
```

### After

```go
summary := collection.New([]int{1, 2, 3}).Reduce("", func(out string, value int) string {
	return out + strconv.Itoa(value)
})
```

## Built-in result boundaries

Operations that end a fluent collection step now return built-in maps and slices. Wrap a result with `New` only when another collection method is needed.

### GroupBy returns map values as built-in slices

`GroupBy` is a generic method returning `map[K][]T`.

### Before

```go
groups := collection.GroupBy(values, func(value int) int { return value % 2 })
evens := groups[0].Items()
```

### After

```go
groups := values.GroupBy(func(value int) int { return value % 2 })
evens := groups[0]
```

### GroupBySlice is removed

`GroupBy` now has the built-in slice result that `GroupBySlice` provided.

### Before

```go
groups := values.GroupBySlice(func(value int) int { return value % 2 })
```

### After

```go
groups := values.GroupBy(func(value int) int { return value % 2 })
```

### Partition returns built-in slices

### Before

```go
matched, unmatched := values.Partition(func(value int) bool { return value%2 == 0 })
consume(matched.Items())
```

### After

```go
matched, unmatched := values.Partition(func(value int) bool { return value%2 == 0 })
consume(matched)
```

### Chunk returns capacity-capped built-in slices

`Chunk` still returns `[][]T`, but each chunk is now a view whose capacity equals its length. Appending to one chunk cannot overwrite the next chunk in the source.

### Before

```go
chunks := values.Chunk(2)
chunks[0] = append(chunks[0], 99) // could overwrite the source's next item
```

### After

```go
chunks := values.Chunk(2)
chunks[0] = append(chunks[0], 99) // allocates instead of overwriting the next item
```

### Window is a method returning built-in slices

Each window is also capacity-capped to its length.

### Before

```go
windows := collection.Window(values, 3, 1)
first := windows.Items()[0]
```

### After

```go
windows := values.Window(3, 1)
first := windows[0]
```

## Views

The following methods still return zero-copy views. Their result capacity is now capped at the result length, so an append cannot overwrite elements just beyond the view. Use `Clone` when later element mutation must also be isolated.

### Take returns a capped view

### Before

```go
head := values.Take(2)
raw := head.Items()
raw = append(raw, 99) // could overwrite the source's next item
```

### After

```go
head := values.Take(2)
head = append(head, 99) // allocates because cap(head) is 2
```

### Take no longer accepts negative counts

Use `TakeLast` when selecting from the end.

### Before

```go
tail := values.Take(-2)
```

### After

```go
tail := values.TakeLast(2)
```

### TakeLast returns a capped view

### Before

```go
tail := values.TakeLast(2)
raw := tail.Items()
raw = append(raw, 99)
```

### After

```go
tail := values.TakeLast(2)
tail = append(tail, 99) // cannot overwrite storage beyond the view
```

### Skip returns a capped view

### Before

```go
tail := values.Skip(2)
raw := tail.Items()
raw = append(raw, 99) // could overwrite storage beyond the view
```

### After

```go
tail := values.Skip(2)
tail = append(tail, 99) // allocates instead of overwriting the source
```

### SkipLast returns a capped view

### Before

```go
head := values.SkipLast(2)
raw := head.Items()
raw = append(raw, 99) // could overwrite the source's skipped items
```

### After

```go
head := values.SkipLast(2)
head = append(head, 99) // allocates instead of overwriting the source
```

### After returns a capped view

### Before

```go
tail := values.After(func(value int) bool { return value == 2 })
raw := tail.Items()
raw = append(raw, 99)
```

### After

```go
tail := values.After(func(value int) bool { return value == 2 })
tail = append(tail, 99) // cannot overwrite storage beyond the view
```

### TakeUntil replaces TakeUntilFn

The predicate method is renamed and returns a capacity-capped view.

### Before

```go
prefix := values.TakeUntilFn(func(value int) bool { return value >= 3 })
```

### After

```go
prefix := values.TakeUntil(func(value int) bool { return value >= 3 })
```

### Before is replaced by TakeUntil

### Before

```go
prefix := values.Before(func(value int) bool { return value >= 3 })
```

### After

```go
prefix := values.TakeUntil(func(value int) bool { return value >= 3 })
```

### Comparable-value TakeUntil is removed

Use the predicate method for equality checks too.

### Before

```go
prefix := collection.TakeUntil(values, 3)
```

### After

```go
prefix := values.TakeUntil(func(value int) bool { return value == 3 })
```

## Zip and Pair

### Zip is a method

`Zip` moves from a free function to a generic method and accepts a built-in slice as its other input.

### Before

```go
pairs := collection.Zip(left, right)
```

### After

```go
pairs := left.Zip(right)
```

### Pair replaces Tuple

`Tuple[A, B]` is removed. `Zip` returns built-in `[]Pair[A, B]`, using `First` and `Second` fields.

### Before

```go
var pair collection.Tuple[int, string]
fmt.Println(pair.First, pair.Second)
```

### After

```go
var pair collection.Pair[int, string]
fmt.Println(pair.First, pair.Second)
```

### Pair no longer constrains its first type

Both `Pair` type parameters accept any type. Map keys remain comparable where a map operation requires it.

### Before

```go
// This did not compile because the first type had to be comparable.
// var pair collection.Pair[[]int, string]
```

### After

```go
var pair collection.Pair[[]int, string]
```

### FromMap uses First and Second

### Before

```go
pairs := collection.FromMap(map[string]int{"one": 1})
fmt.Println(pairs.Items()[0].Key, pairs.Items()[0].Value)
```

### After

```go
pairs := collection.FromMap(map[string]int{"one": 1})
fmt.Println(pairs[0].First, pairs[0].Second)
```

### ToMapKV is removed

Use the generic `ToMap` method and the new `Pair` fields.

### Before

```go
values := collection.ToMapKV(pairs)
```

### After

```go
values := pairs.ToMap(
	func(pair collection.Pair[string, int]) string { return pair.First },
	func(pair collection.Pair[string, int]) int { return pair.Second },
)
```

## Duplicate free functions

Generic methods make the following free-function forms unnecessary. The method forms preserve type inference and fluent calls.

### MinBy free function is removed

### Before

```go
shortest, ok := collection.MinBy(words, func(word string) int { return len(word) })
```

### After

```go
shortest, ok := words.MinBy(func(word string) int { return len(word) })
```

### MaxBy free function is removed

### Before

```go
longest, ok := collection.MaxBy(words, func(word string) int { return len(word) })
```

### After

```go
longest, ok := words.MaxBy(func(word string) int { return len(word) })
```

### ToMap free function is removed

### Before

```go
lengths := collection.ToMap(words,
	func(word string) string { return word },
	func(word string) int { return len(word) },
)
```

### After

```go
lengths := words.ToMap(
	func(word string) string { return word },
	func(word string) int { return len(word) },
)
```

### GroupBy free function is removed

### Before

```go
groups := collection.GroupBy(values, func(value int) int { return value % 2 })
```

### After

```go
groups := values.GroupBy(func(value int) int { return value % 2 })
```

### CountBy free function is removed

### Before

```go
counts := collection.CountBy(values, func(value int) int { return value % 2 })
```

### After

```go
counts := values.CountBy(func(value int) int { return value % 2 })
```

### CountByValue accepts slices directly

### Before

```go
counts := collection.CountByValue(collection.New(values))
```

### After

```go
counts := collection.CountByValue(values)
```

### UniqueBy free function is removed

### Before

```go
unique := collection.UniqueBy(words, func(word string) int { return len(word) })
```

### After

```go
unique := words.UniqueBy(func(word string) int { return len(word) })
```

### ZipWith free function is removed

### Before

```go
sums := collection.ZipWith(left, right, func(a, b int) int { return a + b })
```

### After

```go
sums := left.ZipWith(right, func(a, b int) int { return a + b })
```

## Native replacements

### Count is removed

### Before

```go
count := values.Count()
```

### After

```go
count := len(values)
```

### IsEmpty is removed

### Before

```go
empty := values.IsEmpty()
```

### After

```go
empty := len(values) == 0
```

### Contains is removed

Use `slices.Contains` for comparable elements or `slices.ContainsFunc` for a custom equality rule.

### Before

```go
found := collection.Contains(values, 3)
```

### After

```go
found := slices.Contains(values, 3)
```

### Append is removed

Use `Concat` when the v2 operation's independent backing storage matters. Use the built-in `append` when ordinary slice capacity reuse is acceptable, and capture its result.

### Before

```go
values = values.Append(4, 5)
```

### After

```go
values = values.Concat([]int{4, 5})
```

For native append semantics:

```go
values = append(values, 4, 5)
```

### Pop is removed

Use indexing and slicing. Clear the removed slot when `T` may retain resources.

### Before

```go
last, ok := values.Pop()
```

### After

```go
var last int
ok := len(values) != 0
if ok {
	last = values[len(values)-1]
	clear(values[len(values)-1:])
	values = values[:len(values)-1]
}
```

### PopN is removed

Use a capped tail view and shorten the source. Clone the tail first if it must be independent of the source's storage.

### Before

```go
popped := values.PopN(2)
```

### After

```go
n := min(2, len(values))
split := len(values) - n
popped := values[split:len(values):len(values)]
values = values[:split]
```

### ToJSON is removed

Use `encoding/json` directly. Convert the returned bytes to a string only when needed.

### Before

```go
text, err := values.ToJSON()
```

### After

```go
data, err := json.Marshal(values)
text := string(data)
```

### ToPrettyJSON is removed

### Before

```go
text, err := values.ToPrettyJSON()
```

### After

```go
data, err := json.MarshalIndent(values, "", "  ")
text := string(data)
```

### Pipe is removed

Both the method and free function are replaced by an ordinary function call.

### Before

```go
total := values.Pipe(func(input *collection.Collection[int]) int {
	return input.Reduce(0, func(sum, value int) int { return sum + value })
})
```

### After

```go
sum := func(input collection.Slice[int]) int {
	return input.Reduce(0, func(total, value int) int { return total + value })
}
total := sum(values)
```

## Concatenation and mutation

### Concat is always pure

`Concat` accepts one or more slices and always allocates independent backing storage for its result.

### Before

```go
combined := values.Concat([]int{4, 5})
// combined could reuse values' spare capacity
```

### After

```go
combined := values.Concat([]int{4, 5}, []int{6, 7})
// combined does not share backing storage with values
```

### Prepend is pure

`Prepend` now returns an independent result instead of changing the receiver. Capture the returned slice.

### Before

```go
values.Prepend(1, 2)
// values was changed
```

### After

```go
prefixed := values.Prepend(1, 2)
// values is unchanged
```

### Merge is removed

Use `Concat` for ordered slice concatenation. Handle keyed map merge rules explicitly before converting map data with `FromMap`.

### Before

```go
combined := values.Merge([]int{4, 5})
```

### After

```go
combined := values.Concat([]int{4, 5})
```

### Shuffle uses the package-level random source

`Shuffle` remains an in-place, chainable operation. It now uses the package-level `math/rand` source rather than a collection-owned source.

### Before

```go
values.Shuffle()
```

### After

```go
shuffled := values.Shuffle()
```

## Iterator interop

### Use slices.Values for iter.Seq

`Slice` works directly with the standard `slices` iterator adapters when an `iter.Seq[T]` is useful. `Each` remains available for fluent eager traversal.

### Before

```go
values.Each(func(value int) {
	fmt.Println(value)
})
```

### After

```go
sequence := slices.Values(values)
for value := range sequence {
	fmt.Println(value)
}
```

## Slice-like free functions

Numeric and comparable set operations now accept ordinary slices and named slice types directly. Their results remain `Slice` values where chaining is useful.

### Numeric operations accept slices

This applies to `Sum`, `Avg`, `Min`, `Max`, `Median`, and `Mode`.

### Before

```go
numbers := collection.NewNumeric([]int{1, 2, 3})
smallest, ok := numbers.Min()
```

### After

```go
numbers := []int{1, 2, 3}
smallest, ok := collection.Min(numbers)
```

### Set operations accept slices

This applies to `Union`, `Intersect`, `Difference`, and `SymmetricDifference`.

### Before

```go
combined := collection.Union(collection.New(left), collection.New(right))
```

### After

```go
combined := collection.Union(left, right)
```

### UniqueComparable accepts slices

### Before

```go
unique := collection.UniqueComparable(collection.New(values))
```

### After

```go
unique := collection.UniqueComparable(values)
```
