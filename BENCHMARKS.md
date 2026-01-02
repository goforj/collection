# Benchmarks

Raw results for `collection.New` (borrowed) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 254ns / 238ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Any** | 253ns / 248ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Chunk** | 139ns / 1.0µs | **7.55x** | 1.3KB / 9.3KB | **7.12x less** | 2 / 51 |
| **Contains** | 250ns / 235ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **CountBy** | 8.4µs / 8.5µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **CountByValue** | 8.4µs / 8.3µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **Difference** | 19.7µs / 44.5µs | **2.26x** | 82.2KB / 108.8KB | **1.32x less** | 14 / 43 |
| **Each** | 254ns / 237ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Filter** | 658ns / 1.0µs | **1.56x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **First** | 12ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 251ns / 236ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **GroupBySlice** | 8.3µs / 8.7µs | ≈ | 21.0KB / 21.0KB | ≈ | 84 / 83 |
| **IndexWhere** | 252ns / 236ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Intersect** | 11.3µs / 10.9µs | ≈ | 11.5KB / 11.4KB | ≈ | 22 / 19 |
| **Last** | 11ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Map** | 826ns / 791ns | ≈ | 8.2KB / 8.2KB | ≈ | 2 / 1 |
| **Max** | 256ns / 237ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **Min** | 270ns / 233ns | 0.87x | 32B / 0B | ∞x more | 2 / 0 |
| **None** | 252ns / 236ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 758ns / 1.3µs | **1.66x** | 4.1KB / 12.3KB | **2.97x less** | 3 / 2 |
| **Reduce (sum)** | 248ns / 232ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Reverse** | 240ns / 240ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 4.0µs / 5.3µs | **1.33x** | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Skip** | 12ns / 739ns | **64.23x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **SkipLast** | 12ns / 765ns | **63.73x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Sum** | 266ns / 234ns | 0.88x | 32B / 0B | ∞x more | 2 / 0 |
| **Take** | 22ns / <1ns | ≈ | 48B / 0B | ∞x more | 2 / 0 |
| **ToMap** | 7.8µs / 8.0µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Union** | 17.2µs / 17.9µs | ≈ | 90.3KB / 90.3KB | ≈ | 13 / 10 |
| **Unique** | 6.3µs / 6.3µs | ≈ | 45.2KB / 45.1KB | ≈ | 7 / 6 |
| **UniqueBy** | 6.7µs / 6.4µs | ≈ | 45.2KB / 45.1KB | ≈ | 8 / 6 |
| **Zip** | 1.4µs / 3.3µs | **2.28x** | 16.4KB / 16.4KB | ≈ | 3 / 1 |
| **ZipWith** | 1.1µs / 3.2µs | **3.00x** | 8.2KB / 8.2KB | ≈ | 3 / 1 |
