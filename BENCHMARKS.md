# Benchmarks

Methodology: go1.27.0 on linux/arm64, GOMAXPROCS=16; median of 7 paired samples at 100ms each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±10% equivalence band in the same direction. Medians inside the band are labeled `≈`; medians outside it without consistent paired evidence are labeled `inconclusive`. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 258.0ns / 258.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 251.5ns / 256.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 550.4ns / 3.2µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **Contains** | 262.4ns / 260.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 8.1µs / 8.5µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.3µs / 8.3µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 257.6ns / 258.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 1.9µs / 2.1µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 262.4ns / 261.8ns | same loop | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 11.2µs / 11.5µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 259.1ns / 260.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.4µs / 12.3µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.6µs / 1.6µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 499.4ns / 508.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 499.8ns / 509.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 246.8ns / 243.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.8µs / 2.7µs | ≈ | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 257.2ns / 257.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 363.8ns / 368.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 221.8ns / 240.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.4µs / 5.3µs | **faster** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.6µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.3µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 259.8ns / 260.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 12.6µs / 13.2µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 349.3ns / 347.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 28.2µs / 29.6µs | ≈ | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 11.6µs / 11.9µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 11.9µs / 12.2µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.1µs / 4.9µs | **faster** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.4µs / 4.2µs | **faster** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.5µs / 260.4ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.5µs / 263.2ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 1.8µs / 3.1µs | **faster** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 1.4µs / 261.2ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 10.2µs / 8.4µs | inconclusive | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **CountByValue** | 10.1µs / 8.2µs | slower | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 1.5µs / 259.5ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 3.0µs / 1.8µs | slower | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **First** | 1.3µs / 1.8ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.4µs / 263.3ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 13.4µs / 11.6µs | inconclusive | 29.5KB / 21.4KB | 0.72x more | 86 / 85 |
| **IndexWhere** | 1.5µs / 262.4ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 15.4µs / 12.3µs | slower | 27.7KB / 11.3KB | 0.41x more | 18 / 16 |
| **Last** | 1.3µs / 1.8ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 2.7µs / 1.3µs | slower | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **Max** | 1.8µs / 509.5ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.7µs / 509.9ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.5µs / 261.1ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 4.2µs / 2.8µs | slower | 20.5KB / 12.3KB | 0.60x more | 3 / 2 |
| **Reduce (sum)** | 1.5µs / 260.1ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Retain** | 1.7µs / 368.4ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.4µs / 253.8ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 2.8µs / 5.3µs | **faster** | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.2µs / 1.2µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.1µs / 1.3µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.5µs / 259.9ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.3µs / 1.9ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 13.5µs / 12.3µs | ≈ | 45.2KB / 37.0KB | 0.82x more | 7 / 6 |
| **Transform** | 1.7µs / 350.0ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Union** | 30.0µs / 30.5µs | ≈ | 106.7KB / 90.3KB | 0.85x more | 12 / 10 |
| **Unique** | 12.5µs / 11.7µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **UniqueBy** | 13.3µs / 12.7µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **Zip** | 4.2µs / 4.6µs | ≈ | 32.8KB / 16.4KB | 0.50x more | 3 / 1 |
| **ZipWith** | 3.7µs / 4.0µs | ≈ | 24.6KB / 8.2KB | 0.33x more | 3 / 1 |
