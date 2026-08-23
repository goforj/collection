# Benchmarks

Methodology: go1.27.0 on linux/arm64, GOMAXPROCS=16; median of 7 paired samples at 100ms each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±10% raw equivalence band in the same direction. The condensed read-only scalar table uses a ±15% band. Both tables label timings `≈` when both implementations are below 50ns, where ratios amplify sub-nanosecond noise. Medians outside the applicable band without consistent paired evidence are labeled `inconclusive`. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 261.8ns / 262.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 262.3ns / 262.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 501.9ns / 2.4µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **Contains** | 264.9ns / 261.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **CountBy** | 8.2µs / 8.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.1µs / 8.1µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 259.4ns / 255.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 1.9µs / 1.9µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 263.6ns / 263.1ns | same loop | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 10.7µs / 10.7µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 262.3ns / 262.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.1µs / 12.2µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.4µs / 1.6µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 502.6ns / 509.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 501.7ns / 510.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 261.4ns / 261.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.7µs / 2.5µs | ≈ | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 260.1ns / 259.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 366.8ns / 411.1ns | inconclusive | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 221.4ns / 239.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.4µs / 5.3µs | **faster** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.4µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.2µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 260.6ns / 260.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 12.2µs / 12.4µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 350.2ns / 349.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 28.2µs / 29.1µs | ≈ | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **Unique** | 11.8µs / 11.5µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueBy** | 11.1µs / 11.3µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 1.8µs / 4.8µs | **faster** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.2µs / 4.1µs | **faster** | 8.2KB / 8.2KB | ≈ | 1 / 1 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.4µs / 261.9ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.4µs / 262.2ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 1.7µs / 2.3µs | **faster** | 9.5KB / 9.3KB | ≈ | 2 / 51 |
| **Contains** | 1.4µs / 262.6ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **CountBy** | 9.8µs / 8.1µs | slower | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **CountByValue** | 9.6µs / 8.1µs | inconclusive | 17.7KB / 9.5KB | 0.54x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 1.4µs / 260.4ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 2.6µs / 1.7µs | slower | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **First** | 1.3µs / 1.8ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.3µs / 263.7ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 11.6µs / 10.6µs | ≈ | 29.5KB / 21.4KB | 0.72x more | 86 / 85 |
| **IndexWhere** | 1.3µs / 253.0ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 14.9µs / 11.9µs | slower | 27.7KB / 11.3KB | 0.41x more | 18 / 16 |
| **Last** | 1.1µs / 1.8ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 2.4µs / 1.2µs | slower | 16.4KB / 8.2KB | 0.50x more | 2 / 1 |
| **Max** | 1.5µs / 511.2ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.6µs / 510.9ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.5µs / 261.4ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 4.0µs / 2.3µs | slower | 20.5KB / 12.3KB | 0.60x more | 3 / 2 |
| **Reduce (sum)** | 1.4µs / 260.3ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Retain** | 1.5µs / 368.9ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.2µs / 253.0ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 2.8µs / 5.3µs | **faster** | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.1µs / 1.2µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.1µs / 1.2µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.4µs / 261.1ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.2µs / 1.9ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 13.4µs / 11.9µs | inconclusive | 45.2KB / 37.0KB | 0.82x more | 7 / 6 |
| **Transform** | 1.5µs / 349.7ns | slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Union** | 29.2µs / 30.0µs | ≈ | 106.7KB / 90.3KB | 0.85x more | 12 / 10 |
| **Unique** | 13.3µs / 12.2µs | ≈ | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **UniqueBy** | 13.1µs / 11.5µs | inconclusive | 53.3KB / 45.1KB | 0.85x more | 7 / 6 |
| **Zip** | 4.1µs / 4.6µs | inconclusive | 32.8KB / 16.4KB | 0.50x more | 3 / 1 |
| **ZipWith** | 3.6µs / 4.0µs | inconclusive | 24.6KB / 8.2KB | 0.33x more | 3 / 1 |
