# Benchmarks

Methodology: go1.27.0 on linux/arm64, GOMAXPROCS=16; median of 7 paired samples at 100ms each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±10% raw equivalence band in the same direction. The condensed read-only scalar table uses a ±15% band. Both tables label results `below floor` when both timings are below 50ns rather than drawing a relative conclusion from sub-nanosecond deltas. Medians outside the applicable band without consistent paired evidence are labeled `inconclusive`. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 263.1ns / 263.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 264.1ns / 263.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 530.6ns / 2.8µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **Contains** | 263.0ns / 262.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 8.2µs / 8.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.4µs / 8.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 263.1ns / 262.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 1.9µs / 1.9µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 264.5ns / 265.2ns | same loop | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 11.4µs / 11.1µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 263.1ns / 261.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.4µs / 12.4µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.7µs / 1.6µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 504.1ns / 512.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 502.7ns / 513.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 260.9ns / 261.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.9µs / 2.7µs | ≈ | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 260.5ns / 260.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 367.1ns / 368.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 222.7ns / 255.3ns | inconclusive | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.4µs / 5.3µs | **3.7x faster** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.5µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.4µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 261.9ns / 261.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 11.4µs / 12.1µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 351.6ns / 351.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 27.7µs / 29.4µs | ≈ | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 11.3µs / 11.3µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 11.6µs / 11.6µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.3µs / 4.8µs | **2.1x faster** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.3µs / 4.1µs | **3.1x faster** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.5µs / 261.7ns | 5.6x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.6µs / 262.5ns | 6.1x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 1.8µs / 2.9µs | **1.6x faster** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 1.5µs / 263.5ns | 5.7x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 9.8µs / 8.2µs | 1.2x slower | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **CountByValue** | 10.1µs / 8.3µs | inconclusive | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 1.6µs / 262.5ns | 6.1x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 2.8µs / 1.8µs | 1.5x slower | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **First** | 1.3µs / 1.8ns | 696.3x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.5µs / 265.7ns | 5.5x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 13.6µs / 11.3µs | inconclusive | 29.5KB / 21.4KB | 0.72x more | 86 / 85 |
| **IndexWhere** | 1.5µs / 264.9ns | 5.6x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 15.6µs / 12.8µs | 1.2x slower | 27.7KB / 11.3KB | 0.41x more | 18 / 16 |
| **Last** | 1.3µs / 1.8ns | 703.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 2.8µs / 1.3µs | 2.1x slower | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **Max** | 1.8µs / 518.9ns | 3.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.8µs / 518.4ns | 3.5x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.6µs / 260.1ns | 6.0x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 3.9µs / 2.7µs | 1.4x slower | 20.5KB / 12.3KB | 0.60x more | 3 / 2 |
| **Reduce (sum)** | 1.3µs / 263.0ns | 5.1x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Retain** | 1.5µs / 371.1ns | 4.1x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.3µs / 252.4ns | 5.3x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 2.9µs / 5.4µs | **1.8x faster** | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.2µs / 1.3µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.2µs / 1.3µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.8µs / 266.8ns | 6.6x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.2µs / 1.9ns | 621.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 13.9µs / 13.4µs | ≈ | 45.2KB / 37.0KB | 0.82x more | 7 / 6 |
| **Transform** | 1.6µs / 351.0ns | 4.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Union** | 31.1µs / 30.1µs | ≈ | 106.7KB / 90.3KB | 0.85x more | 12 / 10 |
| **Unique** | 12.5µs / 11.8µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **UniqueBy** | 12.5µs / 11.9µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **Zip** | 4.3µs / 4.6µs | ≈ | 32.8KB / 16.4KB | 0.50x more | 3 / 1 |
| **ZipWith** | 3.8µs / 4.2µs | inconclusive | 24.6KB / 8.2KB | 0.33x more | 3 / 1 |
