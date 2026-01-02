# Benchmarks

Raw results for `collection.New` (borrowed) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 255ns / 238ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Any** | 254ns / 241ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Chunk** | 141ns / 1.1µs | **7.84x** | 1.3KB / 9.3KB | **7.12x less** | 2 / 51 |
| **Contains** | 253ns / 239ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **CountBy** | 8.3µs / 8.6µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **CountByValue** | 8.3µs / 8.3µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **Difference** | 19.8µs / 44.9µs | **2.27x** | 82.2KB / 108.8KB | **1.32x less** | 14 / 43 |
| **Each** | 258ns / 236ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Filter** | 663ns / 1.0µs | **1.55x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **First** | 12ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 257ns / 241ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **GroupBySlice** | 8.2µs / 9.0µs | ≈ | 21.0KB / 21.0KB | ≈ | 84 / 83 |
| **IndexWhere** | 254ns / 240ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Intersect** | 11.1µs / 10.8µs | ≈ | 11.5KB / 11.4KB | ≈ | 22 / 19 |
| **Last** | 12ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Map** | 366ns / 800ns | **2.18x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Max** | 258ns / 235ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **Min** | 256ns / 233ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **None** | 256ns / 240ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 724ns / 1.3µs | **1.74x** | 48B / 12.3KB | **256.00x less** | 2 / 2 |
| **Reduce (sum)** | 259ns / 238ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Reverse** | 232ns / 239ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 4.0µs / 5.7µs | **1.43x** | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Skip** | 11ns / 716ns | **62.99x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **SkipLast** | 12ns / 716ns | **62.29x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Sum** | 258ns / 233ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **Take** | 23ns / <1ns | ≈ | 48B / 0B | ∞x more | 2 / 0 |
| **ToMap** | 7.7µs / 7.9µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Union** | 17.5µs / 18.0µs | ≈ | 90.3KB / 90.3KB | ≈ | 13 / 10 |
| **Unique** | 6.5µs / 6.4µs | ≈ | 45.2KB / 45.1KB | ≈ | 7 / 6 |
| **UniqueBy** | 6.8µs / 6.5µs | ≈ | 45.2KB / 45.1KB | ≈ | 8 / 6 |
| **Zip** | 1.4µs / 3.3µs | **2.27x** | 16.4KB / 16.4KB | ≈ | 3 / 1 |
| **ZipWith** | 1.0µs / 3.3µs | **3.22x** | 8.2KB / 8.2KB | ≈ | 3 / 1 |

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 916ns / 239ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Any** | 906ns / 240ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Chunk** | 846ns / 1.1µs | **1.27x** | 9.5KB / 9.3KB | ≈ | 4 / 51 |
| **Contains** | 905ns / 241ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **CountBy** | 9.1µs / 8.5µs | ≈ | 17.6KB / 9.4KB | 0.53x more | 14 / 11 |
| **CountByValue** | 9.1µs / 8.5µs | ≈ | 17.6KB / 9.4KB | 0.53x more | 14 / 11 |
| **Difference** | 20.8µs / 44.9µs | **2.16x** | 98.6KB / 108.8KB | **1.10x less** | 18 / 43 |
| **Each** | 900ns / 233ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Filter** | 1.2µs / 1.0µs | 0.90x | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **First** | 738ns / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **FirstWhere** | 913ns / 241ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **GroupBySlice** | 8.9µs / 9.2µs | ≈ | 29.3KB / 21.0KB | 0.72x more | 86 / 83 |
| **IndexWhere** | 898ns / 240ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Intersect** | 12.3µs / 11.1µs | ≈ | 27.9KB / 11.4KB | 0.41x more | 26 / 19 |
| **Last** | 726ns / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Map** | 920ns / 803ns | 0.87x | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **Max** | 924ns / 233ns | 0.25x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **Min** | 918ns / 232ns | 0.25x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **None** | 901ns / 239ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Pipeline F→M→T→R** | 1.2µs / 1.3µs | ≈ | 8.3KB / 12.3KB | **1.49x less** | 4 / 2 |
| **Reduce (sum)** | 899ns / 238ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Reverse** | 840ns / 239ns | 0.28x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Shuffle** | 4.5µs / 5.7µs | **1.26x** | 16.5KB / 0B | ∞x more | 5 / 0 |
| **Skip** | 720ns / 725ns | ≈ | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **SkipLast** | 742ns / 729ns | ≈ | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **Sum** | 917ns / 234ns | 0.25x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **Take** | 738ns / <1ns | 0.00x | 8.3KB / 0B | ∞x more | 4 / 0 |
| **ToMap** | 8.2µs / 8.0µs | ≈ | 45.2KB / 37.0KB | 0.82x more | 8 / 6 |
| **Union** | 18.7µs / 18.1µs | ≈ | 106.8KB / 90.3KB | 0.85x more | 17 / 10 |
| **Unique** | 7.1µs / 6.4µs | 0.89x | 53.4KB / 45.1KB | 0.85x more | 9 / 6 |
| **UniqueBy** | 7.4µs / 6.5µs | 0.87x | 53.4KB / 45.1KB | 0.85x more | 10 / 6 |
| **Zip** | 2.8µs / 3.3µs | **1.19x** | 32.9KB / 16.4KB | 0.50x more | 7 / 1 |
| **ZipWith** | 2.4µs / 3.3µs | **1.37x** | 24.7KB / 8.2KB | 0.33x more | 7 / 1 |
