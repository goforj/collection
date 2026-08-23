# Benchmarks

Methodology: go1.27.0 on linux/arm64, GOMAXPROCS=16; median of 7 paired samples at 100ms each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±10% raw equivalence band in the same direction. The condensed read-only scalar table uses a ±15% band. Both tables label results `below floor` when both timings are below 50ns rather than drawing a relative conclusion from sub-nanosecond deltas. Medians outside the applicable band without consistent paired evidence are labeled `inconclusive`. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 261.6ns / 262.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 261.7ns / 262.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 496.3ns / 2.2µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **Contains** | 264.0ns / 261.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 8.1µs / 8.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.2µs / 8.0µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 260.8ns / 260.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 2.0µs / 2.0µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 264.1ns / 262.4ns | same loop | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 10.5µs / 10.2µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 262.9ns / 263.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.3µs / 12.6µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.5µs / 1.4µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 502.1ns / 510.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 503.3ns / 511.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 262.0ns / 261.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.5µs / 2.6µs | ≈ | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 260.6ns / 260.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 414.6ns / 374.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 222.1ns / 239.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.4µs / 5.3µs | **faster** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.3µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.4µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 261.3ns / 261.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 12.1µs / 12.4µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 351.0ns / 350.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 27.2µs / 29.8µs | ≈ | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 11.0µs / 11.9µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 11.5µs / 11.6µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.0µs / 4.7µs | **faster** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.2µs / 4.0µs | **faster** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.2µs / 262.1ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.5µs / 262.7ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 1.6µs / 2.2µs | **faster** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 1.5µs / 262.1ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 9.9µs / 8.1µs | slower | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **CountByValue** | 9.6µs / 8.0µs | slower | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 1.3µs / 260.6ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 2.8µs / 1.9µs | inconclusive | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **First** | 1.3µs / 1.8ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.3µs / 263.1ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 12.3µs / 10.6µs | inconclusive | 29.5KB / 21.4KB | 0.72x more | 86 / 85 |
| **IndexWhere** | 1.4µs / 263.2ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 15.0µs / 12.2µs | slower | 27.7KB / 11.3KB | 0.41x more | 18 / 16 |
| **Last** | 1.2µs / 1.8ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 2.5µs / 1.2µs | slower | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **Max** | 1.6µs / 489.3ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.5µs / 511.3ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.4µs / 261.8ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 3.7µs / 2.4µs | slower | 20.5KB / 12.3KB | 0.60x more | 3 / 2 |
| **Reduce (sum)** | 1.2µs / 260.8ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Retain** | 1.4µs / 362.8ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.2µs / 253.6ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 2.8µs / 5.3µs | **faster** | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.1µs / 1.1µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.1µs / 1.1µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.4µs / 261.4ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.2µs / 1.9ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 13.4µs / 12.2µs | ≈ | 45.2KB / 37.0KB | 0.82x more | 7 / 6 |
| **Transform** | 1.4µs / 349.9ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Union** | 29.8µs / 28.7µs | ≈ | 106.7KB / 90.3KB | 0.85x more | 12 / 10 |
| **Unique** | 12.1µs / 11.6µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **UniqueBy** | 12.9µs / 11.6µs | inconclusive | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **Zip** | 3.9µs / 4.6µs | inconclusive | 32.8KB / 16.4KB | 0.50x more | 3 / 1 |
| **ZipWith** | 3.7µs / 3.9µs | ≈ | 24.6KB / 8.2KB | 0.33x more | 3 / 1 |
