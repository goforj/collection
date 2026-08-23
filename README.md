<p align="center">
  <img src="https://raw.githubusercontent.com/goforj/collection/main/docs/assets/logo.png" width="300" alt="goforj/collection logo">
</p>

<p align="center">
  Fluent collections for Go. Iterate, filter, transform, sort, reduce, group, and debug data with a tiny dependency footprint.
</p>

<p align="center">
    <a href="https://pkg.go.dev/github.com/goforj/collection/v3"><img src="https://pkg.go.dev/badge/github.com/goforj/collection/v3.svg" alt="Go Reference"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License: MIT"></a>
    <a href="https://github.com/goforj/collection/actions"><img src="https://github.com/goforj/collection/actions/workflows/test.yml/badge.svg" alt="Go Test"></a>
    <a href="https://golang.org"><img src="https://img.shields.io/badge/go-1.27+-blue?logo=go" alt="Go version"></a>
    <img src="https://img.shields.io/github/v/tag/goforj/collection?label=version&sort=semver" alt="Latest tag">
    <a href="https://codecov.io/gh/goforj/collection" ><img src="https://codecov.io/github/goforj/collection/graph/badge.svg?token=3KFTK96U8C"/></a>
<!-- test-count:embed:start -->
    <img src="https://img.shields.io/badge/tests-357-brightgreen" alt="Tests">
<!-- test-count:embed:end -->
</p>

## Features

- **Fluent chaining** - pipeline your operations like Laravel Collections
- **Fully generic** (`Slice[T]`) - no reflection, no `interface{}`
- **Tiny dependency footprint** - only `godump` for debugging helpers
- **Go-native slice access** - use `len`, indexing, `range`, and ordinary slice conversions directly
- **Explicit allocation behavior** - slice views and independent results are documented per operation
- **Map / Filter / Reduce** - clean functional transforms
- **Generic methods** - type-changing transforms remain fluent on Go 1.27+
- **First / Last / FirstWhere / IndexWhere** helpers
- **Sort, GroupBy, Chunk**, and more
- **Borrow-by-default** - no defensive copies unless you ask for them
- **Standard-library interop** - use `slices`, `iter`, and `encoding/json` directly
- **Developer-friendly debug helpers** (`Dump()`, `Dd()`, `DumpStr()`)
- **Works with any Go type**, including structs, pointers, and deeply nested composites

## Fluent Chaining

Many methods return a `Slice`, allowing fluent method chaining without a wrapper object.

Some methods may be limited due to Go's generic constraints.

> **Fluent example:**  
> [`examples/chaining/main.go`](./examples/chaining/main.go)

```go
events := []DeviceEvent{
    {Device: "router-1", Region: "us-east", Errors: 3},
    {Device: "router-2", Region: "us-east", Errors: 15},
    {Device: "router-3", Region: "us-west", Errors: 22},
}

// Fluent slice pipeline
collection.
    New(events). // Construction
    Shuffle(). // Ordering
    Filter(func(e DeviceEvent) bool { return e.Errors > 5 }). // Slicing
    Sort(func(a, b DeviceEvent) bool { return a.Errors > b.Errors }). // Ordering
    Take(5). // Slicing
    TakeUntil(func(e DeviceEvent) bool { return e.Errors < 10 }). // Slicing (stop when predicate becomes true)
    SkipLast(1). // Slicing
    Dump() // Debugging

// #[]main.DeviceEvent [
//  0 => #main.DeviceEvent {
//    +Device => "router-3" #string
//    +Region => "us-west" #string
//    +Errors => 22 #int
//  }
// ]
```

Go 1.27 generic methods keep type-changing pipelines fluent:

```go
namesByLength := collection.New(users).
    Map(func(user User) string { return user.Name }).
    UniqueBy(strings.ToLower).
    GroupBy(func(name string) int { return len(name) })
```

### Performance Benchmarks

> **tl;dr**: *lo* is excellent. We solve a different problem - and in chained pipelines, that difference matters.

`lo` is a fantastic library and a major inspiration for this project. It is battle-tested, idiomatic, and often the right choice when you want small, standalone helpers that operate on slices in isolation.

`collection` takes a different approach.

Rather than treating each operation as an independent transformation, `collection` is built around **explicit, fluent pipelines** over a named slice. Operations document whether they return an independent allocation, a view, or the same backing storage.

That design choice doesn't matter much for some single operations. It matters a *lot* once you start chaining and especially in hot paths.

The below tables are automatically generated from [`./docs/bench/main.go`](./docs/bench/main.go).

Matched v2/v3 regression benchmarks for mutating, copied, and pipeline workloads live in [`./docs/regression`](./docs/regression).

<!-- bench:embed:start -->

Full raw tables: see `BENCHMARKS.md`.

#### Read-only scalar ops

| Op | Speed vs lo | Memory | Allocs |
|---:|:-----------:|:------:|:------:|
| **All** | ≈ | ≈ | ≈ |
| **Any** | ≈ | ≈ | ≈ |
| **None** | ≈ | ≈ | ≈ |
| **First** | ≈ | ≈ | ≈ |
| **Last** | ≈ | ≈ | ≈ |
| **FirstWhere** | same loop | ≈ | ≈ |
| **IndexWhere** | ≈ | ≈ | ≈ |
| **Contains** | ≈ | ≈ | ≈ |
| **Reduce (sum)** | ≈ | ≈ | ≈ |
| **Sum** | ≈ | ≈ | ≈ |
| **Min** | ≈ | ≈ | ≈ |
| **Max** | ≈ | ≈ | ≈ |
| **Each** | ≈ | ≈ | ≈ |

#### Transforming ops

| Op | Speed vs lo | Memory | Allocs |
|---:|:-----------:|:------:|:------:|
| **Chunk** | view trade-off | ownership trade-off | ownership trade-off |
| **Filter** | 0.89x | ≈ | ≈ |
| **Map** | ≈ | ≈ | ≈ |
| **Take** | ≈ | ≈ | ≈ |
| **Skip** | view trade-off | ownership trade-off | ownership trade-off |
| **SkipLast** | view trade-off | ownership trade-off | ownership trade-off |
| **Zip** | **1.56x** | ≈ | ≈ |
| **ZipWith** | **3.10x** | ≈ | ≈ |
| **Unique** | ≈ | ≈ | ≈ |
| **UniqueBy** | ≈ | ≈ | ≈ |
| **Union** | ≈ | ≈ | ≈ |
| **Intersect** | ≈ | ≈ | ≈ |
| **Difference** | different work | API trade-off | API trade-off |
| **GroupBy** | ≈ | ≈ | ≈ |
| **CountBy** | ≈ | ≈ | ≈ |
| **CountByValue** | ≈ | ≈ | ≈ |
| **ToMap** | ≈ | ≈ | ≈ |

#### Pipelines

| Op | Speed vs lo | Memory | Allocs |
|---:|:-----------:|:------:|:------:|
| **Pipeline F→M→T→R** | 0.88x | ≈ | ≈ |

#### Mutating ops

| Op | Speed vs lo | Memory | Allocs |
|---:|:-----------:|:------:|:------:|
| **Retain** | 0.89x | ≈ | ≈ |
| **Reverse** | ≈ | ≈ | ≈ |
| **Shuffle** | **3.63x** | ≈ | ≈ |
| **Transform** | ≈ | ≈ | ≈ |
<!-- bench:embed:end -->

## How to read the benchmarks

- **≈** means the two libraries are effectively equivalent
- **same loop** means both implementations compile to the same machine loop, so binary-placement skew is not presented as a library difference
- Explicit memory deltas show allocation differences for equivalent work; ownership and API trade-offs are labeled separately
- Single-operation helpers are expected to be close when they perform equivalent work
- Multi-step pipelines show the cost of the selected ownership model

If you prefer immutable, one-off helpers - `lo` is outstanding.
If you write **expressive, chained data pipelines** and care about hot-path performance - `collection` is built for that job.


## Choosing pure or in-place pipelines

Most functional helpers (including `lo`) operate like this:

```
input -> Map -> new slice -> Filter -> new slice -> Take -> slice view
```

That model is simple and safe. Transforming stages such as `Map` and `Filter`
typically allocate; `Subset`, used here for `Take`, returns a view.

Version 3 gives the same pure behavior through fluent generic methods. Hot paths
can opt into the explicitly named in-place operations instead:

```
input -> Retain (compact in place) -> Transform (in place) -> Take (slice view)
```

The representation stays Go-native throughout: every `Slice` supports `len`,
indexing, slicing, and `range`. `Map` and `Filter` allocate independent results.
`Transform` mutates elements in place, while `Retain` compacts and clears the
existing backing array; its shortened return value must be captured. `Concat`
and `Prepend` always return independent results. `Sort`, `Reverse`, and
`Shuffle` mutate elements in place.

- Pure operations make ownership easy to reason about
- In-place operations avoid allocations when the caller owns the input
- Capped views prevent append from overwriting elements beyond the view
- `Clone` makes an intentional ownership boundary before mutation

The benchmark tables compare equivalent pure operations separately from
`Retain` and `Transform`. `Chunk`, `Skip`, and `SkipLast` are ownership
trade-offs because collection returns capacity-capped views while lo returns
copied slices.

## Explicit branching with `Clone`

Fluent pipelines don't mean you're locked into mutation.

`New` borrows slices by default. Use `Clone()` before an in-place operation when
the original data must remain unchanged.

When you want to branch a pipeline or preserve the original data, `Clone()` creates a shallow copy of the collection so subsequent operations are isolated and predictable.

```go
events := collection.New(deviceEvents)

// Fast alerting path: cheap filters, early exit
alerts := events.
    Clone().
    Retain(func(e DeviceEvent) bool { return e.Severity >= Critical }).
    Take(10)

// Deeper analysis path: heavier work, full ordering
report := events.
    Clone().
    Filter(func(e DeviceEvent) bool { return e.Region == "us-east" }).
    Sort(func(a, b DeviceEvent) bool { return a.Timestamp.Before(b.Timestamp) })
```

This makes divergence points explicit and intentional.

No hidden copies. No surprises.

## Design Principles

- **Type-safe**: no reflection
- **Explicit semantics**: order, mutation, and allocation are documented
- **Go-native**: respects generics and stdlib patterns
- **Eager evaluation**: no lazy pipelines or hidden concurrency
- **Maps are boundaries**: unordered data is handled explicitly

## What this library is not

- Not a lazy or streaming library
- Not concurrency-aware
- Not immutable-by-default
- Not a replacement for idiomatic loops in simple cases
- Not designed to hide allocation, mutation, or ordering semantics

## Working with maps

Maps are unordered in Go. This library does not pretend otherwise.

Instead, map interaction is explicit and intentional:

- `FromMap` materializes key/value pairs into an ordered workflow
- `ToMap` reduces collections back into maps explicitly

This makes transitions between unordered and ordered data visible and honest.

## Behavior semantics

Each method declares how it interacts with the collection:

- **readonly** - reads data only and returns a derived value
- **immutable** - returns a value without mutating the receiver; individual method docs state whether it allocates or returns a view
- **mutable** - may modify elements in the receiver's backing array
- **terminal** - ends the fluent pipeline and returns a non-collection result

These annotations describe **observable behavior**, not implementation details.

Terminal operations do not return a `Slice` and cannot be chained further.
They are designed to be allocation-free under `New()` where possible.

Allocation and copying are **explicitly documented per method**.
Some readonly or immutable operations may allocate internally when required
(e.g. grouping, chunking, scratch copies), but never mutate the receiver.

Borrowed slices, independent results, in-place element mutation, and view semantics are intentional and visible.

## Native slice interoperability

`Slice[T]` is a named slice, so ordinary Go operations work directly:

```go
values := collection.New([]int{10, 20, 30})

fmt.Println(len(values))
// 3
fmt.Println(values[1])
// 20

for _, value := range values {
    fmt.Println(value)
}
// 10
// 20
// 30
```

Use `slices.Values(values)` or `slices.All(values)` when an iterator is useful;
no collection-specific lazy wrapper is required.

## Runnable examples

Every function has a corresponding runnable example under [`./examples`](./examples).

These examples are **generated directly from the documentation blocks** of each function, ensuring the docs and code never drift. These are the same examples you see here in the README and GoDoc.

An automated test executes **every example** to verify it builds and runs successfully.  

This guarantees all examples are valid, up-to-date, and remain functional as the API evolves.

# Installation

This package requires Go 1.27 or newer. Consumers that cannot upgrade their
toolchain can remain on v2.

Existing users should read [Migrating to v3](./MIGRATING_TO_V3.md) for the complete API and ownership changes.

```bash
go get github.com/goforj/collection/v3
```

<!-- api:embed:start -->

# API Index

| Group | Functions |
|------:|-----------|
| **Aggregation** | [Avg](#avg) - [CountBy](#countby) - [CountByValue](#countbyvalue) - [Max](#max) - [MaxBy](#maxby) - [Median](#median) - [Min](#min) - [MinBy](#minby) - [Mode](#mode) - [Reduce](#reduce) - [Sum](#sum) |
| **Construction** | [Clone](#clone) - [New](#new) |
| **Debugging** | [Dd](#dd) - [Dump](#dump) - [DumpStr](#dumpstr) - [Slice.Dump](#slice.dump) |
| **Grouping** | [GroupBy](#groupby) |
| **Maps** | [FromMap](#frommap) - [ToMap](#tomap) |
| **Ordering** | [After](#after) - [Reverse](#reverse) - [Shuffle](#shuffle) - [Sort](#sort) |
| **Querying** | [All](#all) - [Any](#any) - [At](#at) - [First](#first) - [FirstWhere](#firstwhere) - [IndexWhere](#indexwhere) - [Last](#last) - [LastWhere](#lastwhere) - [None](#none) |
| **Set Operations** | [Difference](#difference) - [Intersect](#intersect) - [SymmetricDifference](#symmetricdifference) - [Union](#union) - [Unique](#unique) - [UniqueBy](#uniqueby) - [UniqueComparable](#uniquecomparable) |
| **Slicing** | [Chunk](#chunk) - [Filter](#filter) - [Partition](#partition) - [Retain](#retain) - [Skip](#skip) - [SkipLast](#skiplast) - [Take](#take) - [TakeLast](#takelast) - [TakeUntil](#takeuntil) - [Window](#window) |
| **Transformation** | [Concat](#concat) - [Each](#each) - [Map](#map) - [Multiply](#multiply) - [Prepend](#prepend) - [Tap](#tap) - [Times](#times) - [Transform](#transform) - [Zip](#zip) - [ZipWith](#zipwith) |


## Aggregation

### <a id="avg"></a>Avg - readonly - terminal

Avg returns the average of the numeric slice values as a float64.
If the slice is empty, Avg returns 0.

_Example: integers_

```go
values := []int{2, 4, 6}
collection.Dump(collection.Avg(values))
// 4.000000 #float64
```

_Example: float_

```go
values2 := []float64{1.5, 2.5, 3.0}
collection.Dump(collection.Avg(values2))
// 2.333333 #float64
```

### <a id="countby"></a>CountBy - readonly - terminal

CountBy returns occurrence counts keyed by the extracted value.

```go
numbers := collection.New([]int{1, 2, 3, 5})
counts := numbers.CountBy(func(number int) string {
	if number%2 == 0 {
		return "even"
	}
	return "odd"
})
collection.Dump(counts)
// #map[string]int {
//   even => 1 #int
//   odd => 3 #int
// }
```

### <a id="countbyvalue"></a>CountByValue - readonly - terminal

CountByValue returns the number of occurrences of each distinct item in c.

```go
values := []string{"go", "forj", "go"}
counts := collection.CountByValue(values)
collection.Dump(counts)
// #map[string]int {
//   forj => 1 #int
//   go => 2 #int
// }
```

### <a id="max"></a>Max - readonly - terminal

Max returns the largest item in a numeric slice.
The second return value is false if the slice is empty.

_Example: integers_

```go
values := []int{3, 1, 2}

max1, ok1 := collection.Max(values)
collection.Dump(max1, ok1)
// 3 #int
// true #bool
```

_Example: floats_

```go
values2 := []float64{1.5, 9.2, 4.4}

max2, ok2 := collection.Max(values2)
collection.Dump(max2, ok2)
// 9.200000 #float64
// true #bool
```

_Example: empty numeric slice_

```go
empty := []int{}

max3, ok3 := collection.Max(empty)
collection.Dump(max3, ok3)
// 0 #int
// false #bool
```

### <a id="maxby"></a>MaxBy - readonly - terminal

MaxBy returns the item whose extracted key is the largest.

```go
words := collection.New([]string{"pear", "fig", "banana"})
longest, ok := words.MaxBy(func(word string) int {
	return len(word)
})
collection.Dump(longest, ok)
// "banana" #string
// true #bool
```

### <a id="median"></a>Median - readonly - terminal

Median returns the statistical median of a numeric slice as float64.
It returns (0, false) if the slice is empty.

_Example: integers - odd number of items_

```go
values := []int{3, 1, 2}

median1, ok1 := collection.Median(values)
collection.Dump(median1, ok1)
// 2.000000 #float64
// true #bool
```

_Example: integers - even number of items_

```go
values2 := []int{10, 2, 4, 6}

median2, ok2 := collection.Median(values2)
collection.Dump(median2, ok2)
// 5.000000 #float64
// true #bool
```

_Example: floats_

```go
values3 := []float64{1.1, 9.9, 3.3}

median3, ok3 := collection.Median(values3)
collection.Dump(median3, ok3)
// 3.300000 #float64
// true #bool
```

_Example: integers - empty numeric slice_

```go
empty := []int{}

median4, ok4 := collection.Median(empty)
collection.Dump(median4, ok4)
// 0.000000 #float64
// false #bool
```

### <a id="min"></a>Min - readonly - terminal

Min returns the smallest item in a numeric slice.
The second return value is false if the slice is empty.

_Example: integers_

```go
values := []int{3, 1, 2}
min, ok := collection.Min(values)
collection.Dump(min, ok)
// 1 #int
// true #bool
```

_Example: floats_

```go
values2 := []float64{2.5, 9.1, 1.2}
min2, ok2 := collection.Min(values2)
collection.Dump(min2, ok2)
// 1.200000 #float64
// true #bool
```

_Example: integers - empty collection_

```go
empty := []int{}
min3, ok3 := collection.Min(empty)
collection.Dump(min3, ok3)
// 0 #int
// false #bool
```

### <a id="minby"></a>MinBy - readonly - terminal

MinBy returns the item whose extracted key is the smallest.

```go
words := collection.New([]string{"pear", "fig", "banana"})
shortest, ok := words.MinBy(func(word string) int {
	return len(word)
})
collection.Dump(shortest, ok)
// "fig" #string
// true #bool
```

### <a id="mode"></a>Mode - readonly - terminal

Mode returns the most frequent numeric value or values in a slice.
If multiple values tie for highest frequency, all are returned
in first-seen order.

_Example: integers - single mode_

```go
values := []int{1, 2, 2, 3}
mode := collection.Mode(values)
collection.Dump(mode)
// #[]int [
//   0 => 2 #int
// ]
```

_Example: integers - tie for mode_

```go
values2 := []int{1, 2, 1, 2}
mode2 := collection.Mode(values2)
collection.Dump(mode2)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
// ]
```

_Example: floats_

```go
values3 := []float64{1.1, 2.2, 1.1, 3.3}
mode3 := collection.Mode(values3)
collection.Dump(mode3)
// #[]float64 [
//   0 => 1.100000 #float64
// ]
```

_Example: integers - empty collection_

```go
empty := []int{}
mode4 := collection.Mode(empty)
collection.Dump(mode4)
// []int(nil)
```

### <a id="reduce"></a>Reduce - readonly - terminal

Reduce collapses the collection into a single accumulated value.
The accumulator may have a different type R from the collection's elements.

_Example: integers - sum_

```go
sum := collection.New([]int{1, 2, 3}).Reduce(0, func(acc, n int) int {
	return acc + n
})
collection.Dump(sum)
// 6 #int
```

_Example: strings_

```go
joined := collection.New([]string{"a", "b", "c"}).Reduce("", func(acc, s string) string {
	return acc + s
})
collection.Dump(joined)
// "abc" #string
```

_Example: structs_

```go
type Stats struct {
	Count int
	Sum   int
}

stats := collection.New([]Stats{
	{Count: 1, Sum: 10},
	{Count: 1, Sum: 20},
	{Count: 1, Sum: 30},
})

total := stats.Reduce(Stats{}, func(acc, s Stats) Stats {
	acc.Count += s.Count
	acc.Sum += s.Sum
	return acc
})

collection.Dump(total)
// #main.Stats {
//   +Count => 3 #int
//   +Sum   => 60 #int
// }
```

### <a id="sum"></a>Sum - readonly - terminal

Sum returns the sum of all items in a numeric slice.
If the slice is empty, Sum returns the zero value of T.

_Example: integers_

```go
values := []int{1, 2, 3}
total := collection.Sum(values)
collection.Dump(total)
// 6 #int
```

_Example: floats_

```go
values2 := []float64{1.5, 2.5}
total2 := collection.Sum(values2)
collection.Dump(total2)
// 4.000000 #float64
```

_Example: integers - empty collection_

```go
empty := []int{}
total3 := collection.Sum(empty)
collection.Dump(total3)
// 0 #int
```

## Construction

### <a id="clone"></a>Clone - immutable - chainable

Clone returns a copy of the collection.

The returned collection has its own backing slice, so subsequent mutations
do not affect the original collection.

Clone is intended to be used when branching a pipeline while preserving
the original collection.

_Example: basic cloning_

```go
c := collection.New([]int{1, 2, 3})
clone := c.Clone()

clone.Transform(func(value int) int { return value * 10 })

collection.Dump(c)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
// ]

collection.Dump(clone)
// #[]int [
//   0 => 10 #int
//   1 => 20 #int
//   2 => 30 #int
// ]
```

_Example: branching pipelines_

```go
base := collection.New([]int{1, 2, 3, 4, 5})

evens := base.Clone().Retain(func(v int) bool {
	return v%2 == 0
})

odds := base.Clone().Retain(func(v int) bool {
	return v%2 != 0
})

collection.Dump(base)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
//   3 => 4 #int
//   4 => 5 #int
// ]

collection.Dump(evens)
// #[]int [
//   0 => 2 #int
//   1 => 4 #int
// ]

collection.Dump(odds)
// #[]int [
//   0 => 1 #int
//   1 => 3 #int
//   2 => 5 #int
// ]
```

### <a id="new"></a>New - immutable - chainable

New creates a Slice from items and borrows their backing array.

```go
values := collection.New([]int{10, 20, 30})
fmt.Println(len(values))
// 3
fmt.Println(values[1])
// 20

total := 0
for _, value := range values {
	total += value
}
fmt.Println(total)
// 60
```

## Debugging

### <a id="dd"></a>Dd - readonly - terminal

Dd prints items then terminates execution.
Like Laravel's dd(), this is intended for debugging and
should not be used in production control flow.

```go
c := collection.New([]string{"a", "b"})
c.Dd()
// #[]string [
//   0 => "a" #string
//   1 => "b" #string
// ]
// Process finished with the exit code 1
```

### <a id="dump"></a>Dump - readonly - terminal

Dump is a convenience function that calls godump.Dump.

```go
c2 := collection.New([]int{1, 2, 3})
collection.Dump(c2)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
// ]
```

### <a id="dumpstr"></a>DumpStr - readonly - terminal

DumpStr returns the pretty-printed dump of the items as a string,
without printing or exiting.
Useful for logging, snapshot testing, and non-interactive debugging.

```go
c := collection.New([]int{10, 20})
s := c.DumpStr()
fmt.Println(s)
// #[]int [
//   0 => 10 #int
//   1 => 20 #int
// ]
```

### <a id="slice.dump"></a>Slice.Dump - readonly - chainable

Dump prints items with godump and returns the same collection.
This is a no-op on the collection itself and never panics.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3})
c.Dump()
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
// ]
```

_Example: integers - chaining_

```go
collection.New([]int{1, 2, 3}).
	Filter(func(v int) bool { return v > 1 }).
	Dump()
// #[]int [
//   0 => 2 #int
//   1 => 3 #int
// ]
```

## Grouping

### <a id="groupby"></a>GroupBy - readonly - terminal

GroupBy partitions this Slice into independent built-in slices keyed by the
extracted value.

```go
numbers := collection.New([]int{1, 2, 3, 4})
groups := numbers.GroupBy(func(number int) string {
	if number%2 == 0 {
		return "even"
	}
	return "odd"
})
collection.Dump(groups["even"], groups["odd"])
// #[]int [
//   0 => 2 #int
//   1 => 4 #int
// ]
// #[]int [
//   0 => 1 #int
//   1 => 3 #int
// ]
fmt.Println(len(groups["even"]))
// 2
fmt.Println(groups["odd"][0])
// 1
collection.Dump(groups["even"][:1])
// #[]int [
//   0 => 2 #int
// ]
```

## Maps

### <a id="frommap"></a>FromMap - immutable - chainable

FromMap materializes a map into a collection of key/value pairs.

_Example: basic usage_

```go
m := map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
}

c := collection.FromMap(m)
c.Sort(func(a, b collection.Pair[string, int]) bool {
	return a.First < b.First
})
collection.Dump(c)
// #[]collection.Pair[string,int] [
//   0 => #collection.Pair[string,int] {
//     +First  => "a" #string
//     +Second => 1 #int
//   }
//   1 => #collection.Pair[string,int] {
//     +First  => "b" #string
//     +Second => 2 #int
//   }
//   2 => #collection.Pair[string,int] {
//     +First  => "c" #string
//     +Second => 3 #int
//   }
// ]
```

_Example: filtering map entries_

```go
type Config struct {
	Enabled bool
	Timeout int
}

configs := map[string]Config{
	"router-1": {Enabled: true, Timeout: 30},
	"router-2": {Enabled: false, Timeout: 10},
	"router-3": {Enabled: true, Timeout: 45},
}

out := collection.
	FromMap(configs).
	Filter(func(p collection.Pair[string, Config]) bool {
		return p.Second.Enabled
	}).
	Sort(func(a, b collection.Pair[string, Config]) bool {
		return a.First < b.First
	})

collection.Dump(out)
// #[]collection.Pair[string,main.Config·1] [
//   0 => #collection.Pair[string,main.Config·1] {
//     +First     => "router-1" #string
//     +Second    => #main.Config {
//       +Enabled => true #bool
//       +Timeout => 30 #int
//     }
//   }
//   1 => #collection.Pair[string,main.Config·1] {
//     +First     => "router-3" #string
//     +Second    => #main.Config {
//       +Enabled => true #bool
//       +Timeout => 45 #int
//     }
//   }
// ]
```

### <a id="tomap"></a>ToMap - readonly - terminal

ToMap reduces this collection into a map using the provided key and value functions.

```go
words := collection.New([]string{"go", "forj"})
lengths := words.ToMap(
	func(word string) string { return word },
	func(word string) int { return len(word) },
)
collection.Dump(lengths)
// #map[string]int {
//   forj => 4 #int
//   go => 2 #int
// }
```

## Ordering

### <a id="after"></a>After - immutable - chainable

After returns all items after the first element for which pred returns true.
If no element matches, an empty collection is returned.

```go
c := collection.New([]int{1, 2, 3, 4, 5})
c.After(func(v int) bool { return v == 3 }).Dump()
// #[]int [
//  0 => 4 #int
//  1 => 5 #int
// ]
```

### <a id="reverse"></a>Reverse - mutable - chainable

Reverse reverses the order of items in the collection in place
and returns the same collection for chaining.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4})
c.Reverse()
collection.Dump(c)
// #[]int [
//   0 => 4 #int
//   1 => 3 #int
//   2 => 2 #int
//   3 => 1 #int
// ]
```

_Example: strings - chaining_

```go
out := collection.New([]string{"a", "b", "c"}).
	Reverse().
	Concat([]string{"d"})

collection.Dump(out)
// #[]string [
//   0 => "c" #string
//   1 => "b" #string
//   2 => "a" #string
//   3 => "d" #string
// ]
```

_Example: structs_

```go
type User struct {
	ID int
}

users := collection.New([]User{
	{ID: 1},
	{ID: 2},
	{ID: 3},
})

users.Reverse()
collection.Dump(users)
// #[]main.User [
//   0 => #main.User {
//     +ID => 3 #int
//   }
//   1 => #main.User {
//     +ID => 2 #int
//   }
//   2 => #main.User {
//     +ID => 1 #int
//   }
// ]
```

### <a id="shuffle"></a>Shuffle - mutable - chainable

Shuffle shuffles the collection in place and returns the same collection.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4, 5})
c.Shuffle()
fmt.Println(len(c), collection.Sum(c))
// 5 15
```

_Example: strings - chaining_

```go
out2 := collection.New([]string{"a", "b", "c"}).
	Shuffle().
	Concat([]string{"d"})

fmt.Println(len(out2))
// 4
```

_Example: structs_

```go
type User struct {
	ID int
}

users := collection.New([]User{
	{ID: 1},
	{ID: 2},
	{ID: 3},
	{ID: 4},
})

users.Shuffle()
fmt.Println(len(users))
// 4
```

### <a id="sort"></a>Sort - mutable - chainable

Sort sorts the collection in place using the provided comparison function and
returns the same collection for chaining.

_Example: integers_

```go
c := collection.New([]int{5, 1, 4, 2})
c.Sort(func(a, b int) bool { return a < b })
collection.Dump(c)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 4 #int
//   3 => 5 #int
// ]
```

_Example: strings (descending)_

```go
c2 := collection.New([]string{"apple", "banana", "cherry"})
c2.Sort(func(a, b string) bool { return a > b })
collection.Dump(c2)
// #[]string [
//   0 => "cherry" #string
//   1 => "banana" #string
//   2 => "apple" #string
// ]
```

_Example: structs_

```go
type User struct {
	Name string
	Age  int
}

users := collection.New([]User{
	{Name: "Alice", Age: 30},
	{Name: "Bob", Age: 25},
	{Name: "Carol", Age: 40},
})

// Sort by age ascending
users.Sort(func(a, b User) bool {
	return a.Age < b.Age
})
collection.Dump(users)
// #[]main.User [
//   0 => #main.User {
//     +Name => "Bob" #string
//     +Age  => 25 #int
//   }
//   1 => #main.User {
//     +Name => "Alice" #string
//     +Age  => 30 #int
//   }
//   2 => #main.User {
//     +Name => "Carol" #string
//     +Age  => 40 #int
//   }
// ]
```

## Querying

### <a id="all"></a>All - readonly - terminal

All returns true if fn returns true for every item in the collection.
If the collection is empty, All returns true (vacuously true).

_Example: integers - all even_

```go
c := collection.New([]int{2, 4, 6})
allEven := c.All(func(v int) bool { return v%2 == 0 })
collection.Dump(allEven)
// true #bool
```

_Example: integers - not all even_

```go
c2 := collection.New([]int{2, 3, 4})
allEven2 := c2.All(func(v int) bool { return v%2 == 0 })
collection.Dump(allEven2)
// false #bool
```

_Example: strings - all non-empty_

```go
c3 := collection.New([]string{"a", "b", "c"})
allNonEmpty := c3.All(func(s string) bool { return s != "" })
collection.Dump(allNonEmpty)
// true #bool
```

_Example: empty collection (vacuously true)_

```go
empty := collection.New([]int{})
all := empty.All(func(v int) bool { return v > 0 })
collection.Dump(all)
// true #bool
```

### <a id="any"></a>Any - readonly - terminal

Any returns true if at least one item satisfies fn.

```go
c := collection.New([]int{1, 2, 3, 4})
has := c.Any(func(v int) bool { return v%2 == 0 }) // true
collection.Dump(has)
// true #bool
```

### <a id="at"></a>At - readonly - terminal

At returns the item at the given index and a boolean indicating
whether the index was within bounds.

_Example: integers_

```go
c := collection.New([]int{10, 20, 30})
v, ok := c.At(1)
collection.Dump(v, ok)
// 20 #int
// true #bool
```

_Example: out of bounds_

```go
v2, ok2 := c.At(10)
collection.Dump(v2, ok2)
// 0 #int
// false #bool
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
})

u, ok3 := users.At(0)
collection.Dump(u, ok3)
// #main.User {
//   +ID   => 1 #int
//   +Name => "Alice" #string
// }
// true #bool
```

### <a id="first"></a>First - readonly - terminal

First returns the first element in the collection.
If the collection is empty, ok will be false.

_Example: integers_

```go
c := collection.New([]int{10, 20, 30})

v, ok := c.First()
collection.Dump(v, ok)
// 10 #int
// true #bool
```

_Example: strings_

```go
c2 := collection.New([]string{"alpha", "beta", "gamma"})

v2, ok2 := c2.First()
collection.Dump(v2, ok2)
// "alpha" #string
// true #bool
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
})

u, ok3 := users.First()
collection.Dump(u, ok3)
// #main.User {
//   +ID   => 1 #int
//   +Name => "Alice" #string
// }
// true #bool
```

_Example: integers - empty collection_

```go
c3 := collection.New([]int{})
v3, ok4 := c3.First()
collection.Dump(v3, ok4)
// 0 #int
// false #bool
```

### <a id="firstwhere"></a>FirstWhere - readonly - terminal

FirstWhere returns the first item in the collection for which the provided
predicate function returns true. If no items match, ok=false is returned
along with the zero value of T.

```go
nums := collection.New([]int{1, 2, 3, 4, 5})
v, ok := nums.FirstWhere(func(n int) bool {
	return n%2 == 0
})
collection.Dump(v, ok)
// 2 #int
// true #bool

v, ok = nums.FirstWhere(func(n int) bool {
	return n > 10
})
collection.Dump(v, ok)
// 0 #int
// false #bool
```

### <a id="indexwhere"></a>IndexWhere - readonly - terminal

IndexWhere returns the index of the first item in the collection
for which the provided predicate function returns true.
If no item matches, it returns (0, false).

_Example: integers_

```go
c := collection.New([]int{10, 20, 30, 40})
idx, ok := c.IndexWhere(func(v int) bool { return v == 30 })
collection.Dump(idx, ok)
// 2 #int
// true #bool
```

_Example: not found_

```go
idx2, ok2 := c.IndexWhere(func(v int) bool { return v == 99 })
collection.Dump(idx2, ok2)
// 0 #int
// false #bool
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Carol"},
})

idx3, ok3 := users.IndexWhere(func(u User) bool {
	return u.Name == "Bob"
})

collection.Dump(idx3, ok3)
// 1 #int
// true #bool
```

### <a id="last"></a>Last - readonly - terminal

Last returns the last element in the collection.
If the collection is empty, ok will be false.

_Example: integers_

```go
c := collection.New([]int{10, 20, 30})

v, ok := c.Last()
collection.Dump(v, ok)
// 30 #int
// true #bool
```

_Example: strings_

```go
c2 := collection.New([]string{"alpha", "beta", "gamma"})

v2, ok2 := c2.Last()
collection.Dump(v2, ok2)
// "gamma" #string
// true #bool
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Charlie"},
})

u, ok3 := users.Last()
collection.Dump(u, ok3)
// #main.User {
//   +ID   => 3 #int
//   +Name => "Charlie" #string
// }
// true #bool
```

_Example: empty collection_

```go
c3 := collection.New([]int{})

v3, ok4 := c3.Last()
collection.Dump(v3, ok4)
// 0 #int
// false #bool
```

### <a id="lastwhere"></a>LastWhere - readonly - terminal

LastWhere returns the last element in the collection that satisfies the predicate fn.
If fn is nil, LastWhere returns the final element in the underlying slice.
If the collection is empty or no element matches, ok will be false.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4})

v, ok := c.LastWhere(func(v int, i int) bool {
	return v < 3
})
collection.Dump(v, ok)
// 2 #int
// true #bool
```

_Example: integers without predicate (equivalent to Last())_

```go
c2 := collection.New([]int{10, 20, 30, 40})

v2, ok2 := c2.LastWhere(nil)
collection.Dump(v2, ok2)
// 40 #int
// true #bool
```

_Example: strings_

```go
c3 := collection.New([]string{"alpha", "beta", "gamma", "delta"})

v3, ok3 := c3.LastWhere(func(s string, i int) bool {
	return strings.HasPrefix(s, "g")
})
collection.Dump(v3, ok3)
// "gamma" #string
// true #bool
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Alex"},
	{ID: 4, Name: "Brian"},
})

u, ok4 := users.LastWhere(func(u User, i int) bool {
	return strings.HasPrefix(u.Name, "A")
})
collection.Dump(u, ok4)
// #main.User {
//   +ID   => 3 #int
//   +Name => "Alex" #string
// }
// true #bool
```

_Example: no matching element_

```go
c4 := collection.New([]int{5, 6, 7})

v4, ok5 := c4.LastWhere(func(v int, i int) bool {
	return v > 10
})
collection.Dump(v4, ok5)
// 0 #int
// false #bool
```

_Example: empty collection_

```go
c5 := collection.New([]int{})

v5, ok6 := c5.LastWhere(nil)
collection.Dump(v5, ok6)
// 0 #int
// false #bool
```

### <a id="none"></a>None - readonly - terminal

None returns true if fn returns false for every item in the collection.
If the collection is empty, None returns true.

_Example: integers - none even_

```go
c := collection.New([]int{1, 3, 5})
noneEven := c.None(func(v int) bool { return v%2 == 0 })
collection.Dump(noneEven)
// true #bool
```

_Example: integers - some even_

```go
c2 := collection.New([]int{1, 2, 3})
noneEven2 := c2.None(func(v int) bool { return v%2 == 0 })
collection.Dump(noneEven2)
// false #bool
```

_Example: empty collection_

```go
empty := collection.New([]int{})
none := empty.None(func(v int) bool { return v > 0 })
collection.Dump(none)
// true #bool
```

## Set Operations

### <a id="difference"></a>Difference - immutable - chainable

Difference returns a new collection containing elements from the first collection
that are not present in the second. Order follows the first collection, and
duplicates are removed.

_Example: integers_

```go
a := collection.New([]int{1, 2, 2, 3, 4})
b := collection.New([]int{2, 4})

out := collection.Difference(a, b)
collection.Dump(out)
// #[]int [
//   0 => 1 #int
//   1 => 3 #int
// ]
```

_Example: strings_

```go
left := collection.New([]string{"apple", "banana", "cherry"})
right := collection.New([]string{"banana"})

out2 := collection.Difference(left, right)
collection.Dump(out2)
// #[]string [
//   0 => "apple" #string
//   1 => "cherry" #string
// ]
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

groupA := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Carol"},
})

groupB := collection.New([]User{
	{ID: 2, Name: "Bob"},
})

out3 := collection.Difference(groupA, groupB)
collection.Dump(out3)
// #[]main.User [
//   0 => #main.User {
//     +ID   => 1 #int
//     +Name => "Alice" #string
//   }
//   1 => #main.User {
//     +ID   => 3 #int
//     +Name => "Carol" #string
//   }
// ]
```

### <a id="intersect"></a>Intersect - immutable - chainable

Intersect returns a new collection containing elements from the second
collection that are also present in the first.

_Example: integers_

```go
a := collection.New([]int{1, 2, 2, 3, 4})
b := collection.New([]int{2, 4, 4, 5})

out := collection.Intersect(a, b)
collection.Dump(out)
// #[]int [
//   0 => 2 #int
//   1 => 4 #int
//   2 => 4 #int
// ]
```

_Example: strings_

```go
left := collection.New([]string{"apple", "banana", "cherry"})
right := collection.New([]string{"banana", "date", "cherry", "banana"})

out2 := collection.Intersect(left, right)
collection.Dump(out2)
// #[]string [
//   0 => "banana" #string
//   1 => "cherry" #string
//   2 => "banana" #string
// ]
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

groupA := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Carol"},
})

groupB := collection.New([]User{
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Carol"},
	{ID: 4, Name: "Dave"},
})

out3 := collection.Intersect(groupA, groupB)
collection.Dump(out3)
// #[]main.User [
//   0 => #main.User {
//     +ID   => 2 #int
//     +Name => "Bob" #string
//   }
//   1 => #main.User {
//     +ID   => 3 #int
//     +Name => "Carol" #string
//   }
// ]
```

### <a id="symmetricdifference"></a>SymmetricDifference - immutable - chainable

SymmetricDifference returns a new collection containing elements that appear
in exactly one of the two collections. Order follows the first collection for
its unique items, then the second for its unique items. Duplicates are removed.

_Example: integers_

```go
a := collection.New([]int{1, 2, 3, 3})
b := collection.New([]int{3, 4, 4, 5})

out := collection.SymmetricDifference(a, b)
collection.Dump(out)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 4 #int
//   3 => 5 #int
// ]
```

_Example: strings_

```go
left := collection.New([]string{"apple", "banana"})
right := collection.New([]string{"banana", "date"})

out2 := collection.SymmetricDifference(left, right)
collection.Dump(out2)
// #[]string [
//   0 => "apple" #string
//   1 => "date" #string
// ]
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

groupA := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
})

groupB := collection.New([]User{
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Carol"},
})

out3 := collection.SymmetricDifference(groupA, groupB)
collection.Dump(out3)
// #[]main.User [
//   0 => #main.User {
//     +ID   => 1 #int
//     +Name => "Alice" #string
//   }
//   1 => #main.User {
//     +ID   => 3 #int
//     +Name => "Carol" #string
//   }
// ]
```

### <a id="union"></a>Union - immutable - chainable

Union returns a new collection containing the unique elements from both collections.
Items from the first collection are kept in order, followed by items from the second
that were not already present.

_Example: integers_

```go
a := collection.New([]int{1, 2, 2, 3})
b := collection.New([]int{3, 4, 4, 5})

out := collection.Union(a, b)
collection.Dump(out)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
//   3 => 4 #int
//   4 => 5 #int
// ]
```

_Example: strings_

```go
left := collection.New([]string{"apple", "banana"})
right := collection.New([]string{"banana", "date"})

out2 := collection.Union(left, right)
collection.Dump(out2)
// #[]string [
//   0 => "apple" #string
//   1 => "banana" #string
//   2 => "date" #string
// ]
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

groupA := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
})

groupB := collection.New([]User{
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Carol"},
})

out3 := collection.Union(groupA, groupB)
collection.Dump(out3)
// #[]main.User [
//   0 => #main.User {
//     +ID   => 1 #int
//     +Name => "Alice" #string
//   }
//   1 => #main.User {
//     +ID   => 2 #int
//     +Name => "Bob" #string
//   }
//   2 => #main.User {
//     +ID   => 3 #int
//     +Name => "Carol" #string
//   }
// ]
```

### <a id="unique"></a>Unique - immutable - chainable

Unique returns a new collection with duplicate items removed, based on the
equality function `eq`. The first occurrence of each unique value is kept,
and order is preserved.

_Example: integers_

```go
c1 := collection.New([]int{1, 2, 2, 3, 4, 4, 5})
out1 := c1.Unique(func(a, b int) bool { return a == b })
collection.Dump(out1)
// #[]int [
//	0 => 1 #int
//	1 => 2 #int
//	2 => 3 #int
//	3 => 4 #int
//	4 => 5 #int
// ]
```

_Example: strings (case-insensitive uniqueness)_

```go
c2 := collection.New([]string{"A", "a", "B", "b", "A"})
out2 := c2.Unique(func(a, b string) bool {
	return strings.EqualFold(a, b)
})
collection.Dump(out2)
// #[]string [
//	0 => "A" #string
//	1 => "B" #string
// ]
```

_Example: structs (unique by ID)_

```go
type User struct {
	ID   int
	Name string
}

c3 := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 1, Name: "Alice Duplicate"},
})

out3 := c3.Unique(func(a, b User) bool {
	return a.ID == b.ID
})

collection.Dump(out3)
// #[]main.User [
//  0 => #main.User {
//    +ID   => 1 #int
//    +Name => "Alice" #string
//  }
//  1 => #main.User {
//    +ID   => 2 #int
//    +Name => "Bob" #string
//  }
// ]
```

### <a id="uniqueby"></a>UniqueBy - immutable - chainable

UniqueBy returns a collection containing the first item for each extracted key.

```go
words := collection.New([]string{"go", "up", "forj", "code"})
unique := words.UniqueBy(func(word string) int {
	return len(word)
})
collection.Dump(unique)
// #[]string [
//   0 => "go" #string
//   1 => "forj" #string
// ]
```

### <a id="uniquecomparable"></a>UniqueComparable - immutable - chainable

UniqueComparable returns a new collection with duplicate comparable items removed.
The first occurrence of each value is kept, and order is preserved.
This is a faster, allocation-friendly path for comparable types.

_Example: integers_

```go
c := collection.New([]int{1, 2, 2, 3, 4, 4, 5})
out := collection.UniqueComparable(c)
collection.Dump(out)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
//   3 => 4 #int
//   4 => 5 #int
// ]
```

_Example: strings_

```go
c2 := collection.New([]string{"A", "a", "B", "B"})
out2 := collection.UniqueComparable(c2)
collection.Dump(out2)
// #[]string [
//   0 => "A" #string
//   1 => "a" #string
//   2 => "B" #string
// ]
```

## Slicing

### <a id="chunk"></a>Chunk - readonly - terminal

Chunk splits the collection into chunks of the given size.
The final chunk may be smaller if len(items) is not divisible by size.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4, 5}).Chunk(2)
collection.Dump(c)
// #[][]int [
//  0 => #[]int [
//    0 => 1 #int
//    1 => 2 #int
//  ]
//  1 => #[]int [
//    0 => 3 #int
//    1 => 4 #int
//  ]
//  2 => #[]int [
//    0 => 5 #int
//  ]
//]
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

users := []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Carol"},
	{ID: 4, Name: "Dave"},
}

userChunks := collection.New(users).Chunk(2)
collection.Dump(userChunks)
// #[][]main.User [
//  0 => #[]main.User [
//    0 => #main.User {
//      +ID   => 1 #int
//      +Name => "Alice" #string
//    }
//    1 => #main.User {
//      +ID   => 2 #int
//      +Name => "Bob" #string
//    }
//  ]
//  1 => #[]main.User [
//    0 => #main.User {
//      +ID   => 3 #int
//      +Name => "Carol" #string
//    }
//    1 => #main.User {
//      +ID   => 4 #int
//      +Name => "Dave" #string
//    }
//  ]
//]
```

### <a id="filter"></a>Filter - immutable - chainable

Filter keeps only the elements for which fn returns true.

Filter allocates a new Slice and leaves c and its backing storage unchanged.

_Example: integers_

```go
source := collection.New([]int{1, 2, 3, 4})
filtered := source.Filter(func(v int) bool {
	return v%2 == 0
})
collection.Dump(filtered)
// #[]int [
//   0 => 2 #int
//   1 => 4 #int
// ]
fmt.Println(source[0])
// 1
```

_Example: strings_

```go
c2 := collection.New([]string{"apple", "banana", "cherry", "avocado"})
c2 = c2.Filter(func(v string) bool {
	return strings.HasPrefix(v, "a")
})
collection.Dump(c2)
// #[]string [
//   0 => "apple" #string
//   1 => "avocado" #string
// ]
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Andrew"},
	{ID: 4, Name: "Carol"},
})

users = users.Filter(func(u User) bool {
	return strings.HasPrefix(u.Name, "A")
})

collection.Dump(users)
// #[]main.User [
//   0 => #main.User {
//     +ID   => 1 #int
//     +Name => "Alice" #string
//   }
//   1 => #main.User {
//     +ID   => 3 #int
//     +Name => "Andrew" #string
//   }
// ]
```

### <a id="partition"></a>Partition - immutable - terminal

Partition splits the collection into two new slices based on predicate fn.
The first slice contains items where fn returns true; the second contains
items where fn returns false. Order is preserved within each partition.

_Example: integers - even/odd_

```go
nums := collection.New([]int{1, 2, 3, 4, 5})
evens, odds := nums.Partition(func(n int) bool {
	return n%2 == 0
})
collection.Dump(evens, odds)
// #[]int [
//   0 => 2 #int
//   1 => 4 #int
// ]
// #[]int [
//   0 => 1 #int
//   1 => 3 #int
//   2 => 5 #int
// ]
```

_Example: strings - prefix match_

```go
words := collection.New([]string{"go", "gopher", "rust", "ruby"})
goWords, other := words.Partition(func(s string) bool {
	return strings.HasPrefix(s, "go")
})
collection.Dump(goWords, other)
// #[]string [
//   0 => "go" #string
//   1 => "gopher" #string
// ]
// #[]string [
//   0 => "rust" #string
//   1 => "ruby" #string
// ]
```

_Example: structs - active vs inactive_

```go
type User struct {
	Name   string
	Active bool
}

users := collection.New([]User{
	{Name: "Alice", Active: true},
	{Name: "Bob", Active: false},
	{Name: "Carol", Active: true},
})

active, inactive := users.Partition(func(u User) bool {
	return u.Active
})

collection.Dump(active, inactive)
// #[]main.User [
//   0 => #main.User {
//     +Name   => "Alice" #string
//     +Active => true #bool
//   }
//   1 => #main.User {
//     +Name   => "Carol" #string
//     +Active => true #bool
//   }
// ]
// #[]main.User [
//   0 => #main.User {
//     +Name   => "Bob" #string
//     +Active => false #bool
//   }
// ]
```

### <a id="retain"></a>Retain - mutable - chainable

Retain keeps items for which fn returns true in c's existing backing array.

```go
values := collection.New([]int{1, 2, 3, 4})
evens := values.Retain(func(value int) bool { return value%2 == 0 })
collection.Dump(evens)
// #[]int [
//   0 => 2 #int
//   1 => 4 #int
// ]
fmt.Println(values)
// [2 4 0 0]
```

### <a id="skip"></a>Skip - immutable - chainable

Skip returns a new collection with the first n items skipped.
If n is less than or equal to zero, Skip returns the full collection.
If n is greater than or equal to the collection length, Skip returns
an empty collection.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4, 5})
out := c.Skip(2)
collection.Dump(out)
// #[]int [
//   0 => 3 #int
//   1 => 4 #int
//   2 => 5 #int
// ]
```

_Example: skip none_

```go
out2 := c.Skip(0)
collection.Dump(out2)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
//   3 => 4 #int
//   4 => 5 #int
// ]
```

_Example: skip all_

```go
out3 := c.Skip(10)
collection.Dump(out3)
// #[]int [
// ]
```

_Example: structs_

```go
type User struct {
	ID int
}

users := collection.New([]User{
	{ID: 1},
	{ID: 2},
	{ID: 3},
})

out4 := users.Skip(1)
collection.Dump(out4)
// #[]main.User [
//  0 => #main.User {
//    +ID => 2 #int
//  }
//  1 => #main.User {
//    +ID => 3 #int
//  }
// ]
```

### <a id="skiplast"></a>SkipLast - immutable - chainable

SkipLast returns a new collection with the last n items skipped.
If n is less than or equal to zero, SkipLast returns the full collection.
If n is greater than or equal to the collection length, SkipLast returns
an empty collection.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4, 5})
out := c.SkipLast(2)
collection.Dump(out)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
// ]
```

_Example: skip none_

```go
out2 := c.SkipLast(0)
collection.Dump(out2)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
//   3 => 4 #int
//   4 => 5 #int
// ]
```

_Example: skip all_

```go
out3 := c.SkipLast(10)
collection.Dump(out3)
// #[]int [
// ]
```

_Example: structs_

```go
type User struct {
	ID int
}

users := collection.New([]User{
	{ID: 1},
	{ID: 2},
	{ID: 3},
})

out4 := users.SkipLast(1)
collection.Dump(out4)
// #[]main.User [
//  0 => #main.User {
//    +ID => 1 #int
//  }
//  1 => #main.User {
//    +ID => 2 #int
//  }
// ]
```

### <a id="take"></a>Take - immutable - chainable

Take returns a capacity-capped view containing the first n items.

If n exceeds the collection length, the entire collection is returned.
If n == 0, an empty collection is returned.

NOTE: returns a view (shares backing array). Use Clone() to detach.

_Example: integers - take first 3_

```go
c1 := collection.New([]int{0, 1, 2, 3, 4, 5})
out1 := c1.Take(3)
collection.Dump(out1)
// #[]int [
//	0 => 0 #int
//	1 => 1 #int
//	2 => 2 #int
// ]
```

_Example: integers - n exceeds length → whole collection_

```go
c3 := collection.New([]int{10, 20})
out3 := c3.Take(10)
collection.Dump(out3)
// #[]int [
//	0 => 10 #int
//	1 => 20 #int
// ]
```

_Example: integers - zero → empty_

```go
c4 := collection.New([]int{1, 2, 3})
out4 := c4.Take(0)
collection.Dump(out4)
// #[]int [
// ]
```

### <a id="takelast"></a>TakeLast - immutable - chainable

TakeLast returns a capacity-capped view containing the last n items.
If n is less than or equal to zero, TakeLast returns an empty collection.
If n is greater than or equal to the collection length, TakeLast returns
the full collection.

This operation performs no element allocations; it re-slices the
underlying slice.

NOTE: returns a view (shares backing array). Use Clone() to detach.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4, 5})
out := c.TakeLast(2)
collection.Dump(out)
// #[]int [
//   0 => 4 #int
//   1 => 5 #int
// ]
```

_Example: take none_

```go
out2 := c.TakeLast(0)
collection.Dump(out2)
// #[]int [
// ]
```

_Example: take all_

```go
out3 := c.TakeLast(10)
collection.Dump(out3)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
//   3 => 4 #int
//   4 => 5 #int
// ]
```

_Example: structs_

```go
type User struct {
	ID int
}

users := collection.New([]User{
	{ID: 1},
	{ID: 2},
	{ID: 3},
})

out4 := users.TakeLast(1)
collection.Dump(out4)
// #[]main.User [
//  0 => #main.User {
//    +ID => 3 #int
//  }
// ]
```

### <a id="takeuntil"></a>TakeUntil - immutable - chainable

TakeUntil returns items until the predicate function returns true.
The matching item is NOT included.

_Example: integers - stop when value >= 3_

```go
c1 := collection.New([]int{1, 2, 3, 4})
out1 := c1.TakeUntil(func(v int) bool { return v >= 3 })
collection.Dump(out1)
// #[]int [
//	0 => 1 #int
//	1 => 2 #int
// ]
```

_Example: integers - predicate immediately true → empty result_

```go
c2 := collection.New([]int{10, 20, 30})
out2 := c2.TakeUntil(func(v int) bool { return v < 50 })
collection.Dump(out2)
// #[]int [
// ]
```

_Example: integers - no match → full list returned_

```go
c3 := collection.New([]int{1, 2, 3})
out3 := c3.TakeUntil(func(v int) bool { return v == 99 })
collection.Dump(out3)
// #[]int [
//	0 => 1 #int
//	1 => 2 #int
//	2 => 3 #int
// ]
```

### <a id="window"></a>Window - readonly - terminal

Window returns overlapping (or stepped) windows of the collection.
Each window is a slice of length size; iteration advances by step (default 1 if step <= 0).
Windows that are shorter than size are omitted.

_Example: integers - step 1_

```go
nums := collection.New([]int{1, 2, 3, 4, 5})
win := nums.Window(3, 1)
collection.Dump(win)
// #[][]int [
//   0 => #[]int [
//     0 => 1 #int
//     1 => 2 #int
//     2 => 3 #int
//   ]
//   1 => #[]int [
//     0 => 2 #int
//     1 => 3 #int
//     2 => 4 #int
//   ]
//   2 => #[]int [
//     0 => 3 #int
//     1 => 4 #int
//     2 => 5 #int
//   ]
// ]
```

_Example: strings - step 2_

```go
words := collection.New([]string{"a", "b", "c", "d", "e"})
win2 := words.Window(2, 2)
collection.Dump(win2)
// #[][]string [
//   0 => #[]string [
//     0 => "a" #string
//     1 => "b" #string
//   ]
//   1 => #[]string [
//     0 => "c" #string
//     1 => "d" #string
//   ]
// ]
```

_Example: structs_

```go
type Point struct {
	X int
	Y int
}

points := collection.New([]Point{
	{X: 0, Y: 0},
	{X: 1, Y: 1},
	{X: 2, Y: 4},
	{X: 3, Y: 9},
})

win3 := points.Window(2, 1)
collection.Dump(win3)
// #[][]main.Point [
//   0 => #[]main.Point [
//     0 => #main.Point {
//       +X => 0 #int
//       +Y => 0 #int
//     }
//     1 => #main.Point {
//       +X => 1 #int
//       +Y => 1 #int
//     }
//   ]
//   1 => #[]main.Point [
//     0 => #main.Point {
//       +X => 1 #int
//       +Y => 1 #int
//     }
//     1 => #main.Point {
//       +X => 2 #int
//       +Y => 4 #int
//     }
//   ]
//   2 => #[]main.Point [
//     0 => #main.Point {
//       +X => 2 #int
//       +Y => 4 #int
//     }
//     1 => #main.Point {
//       +X => 3 #int
//       +Y => 9 #int
//     }
//   ]
// ]
```

## Transformation

### <a id="concat"></a>Concat - immutable - chainable

Concat returns an independent collection containing c followed by values.

_Example: strings_

```go
c := collection.New([]string{"John Doe"})
concatenated := c.
	Concat([]string{"Jane Doe"}).
	Concat([]string{"Johnny Doe"})
collection.Dump(concatenated)
// #[]string [
//  0 => "John Doe" #string
//  1 => "Jane Doe" #string
//  2 => "Johnny Doe" #string
// ]
```

_Example: spare capacity_

```go
backing := make([]int, 2, 4)
copy(backing, []int{1, 2})
values := collection.New(backing)
values = values.Concat([]int{3, 4})
fmt.Println(values)
// [1 2 3 4]
```

### <a id="each"></a>Each - readonly - chainable

Each runs fn for every item in the collection and returns the same collection,
so it can be used in chains for side effects (logging, debugging, etc.).

_Example: integers_

```go
c := collection.New([]int{1, 2, 3})

sum := 0
c.Each(func(v int) {
	sum += v
})

collection.Dump(sum)
// 6 #int
```

_Example: strings_

```go
c2 := collection.New([]string{"apple", "banana", "cherry"})

var out []string
c2.Each(func(s string) {
	out = append(out, strings.ToUpper(s))
})

collection.Dump(out)
// #[]string [
//   0 => "APPLE" #string
//   1 => "BANANA" #string
//   2 => "CHERRY" #string
// ]
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Charlie"},
})

var names []string
users.Each(func(u User) {
	names = append(names, u.Name)
})

collection.Dump(names)
// #[]string [
//   0 => "Alice" #string
//   1 => "Bob" #string
//   2 => "Charlie" #string
// ]
```

### <a id="map"></a>Map - immutable - chainable

Map maps this Slice to a newly allocated Slice with a potentially different
element type.

```go
numbers := collection.New([]int{1, 2, 3, 4})
labels := numbers.Map(func(number int) string {
	if number%2 == 0 {
		return "even"
	}
	return "odd"
})
collection.Dump(labels)
// #[]string [
//   0 => "odd" #string
//   1 => "even" #string
//   2 => "odd" #string
//   3 => "even" #string
// ]
fmt.Println(numbers[0])
// 1
```

### <a id="multiply"></a>Multiply - immutable - chainable

Multiply creates `n` copies of all items in the collection
and returns a new collection.

_Example: integers_

```go
ints := collection.New([]int{1, 2})
out := ints.Multiply(3)
collection.Dump(out)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 1 #int
//   3 => 2 #int
//   4 => 1 #int
//   5 => 2 #int
// ]
```

_Example: strings_

```go
strs := collection.New([]string{"a", "b"})
out2 := strs.Multiply(2)
collection.Dump(out2)
// #[]string [
//   0 => "a" #string
//   1 => "b" #string
//   2 => "a" #string
//   3 => "b" #string
// ]
```

_Example: structs_

```go
type User struct {
	Name string
}

users := collection.New([]User{{Name: "Alice"}, {Name: "Bob"}})
out3 := users.Multiply(2)
collection.Dump(out3)
// #[]main.User [
//   0 => #main.User {
//     +Name => "Alice" #string
//   }
//   1 => #main.User {
//     +Name => "Bob" #string
//   }
//   2 => #main.User {
//     +Name => "Alice" #string
//   }
//   3 => #main.User {
//     +Name => "Bob" #string
//   }
// ]
```

_Example: multiplying by zero or negative returns empty_

```go
none := ints.Multiply(0)
collection.Dump(none)
// #[]int [
// ]
```

### <a id="prepend"></a>Prepend - immutable - chainable

Prepend returns an independently backed Slice containing values followed by c.

_Example: integers_

```go
c := collection.New([]int{3, 4})
result := c.Prepend(1, 2)
collection.Dump(result)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
//   3 => 4 #int
// ]
```

_Example: strings_

```go
letters := collection.New([]string{"c", "d"})
result2 := letters.Prepend("a", "b")
collection.Dump(result2)
// #[]string [
//   0 => "a" #string
//   1 => "b" #string
//   2 => "c" #string
//   3 => "d" #string
// ]
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 2, Name: "Bob"},
})

result3 := users.Prepend(User{ID: 1, Name: "Alice"})
collection.Dump(result3)
// #[]main.User [
//   0 => #main.User {
//     +ID   => 1 #int
//     +Name => "Alice" #string
//   }
//   1 => #main.User {
//     +ID   => 2 #int
//     +Name => "Bob" #string
//   }
// ]
```

_Example: integers - Prepending into an empty collection_

```go
empty := collection.New([]int{})
result4 := empty.Prepend(9, 8)
collection.Dump(result4)
// #[]int [
//   0 => 9 #int
//   1 => 8 #int
// ]
```

_Example: integers - Prepending no values → no change_

```go
c2 := collection.New([]int{1, 2})
result5 := c2.Prepend()
collection.Dump(result5)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
// ]
```

### <a id="tap"></a>Tap - mutable - chainable

Tap invokes fn with the Slice value for side effects such as logging,
debugging, or inspection, then returns the Slice to allow chaining.

_Example: integers - capture intermediate state during a chain_

```go
captured1 := []int{}
c1 := collection.New([]int{3, 1, 2}).
	Sort(func(a, b int) bool { return a < b }). // → [1, 2, 3]
	Tap(func(col collection.Slice[int]) {
		captured1 = append([]int(nil), col...) // snapshot copy
	}).
	Filter(func(v int) bool { return v >= 2 }).
	Dump()
	// #[]int [
	//  0 => 2 #int
	//  1 => 3 #int
	// ]

// Use BOTH variables so nothing is "declared and not used"
collection.Dump(c1)
collection.Dump(captured1)
// #[]int [
//  0 => 2 #int
//  1 => 3 #int
// ]
// #[]int [
//  0 => 1 #int
//  1 => 2 #int
//  2 => 3 #int
// ]
```

_Example: integers - tap for debugging without changing flow_

```go
c2 := collection.New([]int{10, 20, 30}).
	Tap(func(col collection.Slice[int]) {
		collection.Dump(col)
		// #[]int [
		//  0 => 10 #int
		//  1 => 20 #int
		//  2 => 30 #int
		// ]
	}).
	Filter(func(v int) bool { return v > 10 })

collection.Dump(c2) // ensures c2 is used
// #[]int [
//  0 => 20 #int
//  1 => 30 #int
// ]
```

_Example: structs - Tap with struct collection_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
})

users2 := users.Tap(func(col collection.Slice[User]) {
	collection.Dump(col)
	// #[]main.User [
	//  0 => #main.User {
	//    +ID   => 1 #int
	//    +Name => "Alice" #string
	//  }
	//  1 => #main.User {
	//    +ID   => 2 #int
	//    +Name => "Bob" #string
	//  }
	// ]
})

collection.Dump(users2) // ensures users2 is used
// #[]main.User [
//  0 => #main.User {
//    +ID   => 1 #int
//    +Name => "Alice" #string
//  }
//  1 => #main.User {
//    +ID   => 2 #int
//    +Name => "Bob" #string
//  }
// ]
```

### <a id="times"></a>Times - immutable - chainable

Times creates a new collection by calling fn(i) for i = 1..count.
This mirrors Laravel's Collection::times(), which is 1-indexed.

_Example: integers - double each index_

```go
cTimes1 := collection.Times(5, func(i int) int {
	return i * 2
})
collection.Dump(cTimes1)
// #[]int [
//	0 => 2 #int
//	1 => 4 #int
//	2 => 6 #int
//	3 => 8 #int
//	4 => 10 #int
// ]
```

_Example: strings_

```go
cTimes2 := collection.Times(3, func(i int) string {
	return fmt.Sprintf("item-%d", i)
})
collection.Dump(cTimes2)
// #[]string [
//	0 => "item-1" #string
//	1 => "item-2" #string
//	2 => "item-3" #string
// ]
```

_Example: structs_

```go
type Point struct {
	X int
	Y int
}

cTimes3 := collection.Times(4, func(i int) Point {
	return Point{X: i, Y: i * i}
})
collection.Dump(cTimes3)
// #[]main.Point [
//	0 => #main.Point {
//		+X => 1 #int
//		+Y => 1 #int
//	}
//	1 => #main.Point {
//		+X => 2 #int
//		+Y => 4 #int
//	}
//	2 => #main.Point {
//		+X => 3 #int
//		+Y => 9 #int
//	}
//	3 => #main.Point {
//		+X => 4 #int
//		+Y => 16 #int
//	}
// ]
```

### <a id="transform"></a>Transform - mutable - chainable

Transform applies a same-type transformation in place and returns the same collection.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3})

c.Transform(func(v int) int {
	return v * 10
})

collection.Dump(c)
// #[]int [
//   0 => 10 #int
//   1 => 20 #int
//   2 => 30 #int
// ]
```

_Example: strings_

```go
c2 := collection.New([]string{"apple", "banana", "cherry"})

upper := c2.Transform(func(s string) string {
	return strings.ToUpper(s)
})

collection.Dump(upper)
// #[]string [
//   0 => "APPLE" #string
//   1 => "BANANA" #string
//   2 => "CHERRY" #string
// ]
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
})

updated := users.Transform(func(u User) User {
	u.Name = strings.ToUpper(u.Name)
	return u
})

collection.Dump(updated)
// #[]main.User [
//   0 => #main.User {
//     +ID   => 1 #int
//     +Name => "ALICE" #string
//   }
//   1 => #main.User {
//     +ID   => 2 #int
//     +Name => "BOB" #string
//   }
// ]
```

### <a id="zip"></a>Zip - immutable - terminal

Zip combines this collection with values element-wise into pairs.
The resulting length is the smaller of the two inputs.

```go
nums := collection.New([]int{1, 2, 3})
words := []string{"one", "two"}

out := nums.Zip(words)
collection.Dump(out)
// #[]collection.Pair[int,string] [
//   0 => #collection.Pair[int,string] {
//     +First  => 1 #int
//     +Second => "one" #string
//   }
//   1 => #collection.Pair[int,string] {
//     +First  => 2 #int
//     +Second => "two" #string
//   }
// ]
```

### <a id="zipwith"></a>ZipWith - immutable - chainable

ZipWith combines this collection with a slice using fn up to the shorter length.

```go
left := collection.New([]int{1, 2, 3})
right := collection.New([]int{10, 20})
sums := left.ZipWith(right, func(a, b int) int {
	return a + b
})
collection.Dump(sums)
// #[]int [
//   0 => 11 #int
//   1 => 22 #int
// ]
```
<!-- api:embed:end -->

## Development

Use `make test` for the root module, `make vet` for static checks, and `make generate` to refresh the generated README API reference. The `docs` and `examples` directories are separate Go modules and can be tested from their own directories when changed.
