# Benchmarks

Methodology: go1.27.0 on linux/arm64, CPU=implementer=0x61, part=0x000, architecture=8, GOMAXPROCS=16, lo=v1.52.0; median of 7 paired samples at 100ms each, alternating implementation order. Timing differences are shown only when every pair falls outside the ±10% raw equivalence band in the same direction. The condensed read-only scalar table uses a ±15% band. Both tables label results `below floor` when both timings are below 50ns rather than drawing a relative conclusion from sub-nanosecond deltas. Medians outside the applicable band without consistent paired evidence are labeled `inconclusive`. Mutable borrowed inputs are restored inside every timed iteration for both implementations.

Raw results for `collection.New` (borrowed) vs `lo`. For Chunk, Skip, and SkipLast, collection returns a view while lo returns a copy; those rows describe an ownership and allocation trade-off, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

FirstWhere compiles to the same scan loop in both implementations. Its ratio is labeled `same loop` because binary placement can dominate the timing of such a small function in this in-process harness.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 266.1ns / 265.6ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Any** | 267.5ns / 267.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Chunk** | 496.7ns / 2.2µs | view trade-off | 1.3KB / 9.3KB | ownership trade-off | ownership trade-off |
| **CountBy** | 8.4µs / 8.3µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **CountByValue** | 8.4µs / 8.3µs | ≈ | 9.5KB / 9.5KB | ≈ | 13 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 266.0ns / 265.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Filter** | 1.9µs / 2.0µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **First** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **FirstWhere** | 263.2ns / 264.0ns | same loop | 0B / 0B | ≈ | 0 / 0 |
| **GroupBy** | 10.8µs / 10.6µs | ≈ | 21.4KB / 21.4KB | ≈ | 85 / 85 |
| **IndexWhere** | 266.4ns / 269.4ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Intersect** | 12.5µs / 12.4µs | ≈ | 11.3KB / 11.3KB | ≈ | 16 / 16 |
| **Last** | 1.8ns / 1.8ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **Map** | 1.6µs / 1.6µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Max** | 510.5ns / 525.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Min** | 513.8ns / 521.0ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **None** | 267.4ns / 269.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Pipeline F→M→T→R** | 2.7µs / 2.7µs | ≈ | 12.3KB / 12.3KB | ≈ | 2 / 2 |
| **Reduce (sum)** | 266.1ns / 264.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Retain** | 427.7ns / 378.8ns | inconclusive | 0B / 0B | ≈ | 0 / 0 |
| **Reverse** | 226.1ns / 228.3ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Shuffle** | 1.4µs / 5.4µs | **3.7x faster** | 0B / 0B | ≈ | 0 / 0 |
| **Skip** | 1.9ns / 1.5µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **SkipLast** | 1.9ns / 1.4µs | view trade-off | 0B / 8.2KB | ownership trade-off | ownership trade-off |
| **Sum** | 264.2ns / 264.5ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Take** | 1.8ns / 1.9ns | below floor | 0B / 0B | ≈ | 0 / 0 |
| **ToMap** | 11.4µs / 12.1µs | ≈ | 37.0KB / 37.0KB | ≈ | 6 / 6 |
| **Transform** | 353.0ns / 342.7ns | ≈ | 0B / 0B | ≈ | 0 / 0 |
| **Union** | 27.7µs / 30.0µs | ≈ | 90.3KB / 90.3KB | ≈ | 10 / 10 |
| **UniqueBy** | 12.7µs / 12.0µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **UniqueComparable** | 11.9µs / 12.4µs | ≈ | 45.1KB / 45.1KB | ≈ | 6 / 6 |
| **Zip** | 2.0µs / 5.0µs | **2.5x faster** | 16.4KB / 16.4KB | ≈ | 1 / 1 |
| **ZipWith** | 1.3µs / 3.9µs | **3.1x faster** | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **slices.Contains** | 265.1ns / 264.9ns | ≈ | 0B / 0B | ≈ | 0 / 0 |

Chunk, Skip, and SkipLast return collection views while lo returns copied slices. Their rows describe ownership and allocation trade-offs, not equal-work speed superiority. Difference returns one-sided output while lo returns both sides, so its rows are an API trade-off.

Raw results for `collection.New().Clone()` (explicit copy) vs `lo`. This section includes collection's explicit input-copy cost.

| Op | ns/op (vs lo) | Timing | bytes/op (vs lo) | × (less memory) | allocs/op (vs lo) |
|---:|----------------|:--:|------------------|:--:|--------------------|
| **All** | 1.4µs / 266.5ns | 5.3x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Any** | 1.5µs / 268.6ns | 5.5x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Chunk** | 1.7µs / 2.2µs | inconclusive | 9.5KB / 9.3KB | 1.02x more | 2 / 51 |
| **CountBy** | 9.6µs / 8.3µs | inconclusive | 17.7KB / 9.5KB | 1.86x more | 14 / 13 |
| **CountByValue** | 10.1µs / 8.6µs | 1.2x slower | 17.7KB / 9.5KB | 1.86x more | 14 / 13 |
| **Difference** | different work | API trade-off | different work | API trade-off | API trade-off |
| **Each** | 1.4µs / 265.2ns | 5.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Filter** | 2.6µs / 1.7µs | 1.5x slower | 16.4KB / 8.2KB | 2.00x more | 2 / 1 |
| **First** | 1.4µs / 1.8ns | 766.2x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **FirstWhere** | 1.5µs / 268.8ns | 5.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **GroupBy** | 12.1µs / 10.3µs | inconclusive | 29.5KB / 21.4KB | 1.38x more | 86 / 85 |
| **IndexWhere** | 1.4µs / 270.0ns | 5.3x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Intersect** | 15.1µs / 12.3µs | 1.2x slower | 27.7KB / 11.3KB | 2.45x more | 18 / 16 |
| **Last** | 1.2µs / 1.8ns | 660.9x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Map** | 2.4µs / 1.2µs | 2.0x slower | 16.4KB / 8.2KB | 2.00x more | 2 / 1 |
| **Max** | 1.8µs / 520.2ns | 3.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Min** | 1.8µs / 520.5ns | 3.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **None** | 1.5µs / 266.7ns | 5.4x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Pipeline F→M→T→R** | 4.1µs / 2.6µs | 1.6x slower | 20.5KB / 12.3KB | 1.67x more | 3 / 2 |
| **Reduce (sum)** | 1.4µs / 267.9ns | 5.1x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Retain** | 1.4µs / 373.7ns | 3.7x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Reverse** | 1.2µs / 228.6ns | 5.2x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Shuffle** | 2.7µs / 5.4µs | **2.0x faster** | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Skip** | 1.2µs / 1.2µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **SkipLast** | 1.1µs / 1.1µs | ≈ | 8.2KB / 8.2KB | ≈ | 1 / 1 |
| **Sum** | 1.5µs / 266.2ns | 5.5x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Take** | 1.2µs / 1.9ns | 628.0x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **ToMap** | 14.2µs / 12.5µs | inconclusive | 45.2KB / 37.0KB | 1.22x more | 7 / 6 |
| **Transform** | 1.4µs / 344.5ns | 4.1x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
| **Union** | 28.4µs / 28.8µs | ≈ | 106.7KB / 90.3KB | 1.18x more | 12 / 10 |
| **UniqueBy** | 12.8µs / 11.6µs | ≈ | 53.3KB / 45.1KB | 1.18x more | 7 / 6 |
| **UniqueComparable** | 12.4µs / 11.3µs | ≈ | 53.3KB / 45.1KB | 1.18x more | 7 / 6 |
| **Zip** | 4.0µs / 4.4µs | inconclusive | 32.8KB / 16.4KB | 2.00x more | 3 / 1 |
| **ZipWith** | 3.6µs / 4.0µs | inconclusive | 24.6KB / 8.2KB | 3.00x more | 3 / 1 |
| **slices.Contains** | 1.7µs / 267.7ns | 6.2x slower | 8.2KB / 0B | ∞x more | 1 / 0 |
