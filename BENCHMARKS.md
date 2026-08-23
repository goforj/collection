# Benchmarks

Methodology: go1.27.0 on linux/arm64, GOMAXPROCS=16; median of 7 paired samples at 100ms each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±10% raw equivalence band in the same direction. The condensed read-only scalar table uses a ±15% band. Both tables label results `below floor` when both timings are below 50ns rather than drawing a relative conclusion from sub-nanosecond deltas. Medians outside the applicable band without consistent paired evidence are labeled `inconclusive`. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 264.8ns / 265.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 265.4ns / 265.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 503.4ns / 2.4µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **Contains** | 268.0ns / 265.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 8.3µs / 8.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.3µs / 8.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 264.4ns / 264.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 2.0µs / 2.0µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 268.5ns / 267.7ns | same loop | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 10.5µs / 10.5µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 267.4ns / 267.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.2µs / 12.2µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.8µs / 1.6µs | inconclusive | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 505.9ns / 514.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 500.8ns / 508.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 264.8ns / 264.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.4µs / 2.8µs | inconclusive | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 266.4ns / 267.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 368.7ns / 369.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 224.5ns / 255.9ns | **1.1x faster** | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.5µs / 5.4µs | **3.7x faster** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.4µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.2µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 260.1ns / 260.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 11.9µs / 12.3µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 353.5ns / 352.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 27.4µs / 29.2µs | ≈ | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 11.6µs / 11.6µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 11.6µs / 11.1µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.1µs / 4.7µs | **2.2x faster** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.4µs / 4.1µs | **3.0x faster** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.5µs / 264.1ns | 5.5x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.4µs / 264.4ns | 5.5x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 1.7µs / 2.2µs | **1.3x faster** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 1.3µs / 259.9ns | 4.9x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 9.8µs / 8.3µs | inconclusive | 17.7KB / 9.5KB | 1.86x more | 14 / 13 |
| **CountByValue** | 9.8µs / 8.1µs | 1.2x slower | 17.7KB / 9.5KB | 1.86x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 1.4µs / 262.7ns | 5.2x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 2.9µs / 1.7µs | 1.7x slower | 16.4KB / 8.2KB | 2.00x more | 2 / 1 |
| **First** | 1.2µs / 1.8ns | 680.3x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.4µs / 265.1ns | 5.3x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 12.4µs / 10.9µs | inconclusive | 29.5KB / 21.4KB | 1.38x more | 86 / 85 |
| **IndexWhere** | 1.5µs / 266.0ns | 5.5x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 15.3µs / 12.3µs | 1.2x slower | 27.7KB / 11.3KB | 2.45x more | 18 / 16 |
| **Last** | 1.3µs / 1.8ns | 725.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 2.5µs / 1.3µs | 1.9x slower | 16.4KB / 8.2KB | 2.00x more | 2 / 1 |
| **Max** | 1.7µs / 517.1ns | 3.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.7µs / 517.6ns | 3.3x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.5µs / 263.5ns | 5.6x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 3.8µs / 2.6µs | 1.5x slower | 20.5KB / 12.3KB | 1.67x more | 3 / 2 |
| **Reduce (sum)** | 1.4µs / 262.7ns | 5.2x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Retain** | 1.5µs / 369.8ns | 4.1x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.3µs / 254.5ns | 5.2x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 2.8µs / 5.3µs | **1.9x faster** | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.2µs / 1.2µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.2µs / 1.2µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.4µs / 265.7ns | 5.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.1µs / 1.9ns | 564.3x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 13.1µs / 12.6µs | ≈ | 45.2KB / 37.0KB | 1.22x more | 7 / 6 |
| **Transform** | 1.5µs / 351.2ns | 4.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Union** | 29.8µs / 29.1µs | ≈ | 106.7KB / 90.3KB | 1.18x more | 12 / 10 |
| **Unique** | 12.1µs / 11.2µs | ≈ | 53.3KB / 45.1KB | 1.18x more | 7 / 6 |
| **UniqueBy** | 13.4µs / 11.5µs | inconclusive | 53.3KB / 45.1KB | 1.18x more | 7 / 6 |
| **Zip** | 4.5µs / 4.5µs | ≈ | 32.8KB / 16.4KB | 2.00x more | 3 / 1 |
| **ZipWith** | 3.9µs / 4.0µs | ≈ | 24.6KB / 8.2KB | 3.00x more | 3 / 1 |
