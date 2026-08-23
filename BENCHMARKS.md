# Benchmarks

Methodology: go1.27.0 on linux/arm64, CPU=implementer=0x61, part=0x000, architecture=8, GOMAXPROCS=16, lo=v1.52.0; median of 7 paired samples at 100ms each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±10% raw equivalence band in the same direction. The condensed read-only scalar table uses a ±15% band. Both tables label results `below floor` when both timings are below 50ns rather than drawing a relative conclusion from sub-nanosecond deltas. Medians outside the applicable band without consistent paired evidence are labeled `inconclusive`. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 262.1ns / 263.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 253.8ns / 255.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 527.8ns / 3.3µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **CountBy** | 8.4µs / 8.2µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.3µs / 8.8µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 263.9ns / 264.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 2.2µs / 2.1µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 265.5ns / 265.8ns | same loop | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 11.1µs / 11.4µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 265.4ns / 264.8ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.5µs / 12.3µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.8µs / 1.7µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 508.9ns / 520.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 506.6ns / 517.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 263.3ns / 262.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.9µs / 2.8µs | ≈ | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 263.5ns / 263.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 616.1ns / 380.0ns | 1.6x slower | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 210.0ns / 233.5ns | **1.1x faster** | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.5µs / 5.4µs | **3.7x faster** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.8ns / 1.8µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.8ns / 1.4µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 263.9ns / 263.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 12.1µs / 12.5µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 342.6ns / 350.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 27.5µs / 30.5µs | inconclusive | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **UniqueBy** | 12.1µs / 11.9µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueComparable** | 11.8µs / 12.1µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.4µs / 4.9µs | inconclusive | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.4µs / 4.0µs | **2.9x faster** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **slices.Contains** | 265.5ns / 264.2ns | ≈ | 0B / 0B | ≈ | 0 / 0 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.
