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
    <a href="https://goreportcard.com/report/github.com/goforj/collection"><img src="https://goreportcard.com/badge/github.com/goforj/collection" alt="Go Report Card"></a>
</p>

<p align="center">
  <code>collection</code> brings an expressive, fluent API to Go.  
  Iterate, filter, transform, sort, reduce, group, and debug your data with zero dependencies.  
  Designed to feel natural to Go developers - and luxurious to everyone else.
</p>

# Features

- 🔗 **Fluent chaining** - pipeline your operations like Laravel Collections
- 🧬 **Fully generic** (`Collection[T]`) - no reflection, no `interface{}`
- ⚡ **Zero dependencies** - pure Go, fast, lightweight
- 🧵 **Minimal allocations** - avoids unnecessary copies; most operations reuse the underlying slice
- 🧹 **Map / Filter / Reduce** - clean functional transforms
- 🔍 **First / Last / Find / Contains** helpers
- 📏 **Sort, GroupBy, Chunk**, and more
- 🧪 **Safe-by-default** - defensive copies where appropriate
- 📜 **Built-in JSON helpers** (`ToJSON()`, `ToPrettyJSON()`)
- 🧰 **Developer-friendly debug helpers** (`Dump()`, `Dd()`, `DumpStr()`)
- 🧱 **Works with any Go type**, including structs, pointers, and deeply nested composites

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

# 📦 Installation

```bash
go get github.com/goforj/collection
```

<!-- api:embed:start -->

### Other
- `After`
- `All`
- `Any`
- `Append`
- `At`
- `Avg`
- `Before`
- `Chunk`
- `Concat`
- `Contains`
- `Count`
- `CountBy`
- `CountByValue`
- `Dd`
- `Dump`
- `DumpStr`
- `Each`
- `Filter`
- `FindWhere`
- `First`
- `FirstWhere`
- `FromMap`
- `GroupBy`
- `IndexWhere`
- `IsEmpty`
- `Items`
- `Last`
- `LastWhere`
- `Map`
- `MapTo`
- `Max`
- `Median`
- `Merge`
- `Min`
- `Mode`
- `Multiply`
- `New`
- `NewNumeric`
- `None`
- `Pipe`
- `Pluck`
- `Pop`
- `PopN`
- `Prepend`
- `Push`
- `Reduce`
- `Reverse`
- `Shuffle`
- `Skip`
- `SkipLast`
- `Sort`
- `Sum`
- `TakeUntil`
- `TakeUntilFn`
- `Tap`
- `Times`
- `ToJSON`
- `ToMap`
- `ToMapKV`
- `ToPrettyJSON`
- `Transform`
- `Unique`
- `UniqueBy`

### Slicing
- `Take`
- `TakeLast`
<!-- api:embed:end -->