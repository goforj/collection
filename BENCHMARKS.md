# Benchmarks

Raw results for `collection.New` (borrowed) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 251ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Any** | 249ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Chunk** | 141ns / 1.1µs | **7.49x** | 1.3KB / 9.3KB | **7.12x less** | 2 / 51 |
| **Contains** | 249ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **CountBy** | 8.2µs / 8.0µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **CountByValue** | 8.2µs / 8.1µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **Difference** | 19.1µs / 44.2µs | **2.31x** | 82.2KB / 108.8KB | **1.32x less** | 14 / 43 |
| **Each** | 253ns / 231ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Filter** | 644ns / 1.1µs | **1.64x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **First** | 11ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 249ns / 234ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **GroupBySlice** | 8.0µs / 8.4µs | ≈ | 21.0KB / 21.0KB | ≈ | 84 / 83 |
| **IndexWhere** | 250ns / 233ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Intersect** | 11.0µs / 10.7µs | ≈ | 11.5KB / 11.4KB | ≈ | 22 / 19 |
| **Last** | 11ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Map** | 358ns / 819ns | **2.29x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Max** | 254ns / 231ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **Min** | 255ns / 236ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **None** | 248ns / 234ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 504ns / 1.3µs | **2.53x** | 48B / 12.3KB | **256.00x less** | 2 / 2 |
| **Reduce (sum)** | 253ns / 236ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Reverse** | 227ns / 235ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 3.5µs / 5.5µs | **1.56x** | 24B / 0B | ∞x more | 1 / 0 |
| **Skip** | 11ns / 726ns | **65.27x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **SkipLast** | 11ns / 730ns | **64.89x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Sum** | 255ns / 236ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **Take** | 22ns / <1ns | ≈ | 48B / 0B | ∞x more | 2 / 0 |
| **ToMap** | 7.6µs / 8.0µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Union** | 17.2µs / 17.9µs | ≈ | 90.3KB / 90.3KB | ≈ | 13 / 10 |
| **Unique** | 6.4µs / 6.4µs | ≈ | 45.2KB / 45.1KB | ≈ | 7 / 6 |
| **UniqueBy** | 6.8µs / 6.4µs | ≈ | 45.2KB / 45.1KB | ≈ | 8 / 6 |
| **Zip** | 1.5µs / 3.2µs | **2.22x** | 16.4KB / 16.4KB | ≈ | 3 / 1 |
| **ZipWith** | 1.0µs / 3.0µs | **2.87x** | 8.2KB / 8.2KB | ≈ | 3 / 1 |

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 916ns / 234ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Any** | 920ns / 235ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Chunk** | 862ns / 1.1µs | **1.26x** | 9.5KB / 9.3KB | ≈ | 4 / 51 |
| **Contains** | 924ns / 234ns | 0.25x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **CountBy** | 8.9µs / 8.6µs | ≈ | 17.6KB / 9.4KB | 0.53x more | 14 / 11 |
| **CountByValue** | 9.0µs / 15.2µs | **1.68x** | 17.6KB / 9.4KB | 0.53x more | 14 / 11 |
| **Difference** | 20.6µs / 44.8µs | **2.17x** | 98.6KB / 108.8KB | **1.10x less** | 18 / 43 |
| **Each** | 913ns / 232ns | 0.25x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Filter** | 1.2µs / 1.0µs | 0.89x | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **First** | 737ns / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **FirstWhere** | 934ns / 236ns | 0.25x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **GroupBySlice** | 8.7µs / 8.7µs | ≈ | 29.3KB / 21.0KB | 0.72x more | 86 / 83 |
| **IndexWhere** | 925ns / 234ns | 0.25x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Intersect** | 12.2µs / 10.9µs | 0.90x | 27.9KB / 11.4KB | 0.41x more | 26 / 19 |
| **Last** | 740ns / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Map** | 958ns / 816ns | 0.85x | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **Max** | 930ns / 231ns | 0.25x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **Min** | 925ns / 236ns | 0.26x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **None** | 916ns / 236ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Pipeline F→M→T→R** | 1.1µs / 1.3µs | **1.21x** | 8.3KB / 12.3KB | **1.49x less** | 4 / 2 |
| **Reduce (sum)** | 911ns / 236ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Reverse** | 841ns / 240ns | 0.29x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Shuffle** | 4.2µs / 5.7µs | **1.34x** | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Skip** | 771ns / 982ns | **1.27x** | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **SkipLast** | 2.2µs / 736ns | 0.33x | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **Sum** | 918ns / 236ns | 0.26x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **Take** | 746ns / <1ns | 0.00x | 8.3KB / 0B | ∞x more | 4 / 0 |
| **ToMap** | 8.3µs / 8.0µs | ≈ | 45.2KB / 37.0KB | 0.82x more | 8 / 6 |
| **Union** | 18.6µs / 18.1µs | ≈ | 106.8KB / 90.3KB | 0.85x more | 17 / 10 |
| **Unique** | 7.1µs / 6.4µs | ≈ | 53.4KB / 45.1KB | 0.85x more | 9 / 6 |
| **UniqueBy** | 7.5µs / 6.5µs | 0.86x | 53.4KB / 45.1KB | 0.85x more | 10 / 6 |
| **Zip** | 2.8µs / 3.3µs | **1.17x** | 32.9KB / 16.4KB | 0.50x more | 7 / 1 |
| **ZipWith** | 2.4µs / 3.0µs | **1.26x** | 24.7KB / 8.2KB | 0.33x more | 7 / 1 |
