# Benchmarks

Methodology: go1.27.0 on linux/arm64, GOMAXPROCS=16; median of 7 paired samples at 100ms each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±10% equivalence band in the same direction. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 257.6ns / 257.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 257.7ns / 257.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 501.9ns / 2.8µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **Contains** | 262.4ns / 258.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 8.2µs / 8.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.2µs / 8.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 255.9ns / 256.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 2.1µs / 2.0µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 260.7ns / 260.2ns | same loop | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 11.3µs / 11.4µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 257.5ns / 257.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.3µs / 12.4µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.6µs / 1.6µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 498.8ns / 508.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 497.1ns / 507.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 257.9ns / 257.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.7µs / 2.8µs | ≈ | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 257.4ns / 256.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 364.5ns / 365.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 221.6ns / 238.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.4µs / 5.3µs | **3.71x** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.4µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.4µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 259.0ns / 259.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 11.9µs / 12.3µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 349.2ns / 347.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 27.0µs / 30.1µs | ≈ | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 11.4µs / 11.5µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 11.6µs / 11.4µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 1.8µs / 4.6µs | **2.54x** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.2µs / 4.1µs | **3.36x** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | × (faster) | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.5µs / 259.4ns | 0.17x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.4µs / 260.1ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 1.7µs / 2.8µs | **1.66x** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 1.4µs / 260.8ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 9.7µs / 8.2µs | 0.85x | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **CountByValue** | 9.5µs / 8.4µs | ≈ | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 1.4µs / 258.3ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 2.8µs / 1.7µs | 0.60x | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **First** | 1.2µs / 1.8ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.4µs / 262.3ns | 0.19x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 12.6µs / 10.8µs | ≈ | 29.5KB / 21.4KB | 0.72x more | 86 / 85 |
| **IndexWhere** | 1.4µs / 262.7ns | 0.19x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 15.4µs / 12.2µs | 0.79x | 27.7KB / 11.3KB | 0.41x more | 18 / 16 |
| **Last** | 1.3µs / 1.8ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 2.5µs / 1.3µs | 0.51x | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **Max** | 1.6µs / 509.2ns | 0.33x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.7µs / 509.2ns | 0.30x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.7µs / 261.7ns | 0.16x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 3.9µs / 2.4µs | 0.62x | 20.5KB / 12.3KB | 0.60x more | 3 / 2 |
| **Reduce (sum)** | 1.5µs / 259.0ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Retain** | 1.6µs / 409.0ns | 0.26x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.3µs / 239.1ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 2.8µs / 5.3µs | **1.92x** | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.2µs / 1.2µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.2µs / 1.3µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.5µs / 260.2ns | 0.18x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.2µs / 1.9ns | 0.00x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 12.9µs / 12.1µs | ≈ | 45.2KB / 37.0KB | 0.82x more | 7 / 6 |
| **Transform** | 1.4µs / 349.6ns | 0.25x | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Union** | 28.8µs / 28.9µs | ≈ | 106.7KB / 90.3KB | 0.85x more | 12 / 10 |
| **Unique** | 12.4µs / 11.6µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **UniqueBy** | 12.2µs / 10.9µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **Zip** | 4.4µs / 4.4µs | ≈ | 32.8KB / 16.4KB | 0.50x more | 3 / 1 |
| **ZipWith** | 3.8µs / 4.1µs | ≈ | 24.6KB / 8.2KB | 0.33x more | 3 / 1 |
