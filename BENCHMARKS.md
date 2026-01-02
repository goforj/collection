# Benchmarks

Raw results for `collection.New` (borrowed) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 249ns / 232ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Any** | 249ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Chunk** | 142ns / 1.0µs | **7.39x** | 1.3KB / 9.3KB | **7.12x less** | 2 / 51 |
| **Contains** | 250ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **CountBy** | 8.3µs / 8.3µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **CountByValue** | 8.2µs / 8.0µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **Difference** | 20.8µs / 49.1µs | **2.37x** | 82.2KB / 134.0KB | **1.63x less** | 14 / 55 |
| **Each** | 248ns / 232ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Filter** | 757ns / 1.1µs | **1.41x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **First** | 11ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 249ns / 234ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **GroupBySlice** | 9.8µs / 10.4µs | ≈ | 22.2KB / 22.1KB | ≈ | 140 / 139 |
| **IndexWhere** | 249ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Intersect** | 11.0µs / 10.7µs | ≈ | 9.4KB / 9.4KB | ≈ | 15 / 12 |
| **Last** | 11ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Map** | 295ns / 790ns | **2.68x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Max** | 264ns / 238ns | 0.90x | 32B / 0B | ∞x more | 2 / 0 |
| **Min** | 255ns / 234ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **None** | 248ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 503ns / 1.2µs | **2.47x** | 48B / 12.3KB | **256.00x less** | 2 / 2 |
| **Reduce (sum)** | 253ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Reverse** | 237ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 3.9µs / 5.6µs | **1.44x** | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Skip** | 11ns / 712ns | **63.96x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **SkipLast** | 11ns / 708ns | **63.64x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Sum** | 255ns / 232ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **Take** | 21ns / <1ns | ≈ | 48B / 0B | ∞x more | 2 / 0 |
| **ToMap** | 7.9µs / 8.1µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Union** | 18.0µs / 18.6µs | ≈ | 90.3KB / 90.3KB | ≈ | 13 / 10 |
| **Unique** | 6.3µs / 6.3µs | ≈ | 45.2KB / 45.1KB | ≈ | 7 / 6 |
| **UniqueBy** | 6.6µs / 6.3µs | ≈ | 45.2KB / 45.1KB | ≈ | 8 / 6 |
| **Zip** | 1.4µs / 3.2µs | **2.28x** | 16.4KB / 16.4KB | ≈ | 3 / 1 |
| **ZipWith** | 1.0µs / 3.0µs | **2.96x** | 8.2KB / 8.2KB | ≈ | 3 / 1 |
