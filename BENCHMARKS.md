# Benchmarks

Raw results for `collection.New` (borrowed) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 248ns / 232ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Any** | 248ns / 235ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Chunk** | 138ns / 1.1µs | **7.64x** | 1.3KB / 9.3KB | **7.12x less** | 2 / 51 |
| **Contains** | 250ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **CountBy** | 8.6µs / 8.4µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **CountByValue** | 8.3µs / 8.3µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **Difference** | 19.5µs / 45.6µs | **2.33x** | 82.2KB / 108.8KB | **1.32x less** | 14 / 43 |
| **Each** | 256ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Filter** | 734ns / 1.0µs | **1.41x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **First** | 10ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 251ns / 244ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **GroupBySlice** | 8.1µs / 8.3µs | ≈ | 21.0KB / 21.0KB | ≈ | 84 / 83 |
| **IndexWhere** | 247ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Intersect** | 11.1µs / 10.8µs | ≈ | 11.5KB / 11.4KB | ≈ | 22 / 19 |
| **Last** | 10ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Map** | 358ns / 790ns | **2.21x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Max** | 263ns / 230ns | 0.87x | 32B / 0B | ∞x more | 2 / 0 |
| **Min** | 253ns / 230ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **None** | 248ns / 231ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 690ns / 1.2µs | **1.79x** | 48B / 12.3KB | **256.00x less** | 2 / 2 |
| **Reduce (sum)** | 262ns / 233ns | 0.89x | 24B / 0B | ∞x more | 1 / 0 |
| **Reverse** | 230ns / 230ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 3.6µs / 5.6µs | **1.57x** | 24B / 0B | ∞x more | 1 / 0 |
| **Skip** | 10ns / 714ns | **71.40x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **SkipLast** | 10ns / 719ns | **71.90x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Sum** | 267ns / 231ns | 0.87x | 32B / 0B | ∞x more | 2 / 0 |
| **Take** | 21ns / <1ns | ≈ | 48B / 0B | ∞x more | 2 / 0 |
| **ToMap** | 7.8µs / 7.9µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Union** | 17.4µs / 17.9µs | ≈ | 90.3KB / 90.3KB | ≈ | 13 / 10 |
| **Unique** | 6.5µs / 6.2µs | ≈ | 45.2KB / 45.1KB | ≈ | 7 / 6 |
| **UniqueBy** | 6.6µs / 6.4µs | ≈ | 45.2KB / 45.1KB | ≈ | 8 / 6 |
| **Zip** | 1.4µs / 3.3µs | **2.35x** | 16.4KB / 16.4KB | ≈ | 3 / 1 |
| **ZipWith** | 1.0µs / 3.2µs | **3.15x** | 8.2KB / 8.2KB | ≈ | 3 / 1 |

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 921ns / 245ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Any** | 908ns / 245ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Chunk** | 855ns / 1.1µs | **1.28x** | 9.5KB / 9.3KB | ≈ | 4 / 51 |
| **Contains** | 912ns / 244ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **CountBy** | 9.2µs / 8.6µs | ≈ | 17.6KB / 9.4KB | 0.53x more | 14 / 11 |
| **CountByValue** | 9.2µs / 8.6µs | ≈ | 17.6KB / 9.4KB | 0.53x more | 14 / 11 |
| **Difference** | 21.1µs / 45.9µs | **2.17x** | 98.6KB / 108.8KB | **1.10x less** | 18 / 43 |
| **Each** | 887ns / 235ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Filter** | 1.2µs / 1.0µs | 0.85x | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **First** | 707ns / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **FirstWhere** | 903ns / 247ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **GroupBySlice** | 9.4µs / 8.7µs | ≈ | 29.3KB / 21.0KB | 0.72x more | 86 / 83 |
| **IndexWhere** | 887ns / 236ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Intersect** | 12.6µs / 11.2µs | 0.89x | 27.9KB / 11.4KB | 0.41x more | 26 / 19 |
| **Last** | 730ns / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Map** | 903ns / 782ns | 0.87x | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **Max** | 928ns / 236ns | 0.25x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **Min** | 931ns / 235ns | 0.25x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **None** | 882ns / 232ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Pipeline F→M→T→R** | 1.2µs / 1.2µs | ≈ | 8.3KB / 12.3KB | **1.49x less** | 4 / 2 |
| **Reduce (sum)** | 887ns / 231ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Reverse** | 904ns / 235ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Shuffle** | 4.1µs / 5.7µs | **1.39x** | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Skip** | 730ns / 955ns | **1.31x** | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **SkipLast** | 768ns / 772ns | ≈ | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **Sum** | 944ns / 239ns | 0.25x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **Take** | 727ns / <1ns | 0.00x | 8.3KB / 0B | ∞x more | 4 / 0 |
| **ToMap** | 8.4µs / 8.1µs | ≈ | 45.2KB / 37.0KB | 0.82x more | 8 / 6 |
| **Union** | 19.6µs / 18.3µs | ≈ | 106.8KB / 90.3KB | 0.85x more | 17 / 10 |
| **Unique** | 7.4µs / 7.5µs | ≈ | 53.4KB / 45.1KB | 0.85x more | 9 / 6 |
| **UniqueBy** | 10.4µs / 6.6µs | 0.64x | 53.4KB / 45.1KB | 0.85x more | 10 / 6 |
| **Zip** | 2.9µs / 3.3µs | **1.15x** | 32.9KB / 16.4KB | 0.50x more | 7 / 1 |
| **ZipWith** | 2.5µs / 3.1µs | **1.27x** | 24.7KB / 8.2KB | 0.33x more | 7 / 1 |
