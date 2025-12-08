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
    <a href="https://goreportcard.com/report/github.com/goforj/collection"><img src="https://goreportcard.com/badge/github.com/goforj/collection" alt="Go Report Card"></a>
</p>

<p align="center">
  <code>collection</code> brings an expressive, fluent API to Go.  
  Iterate, filter, transform, sort, reduce, group, and debug your data with zero dependencies.  
  Designed to feel natural to Go developers - and luxurious to everyone else.
</p>

# Features

- 🔗 **Fluent chaining** - pipeline your operations like Laravel Collections
- 🧬 **Fully generic** (`Collection[T]`) - no reflection, no interface{}
- ⚡ **Zero dependencies** - pure Go, fast, lightweight
- 🧹 **Map / Filter / Reduce** - clean functional transforms
- 🔍 **First / Last / Find / Contains** helpers
- 📏 **Sort, GroupBy, Chunk**, and more
- 🧪 **Safe-by-default** - defensive copies where appropriate
- 📜 **Built-in JSON helpers** (`ToJSON()`, `ToPrettyJSON()`)
- 🧰 **Developer-friendly debug helpers** (`Dump()`, `Dd()`, `DumpStr()`, `DdStr()`)
- 🧱 **Works with any Go type**, including structs, pointers, and deeply nested composites

# 📦 Installation

```bash
go get github.com/goforj/collection
```

