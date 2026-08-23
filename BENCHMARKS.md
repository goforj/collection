# Benchmarks

Raw results for `collection.New` (borrowed) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 269ns / 267ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 269ns / 270ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 970ns / 3.5µs | **3.62x** | 1.3KB / 9.3KB | **7.25x less** | 1 / 51 |
| **Contains** | 265ns / 263ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 8.1µs / 8.2µs | ≈ | 9.4KB / 9.4KB | ≈ | 11 / 11 |
| **CountByValue** | 8.0µs / 8.0µs | ≈ | 9.4KB / 9.4KB | ≈ | 11 / 11 |
| **Difference** | 30.3µs / 58.1µs | **1.92x** | 82.1KB / 108.8KB | **1.33x less** | 11 / 40 |
| **Each** | 267ns / 265ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 615ns / 2.3µs | **3.70x** | 0B / 8.2KB | **∞x less** | 0 / 1 |
| **First** | <1ns / <1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 263ns / 264ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 11.7µs / 11.3µs | ≈ | 21.0KB / 21.0KB | ≈ | 83 / 83 |
| **IndexWhere** | 271ns / 268ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 13.0µs / 12.2µs | ≈ | 11.4KB / 11.3KB | ≈ | 19 / 16 |
| **Last** | <1ns / <1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 348ns / 1.8µs | **5.17x** | 0B / 8.2KB | **∞x less** | 0 / 1 |
| **Max** | 262ns / 256ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 261ns / 262ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 270ns / 267ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 540ns / 3.1µs | **5.66x** | 0B / 12.3KB | **∞x less** | 0 / 2 |
| **Reduce (sum)** | 264ns / 265ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 210ns / 227ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 3.8µs / 5.4µs | **1.43x** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | <1ns / 1.6µs | ∞ | 0B / 8.2KB | **∞x less** | 0 / 1 |
| **SkipLast** | <1ns / 1.4µs | ∞ | 0B / 8.2KB | **∞x less** | 0 / 1 |
| **Sum** | 262ns / 262ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1ns / <1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 11.5µs / 12.2µs | ≈ | 36.9KB / 37.0KB | ≈ | 5 / 6 |
| **Union** | 27.3µs / 29.6µs | ≈ | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 11.8µs / 11.7µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 11.8µs / 11.8µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.0µs / 4.7µs | **2.30x** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.4µs / 4.0µs | **2.81x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.5µs / 262ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.5µs / 263ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 2.0µs / 3.3µs | **1.69x** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 1.5µs / 263ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 9.9µs / 8.2µs | 0.82x | 17.5KB / 9.4KB | 0.53x more | 12 / 11 |
| **CountByValue** | 9.6µs / 8.0µs | 0.83x | 17.5KB / 9.4KB | 0.53x more | 12 / 11 |
| **Difference** | 32.3µs / 57.0µs | **1.77x** | 98.5KB / 108.8KB | **1.10x less** | 13 / 40 |
| **Each** | 1.5µs / 257ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 1.8µs / 1.9µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.3µs / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.5µs / 262ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 13.0µs / 11.8µs | ≈ | 29.2KB / 21.0KB | 0.72x more | 84 / 83 |
| **IndexWhere** | 1.5µs / 262ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 15.5µs / 12.5µs | 0.81x | 27.8KB / 11.3KB | 0.41x more | 21 / 16 |
| **Last** | 1.6µs / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 1.6µs / 1.5µs | 0.90x | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 1.6µs / 255ns | 0.16x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.6µs / 254ns | 0.16x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.5µs / 261ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 2.0µs / 2.4µs | **1.17x** | 8.2KB / 12.3KB | **1.50x less** | 1 / 2 |
| **Reduce (sum)** | 1.5µs / 255ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.4µs / 226ns | 0.16x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 5.0µs / 5.4µs | ≈ | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.4µs / 1.4µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.3µs / 1.3µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.5µs / 253ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.4µs / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 13.2µs / 11.5µs | 0.87x | 45.1KB / 37.0KB | 0.82x more | 6 / 6 |
| **Union** | 29.8µs / 30.7µs | ≈ | 106.7KB / 90.3KB | 0.85x more | 12 / 10 |
| **Unique** | 12.7µs / 12.1µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **UniqueBy** | 12.9µs / 11.8µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **Zip** | 4.7µs / 4.6µs | ≈ | 32.8KB / 16.4KB | 0.50x more | 3 / 1 |
| **ZipWith** | 4.0µs / 4.1µs | ≈ | 24.6KB / 8.2KB | 0.33x more | 3 / 1 |
