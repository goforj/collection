# Benchmarks

Methodology: go1.27.0 on linux/arm64, CPU=implementer=0x61, part=0x000, architecture=8, GOMAXPROCS=16, lo=v1.53.0; median of 7 paired samples at 100ms each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±10% raw equivalence band in the same direction. The condensed read-only scalar table uses a ±15% band. Both tables label results `below floor` when both timings are below 50ns rather than drawing a relative conclusion from sub-nanosecond deltas. Medians outside the applicable band without consistent paired evidence are labeled `inconclusive`. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 261.4ns / 262.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 262.0ns / 261.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 888.9ns / 5.4µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **CountBy** | 10.1µs / 10.7µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 10.6µs / 10.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 261.0ns / 261.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 3.0µs / 3.5µs | inconclusive | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 265.2ns / 265.3ns | same loop | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 17.8µs / 18.4µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 263.8ns / 265.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 15.3µs / 24.0µs | **1.6x faster** | 11.3KB / 45.1KB | **3.98x less** | 16 / 6 |
| **Last** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 2.2µs / 3.1µs | inconclusive | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 505.8ns / 514.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 507.2ns / 516.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 262.5ns / 261.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 5.3µs / 3.4µs | 1.5x slower | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 262.0ns / 262.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 432.1ns / 390.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 226.2ns / 229.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.5µs / 5.4µs | **3.7x faster** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.6µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.8µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 263.4ns / 262.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 16.8µs / 15.9µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 352.9ns / 343.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 37.9µs / 43.7µs | inconclusive | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **UniqueBy** | 15.3µs / 16.3µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueComparable** | 16.0µs / 15.3µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.7µs / 3.3µs | inconclusive | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.8µs / 3.6µs | **1.9x faster** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **slices.Contains** | 264.6ns / 264.1ns | ≈ | 0B / 0B | ≈ | 0 / 0 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.
