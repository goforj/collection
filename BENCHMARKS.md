# Benchmarks

Raw results for `collection.New` (borrowed) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 262ns / 240ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Any** | 256ns / 236ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Chunk** | 140ns / 1.1µs | **7.57x** | 1.3KB / 9.3KB | **7.12x less** | 2 / 51 |
| **Contains** | 257ns / 237ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **CountBy** | 8.3µs / 8.5µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **CountByValue** | 8.6µs / 8.6µs | ≈ | 9.4KB / 9.4KB | ≈ | 12 / 11 |
| **Difference** | 19.9µs / 45.1µs | **2.27x** | 82.2KB / 108.8KB | **1.32x less** | 14 / 43 |
| **Each** | 258ns / 236ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Filter** | 658ns / 1.0µs | **1.57x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **First** | 11ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 254ns / 236ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **GroupBySlice** | 8.2µs / 8.7µs | ≈ | 21.0KB / 21.0KB | ≈ | 84 / 83 |
| **IndexWhere** | 278ns / 254ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Intersect** | 11.6µs / 10.9µs | ≈ | 11.5KB / 11.4KB | ≈ | 22 / 19 |
| **Last** | 12ns / <1ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Map** | 366ns / 796ns | **2.18x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Max** | 260ns / 244ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **Min** | 266ns / 247ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **None** | 252ns / 237ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 502ns / 1.4µs | **2.70x** | 48B / 12.3KB | **256.00x less** | 2 / 2 |
| **Reduce (sum)** | 256ns / 241ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Reverse** | 236ns / 239ns | ≈ | 24B / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 4.1µs / 5.7µs | **1.41x** | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Skip** | 12ns / 727ns | **59.96x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **SkipLast** | 12ns / 717ns | **61.68x** | 24B / 8.2KB | **341.33x less** | 1 / 1 |
| **Sum** | 262ns / 241ns | ≈ | 32B / 0B | ∞x more | 2 / 0 |
| **Take** | 23ns / <1ns | ≈ | 48B / 0B | ∞x more | 2 / 0 |
| **ToMap** | 7.9µs / 8.0µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Union** | 18.0µs / 18.2µs | ≈ | 90.3KB / 90.3KB | ≈ | 13 / 10 |
| **Unique** | 6.7µs / 6.7µs | ≈ | 45.2KB / 45.1KB | ≈ | 7 / 6 |
| **UniqueBy** | 6.9µs / 6.7µs | ≈ | 45.2KB / 45.1KB | ≈ | 8 / 6 |
| **Zip** | 1.5µs / 3.6µs | **2.44x** | 16.4KB / 16.4KB | ≈ | 3 / 1 |
| **ZipWith** | 1.1µs / 3.3µs | **3.09x** | 8.2KB / 8.2KB | ≈ | 3 / 1 |

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 940ns / 242ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Any** | 906ns / 241ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Chunk** | 856ns / 1.1µs | **1.25x** | 9.5KB / 9.3KB | ≈ | 4 / 51 |
| **Contains** | 904ns / 242ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **CountBy** | 9.1µs / 8.5µs | ≈ | 17.6KB / 9.4KB | 0.53x more | 14 / 11 |
| **CountByValue** | 9.1µs / 8.4µs | ≈ | 17.6KB / 9.4KB | 0.53x more | 14 / 11 |
| **Difference** | 21.1µs / 45.1µs | **2.14x** | 98.6KB / 108.8KB | **1.10x less** | 18 / 43 |
| **Each** | 911ns / 234ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Filter** | 1.2µs / 1.0µs | 0.90x | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **First** | 739ns / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **FirstWhere** | 913ns / 241ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **GroupBySlice** | 8.9µs / 8.9µs | ≈ | 29.3KB / 21.0KB | 0.72x more | 86 / 83 |
| **IndexWhere** | 916ns / 244ns | 0.27x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Intersect** | 12.3µs / 10.9µs | 0.89x | 27.9KB / 11.4KB | 0.41x more | 26 / 19 |
| **Last** | 729ns / <1ns | 0.00x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Map** | 924ns / 803ns | 0.87x | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **Max** | 923ns / 236ns | 0.26x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **Min** | 930ns / 241ns | 0.26x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **None** | 946ns / 245ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Pipeline F→M→T→R** | 1.1µs / 1.3µs | **1.15x** | 8.3KB / 12.3KB | **1.49x less** | 4 / 2 |
| **Reduce (sum)** | 908ns / 239ns | 0.26x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Reverse** | 825ns / 238ns | 0.29x | 8.2KB / 0B | ∞x more | 3 / 0 |
| **Shuffle** | 4.5µs / 5.7µs | **1.26x** | 16.5KB / 0B | ∞x more | 5 / 0 |
| **Skip** | 722ns / 723ns | ≈ | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **SkipLast** | 728ns / 723ns | ≈ | 8.2KB / 8.2KB | ≈ | 3 / 1 |
| **Sum** | 927ns / 241ns | 0.26x | 8.3KB / 0B | ∞x more | 5 / 0 |
| **Take** | 734ns / <1ns | 0.00x | 8.3KB / 0B | ∞x more | 4 / 0 |
| **ToMap** | 8.5µs / 8.2µs | ≈ | 45.2KB / 37.0KB | 0.82x more | 8 / 6 |
| **Union** | 18.9µs / 18.2µs | ≈ | 106.8KB / 90.3KB | 0.85x more | 17 / 10 |
| **Unique** | 7.2µs / 6.4µs | 0.89x | 53.4KB / 45.1KB | 0.85x more | 9 / 6 |
| **UniqueBy** | 7.5µs / 6.5µs | 0.87x | 53.4KB / 45.1KB | 0.85x more | 10 / 6 |
| **Zip** | 2.8µs / 3.3µs | **1.18x** | 32.9KB / 16.4KB | 0.50x more | 7 / 1 |
| **ZipWith** | 2.4µs / 3.2µs | **1.33x** | 24.7KB / 8.2KB | 0.33x more | 7 / 1 |
