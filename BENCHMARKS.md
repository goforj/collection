# Benchmarks

Methodology: go1.27.0 on linux/arm64, GOMAXPROCS=16; median of 5 samples at 100ms each. Mutable scratch initialization is outside timed regions.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 267.2ns / 269.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 267.8ns / 265.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 544.6ns / 3.3µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **Contains** | 266.0ns / 265.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 8.4µs / 8.4µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.4µs / 8.4µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 264.1ns / 264.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 2.2µs / 2.2µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 267.1ns / 268.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 11.7µs / 11.9µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 266.7ns / 268.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.7µs / 12.4µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.6µs / 1.8µs | **1.13x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 509.7ns / 519.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 514.5ns / 518.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 266.4ns / 264.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.7µs / 2.9µs | ≈ | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 264.0ns / 263.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 357.1ns / 356.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 142.9ns / 165.0ns | **1.16x** | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 5.0µs / 5.3µs | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.7µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.8µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 264.4ns / 264.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 12.3µs / 12.5µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 269.9ns / 269.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 27.0µs / 29.9µs | **1.11x** | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 12.0µs / 12.1µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 12.6µs / 11.4µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.8µs / 5.0µs | **1.77x** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.4µs / 4.2µs | **3.01x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.6µs / 265.4ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.7µs / 266.0ns | 0.16x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 1.9µs / 3.4µs | **1.83x** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 1.5µs / 265.9ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 10.2µs / 8.4µs | 0.83x | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **CountByValue** | 10.6µs / 8.3µs | 0.79x | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 1.7µs / 267.7ns | 0.16x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 3.2µs / 1.9µs | 0.61x | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **First** | 1.2µs / 1.8ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.6µs / 271.4ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 14.0µs / 11.3µs | 0.81x | 29.5KB / 21.4KB | 0.72x more | 86 / 85 |
| **IndexWhere** | 1.5µs / 271.2ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 15.6µs / 12.7µs | 0.81x | 27.7KB / 11.3KB | 0.41x more | 18 / 16 |
| **Last** | 1.3µs / 1.8ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 2.9µs / 1.5µs | 0.51x | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **Max** | 1.8µs / 517.9ns | 0.29x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.8µs / 519.0ns | 0.28x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.4µs / 266.3ns | 0.19x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 4.5µs / 2.4µs | 0.54x | 20.5KB / 12.3KB | 0.60x more | 3 / 2 |
| **Reduce (sum)** | 1.6µs / 266.9ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Retain** | 1.5µs / 356.0ns | 0.23x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.3µs / 164.1ns | 0.12x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 6.0µs / 5.3µs | 0.88x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.4µs / 1.3µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.2µs / 1.4µs | **1.15x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.5µs / 264.1ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.1µs / 1.9ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 13.5µs / 12.7µs | ≈ | 45.2KB / 37.0KB | 0.82x more | 7 / 6 |
| **Transform** | 1.5µs / 270.4ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Union** | 30.3µs / 30.1µs | ≈ | 106.7KB / 90.3KB | 0.85x more | 12 / 10 |
| **Unique** | 12.9µs / 12.1µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **UniqueBy** | 13.6µs / 12.2µs | 0.90x | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **Zip** | 4.3µs / 4.7µs | ≈ | 32.8KB / 16.4KB | 0.50x more | 3 / 1 |
| **ZipWith** | 4.1µs / 4.2µs | ≈ | 24.6KB / 8.2KB | 0.33x more | 3 / 1 |
