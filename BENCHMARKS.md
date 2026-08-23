# Benchmarks

Methodology: go1.27.0 on linux/arm64, GOMAXPROCS=1; median of 5 samples at 100ms each. Mutable-input restoration is outside timed regions.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 265.6ns / 265.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 267.0ns / 266.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 202.3ns / 1.5µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **Contains** | 266.9ns / 266.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 22.4µs / 22.3µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 22.3µs / 22.5µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 265.0ns / 520.5ns | **1.96x** | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 513.4ns / 1.2µs | **2.35x** | 0B / 8.2KB | **∞x less** | 0 / 1 |
| **First** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 267.4ns / 267.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 9.1µs / 9.1µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 267.8ns / 267.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.2µs / 12.1µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 276.0ns / 707.1ns | **2.56x** | 0B / 8.2KB | **∞x less** | 0 / 1 |
| **Max** | 511.5ns / 520.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 511.6ns / 520.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 266.9ns / 267.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 490.2ns / 1.4µs | **2.78x** | 0B / 12.3KB | **∞x less** | 0 / 2 |
| **Reduce (sum)** | 265.0ns / 265.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 152.2ns / 176.6ns | **1.16x** | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 3.6µs / 5.3µs | **1.48x** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 541.0ns | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 544.8ns | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 265.4ns / 265.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 2.0ns / 1.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 15.3µs / 17.8µs | **1.16x** | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Union** | 15.8µs / 18.0µs | **1.14x** | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 5.9µs / 6.1µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 6.1µs / 6.0µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 916.5ns / 3.3µs | **3.63x** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 710.1ns / 3.5µs | **4.96x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 786.7ns / 266.3ns | 0.34x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 786.5ns / 267.3ns | 0.34x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 688.3ns / 1.5µs | **2.25x** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 788.5ns / 266.8ns | 0.34x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 23.1µs / 22.4µs | ≈ | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **CountByValue** | 23.0µs / 22.4µs | ≈ | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 805.9ns / 521.7ns | 0.65x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 1.1µs / 1.2µs | **1.15x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 542.8ns / 1.8ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 785.2ns / 267.4ns | 0.34x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 9.7µs / 9.1µs | ≈ | 29.5KB / 21.4KB | 0.72x more | 86 / 85 |
| **IndexWhere** | 786.4ns / 268.7ns | 0.34x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 13.4µs / 12.0µs | 0.90x | 27.7KB / 11.3KB | 0.41x more | 18 / 16 |
| **Last** | 542.5ns / 1.9ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 794.9ns / 709.2ns | 0.89x | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 986.1ns / 519.8ns | 0.53x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 984.6ns / 520.3ns | 0.53x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 788.8ns / 266.3ns | 0.34x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 1.3µs / 1.4µs | ≈ | 8.2KB / 12.3KB | **1.50x less** | 1 / 2 |
| **Reduce (sum)** | 803.0ns / 265.4ns | 0.33x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 672.4ns / 177.4ns | 0.26x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 4.3µs / 5.3µs | **1.25x** | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 478.2ns / 540.8ns | **1.13x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 479.5ns / 544.0ns | **1.13x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 746.5ns / 263.9ns | 0.35x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 480.0ns / 1.9ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 15.9µs / 17.6µs | **1.11x** | 45.2KB / 37.0KB | 0.82x more | 7 / 6 |
| **Union** | 17.6µs / 18.3µs | ≈ | 106.7KB / 90.3KB | 0.85x more | 12 / 10 |
| **Unique** | 6.7µs / 6.1µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **UniqueBy** | 7.0µs / 6.1µs | 0.87x | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **Zip** | 2.0µs / 3.3µs | **1.67x** | 32.8KB / 16.4KB | 0.50x more | 3 / 1 |
| **ZipWith** | 1.8µs / 3.5µs | **1.95x** | 24.6KB / 8.2KB | 0.33x more | 3 / 1 |
