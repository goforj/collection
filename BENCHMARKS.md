# Benchmarks

Methodology: go1.27.0 on linux/arm64, GOMAXPROCS=16; median of 7 paired samples at 100ms each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±10% raw equivalence band in the same direction. The condensed read-only scalar table uses a ±15% band. Both tables label results `below floor` when both timings are below 50ns rather than drawing a relative conclusion from sub-nanosecond deltas. Medians outside the applicable band without consistent paired evidence are labeled `inconclusive`. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 263.3ns / 263.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 264.2ns / 263.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 540.6ns / 3.3µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **Contains** | 269.1ns / 264.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 8.3µs / 8.6µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.2µs / 8.3µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 262.0ns / 262.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 2.2µs / 2.1µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 266.5ns / 265.8ns | same loop | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 11.3µs / 11.1µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 264.1ns / 264.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.4µs / 12.3µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.7µs / 1.8µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 505.3ns / 515.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 505.2ns / 514.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 254.6ns / 255.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.9µs / 3.0µs | ≈ | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 262.9ns / 262.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 368.6ns / 371.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 222.4ns / 254.2ns | inconclusive | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.4µs / 5.4µs | **3.7x faster** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.7µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.5µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 263.0ns / 263.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 11.5µs / 12.5µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 352.2ns / 351.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 27.2µs / 29.8µs | ≈ | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 12.1µs / 11.9µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 11.7µs / 11.7µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.1µs / 4.8µs | **2.3x faster** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.5µs / 4.0µs | **2.8x faster** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.6µs / 260.6ns | 6.0x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.6µs / 265.1ns | 5.9x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 1.9µs / 3.4µs | **1.8x faster** | 9.5KB / 9.3KB | 1.02x more | 2 / 51 |
| **Contains** | 1.7µs / 265.0ns | 6.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 9.8µs / 8.2µs | inconclusive | 17.7KB / 9.5KB | 1.86x more | 14 / 13 |
| **CountByValue** | 9.9µs / 8.3µs | 1.2x slower | 17.7KB / 9.5KB | 1.86x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 1.6µs / 264.2ns | 6.1x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 3.1µs / 1.8µs | 1.7x slower | 16.4KB / 8.2KB | 2.00x more | 2 / 1 |
| **First** | 1.4µs / 1.8ns | 765.2x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.6µs / 266.4ns | 6.1x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 13.6µs / 11.7µs | inconclusive | 29.5KB / 21.4KB | 1.38x more | 86 / 85 |
| **IndexWhere** | 1.6µs / 266.0ns | 6.2x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 15.5µs / 12.3µs | 1.3x slower | 27.7KB / 11.3KB | 2.45x more | 18 / 16 |
| **Last** | 1.5µs / 1.8ns | 836.6x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 2.9µs / 1.5µs | 2.0x slower | 16.4KB / 8.2KB | 2.00x more | 2 / 1 |
| **Max** | 1.9µs / 516.8ns | 3.7x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.7µs / 516.4ns | 3.3x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.7µs / 263.7ns | 6.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 4.3µs / 2.6µs | 1.6x slower | 20.5KB / 12.3KB | 1.67x more | 3 / 2 |
| **Reduce (sum)** | 1.5µs / 263.0ns | 5.8x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Retain** | 1.7µs / 372.3ns | 4.5x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.6µs / 254.2ns | 6.2x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 2.9µs / 5.4µs | **1.9x faster** | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.2µs / 1.4µs | inconclusive | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.4µs / 1.3µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.6µs / 264.2ns | 6.2x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.4µs / 1.9ns | 710.1x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 13.1µs / 12.0µs | ≈ | 45.2KB / 37.0KB | 1.22x more | 7 / 6 |
| **Transform** | 1.7µs / 351.5ns | 4.8x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Union** | 29.5µs / 28.3µs | ≈ | 106.7KB / 90.3KB | 1.18x more | 12 / 10 |
| **Unique** | 13.0µs / 11.9µs | ≈ | 53.3KB / 45.1KB | 1.18x more | 7 / 6 |
| **UniqueBy** | 12.8µs / 12.0µs | ≈ | 53.3KB / 45.1KB | 1.18x more | 7 / 6 |
| **Zip** | 4.9µs / 4.6µs | ≈ | 32.8KB / 16.4KB | 2.00x more | 3 / 1 |
| **ZipWith** | 4.4µs / 4.1µs | ≈ | 24.6KB / 8.2KB | 3.00x more | 3 / 1 |
