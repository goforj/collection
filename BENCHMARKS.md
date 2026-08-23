# Benchmarks

Methodology: go1.27.0 on linux/arm64, GOMAXPROCS=16; median of 5 samples at 100ms each. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 259.6ns / 259.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 259.7ns / 260.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 558.4ns / 3.5µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **Contains** | 262.1ns / 261.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 8.3µs / 8.1µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.2µs / 8.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 258.8ns / 258.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 2.1µs / 2.3µs | **1.11x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 507.0ns / 262.8ns | 0.52x | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 11.8µs / 11.8µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 260.9ns / 260.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.2µs / 12.4µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.6µs / 1.7µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 499.9ns / 509.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 500.8ns / 507.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 259.5ns / 259.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.8µs / 3.0µs | ≈ | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 259.2ns / 258.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 369.6ns / 421.3ns | **1.14x** | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 221.7ns / 238.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.4µs / 5.3µs | **3.71x** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.6µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.5µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 259.4ns / 259.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 11.7µs / 11.8µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 349.8ns / 348.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 28.7µs / 30.1µs | ≈ | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 11.9µs / 11.9µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 11.8µs / 12.2µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.8µs / 4.6µs | **1.67x** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.4µs / 4.2µs | **3.04x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.5µs / 260.3ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.6µs / 260.7ns | 0.16x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 2.0µs / 3.1µs | **1.58x** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 1.5µs / 261.8ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 9.8µs / 8.1µs | 0.83x | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **CountByValue** | 10.1µs / 8.5µs | 0.84x | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 1.8µs / 259.4ns | 0.15x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 3.3µs / 2.0µs | 0.60x | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **First** | 1.3µs / 1.8ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.4µs / 262.5ns | 0.19x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 14.0µs / 11.8µs | 0.84x | 29.5KB / 21.4KB | 0.72x more | 86 / 85 |
| **IndexWhere** | 1.6µs / 262.7ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 16.2µs / 12.1µs | 0.75x | 27.7KB / 11.3KB | 0.41x more | 18 / 16 |
| **Last** | 1.6µs / 1.8ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 3.0µs / 1.3µs | 0.45x | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **Max** | 1.7µs / 510.1ns | 0.29x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.7µs / 509.2ns | 0.30x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.5µs / 260.2ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 4.2µs / 2.5µs | 0.61x | 20.5KB / 12.3KB | 0.60x more | 3 / 2 |
| **Reduce (sum)** | 1.6µs / 259.6ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Retain** | 1.5µs / 370.1ns | 0.25x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.4µs / 238.6ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 2.9µs / 5.3µs | **1.84x** | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.2µs / 1.3µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.2µs / 1.3µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.6µs / 259.3ns | 0.16x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.2µs / 1.9ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 13.2µs / 12.1µs | ≈ | 45.2KB / 37.0KB | 0.82x more | 7 / 6 |
| **Transform** | 1.5µs / 349.6ns | 0.23x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Union** | 31.2µs / 29.8µs | ≈ | 106.7KB / 90.3KB | 0.85x more | 12 / 10 |
| **Unique** | 12.8µs / 12.1µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **UniqueBy** | 13.3µs / 11.8µs | 0.88x | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **Zip** | 4.8µs / 4.3µs | 0.90x | 32.8KB / 16.4KB | 0.50x more | 3 / 1 |
| **ZipWith** | 4.2µs / 4.1µs | ≈ | 24.6KB / 8.2KB | 0.33x more | 3 / 1 |
