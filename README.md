<p align="center">
  <img src="./docs/assets/logo.png" width="600" alt="goforj/collection logo">
</p>

<p align="center">
    Fluent, Laravel-style Collections for Go - with generics, chainable pipelines, and expressive data transforms.
</p>

<p align="center">
    <a href="https://pkg.go.dev/github.com/goforj/collection"><img src="https://pkg.go.dev/badge/github.com/goforj/collection.svg" alt="Go Reference"></a>
    <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License: MIT"></a>
    <a href="https://github.com/goforj/collection/actions"><img src="https://github.com/goforj/collection/actions/workflows/test.yml/badge.svg" alt="Go Test"></a>
    <a href="https://golang.org"><img src="https://img.shields.io/badge/go-1.21+-blue?logo=go" alt="Go version"></a>
    <img src="https://img.shields.io/github/v/tag/goforj/collection?label=version&sort=semver" alt="Latest tag">
    <a href="https://codecov.io/gh/goforj/collection" ><img src="https://codecov.io/github/goforj/collection/graph/badge.svg?token=3KFTK96U8C"/></a>
<!-- test-count:embed:start -->
    <img src="https://img.shields.io/badge/tests-440-brightgreen" alt="Tests">
<!-- test-count:embed:end -->
    <a href="https://goreportcard.com/report/github.com/goforj/collection"><img src="https://goreportcard.com/badge/github.com/goforj/collection" alt="Go Report Card"></a>
</p>

<p align="center">
  <code>collection</code> brings an expressive, fluent API to Go.  
  Iterate, filter, transform, sort, reduce, group, and debug your data with zero dependencies.  
  Designed to feel natural to Go developers - and luxurious to everyone else.
</p>

# Features

- **Fluent chaining** - pipeline your operations like Laravel Collections
- **Fully generic** (`Collection[T]`) - no reflection, no `interface{}`
- **Minimal dependencies** - small footprint (godump for debugging)
- **Minimal allocations** - views for slice selections; in-place ops reuse backing when possible
- **Map / Filter / Reduce** - clean functional transforms
- **First / Last / FirstWhere / IndexWhere / Contains** helpers
- **Sort, GroupBy, Chunk**, and more
- **Safe-by-default** - defensive copies where appropriate
- **Built-in JSON helpers** (`ToJSON()`, `ToPrettyJSON()`)
- **Developer-friendly debug helpers** (`Dump()`, `Dd()`, `DumpStr()`)
- **Works with any Go type**, including structs, pointers, and deeply nested composites

## Fluent Chaining

Many methods return the collection itself, allowing for fluent method chaining.

Some methods maybe limited to due to go's generic constraints. 

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
    TakeUntilFn(func(e DeviceEvent) bool { return e.Errors < 10 }). // Slicing (stop when predicate becomes true)
    SkipLast(1). // Slicing
    Dump() // Debugging

// []main.DeviceEvent [
//  0 => #main.DeviceEvent {
//    +Device => "router-3" #string
//    +Region => "us-west" #string
//    +Errors => 22 #int
//  }
// ]
```

<!-- bench:embed:start -->

### Performance Benchmarks

> **tl;dr**: *lo* is excellent. We solve a different problem - and in chained pipelines, that difference matters.

`lo` is a fantastic library and a major inspiration for this project. It is battle-tested, idiomatic, and often the right choice when you want small, standalone helpers that operate on slices in isolation.

`collection` takes a different approach.

Rather than treating each operation as an independent transformation, `collection` is built around **explicit, fluent pipelines**. Many operations are designed to **mutate the same backing slice intentionally**, allowing chained workflows to avoid intermediate allocations and unnecessary copying - while still making that behavior visible and documented.

That design choice doesn't matter much for some single operations. It matters a *lot* once you start chaining and especially in hot paths.

| Op | ns/op (vs lo) | × | bytes/op (vs lo) | × | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 945ns / 278ns | 0.29x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 812ns / 281ns | 0.35x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 772ns / 1.0µs | **1.35x** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 777ns / 237ns | 0.31x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 8.7µs / 8.5µs | ≈ | 17.5KB / 9.4KB | 0.53x more | 12 / 11 |
| **CountByValue** | 8.6µs / 8.4µs | ≈ | 17.5KB / 9.4KB | 0.53x more | 12 / 11 |
| **Difference** | 20.1µs / 43.3µs | **2.16x** | 98.5KB / 108.8KB | **1.10x less** | 14 / 43 |
| **Each** | 898ns / 281ns | 0.31x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 1.1µs / 1.1µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 653ns / 0ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 793ns / 238ns | 0.30x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBySlice** | 8.7µs / 8.4µs | ≈ | 29.2KB / 21.0KB | 0.72x more | 84 / 83 |
| **IndexWhere** | 821ns / 287ns | 0.35x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 11.8µs / 10.7µs | 0.90x | 27.8KB / 11.4KB | 0.41x more | 22 / 19 |
| **Last** | 656ns / 0ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 1.5µs / 807ns | 0.52x | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **Max** | 780ns / 229ns | 0.29x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 792ns / 229ns | 0.29x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 800ns / 254ns | 0.32x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 1.4µs / 1.4µs | ≈ | 12.3KB / 12.3KB | ≈ | 3 / 2 |
| **Reduce (sum)** | 773ns / 231ns | 0.30x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 785ns / 232ns | 0.30x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 4.4µs / 5.3µs | **1.21x** | 16.4KB / 0B | ∞x more | 3 / 0 |
| **Skip** | 657ns / 696ns | **1.06x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 657ns / 703ns | **1.07x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 777ns / 229ns | 0.29x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 677ns / 0ns | 0.00x | 8.2KB / 0B | ∞x more | 2 / 0 |
| **ToMap** | 8.2µs / 7.9µs | ≈ | 45.1KB / 37.0KB | 0.82x more | 6 / 6 |
| **Union** | 18.1µs / 17.7µs | ≈ | 106.7KB / 90.3KB | 0.85x more | 13 / 10 |
| **Unique** | 6.9µs / 6.3µs | 0.90x | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **UniqueBy** | 7.2µs / 6.3µs | 0.88x | 53.4KB / 45.1KB | 0.85x more | 8 / 6 |
| **Zip** | 2.6µs / 3.2µs | **1.24x** | 32.8KB / 16.4KB | 0.50x more | 3 / 1 |
| **ZipWith** | 2.2µs / 3.1µs | **1.38x** | 24.6KB / 8.2KB | 0.33x more | 3 / 1 |
<!-- bench:embed:end -->

## How to read the benchmarks

* **≈** means the two libraries are effectively equivalent
* **∞x less** means one side allocates while the other allocates nothing
* Single-operation helpers are intentionally close in performance if not exceeds
* Multi-step pipelines highlight the architectural difference

If you prefer immutable, one-off helpers - `lo` is outstanding.
If you write **expressive, chained data pipelines** and care about hot-path performance - `collection` is built for that job.


## Why chaining changes the performance story

Most functional helpers (including `lo`) operate like this:

```
input → Map → new slice → Filter → new slice → Take → new slice
```

That model is simple and safe - but each step typically allocates.

`collection` pipelines are designed to look more like this:

```
input → Filter (in place) → Sort (in place) → Take (slice view)
```

When you opt into mutation, **the pipeline stays on the same backing array** unless an operation explicitly documents that it allocates. The result is:

* **Fewer allocations**
* **Less GC pressure**
* **Lower end-to-end latency in hot paths**
* **Much stronger scaling in multi-step pipelines**

That's why the biggest deltas appear in benchmarks like:

* `Pipeline F→M→T→R`
* `Map`
* `Filter`
* `Chunk`
* `Zip / ZipWith`
* `Skip / SkipLast`

In these cases, `collection` can be **2×–30× faster** and often reduce allocations to **zero**, not by doing "clever tricks", but by making mutation *explicit and opt-in*.

## Explicit branching with `Clone`

Fluent pipelines don't mean you're locked into mutation.

When you want to branch a pipeline or preserve the original data, `Clone()` creates a shallow copy of the collection so subsequent operations are isolated and predictable.

```go
events := collection.New(deviceEvents)

// Fast alerting path: cheap filters, early exit
alerts := events.
    Clone().
    Filter(func(e DeviceEvent) bool { return e.Severity >= Critical }).
    Take(10)

// Deeper analysis path: heavier work, full ordering
report := events.
    Clone().
    Filter(func(e DeviceEvent) bool { return e.Region == "us-east" }).
    Sort(func(a, b DeviceEvent) bool { return a.Timestamp.Before(b.Timestamp) })
```

This keeps the performance benefits of in-place operations **where they matter**, while making divergence points explicit and intentional.

No hidden copies. No surprises.

## Design Principles

- **Type-safe**: no reflection, no `any` leaks
- **Explicit semantics**: order, mutation, and allocation are documented
- **Go-native**: respects generics and stdlib patterns
- **Eager evaluation**: no lazy pipelines or hidden concurrency
- **Maps are boundaries**: unordered data is handled explicitly

## What this library is not

- Not a lazy or streaming library
- Not concurrency-aware
- Not immutable-by-default
- Not a replacement for idiomatic loops in simple cases

## Working with maps

Maps are unordered in Go. This library does not pretend otherwise.

Instead, map interaction is explicit and intentional:

- `FromMap` materializes key/value pairs into an ordered workflow
- `ToMap` reduces collections back into maps explicitly
- `ToMapKV` provides a convenience for `Pair[K,V]`

This makes transitions between unordered and ordered data visible and honest.

## Behavior semantics

Each method declares how it interacts with the collection:

- **readonly** – reads data only, returns a derived value
- **immutable** – returns a new collection, original unchanged
- **mutable** – modifies the collection in place

Annotations describe **observable behavior**, not implementation details.

## Runnable examples

Every function has a corresponding runnable example under [`./examples`](./examples).

These examples are **generated directly from the documentation blocks** of each function, ensuring the docs and code never drift. These are the same examples you see here in the README and GoDoc.

An automated test executes **every example** to verify it builds and runs successfully.  

This guarantees all examples are valid, up-to-date, and remain functional as the API evolves.

# Installation

```bash
go get github.com/goforj/collection
```

<!-- api:embed:start -->

## API Index

| Group | Functions |
|------:|-----------|
| **Access** | [Items](#items) [ItemsCopy](#itemscopy) |
| **Aggregation** | [Avg](#avg) [Count](#count) [CountBy](#countby) [CountByValue](#countbyvalue) [Max](#max) [MaxBy](#maxby) [Median](#median) [Min](#min) [MinBy](#minby) [Mode](#mode) [Reduce](#reduce) [Sum](#sum) |
| **Construction** | [Attach](#attach) [AttachNumeric](#attachnumeric) [Clone](#clone) [New](#new) [NewNumeric](#newnumeric) |
| **Debugging** | [Dd](#dd) [Dump](#dump) [DumpStr](#dumpstr) |
| **Grouping** | [GroupBy](#groupby) [GroupBySlice](#groupbyslice) |
| **Maps** | [FromMap](#frommap) [ToMap](#tomap) [ToMapKV](#tomapkv) |
| **Ordering** | [After](#after) [Before](#before) [Reverse](#reverse) [Shuffle](#shuffle) [Sort](#sort) |
| **Querying** | [All](#all) [Any](#any) [At](#at) [Contains](#contains) [First](#first) [FirstWhere](#firstwhere) [IndexWhere](#indexwhere) [IsEmpty](#isempty) [Last](#last) [LastWhere](#lastwhere) [None](#none) |
| **Serialization** | [ToJSON](#tojson) [ToPrettyJSON](#toprettyjson) |
| **Set Operations** | [Difference](#difference) [Intersect](#intersect) [SymmetricDifference](#symmetricdifference) [Union](#union) [Unique](#unique) [UniqueBy](#uniqueby) [UniqueComparable](#uniquecomparable) |
| **Slicing** | [Chunk](#chunk) [Filter](#filter) [Partition](#partition) [Pop](#pop) [PopN](#popn) [Skip](#skip) [SkipLast](#skiplast) [Take](#take) [TakeLast](#takelast) [TakeUntil](#takeuntil) [TakeUntilFn](#takeuntilfn) [Window](#window) |
| **Transformation** | [Append](#append) [Concat](#concat) [Each](#each) [Map](#map) [MapTo](#mapto) [Merge](#merge) [Multiply](#multiply) [Pipe](#pipe) [Prepend](#prepend) [Tap](#tap) [Times](#times) [Transform](#transform) [Zip](#zip) [ZipWith](#zipwith) |


## Access

### <a id="items"></a>Items · readonly · terminal

Items returns the backing slice of items.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3})
items := c.Items()
collection.Dump(items)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
// ]
```

_Example: strings_

```go
c2 := collection.New([]string{"apple", "banana"})
items2 := c2.Items()
collection.Dump(items2)
// #[]string [
//   0 => "apple" #string
//   1 => "banana" #string
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

out := users.Items()
collection.Dump(out)
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

### <a id="itemscopy"></a>ItemsCopy · readonly · terminal

ItemsCopy returns a copy of the collection's items.

```go
c := collection.New([]int{1, 2, 3})
items := c.ItemsCopy()
collection.Dump(items)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
// ]
```

## Aggregation

### <a id="avg"></a>Avg · readonly · terminal

Avg returns the average of the collection values as a float64.
If the collection is empty, Avg returns 0.

_Example: integers_

```go
c := collection.NewNumeric([]int{2, 4, 6})
collection.Dump(c.Avg())
// 4.000000 #float64
```

_Example: float_

```go
c2 := collection.NewNumeric([]float64{1.5, 2.5, 3.0})
collection.Dump(c2.Avg())
// 2.333333 #float64
```

### <a id="count"></a>Count · readonly · terminal

Count returns the total number of items in the collection.

```go
count := collection.New([]int{1, 2, 3, 4}).Count()
collection.Dump(count)
// 4 #int
```

### <a id="countby"></a>CountBy · readonly · terminal

CountBy returns a map of keys extracted by fn to their occurrence counts.
K must be comparable.

_Example: integers_

```go
c := collection.New([]int{1, 2, 2, 3, 3, 3})
counts := collection.CountBy(c, func(v int) int {
	return v
})
collection.Dump(counts)
// map[int]int {
//   1: 1 #int
//   2: 2 #int
//   3: 3 #int
// }
```

_Example: strings_

```go
c2 := collection.New([]string{"apple", "banana", "apple", "cherry", "banana"})
counts2 := collection.CountBy(c2, func(v string) string {
	return v
})
collection.Dump(counts2)
// map[string]int {
//   "apple":  2 #int
//   "banana": 2 #int
//   "cherry": 1 #int
// }
```

_Example: structs_

```go
type User struct {
	Name string
	Role string
}

users := collection.New([]User{
	{Name: "Alice", Role: "admin"},
	{Name: "Bob", Role: "user"},
	{Name: "Carol", Role: "admin"},
	{Name: "Dave", Role: "user"},
	{Name: "Eve", Role: "admin"},
})

roleCounts := collection.CountBy(users, func(u User) string {
	return u.Role
})

collection.Dump(roleCounts)
// map[string]int {
//   "admin": 3 #int
//   "user":  2 #int
// }
```

### <a id="countbyvalue"></a>CountByValue · readonly · terminal

CountByValue returns a map where each distinct item in the collection
is mapped to the number of times it appears.

_Example: strings_

```go
c1 := collection.New([]string{"a", "b", "a"})
counts1 := collection.CountByValue(c1)
collection.Dump(counts1)
// #map[string]int [
//	"a" => 2 #int
//	"b" => 1 #int
// ]
```

_Example: integers_

```go
c2 := collection.New([]int{1, 2, 2, 3, 3, 3})
counts2 := collection.CountByValue(c2)
collection.Dump(counts2)
// #map[int]int [
//	1 => 1 #int
//	2 => 2 #int
//	3 => 3 #int
// ]
```

_Example: structs (comparable)_

```go
type Point struct {
	X int
	Y int
}

c3 := collection.New([]Point{
	{X: 1, Y: 1},
	{X: 2, Y: 2},
	{X: 1, Y: 1},
})

counts3 := collection.CountByValue(c3)
collection.Dump(counts3)
// #map[collection.Point]int [
//	{X:1 Y:1} => 2 #int
//	{X:2 Y:2} => 1 #int
// ]
```

### <a id="max"></a>Max · readonly · terminal

Max returns the largest numeric item in the collection.
The second return value is false if the collection is empty.

_Example: integers_

```go
c := collection.NewNumeric([]int{3, 1, 2})

max1, ok1 := c.Max()
collection.Dump(max1, ok1)
// 3    #int
// true #bool
```

_Example: floats_

```go
c2 := collection.NewNumeric([]float64{1.5, 9.2, 4.4})

max2, ok2 := c2.Max()
collection.Dump(max2, ok2)
// 9.200000 #float64
// true     #bool
```

_Example: empty numeric collection_

```go
c3 := collection.NewNumeric([]int{})

max3, ok3 := c3.Max()
collection.Dump(max3, ok3)
// 0     #int
// false #bool
```

### <a id="maxby"></a>MaxBy · readonly · terminal

MaxBy returns the item whose key (produced by keyFn) is the largest.
The second return value is false if the collection is empty.

_Example: structs - highest score_

```go
type Player struct {
	Name  string
	Score int
}

players := collection.New([]Player{
	{Name: "Alice", Score: 10},
	{Name: "Bob", Score: 25},
	{Name: "Carol", Score: 18},
})

top, ok := collection.MaxBy(players, func(p Player) int {
	return p.Score
})

collection.Dump(top, ok)
// #main.Player {
//   +Name  => "Bob" #string
//   +Score => 25 #int
// }
// true #bool
```

_Example: strings - longest length_

```go
words := collection.New([]string{"go", "collection", "rocks"})

longest, ok := collection.MaxBy(words, func(s string) int {
	return len(s)
})

collection.Dump(longest, ok)
// "collection" #string
// true #bool
```

_Example: empty collection_

```go
empty := collection.New([]int{})
maxVal, ok := collection.MaxBy(empty, func(v int) int { return v })
collection.Dump(maxVal, ok)
// 0 #int
// false #bool
```

### <a id="median"></a>Median · readonly · terminal

Median returns the statistical median of the numeric collection as float64.
Returns (0, false) if the collection is empty.

_Example: integers - odd number of items_

```go
c := collection.NewNumeric([]int{3, 1, 2})

median1, ok1 := c.Median()
collection.Dump(median1, ok1)
// 2.000000 #float64
// true     #bool
```

_Example: integers - even number of items_

```go
c2 := collection.NewNumeric([]int{10, 2, 4, 6})

median2, ok2 := c2.Median()
collection.Dump(median2, ok2)
// 5.000000 #float64
// true     #bool
```

_Example: floats_

```go
c3 := collection.NewNumeric([]float64{1.1, 9.9, 3.3})

median3, ok3 := c3.Median()
collection.Dump(median3, ok3)
// 3.300000 #float64
// true     #bool
```

_Example: integers - empty numeric collection_

```go
c4 := collection.NewNumeric([]int{})

median4, ok4 := c4.Median()
collection.Dump(median4, ok4)
// 0.000000 #float64
// false    #bool
```

### <a id="min"></a>Min · readonly · terminal

Min returns the smallest numeric item in the collection.
The second return value is false if the collection is empty.

_Example: integers_

```go
c := collection.NewNumeric([]int{3, 1, 2})
min, ok := c.Min()
collection.Dump(min, ok)
// 1 #int
// true #bool
```

_Example: floats_

```go
c2 := collection.NewNumeric([]float64{2.5, 9.1, 1.2})
min2, ok2 := c2.Min()
collection.Dump(min2, ok2)
// 1.200000 #float64
// true #bool
```

_Example: integers - empty collection_

```go
empty := collection.NewNumeric([]int{})
min3, ok3 := empty.Min()
collection.Dump(min3, ok3)
// 0 #int
// false #bool
```

### <a id="minby"></a>MinBy · readonly · terminal

MinBy returns the item whose key (produced by keyFn) is the smallest.
The second return value is false if the collection is empty.

_Example: structs - smallest age_

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

minUser, ok := collection.MinBy(users, func(u User) int {
	return u.Age
})

collection.Dump(minUser, ok)
// #main.User {
//   +Name => "Bob" #string
//   +Age  => 25 #int
// }
// true #bool
```

_Example: strings - shortest length_

```go
words := collection.New([]string{"apple", "fig", "banana"})

shortest, ok := collection.MinBy(words, func(s string) int {
	return len(s)
})

collection.Dump(shortest, ok)
// "fig" #string
// true #bool
```

_Example: empty collection_

```go
empty := collection.New([]int{})
minVal, ok := collection.MinBy(empty, func(v int) int { return v })
collection.Dump(minVal, ok)
// 0 #int
// false #bool
```

### <a id="mode"></a>Mode · readonly · terminal

Mode returns the most frequent numeric value(s) in the collection.
If multiple values tie for highest frequency, all are returned
in first-seen order.

_Example: integers – single mode_

```go
c := collection.NewNumeric([]int{1, 2, 2, 3})
mode := c.Mode()
collection.Dump(mode)
// #[]int [
//   0 => 2 #int
// ]
```

_Example: integers – tie for mode_

```go
c2 := collection.NewNumeric([]int{1, 2, 1, 2})
mode2 := c2.Mode()
collection.Dump(mode2)
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
// ]
```

_Example: floats_

```go
c3 := collection.NewNumeric([]float64{1.1, 2.2, 1.1, 3.3})
mode3 := c3.Mode()
collection.Dump(mode3)
// #[]float64 [
//   0 => 1.100000 #float64
// ]
```

_Example: integers - empty collection_

```go
empty := collection.NewNumeric([]int{})
mode4 := empty.Mode()
collection.Dump(mode4)
// <nil>
```

### <a id="reduce"></a>Reduce · readonly · terminal

Reduce collapses the collection into a single accumulated value.
The accumulator has the same type T as the collection's elements.

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
// #main.Stats [
//   +Count => 3 #int
//   +Sum   => 60 #int
// ]
```

### <a id="sum"></a>Sum · readonly · terminal

Sum returns the sum of all numeric items in the NumericCollection.
If the collection is empty, Sum returns the zero value of T.

_Example: integers_

```go
c := collection.NewNumeric([]int{1, 2, 3})
total := c.Sum()
collection.Dump(total)
// 6 #int
```

_Example: floats_

```go
c2 := collection.NewNumeric([]float64{1.5, 2.5})
total2 := c2.Sum()
collection.Dump(total2)
// 4.000000 #float64
```

_Example: integers - empty collection_

```go
c3 := collection.NewNumeric([]int{})
total3 := c3.Sum()
collection.Dump(total3)
// 0 #int
```

## Construction

### <a id="attach"></a>Attach · immutable · fluent

Attach wraps a slice without copying.

```go
items := []int{1, 2, 3}
c := collection.Attach(items)

items[0] = 9
collection.Dump(c.Items())
// #[]int [
//   0 => 9 #int
//   1 => 2 #int
//   2 => 3 #int
// ]
```

### <a id="attachnumeric"></a>AttachNumeric · immutable · fluent

AttachNumeric wraps a slice of numeric types without copying.

```go
items := []int{1, 2, 3}
c := collection.AttachNumeric(items)

items[0] = 9
collection.Dump(c.Items())
// #[]int [
//   0 => 9 #int
//   1 => 2 #int
//   2 => 3 #int
// ]
```

### <a id="clone"></a>Clone · immutable · fluent

Clone returns a copy of the collection.

_Example: basic cloning_

```go
c := collection.New([]int{1, 2, 3})
clone := c.Clone()

clone.Append(4)

collection.Dump(c.Items())
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
// ]

collection.Dump(clone.Items())
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
//   3 => 4 #int
// ]
```

_Example: branching pipelines_

```go
base := collection.New([]int{1, 2, 3, 4, 5})

evens := base.Clone().Filter(func(v int) bool {
	return v%2 == 0
})

odds := base.Clone().Filter(func(v int) bool {
	return v%2 != 0
})

collection.Dump(base.Items())
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
//   3 => 4 #int
//   4 => 5 #int
// ]

collection.Dump(evens.Items())
// #[]int [
//   0 => 2 #int
//   1 => 4 #int
// ]

collection.Dump(odds.Items())
// #[]int [
//   0 => 1 #int
//   1 => 3 #int
//   2 => 5 #int
// ]
```

### <a id="new"></a>New · immutable · fluent

New creates a new Collection from the provided slice.

### <a id="newnumeric"></a>NewNumeric · immutable · fluent

NewNumeric wraps a slice of numeric types in a NumericCollection.

## Debugging

### <a id="dd"></a>Dd · terminal

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

### <a id="dump"></a>Dump · readonly · fluent

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

_Example: integers_

```go
c2 := collection.New([]int{1, 2, 3})
collection.Dump(c2.Items())
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
// ]
```

### <a id="dumpstr"></a>DumpStr · readonly · terminal

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

## Grouping

### <a id="groupby"></a>GroupBy · readonly · terminal

GroupBy partitions the collection into groups keyed by the value
returned from keyFn.

_Example: grouping integers by parity_

```go
values := []int{1, 2, 3, 4, 5}

groups := collection.GroupBy(
	collection.New(values),
	func(v int) string {
		if v%2 == 0 {
			return "even"
		}
		return "odd"
	},
)

collection.Dump(groups["even"].Items())
// []int [
//  0 => 2 #int
//  1 => 4 #int
// ]
collection.Dump(groups["odd"].Items())
// []int [
//  0 => 1 #int
//  1 => 3 #int
//  2 => 5 #int
// ]
```

_Example: grouping structs by field_

```go
type User struct {
	ID   int
	Role string
}

users := []User{
	{ID: 1, Role: "admin"},
	{ID: 2, Role: "user"},
	{ID: 3, Role: "admin"},
}

groups2 := collection.GroupBy(
	collection.New(users),
	func(u User) string { return u.Role },
)

collection.Dump(groups2["admin"].Items())
// []main.User [
//  0 => #main.User {
//    +ID   => 1 #int
//    +Role => "admin" #string
//  }
//  1 => #main.User {
//    +ID   => 3 #int
//    +Role => "admin" #string
//  }
// ]
collection.Dump(groups2["user"].Items())
// []main.User [
//  0 => #main.User {
//    +ID   => 2 #int
//    +Role => "user" #string
//  }
// ]
```

### <a id="groupbyslice"></a>GroupBySlice · readonly · terminal

GroupBySlice partitions the collection into groups keyed by the value
returned from keyFn.

_Example: grouping integers by parity_

```go
values := []int{1, 2, 3, 4, 5}

groups := collection.GroupBySlice(
	collection.New(values),
	func(v int) string {
		if v%2 == 0 {
			return "even"
		}
		return "odd"
	},
)

collection.Dump(groups["even"])
// []int [
//  0 => 2 #int
//  1 => 4 #int
// ]
collection.Dump(groups["odd"])
// []int [
//  0 => 1 #int
//  1 => 3 #int
//  2 => 5 #int
// ]
```

_Example: grouping structs by field_

```go
type User struct {
	ID   int
	Role string
}

users := []User{
	{ID: 1, Role: "admin"},
	{ID: 2, Role: "user"},
	{ID: 3, Role: "admin"},
}

groups2 := collection.GroupBySlice(
	collection.New(users),
	func(u User) string { return u.Role },
)

collection.Dump(groups2["admin"])
// []main.User [
//  0 => #main.User {
//    +ID   => 1 #int
//    +Role => "admin" #string
//  }
//  1 => #main.User {
//    +ID   => 3 #int
//    +Role => "admin" #string
//  }
// ]
collection.Dump(groups2["user"])
// []main.User [
//  0 => #main.User {
//    +ID   => 2 #int
//    +Role => "user" #string
//  }
// ]
```

## Maps

### <a id="frommap"></a>FromMap · immutable · fluent

FromMap materializes a map into a collection of key/value pairs.

_Example: basic usage_

```go
m := map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
}

c := collection.FromMap(m)
collection.Dump(c.Items())

// #[]collection.Pair[string,int] [
//   0 => {Key:"a" Value:1}
//   1 => {Key:"b" Value:2}
//   2 => {Key:"c" Value:3}
// ]
```

_Example: filtering map entries_

```go
type Config struct {
	Enabled bool
	Timeout int
}

configs := map[string]Config{
	"router-1": {Enabled: true,  Timeout: 30},
	"router-2": {Enabled: false, Timeout: 10},
	"router-3": {Enabled: true,  Timeout: 45},
}

out := collection.
	FromMap(configs).
	Filter(func(p collection.Pair[string, Config]) bool {
		return p.Value.Enabled
	}).
	Items()

collection.Dump(out)

// #[]collection.Pair[string,collection.Config] [
//   0 => {Key:"router-1" Value:{Enabled:true Timeout:30}}
//   1 => {Key:"router-3" Value:{Enabled:true Timeout:45}}
// ]
```

_Example: map → collection → map_

```go
users := map[string]int{
	"alice": 1,
	"bob":   2,
}

c2 := collection.FromMap(users)
out2 := collection.ToMapKV(c2)

collection.Dump(out2)

// #map[string]int [
//   "alice" => 1
//   "bob"   => 2
// ]
```

### <a id="tomap"></a>ToMap · readonly · terminal

ToMap reduces a collection into a map using the provided key and value
selector functions.

_Example: basic usage_

```go
users := []string{"alice", "bob", "carol"}

out := collection.ToMap(
	collection.New(users),
	func(name string) string { return name },
	func(name string) int { return len(name) },
)

collection.Dump(out)
```

_Example: re-keying structs_

```go
type User struct {
	ID   int
	Name string
}

users2 := []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

byID := collection.ToMap(
	collection.New(users2),
	func(u User) int { return u.ID },
	func(u User) User { return u },
)

collection.Dump(byID)
```

### <a id="tomapkv"></a>ToMapKV · readonly · terminal

ToMapKV converts a collection of key/value pairs into a map.

_Example: basic usage_

```go
m := map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
}

c := collection.FromMap(m)
out := collection.ToMapKV(c)

collection.Dump(out)

// #map[string]int [
//   "a" => 1
//   "b" => 2
//   "c" => 3
// ]
```

_Example: filtering before conversion_

```go
type Config struct {
	Enabled bool
	Timeout int
}

configs := map[string]Config{
	"router-1": {Enabled: true,  Timeout: 30},
	"router-2": {Enabled: false, Timeout: 10},
	"router-3": {Enabled: true,  Timeout: 45},
}

c2 := collection.
	FromMap(configs).
	Filter(func(p collection.Pair[string, Config]) bool {
		return p.Value.Enabled
	})

out2 := collection.ToMapKV(c2)

collection.Dump(out2)

// #map[string]collection.Config [
//   "router-1" => {Enabled:true Timeout:30}
//   "router-3" => {Enabled:true Timeout:45}
// ]
```

## Ordering

### <a id="after"></a>After · immutable · fluent

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

### <a id="before"></a>Before · immutable · fluent

Before returns a new collection containing all items that appear
*before* the first element for which pred returns true.

_Example: integers_

```go
c1 := collection.New([]int{1, 2, 3, 4, 5})
out1 := c1.Before(func(v int) bool { return v >= 3 })
collection.Dump(out1.Items())
// #[]int [
//	0 => 1 #int
//	1 => 2 #int
// ]
```

_Example: predicate never matches → whole collection returned_

```go
c2 := collection.New([]int{10, 20, 30})
out2 := c2.Before(func(v int) bool { return v == 99 })
collection.Dump(out2.Items())
// #[]int [
//	0 => 10 #int
//	1 => 20 #int
//	2 => 30 #int
// ]
```

_Example: structs: get all users before the first admin_

```go
type User struct {
	Name  string
	Admin bool
}

c3 := collection.New([]User{
	{Name: "Alice", Admin: false},
	{Name: "Bob", Admin: false},
	{Name: "Eve", Admin: true},
	{Name: "Mallory", Admin: false},
})

out3 := c3.Before(func(u User) bool { return u.Admin })
collection.Dump(out3.Items())
// #[]collection.User [
//	0 => {Name:"Alice" Admin:false}  #collection.User
//	1 => {Name:"Bob"   Admin:false}  #collection.User
// ]
```

### <a id="reverse"></a>Reverse · mutable · fluent

Reverse reverses the order of items in the collection in place
and returns the same collection for chaining.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4})
c.Reverse()
collection.Dump(c.Items())
// #[]int [
//   0 => 4 #int
//   1 => 3 #int
//   2 => 2 #int
//   3 => 1 #int
// ]
```

_Example: strings – chaining_

```go
out := collection.New([]string{"a", "b", "c"}).
	Reverse().
	Append("d").
	Items()

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
collection.Dump(users.Items())
// #[]collection.User [
//   0 => {ID:3} #collection.User
//   1 => {ID:2} #collection.User
//   2 => {ID:1} #collection.User
// ]
```

### <a id="shuffle"></a>Shuffle · immutable · fluent

Shuffle returns a shuffled copy of the collection.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4, 5})
out1 := c.Shuffle()
collection.Dump(out1.Items())
```

_Example: strings – chaining_

```go
out2 := collection.New([]string{"a", "b", "c"}).
	Shuffle().
	Append("d").
	Items()

collection.Dump(out2)
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

out3 := users.Shuffle()
collection.Dump(out3.Items())
```

### <a id="sort"></a>Sort · mutable · fluent

Sort sorts the collection in place using the provided comparison function and
returns the same collection for chaining.

_Example: integers_

```go
c := collection.New([]int{5, 1, 4, 2})
c.Sort(func(a, b int) bool { return a < b })
collection.Dump(c.Items())
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
collection.Dump(c2.Items())
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
collection.Dump(users.Items())
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

### <a id="all"></a>All · readonly · terminal

All returns true if fn returns true for every item in the collection.
If the collection is empty, All returns true (vacuously true).

_Example: integers – all even_

```go
c := collection.New([]int{2, 4, 6})
allEven := c.All(func(v int) bool { return v%2 == 0 })
collection.Dump(allEven)
// true #bool
```

_Example: integers – not all even_

```go
c2 := collection.New([]int{2, 3, 4})
allEven2 := c2.All(func(v int) bool { return v%2 == 0 })
collection.Dump(allEven2)
// false #bool
```

_Example: strings – all non-empty_

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

### <a id="any"></a>Any · readonly · terminal

Any returns true if at least one item satisfies fn.

```go
c := collection.New([]int{1, 2, 3, 4})
has := c.Any(func(v int) bool { return v%2 == 0 }) // true
collection.Dump(has)
// true #bool
```

### <a id="at"></a>At · readonly · terminal

At returns the item at the given index and a boolean indicating
whether the index was within bounds.

_Example: integers_

```go
c := collection.New([]int{10, 20, 30})
v, ok := c.At(1)
collection.Dump(v, ok)
// 20 true
```

_Example: out of bounds_

```go
v2, ok2 := c.At(10)
collection.Dump(v2, ok2)
// 0 false
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
// {ID:1 Name:"Alice"} true
```

### <a id="contains"></a>Contains · readonly · terminal

Contains returns true if the collection contains the given value.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4, 5})
hasTwo := collection.Contains(c, 2)
collection.Dump(hasTwo)
// true #bool
```

_Example: strings_

```go
c2 := collection.New([]string{"apple", "banana", "cherry"})
hasBanana := collection.Contains(c2, "banana")
collection.Dump(hasBanana)
// true #bool
```

### <a id="first"></a>First · readonly · terminal

First returns the first element in the collection.
If the collection is empty, ok will be false.

_Example: integers_

```go
c := collection.New([]int{10, 20, 30})

v, ok := c.First()
collection.Dump(v, ok)
// 10   #int
// true #bool
```

_Example: strings_

```go
c2 := collection.New([]string{"alpha", "beta", "gamma"})

v2, ok2 := c2.First()
collection.Dump(v2, ok2)
// "alpha" #string
// true    #bool
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
//   +ID   => 1      #int
//   +Name => "Alice" #string
// }
// true #bool
```

_Example: integers - empty collection_

```go
c3 := collection.New([]int{})
v3, ok4 := c3.First()
collection.Dump(v3, ok4)
// 0    #int
// false #bool
```

### <a id="firstwhere"></a>FirstWhere · readonly · terminal

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

### <a id="indexwhere"></a>IndexWhere · readonly · terminal

IndexWhere returns the index of the first item in the collection
for which the provided predicate function returns true.
If no item matches, it returns (0, false).

_Example: integers_

```go
c := collection.New([]int{10, 20, 30, 40})
idx, ok := c.IndexWhere(func(v int) bool { return v == 30 })
collection.Dump(idx, ok)
// 2 true
```

_Example: not found_

```go
idx2, ok2 := c.IndexWhere(func(v int) bool { return v == 99 })
collection.Dump(idx2, ok2)
// 0 false
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
// 1 true
```

### <a id="isempty"></a>IsEmpty · readonly · terminal

IsEmpty returns true if the collection has no items.

_Example: integers (non-empty)_

```go
c := collection.New([]int{1, 2, 3})

empty := c.IsEmpty()
collection.Dump(empty)
// false #bool
```

_Example: strings (empty)_

```go
c2 := collection.New([]string{})

empty2 := c2.IsEmpty()
collection.Dump(empty2)
// true #bool
```

_Example: structs (non-empty)_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
})

empty3 := users.IsEmpty()
collection.Dump(empty3)
// false #bool
```

_Example: structs (empty)_

```go
none := collection.New([]User{})

empty4 := none.IsEmpty()
collection.Dump(empty4)
// true #bool
```

### <a id="last"></a>Last · readonly · terminal

Last returns the last element in the collection.
If the collection is empty, ok will be false.

_Example: integers_

```go
c := collection.New([]int{10, 20, 30})

v, ok := c.Last()
collection.Dump(v, ok)
// 30   #int
// true #bool
```

_Example: strings_

```go
c2 := collection.New([]string{"alpha", "beta", "gamma"})

v2, ok2 := c2.Last()
collection.Dump(v2, ok2)
// "gamma" #string
// true    #bool
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
//   +ID   => 3         #int
//   +Name => "Charlie" #string
// }
// true #bool
```

_Example: empty collection_

```go
c3 := collection.New([]int{})

v3, ok4 := c3.Last()
collection.Dump(v3, ok4)
// 0     #int
// false #bool
```

### <a id="lastwhere"></a>LastWhere · readonly · terminal

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
// 2    #int
// true #bool
```

_Example: integers without predicate (equivalent to Last())_

```go
c2 := collection.New([]int{10, 20, 30, 40})

v2, ok2 := c2.LastWhere(nil)
collection.Dump(v2, ok2)
// 40   #int
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
// true    #bool
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
//   +ID   => 3        #int
//   +Name => "Alex"  #string
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
// 0     #int
// false #bool
```

_Example: empty collection_

```go
c5 := collection.New([]int{})

v5, ok6 := c5.LastWhere(nil)
collection.Dump(v5, ok6)
// 0     #int
// false #bool
```

### <a id="none"></a>None · readonly · terminal

None returns true if fn returns false for every item in the collection.
If the collection is empty, None returns true.

_Example: integers – none even_

```go
c := collection.New([]int{1, 3, 5})
noneEven := c.None(func(v int) bool { return v%2 == 0 })
collection.Dump(noneEven)
// true #bool
```

_Example: integers – some even_

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

## Serialization

### <a id="tojson"></a>ToJSON · readonly · terminal

ToJSON converts the collection's items into a compact JSON string.

```go
pj1 := collection.New([]string{"a", "b"})
out1, _ := pj1.ToJSON()
fmt.Println(out1)
// ["a","b"]
```

### <a id="toprettyjson"></a>ToPrettyJSON · readonly · terminal

ToPrettyJSON converts the collection's items into a human-readable,
indented JSON string.

```go
pj1 := collection.New([]string{"a", "b"})
out1, _ := pj1.ToPrettyJSON()
fmt.Println(out1)
// [
//  "a",
//  "b"
// ]
```

## Set Operations

### <a id="difference"></a>Difference · immutable · fluent

Difference returns a new collection containing elements from the first collection
that are not present in the second. Order follows the first collection, and
duplicates are removed.

_Example: integers_

```go
a := collection.New([]int{1, 2, 2, 3, 4})
b := collection.New([]int{2, 4})

out := collection.Difference(a, b)
collection.Dump(out.Items())
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
collection.Dump(out2.Items())
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
collection.Dump(out3.Items())
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

### <a id="intersect"></a>Intersect · immutable · fluent

Intersect returns a new collection containing elements from the second
collection that are also present in the first.

_Example: integers_

```go
a := collection.New([]int{1, 2, 2, 3, 4})
b := collection.New([]int{2, 4, 4, 5})

out := collection.Intersect(a, b)
collection.Dump(out.Items())
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
collection.Dump(out2.Items())
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
collection.Dump(out3.Items())
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

### <a id="symmetricdifference"></a>SymmetricDifference · immutable · fluent

SymmetricDifference returns a new collection containing elements that appear
in exactly one of the two collections. Order follows the first collection for
its unique items, then the second for its unique items. Duplicates are removed.

_Example: integers_

```go
a := collection.New([]int{1, 2, 3, 3})
b := collection.New([]int{3, 4, 4, 5})

out := collection.SymmetricDifference(a, b)
collection.Dump(out.Items())
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
collection.Dump(out2.Items())
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
collection.Dump(out3.Items())
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

### <a id="union"></a>Union · immutable · fluent

Union returns a new collection containing the unique elements from both collections.
Items from the first collection are kept in order, followed by items from the second
that were not already present.

_Example: integers_

```go
a := collection.New([]int{1, 2, 2, 3})
b := collection.New([]int{3, 4, 4, 5})

out := collection.Union(a, b)
collection.Dump(out.Items())
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
collection.Dump(out2.Items())
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
collection.Dump(out3.Items())
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

### <a id="unique"></a>Unique · immutable · fluent

Unique returns a new collection with duplicate items removed, based on the
equality function `eq`. The first occurrence of each unique value is kept,
and order is preserved.

_Example: integers_

```go
c1 := collection.New([]int{1, 2, 2, 3, 4, 4, 5})
out1 := c1.Unique(func(a, b int) bool { return a == b })
collection.Dump(out1.Items())
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
collection.Dump(out2.Items())
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

collection.Dump(out3.Items())
// #[]collection.User [
//	0 => {ID:1 Name:"Alice"} #collection.User
//	1 => {ID:2 Name:"Bob"}   #collection.User
// ]
```

### <a id="uniqueby"></a>UniqueBy · immutable · fluent

UniqueBy returns a new collection containing only the first occurrence
of each element as determined by keyFn.

_Example: structs – unique by ID_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 1, Name: "Alice Duplicate"},
})

out := collection.UniqueBy(users, func(u User) int { return u.ID })
collection.Dump(out.Items())
// #[]collection.User [
//   0 => {ID:1 Name:"Alice"} #collection.User
//   1 => {ID:2 Name:"Bob"}   #collection.User
// ]
```

_Example: strings – case-insensitive uniqueness_

```go
values := collection.New([]string{"A", "a", "B", "b", "A"})

out2 := collection.UniqueBy(values, func(s string) string {
	return strings.ToLower(s)
})

collection.Dump(out2.Items())
// #[]string [
//   0 => "A" #string
//   1 => "B" #string
// ]
```

_Example: integers – identity key_

```go
nums := collection.New([]int{3, 1, 2, 1, 3})

out3 := collection.UniqueBy(nums, func(v int) int { return v })
collection.Dump(out3.Items())
// #[]int [
//   0 => 3 #int
//   1 => 1 #int
//   2 => 2 #int
// ]
```

### <a id="uniquecomparable"></a>UniqueComparable · immutable · fluent

UniqueComparable returns a new collection with duplicate comparable items removed.
The first occurrence of each value is kept, and order is preserved.
This is a faster, allocation-friendly path for comparable types.

_Example: integers_

```go
c := collection.New([]int{1, 2, 2, 3, 4, 4, 5})
out := collection.UniqueComparable(c)
collection.Dump(out.Items())
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
collection.Dump(out2.Items())
// #[]string [
//   0 => "A" #string
//   1 => "a" #string
//   2 => "B" #string
// ]
```

## Slicing

### <a id="chunk"></a>Chunk · readonly · terminal

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

// Dump output will show [][]User grouped in size-2 chunks, e.g.:
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

### <a id="filter"></a>Filter · mutable · fluent

Filter keeps only the elements for which fn returns true.
This method mutates the collection in place and returns the same instance.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4})
c.Filter(func(v int) bool {
	return v%2 == 0
})
collection.Dump(c.Items())
// #[]int [
//   0 => 2 #int
//   1 => 4 #int
// ]
```

_Example: strings_

```go
c2 := collection.New([]string{"apple", "banana", "cherry", "avocado"})
c2.Filter(func(v string) bool {
	return strings.HasPrefix(v, "a")
})
collection.Dump(c2.Items())
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

users.Filter(func(u User) bool {
	return strings.HasPrefix(u.Name, "A")
})

collection.Dump(users.Items())
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

### <a id="partition"></a>Partition · immutable · fluent

Partition splits the collection into two new collections based on predicate fn.
The first collection contains items where fn returns true; the second contains
items where fn returns false. Order is preserved within each partition.

_Example: integers - even/odd_

```go
nums := collection.New([]int{1, 2, 3, 4, 5})
evens, odds := nums.Partition(func(n int) bool {
	return n%2 == 0
})
collection.Dump(evens.Items(), odds.Items())
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
collection.Dump(goWords.Items(), other.Items())
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

collection.Dump(active.Items(), inactive.Items())
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

### <a id="pop"></a>Pop · mutable · terminal

Pop removes and returns the last item in the collection.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3})
item, ok := c.Pop()
collection.Dump(item, ok, c.Items())
// 3 #int
// true #bool
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
// ]
```

_Example: strings_

```go
c2 := collection.New([]string{"a", "b", "c"})
item2, ok2 := c2.Pop()
collection.Dump(item2, ok2, c2.Items())
// "c" #string
// true #bool
// #[]string [
//   0 => "a" #string
//   1 => "b" #string
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

item3, ok3 := users.Pop()
collection.Dump(item3, ok3, users.Items())
// #main.User {
//   +ID   => 2 #int
//   +Name => "Bob" #string
// }
// true #bool
// #[]main.User [
//   0 => #main.User {
//     +ID   => 1 #int
//     +Name => "Alice" #string
//   }
// ]
```

_Example: empty collection_

```go
empty := collection.New([]int{})
item4, ok4 := empty.Pop()
collection.Dump(item4, ok4, empty.Items())
// 0 #int
// false #bool
// #[]int [
// ]
```

### <a id="popn"></a>PopN · mutable · terminal

PopN removes and returns the last n items in original order.

_Example: integers – pop 2_

```go
c := collection.New([]int{1, 2, 3, 4})
popped := c.PopN(2)
collection.Dump(popped, c.Items())
// #[]int [
//   0 => 3 #int
//   1 => 4 #int
// ]
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
// ]
```

_Example: strings – pop 1_

```go
c2 := collection.New([]string{"a", "b", "c"})
popped2 := c2.PopN(1)
collection.Dump(popped2, c2.Items())
// #[]string [
//   0 => "c" #string
// ]
// #[]string [
//   0 => "a" #string
//   1 => "b" #string
// ]
```

_Example: structs – pop 2_

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

popped3 := users.PopN(2)
collection.Dump(popped3, users.Items())
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
// #[]main.User [
//   0 => #main.User {
//     +ID   => 1 #int
//     +Name => "Alice" #string
//   }
// ]
```

_Example: integers - n <= 0 → returns nil, no change_

```go
c3 := collection.New([]int{1, 2, 3})
popped4 := c3.PopN(0)
collection.Dump(popped4, c3.Items())
// <nil>
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
// ]
```

_Example: strings - n exceeds length → all items popped, rest empty_

```go
c4 := collection.New([]string{"x", "y"})
popped5 := c4.PopN(10)
collection.Dump(popped5, c4.Items())
// #[]string [
//   0 => "x" #string
//   1 => "y" #string
// ]
// #[]string [
// ]
```

### <a id="skip"></a>Skip · immutable · fluent

Skip returns a new collection with the first n items skipped.
If n is less than or equal to zero, Skip returns the full collection.
If n is greater than or equal to the collection length, Skip returns
an empty collection.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4, 5})
out := c.Skip(2)
collection.Dump(out.Items())
// #[]int [
//   0 => 3 #int
//   1 => 4 #int
//   2 => 5 #int
// ]
```

_Example: skip none_

```go
out2 := c.Skip(0)
collection.Dump(out2.Items())
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
collection.Dump(out3.Items())
// #[]int []
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
collection.Dump(out4.Items())
// []main.User [
//  0 => #main.User {
//    +ID => 2 #int
//  }
//  1 => #main.User {
//    +ID => 3 #int
//  }
// ]
```

### <a id="skiplast"></a>SkipLast · immutable · fluent

SkipLast returns a new collection with the last n items skipped.
If n is less than or equal to zero, SkipLast returns the full collection.
If n is greater than or equal to the collection length, SkipLast returns
an empty collection.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4, 5})
out := c.SkipLast(2)
collection.Dump(out.Items())
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
// ]
```

_Example: skip none_

```go
out2 := c.SkipLast(0)
collection.Dump(out2.Items())
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
collection.Dump(out3.Items())
// #[]int []
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
collection.Dump(out4.Items())
// #[]collection.User [
//   0 => {ID:1} #collection.User
//   1 => {ID:2} #collection.User
// ]
```

### <a id="take"></a>Take · immutable · fluent

Take returns a new collection containing the first `n` items when n > 0,
or the last `|n|` items when n < 0.

_Example: integers - take first 3_

```go
c1 := collection.New([]int{0, 1, 2, 3, 4, 5})
out1 := c1.Take(3)
collection.Dump(out1.Items())
// #[]int [
//	0 => 0 #int
//	1 => 1 #int
//	2 => 2 #int
// ]
```

_Example: integers - take last 2 (negative n)_

```go
c2 := collection.New([]int{0, 1, 2, 3, 4, 5})
out2 := c2.Take(-2)
collection.Dump(out2.Items())
// #[]int [
//	0 => 4 #int
//	1 => 5 #int
// ]
```

_Example: integers - n exceeds length → whole collection_

```go
c3 := collection.New([]int{10, 20})
out3 := c3.Take(10)
collection.Dump(out3.Items())
// #[]int [
//	0 => 10 #int
//	1 => 20 #int
// ]
```

_Example: integers - zero → empty_

```go
c4 := collection.New([]int{1, 2, 3})
out4 := c4.Take(0)
collection.Dump(out4.Items())
// #[]int [
// ]
```

### <a id="takelast"></a>TakeLast · immutable · fluent

TakeLast returns a new collection containing the last n items.
If n is less than or equal to zero, TakeLast returns an empty collection.
If n is greater than or equal to the collection length, TakeLast returns
the full collection.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3, 4, 5})
out := c.TakeLast(2)
collection.Dump(out.Items())
// #[]int [
//   0 => 4 #int
//   1 => 5 #int
// ]
```

_Example: take none_

```go
out2 := c.TakeLast(0)
collection.Dump(out2.Items())
// #[]int []
```

_Example: take all_

```go
out3 := c.TakeLast(10)
collection.Dump(out3.Items())
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
collection.Dump(out4.Items())
// #[]collection.User [
//   0 => {ID:3} #collection.User
// ]
```

### <a id="takeuntil"></a>TakeUntil · immutable · fluent

TakeUntil returns items until the first element equals `value`.
The matching item is NOT included.

_Example: integers - stop at value 3_

```go
c4 := collection.New([]int{1, 2, 3, 4})
out4 := collection.TakeUntil(c4, 3)
collection.Dump(out4.Items())
// #[]int [
//	0 => 1 #int
//	1 => 2 #int
// ]
```

_Example: strings - value never appears → full slice_

```go
c5 := collection.New([]string{"a", "b", "c"})
out5 := collection.TakeUntil(c5, "x")
collection.Dump(out5.Items())
// #[]string [
//	0 => "a" #string
//	1 => "b" #string
//	2 => "c" #string
// ]
```

_Example: integers - match is first item → empty result_

```go
c6 := collection.New([]int{9, 10, 11})
out6 := collection.TakeUntil(c6, 9)
collection.Dump(out6.Items())
// #[]int [
// ]
```

### <a id="takeuntilfn"></a>TakeUntilFn · immutable · fluent

TakeUntilFn returns items until the predicate function returns true.
The matching item is NOT included.

_Example: integers - stop when value >= 3_

```go
c1 := collection.New([]int{1, 2, 3, 4})
out1 := c1.TakeUntilFn(func(v int) bool { return v >= 3 })
collection.Dump(out1.Items())
// #[]int [
//	0 => 1 #int
//	1 => 2 #int
// ]
```

_Example: integers - predicate immediately true → empty result_

```go
c2 := collection.New([]int{10, 20, 30})
out2 := c2.TakeUntilFn(func(v int) bool { return v < 50 })
collection.Dump(out2.Items())
// #[]int [
// ]
```

_Example: integers - no match → full list returned_

```go
c3 := collection.New([]int{1, 2, 3})
out3 := c3.TakeUntilFn(func(v int) bool { return v == 99 })
collection.Dump(out3.Items())
// #[]int [
//	0 => 1 #int
//	1 => 2 #int
//	2 => 3 #int
// ]
```

### <a id="window"></a>Window · allocates · fluent

Window returns overlapping (or stepped) windows of the collection.
Each window is a slice of length size; iteration advances by step (default 1 if step <= 0).
Windows that are shorter than size are omitted.

_Example: integers - step 1_

```go
nums := collection.New([]int{1, 2, 3, 4, 5})
win := collection.Window(nums, 3, 1)
collection.Dump(win.Items())
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
win2 := collection.Window(words, 2, 2)
collection.Dump(win2.Items())
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

win3 := collection.Window(points, 2, 1)
collection.Dump(win3.Items())
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

### <a id="append"></a>Append · immutable · fluent

Append returns a new collection with the given values appended.

_Example: integers_

```go
c := collection.New([]int{1, 2})
c.Append(3, 4).Dump()
// #[]int [
//  0 => 1 #int
//  1 => 2 #int
//  2 => 3 #int
//  3 => 4 #int
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

users.Append(
	User{ID: 3, Name: "Carol"},
	User{ID: 4, Name: "Dave"},
).Dump()

// #[]main.User [
//  0 => #main.User {
//    +ID   => 1 #int
//    +Name => "Alice" #string
//  }
//  1 => #main.User {
//    +ID   => 2 #int
//    +Name => "Bob" #string
//  }
//  2 => #main.User {
//    +ID   => 3 #int
//    +Name => "Carol" #string
//  }
//  3 => #main.User {
//    +ID   => 4 #int
//    +Name => "Dave" #string
//  }
// ]
```

### <a id="concat"></a>Concat · mutable · fluent

Concat appends the values from the given slice onto the end of the collection,

```go
c := collection.New([]string{"John Doe"})
concatenated := c.
	Concat([]string{"Jane Doe"}).
	Concat([]string{"Johnny Doe"}).
	Items()
collection.Dump(concatenated)

// #[]string [
//  0 => "John Doe" #string
//  1 => "Jane Doe" #string
//  2 => "Johnny Doe" #string
// ]
```

### <a id="each"></a>Each · readonly · fluent

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
//   0 => "APPLE"  #string
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
//   0 => "Alice"   #string
//   1 => "Bob"     #string
//   2 => "Charlie" #string
// ]
```

### <a id="map"></a>Map · immutable · fluent

Map applies a same-type transformation and returns a new collection.

_Example: integers_

```go
c := collection.New([]int{1, 2, 3})

mapped := c.Map(func(v int) int {
	return v * 10
})

collection.Dump(mapped.Items())
// #[]int [
//   0 => 10 #int
//   1 => 20 #int
//   2 => 30 #int
// ]
```

_Example: strings_

```go
c2 := collection.New([]string{"apple", "banana", "cherry"})

upper := c2.Map(func(s string) string {
	return strings.ToUpper(s)
})

collection.Dump(upper.Items())
// #[]string [
//   0 => "APPLE"  #string
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

updated := users.Map(func(u User) User {
	u.Name = strings.ToUpper(u.Name)
	return u
})

collection.Dump(updated.Items())
// #[]main.User [
//   0 => #main.User {
//     +ID   => 1        #int
//     +Name => "ALICE"  #string
//   }
//   1 => #main.User {
//     +ID   => 2        #int
//     +Name => "BOB"    #string
//   }
// ]
```

### <a id="mapto"></a>MapTo · immutable · fluent

MapTo maps a Collection[T] to a Collection[R] using fn(T) R.

_Example: integers - extract parity label_

```go
nums := collection.New([]int{1, 2, 3, 4})
parity := collection.MapTo(nums, func(n int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
})
collection.Dump(parity.Items())
// #[]string [
//   0 => "odd" #string
//   1 => "even" #string
//   2 => "odd" #string
//   3 => "even" #string
// ]
```

_Example: strings - length of each value_

```go
words := collection.New([]string{"go", "forj", "rocks"})
lengths := collection.MapTo(words, func(s string) int {
	return len(s)
})
collection.Dump(lengths.Items())
// #[]int [
//   0 => 2 #int
//   1 => 4 #int
//   2 => 5 #int
// ]
```

_Example: structs - MapTo a field_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
})

names := collection.MapTo(users, func(u User) string {
	return u.Name
})

collection.Dump(names.Items())
// #[]string [
//   0 => "Alice" #string
//   1 => "Bob" #string
// ]
```

### <a id="merge"></a>Merge · immutable · fluent

Merge merges the given data into a new collection.

_Example: integers - merging slices_

```go
ints := collection.New([]int{1, 2})
extra := []int{3, 4}
// Merge the extra slice into the ints collection
merged1 := ints.Merge(extra)
collection.Dump(merged1.Items())
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
//   2 => 3 #int
//   3 => 4 #int
// ]
```

_Example: strings - merging another collection_

```go
strs := collection.New([]string{"a", "b"})
more := collection.New([]string{"c", "d"})

merged2 := strs.Merge(more)
collection.Dump(merged2.Items())
// #[]string [
//   0 => "a" #string
//   1 => "b" #string
//   2 => "c" #string
//   3 => "d" #string
// ]
```

_Example: structs - merging struct slices_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
})

moreUsers := []User{
	{ID: 3, Name: "Carol"},
	{ID: 4, Name: "Dave"},
}

merged3 := users.Merge(moreUsers)
collection.Dump(merged3.Items())
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
//   3 => #main.User {
//     +ID   => 4 #int
//     +Name => "Dave" #string
//   }
// ]
```

### <a id="multiply"></a>Multiply · immutable · fluent

Multiply creates `n` copies of all items in the collection
and returns a new collection.

_Example: integers_

```go
ints := collection.New([]int{1, 2})
out := ints.Multiply(3)
collection.Dump(out.Items())
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
collection.Dump(out2.Items())
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
collection.Dump(out3.Items())
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
collection.Dump(none.Items())
// #[]int [
// ]
```

### <a id="pipe"></a>Pipe · readonly · terminal

Pipe passes the entire collection into the given function
and returns the function's result.

_Example: integers – computing a sum_

```go
c := collection.New([]int{1, 2, 3})
sum := collection.Pipe(c, func(col *collection.Collection[int]) int {
	total := 0
	for _, v := range col.Items() {
		total += v
	}
	return total
})
collection.Dump(sum)
// 6 #int
```

_Example: strings – joining values_

```go
c2 := collection.New([]string{"a", "b", "c"})
joined := collection.Pipe(c2, func(col *collection.Collection[string]) string {
	out := ""
	for _, v := range col.Items() {
		out += v
	}
	return out
})
collection.Dump(joined)
// "abc" #string
```

_Example: structs – extracting just the names_

```go
type User struct {
	ID   int
	Name string
}

users := collection.New([]User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
})

names := collection.Pipe(users, func(col *collection.Collection[User]) []string {
	result := make([]string, 0, len(col.Items()))
	for _, u := range col.Items() {
		result = append(result, u.Name)
	}
	return result
})

collection.Dump(names)
// #[]string [
//   0 => "Alice" #string
//   1 => "Bob" #string
// ]
```

### <a id="prepend"></a>Prepend · mutable · fluent

Prepend adds the given values to the beginning of the collection.

_Example: integers_

```go
c := collection.New([]int{3, 4})
c.Prepend(1, 2)
collection.Dump(c.Items())
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
letters.Prepend("a", "b")
collection.Dump(letters.Items())
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

users.Prepend(User{ID: 1, Name: "Alice"})
collection.Dump(users.Items())
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
empty.Prepend(9, 8)
collection.Dump(empty.Items())
// #[]int [
//   0 => 9 #int
//   1 => 8 #int
// ]
```

_Example: integers - Prepending no values → no change_

```go
c2 := collection.New([]int{1, 2})
c2.Prepend()
collection.Dump(c2.Items())
// #[]int [
//   0 => 1 #int
//   1 => 2 #int
// ]
```

### <a id="tap"></a>Tap · immutable · fluent

Tap invokes fn with the collection pointer for side effects (logging, debugging,
inspection) and returns the same collection to allow chaining.

_Example: integers - capture intermediate state during a chain_

```go
captured1 := []int{}
c1 := collection.New([]int{3, 1, 2}).
	Sort(func(a, b int) bool { return a < b }). // → [1, 2, 3]
	Tap(func(col *collection.Collection[int]) {
		captured1 = append([]int(nil), col.Items()...) // snapshot copy
	}).
	Filter(func(v int) bool { return v >= 2 }).
	Dump()
	// #[]int [
	//  0 => 2 #int
	//  1 => 3 #int
	// ]

// Use BOTH variables so nothing is "declared and not used"
collection.Dump(c1.Items())
collection.Dump(captured1)
// c1 → #[]int [2,3]
// captured1 → #[]int [1,2,3]
```

_Example: integers - tap for debugging without changing flow_

```go
c2 := collection.New([]int{10, 20, 30}).
	Tap(func(col *collection.Collection[int]) {
		collection.Dump(col.Items())
	}).
	Filter(func(v int) bool { return v > 10 })

collection.Dump(c2.Items()) // ensures c2 is used
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

users2 := users.Tap(func(col *collection.Collection[User]) {
	collection.Dump(col.Items())
})

collection.Dump(users2.Items()) // ensures users2 is used
```

### <a id="times"></a>Times · immutable · fluent

Times creates a new collection by calling fn(i) for i = 1..count.
This mirrors Laravel's Collection::times(), which is 1-indexed.

_Example: integers - double each index_

```go
cTimes1 := collection.Times(5, func(i int) int {
	return i * 2
})
collection.Dump(cTimes1.Items())
// #[]int [
//	0 => 2  #int
//	1 => 4  #int
//	2 => 6  #int
//	3 => 8  #int
//	4 => 10 #int
// ]
```

_Example: strings_

```go
cTimes2 := collection.Times(3, func(i int) string {
	return fmt.Sprintf("item-%d", i)
})
collection.Dump(cTimes2.Items())
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
collection.Dump(cTimes3.Items())
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

### <a id="transform"></a>Transform · mutable · terminal

Transform applies fn to every item *in place*, mutating the collection.

_Example: integers_

```go
c1 := collection.New([]int{1, 2, 3})
c1.Transform(func(v int) int { return v * 2 })
collection.Dump(c1.Items())
// #[]int [
//	0 => 2 #int
//	1 => 4 #int
//	2 => 6 #int
// ]
```

_Example: strings_

```go
c2 := collection.New([]string{"a", "b", "c"})
c2.Transform(func(s string) string { return strings.ToUpper(s) })
collection.Dump(c2.Items())
// #[]string [
//	0 => "A" #string
//	1 => "B" #string
//	2 => "C" #string
// ]
```

_Example: structs_

```go
type User struct {
	ID   int
	Name string
}

c3 := collection.New([]User{
	{ID: 1, Name: "alice"},
	{ID: 2, Name: "bob"},
})

c3.Transform(func(u User) User {
	u.Name = strings.ToUpper(u.Name)
	return u
})

collection.Dump(c3.Items())
// #[]collection.User [
//	0 => {ID:1 Name:"ALICE"} #collection.User
//	1 => {ID:2 Name:"BOB"}   #collection.User
// ]
```

### <a id="zip"></a>Zip · immutable · fluent

Zip combines two collections element-wise into a collection of tuples.
The resulting length is the smaller of the two inputs.

_Example: integers and strings_

```go
nums := collection.New([]int{1, 2, 3})
words := collection.New([]string{"one", "two"})

out := collection.Zip(nums, words)
collection.Dump(out.Items())
// #[]collection.Tuple[int,string] [
//   0 => #collection.Tuple[int,string] {
//     +First  => 1 #int
//     +Second => "one" #string
//   }
//   1 => #collection.Tuple[int,string] {
//     +First  => 2 #int
//     +Second => "two" #string
//   }
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

roles := collection.New([]string{"admin", "user", "extra"})

out2 := collection.Zip(users, roles)
collection.Dump(out2.Items())
// #[]collection.Tuple[main.User,string] [
//   0 => #collection.Tuple[main.User,string] {
//     +First  => #main.User {
//       +ID   => 1 #int
//       +Name => "Alice" #string
//     }
//     +Second => "admin" #string
//   }
//   1 => #collection.Tuple[main.User,string] {
//     +First  => #main.User {
//       +ID   => 2 #int
//       +Name => "Bob" #string
//     }
//     +Second => "user" #string
//   }
// ]
```

### <a id="zipwith"></a>ZipWith · immutable · fluent

ZipWith combines two collections element-wise using combiner fn.
The resulting length is the smaller of the two inputs.

_Example: sum ints_

```go
a := collection.New([]int{1, 2, 3})
b := collection.New([]int{10, 20})

out := collection.ZipWith(a, b, func(x, y int) int {
	return x + y
})

collection.Dump(out.Items())
// #[]int [
//   0 => 11 #int
//   1 => 22 #int
// ]
```

_Example: format strings_

```go
names := collection.New([]string{"alice", "bob"})
roles := collection.New([]string{"admin", "user", "extra"})

out2 := collection.ZipWith(names, roles, func(name, role string) string {
	return name + ":" + role
})

collection.Dump(out2.Items())
// #[]string [
//   0 => "alice:admin" #string
//   1 => "bob:user" #string
// ]
```

_Example: structs_

```go
type User struct {
	Name string
}

type Role struct {
	Title string
}

users := collection.New([]User{{Name: "Alice"}, {Name: "Bob"}})
roles2 := collection.New([]Role{{Title: "admin"}})

out3 := collection.ZipWith(users, roles2, func(u User, r Role) string {
	return u.Name + " -> " + r.Title
})

collection.Dump(out3.Items())
// #[]string [
//   0 => "Alice -> admin" #string
// ]
```
<!-- api:embed:end -->
