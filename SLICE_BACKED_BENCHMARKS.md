# Slice-backed collection A/B benchmark

This report compares pointer-backed `Collection` commit `96f31245b5c8e941b34b20fe16c5f102eddf85ae` with slice-backed `Slice` commit `09696b8823860c23ab327bde20f7b89dec741214`. In the typed-helper headline matrix, the slice commit has no statistically significant runtime regression above the practical +5% threshold and its time geomean is 14.54% lower. A historical typed global-sink matrix measured a large copy `Reduce` slowdown, but two focused reruns disagree and the latest independent 12-pair rerun is statistically neutral. The historical result and its caveat are preserved below rather than generalized as a stable regression.

## Methodology

The headline harness covers exactly the 33 collection-side labels from `docs/bench/main.go` in both generator modes: 33 borrowed/view cases and 33 copied/isolation cases. Inputs and constants are identical (`n=1000`, duplicate modulus 128, pipeline take 40, chunk 20, skip/take 40, group modulus 10). Borrowed mutating cases copy the canonical input into reusable storage inside every timed iteration; copied cases clone inside every timed iteration. Results use typed locals retained with `runtime.KeepAlive`. Each operation is behind a statically typed `//go:noinline` helper, while compiler diagnostics confirm collection calls inside the helper retain normal inlining. This avoids the generator's dead-code-elimination and indirect-constructor biases, and avoids the monolithic direct-`b.Loop` code-generation artifact described below.

Measured runs used Go 1.27.0 linux/arm64, `GOMAXPROCS=1`, `-cpu=1`, `-benchmem`, six fresh-process samples per commit, and 500 ms per case. Commit order alternated by pair. One 100 ms warm-up per commit was discarded. The separate cold tables are first-process 100 ms samples and are descriptive only (`n=1`). Build and module caches were under `/tmp`. Benchstat uses 95% confidence intervals and alpha 0.05.

Borrow/view and copy/isolation are intentionally separate because they have different aliasing semantics. They must not be combined as though they were the same workload. `GroupBy` maps base `GroupBySlice` to slice `Slice.GroupBy`; values are equivalent but the public result type changes from `map[K][]T` to `map[K]Slice[T]`. Numeric rows map base numeric receiver methods to slice free functions. `Filter` captures the returned value in both harnesses; this is required by the slice value receiver even though the base pointer receiver also updates its header. These API-shape cases cannot be source-identical, but they perform corresponding result-equivalent work.

No benchmark was excluded as result-semantically incomparable. The GroupBy, numeric, and Filter mappings above are the cases that cannot be made source/API-identical; their benchmark outputs and timed work are equivalent under the stated capture and construction rules.

## Correctness and completeness

Both commits pass `go test ./...`; both benchmark modules compile; and the headline preflight emits 66 rows and 33 unique operation labels per commit. No repository branch or original checkout was edited: benchmark-only files were added to detached worktrees under `/tmp/collection-ab-96f3124-vs-09696b8/{base,slice}`. The complete label inventory is: `Pipeline F→M→T→R`, `All`, `Any`, `None`, `First`, `Last`, `IndexWhere`, `Each`, `Map`, `Reduce (sum)`, `Filter`, `Chunk`, `Take`, `Contains`, `FirstWhere`, `GroupBy†`, `CountBy`, `CountByValue`, `Skip`, `SkipLast`, `Reverse`, `Shuffle`, `Zip`, `ZipWith`, `Unique`, `UniqueBy`, `Union`, `Intersect`, `Difference`, `ToMap`, `Sum`, `Min`, `Max`.

## Headline findings

- No statistically significant time regression exceeds +5% in either mode. Borrow `Each` is the largest significant slowdown at +2.79%, below the practical threshold.
- Borrow `Take`, `Skip`, and `SkipLast` improve by about 88%; `ZipWith` improves 23.71%, `Zip` 18.00%, and `Map` 5.53%.
- Copy `Filter` improves 20.03%. Several other copy medians improve materially, but their six-sample intervals are wide and benchstat does not call them significant; they are candidates for additional samples.
- Borrow Pipeline is statistically neutral: 537.1 ns/op versus 542.1 ns/op (`p=0.669`). Copy Pipeline is also neutral: 1038.5 versus 995.4 ns/op (`p=0.981`). Slice removes one 24-byte wrapper allocation in both modes.
- Intersect is time-neutral and improves by one allocation plus 24 bytes in both modes: 17 to 16 allocations borrowed and 19 to 18 copied.

A `REGRESSION` marker requires a slice median above base by more than 5%; statistical significance is shown independently in the `p` column. No headline row has a significant `REGRESSION` marker. Negative deltas favor slice.

### Headline: borrowed/view

| Operation | Base time ±CI | Slice time ±CI | Δ time | p | Base/Slice B | Δ B | Base/Slice allocs | Δ allocs | Flag |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---|
| Pipeline F→M→T→R | 537.1 ns ±7% | 542.1 ns ±5% | +0.93% | p=0.669 n=6 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| All | 264.7 ns ±0% | 266.4 ns ±0% | +0.62% | p=0.004 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Any | 266.8 ns ±0% | 266.0 ns ±0% | -0.30% | p=0.307 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| None | 265.4 ns ±0% | 266.2 ns ±1% | +0.30% | p=0.229 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| First | 1.798 ns ±1% | 1.797 ns ±1% | -0.06% | p=0.589 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Last | 1.796 ns ±0% | 1.797 ns ±3% | +0.03% | p=0.974 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| IndexWhere | 266.2 ns ±1% | 266.5 ns ±1% | +0.11% | p=0.372 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Each | 265.2 ns ±2% | 272.6 ns ±1% | +2.79% | p=0.002 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Map | 373.4 ns ±2% | 352.8 ns ±0% | -5.53% | p=0.002 n=6 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Reduce (sum) | 265.2 ns ±2% | 264.4 ns ±0% | -0.26% | p=0.223 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Filter | 584.1 ns ±8% | 570.0 ns ±6% | -2.42% | p=0.394 n=6 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Chunk | 210.8 ns ±13% | 192.4 ns ±10% | -8.73% | p=0.240 n=6 | 1.25 KiB / 1.25 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Take | 15.8 ns ±33% | 1.859 ns ±3% | -88.26% | p=0.002 n=6 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Contains | 266.1 ns ±0% | 266.6 ns ±1% | +0.17% | p=0.288 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| FirstWhere | 267.7 ns ±1% | 268.9 ns ±1% | +0.43% | p=0.900 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| GroupBy† | 9.072 µs ±10% | 8.665 µs ±7% | -4.48% | p=0.699 n=6 | 20.85 KiB / 20.85 KiB | +0.00% | 85 / 85 | +0.00% | — |
| CountBy | 7.962 µs ±2% | 8.072 µs ±2% | +1.38% | p=0.132 n=6 | 9.32 KiB / 9.32 KiB | +0.00% | 13 / 13 | +0.00% | — |
| CountByValue | 7.949 µs ±2% | 8.098 µs ±2% | +1.87% | p=0.143 n=6 | 9.32 KiB / 9.32 KiB | +0.00% | 13 / 13 | +0.00% | — |
| Skip | 15.0 ns ±34% | 1.805 ns ±1% | -87.94% | p=0.002 n=6 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| SkipLast | 15.1 ns ±33% | 1.794 ns ±1% | -88.12% | p=0.002 n=6 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Reverse | 224.1 ns ±2% | 213.2 ns ±1% | -4.86% | p=0.002 n=6 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Shuffle | 3.796 µs ±2% | 3.678 µs ±1% | -3.13% | p=0.002 n=6 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Zip | 1.129 µs ±13% | 925.8 ns ±10% | -18.00% | p=0.026 n=6 | 16.02 KiB / 16.00 KiB | -0.15% | 2 / 1 | -50.00% | — |
| ZipWith | 911.3 ns ±15% | 695.2 ns ±9% | -23.71% | p=0.002 n=6 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Unique | 6.207 µs ±5% | 6.011 µs ±5% | -3.16% | p=0.387 n=6 | 44.10 KiB / 44.08 KiB | -0.05% | 7 / 6 | -14.29% | — |
| UniqueBy | 6.400 µs ±4% | 6.216 µs ±6% | -2.87% | p=0.240 n=6 | 44.10 KiB / 44.08 KiB | -0.05% | 7 / 6 | -14.29% | — |
| Union | 16.546 µs ±3% | 16.390 µs ±5% | -0.94% | p=0.699 n=6 | 88.18 KiB / 88.16 KiB | -0.03% | 11 / 10 | -9.09% | — |
| Intersect | 12.009 µs ±2% | 11.982 µs ±3% | -0.23% | p=1.000 n=6 | 11.09 KiB / 11.07 KiB | -0.21% | 17 / 16 | -5.88% | — |
| Difference | 19.489 µs ±2% | 19.419 µs ±4% | -0.36% | p=0.818 n=6 | 80.18 KiB / 80.16 KiB | -0.03% | 12 / 11 | -8.33% | — |
| ToMap | 6.803 µs ±3% | 6.820 µs ±3% | +0.25% | p=0.818 n=6 | 36.12 KiB / 36.12 KiB | +0.00% | 6 / 6 | +0.00% | — |
| Sum | 264.5 ns ±1% | 264.7 ns ±1% | +0.06% | p=0.699 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Min | 511.4 ns ±1% | 511.1 ns ±2% | -0.06% | p=0.563 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Max | 512.8 ns ±1% | 510.5 ns ±2% | -0.45% | p=0.132 n=6 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |

### Headline: copied/isolation

| Operation | Base time ±CI | Slice time ±CI | Δ time | p | Base/Slice B | Δ B | Base/Slice allocs | Δ allocs | Flag |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---|
| Pipeline F→M→T→R | 1.038 µs ±13% | 995.4 ns ±8% | -4.16% | p=0.981 n=6 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| All | 824.8 ns ±15% | 761.2 ns ±9% | -7.72% | p=0.394 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Any | 813.4 ns ±14% | 758.1 ns ±10% | -6.79% | p=0.331 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| None | 816.0 ns ±15% | 762.0 ns ±9% | -6.61% | p=0.818 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| First | 600.2 ns ±23% | 544.2 ns ±14% | -9.31% | p=0.485 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Last | 597.1 ns ±22% | 541.7 ns ±14% | -9.29% | p=0.851 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| IndexWhere | 817.1 ns ±15% | 760.2 ns ±9% | -6.96% | p=0.937 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Each | 813.9 ns ±14% | 795.8 ns ±12% | -2.23% | p=0.589 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Map | 919.5 ns ±16% | 740.1 ns ±47% | -19.50% | p=0.240 n=6 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Reduce (sum) | 807.1 ns ±14% | 786.4 ns ±28% | -2.56% | p=0.589 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Filter | 1.184 µs ±12% | 946.9 ns ±8% | -20.03% | p=0.002 n=6 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Chunk | 810.2 ns ±20% | 675.2 ns ±30% | -16.66% | p=0.180 n=6 | 9.25 KiB / 9.25 KiB | +0.00% | 2 / 2 | +0.00% | — |
| Take | 625.5 ns ±24% | 477.4 ns ±16% | -23.66% | p=0.065 n=6 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Contains | 813.2 ns ±15% | 758.9 ns ±9% | -6.68% | p=0.937 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| FirstWhere | 814.4 ns ±15% | 761.3 ns ±10% | -6.52% | p=1.000 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| GroupBy† | 9.691 µs ±11% | 9.352 µs ±6% | -3.49% | p=0.699 n=6 | 28.85 KiB / 28.85 KiB | +0.00% | 86 / 86 | +0.00% | — |
| CountBy | 8.672 µs ±3% | 8.524 µs ±2% | -1.71% | p=0.240 n=6 | 17.32 KiB / 17.32 KiB | +0.00% | 14 / 14 | +0.00% | — |
| CountByValue | 8.637 µs ±3% | 8.519 µs ±1% | -1.37% | p=0.485 n=6 | 17.32 KiB / 17.32 KiB | +0.00% | 14 / 14 | +0.00% | — |
| Skip | 643.0 ns ±26% | 475.3 ns ±15% | -26.08% | p=0.065 n=6 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| SkipLast | 634.1 ns ±23% | 465.4 ns ±15% | -26.60% | p=0.065 n=6 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Reverse | 771.9 ns ±19% | 615.1 ns ±11% | -20.30% | p=0.065 n=6 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Shuffle | 4.277 µs ±3% | 4.197 µs ±2% | -1.88% | p=0.180 n=6 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Zip | 2.393 µs ±19% | 2.010 µs ±11% | -16.01% | p=0.065 n=6 | 32.02 KiB / 32.00 KiB | -0.07% | 4 / 3 | -25.00% | — |
| ZipWith | 2.161 µs ±20% | 1.768 µs ±11% | -18.16% | p=0.065 n=6 | 24.02 KiB / 24.00 KiB | -0.10% | 4 / 3 | -25.00% | — |
| Unique | 6.855 µs ±7% | 6.683 µs ±5% | -2.51% | p=0.818 n=6 | 52.10 KiB / 52.08 KiB | -0.04% | 8 / 7 | -12.50% | — |
| UniqueBy | 7.134 µs ±6% | 6.933 µs ±4% | -2.82% | p=0.394 n=6 | 52.10 KiB / 52.08 KiB | -0.04% | 8 / 7 | -12.50% | — |
| Union | 17.869 µs ±4% | 17.500 µs ±3% | -2.07% | p=0.589 n=6 | 104.18 KiB / 104.16 KiB | -0.02% | 13 / 12 | -7.69% | — |
| Intersect | 13.371 µs ±3% | 13.184 µs ±2% | -1.40% | p=0.818 n=6 | 27.09 KiB / 27.07 KiB | -0.09% | 19 / 18 | -5.26% | — |
| Difference | 20.996 µs ±5% | 20.430 µs ±4% | -2.69% | p=0.485 n=6 | 96.18 KiB / 96.16 KiB | -0.02% | 14 / 13 | -7.14% | — |
| ToMap | 7.510 µs ±3% | 7.460 µs ±3% | -0.67% | p=0.818 n=6 | 44.12 KiB / 44.12 KiB | +0.00% | 7 / 7 | +0.00% | — |
| Sum | 825.4 ns ±14% | 765.2 ns ±9% | -7.28% | p=0.180 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Min | 1.065 µs ±11% | 1.011 µs ±7% | -5.05% | p=1.000 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Max | 1.069 µs ±11% | 1.008 µs ±7% | -5.68% | p=0.331 n=6 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |

## Pipeline and Intersect investigation

The direct-`b.Loop` diagnostic initially showed borrow Pipeline at 556.2 ns/op base versus 753.6 ns/op slice (+35.48%). That result does not reproduce once each operation is isolated behind a typed helper: 537.1 versus 542.1 ns/op, statistically neutral. The monolithic diagnostic dispatcher compiled into 8,304 bytes with a 1,376-byte frame for base and 7,680 bytes with a 1,216-byte frame for slice, while retaining every possible output type across a 33-way switch. Isolating Pipeline produces compact 384-byte base and 352-byte slice helpers; compiler `-m=2` confirms `New`, `Filter`, `Map`, callbacks, and `Reduce` inline inside them and the base initial wrapper does not escape. The earlier slowdown is therefore benchmark-context/code-generation sensitivity, not evidence of a slice implementation regression.

Sequential CPU profiles of the isolated borrowed Pipeline agree: Filter and Map account for roughly 79% of sampled CPU on each commit, input reset (`memmove`) for 12–15%, and tail clearing for 2–4%. Base additionally spends about 3.4% cumulative time in `Take`/its 24-byte wrapper allocation. Both versions clear Filter's rejected tail, so there is no hidden semantic-work mismatch. Focused cumulative stage data is in `stages-benchstat.txt`; because benchmark placement itself changes the base code-generation result, it is diagnostic rather than headline evidence.

The monolithic direct-`b.Loop` and typed global-sink dispatchers reported Intersect at 17 versus 18 allocations borrowed and 19 versus 20 copied, with about 32 extra bytes for slice. A separate ten-sample global-sink benchmark containing only borrowed Intersect reverses that result: base is 11.92 µs, 11.09 KiB, and 17 allocations versus slice at 11.96 µs, 11.07 KiB, and 16 allocations; time is neutral (`p=0.383`). Allocation profiles attribute base's isolated extra object to `collection.New` wrapping Intersect's returned slice. The secondary table's +1 allocation is therefore specific to the 33-way dispatcher retaining many possible result types, whereas the headline and isolated operation both show the intended wrapper removal.

## Secondary: typed global-sink context

These ten-sample, 500 ms rows assign every result to a typed global sink through a monolithic dispatcher. Aggregate results can therefore be retained/escape; scalar results such as `Reduce` do not escape, but still use the same global-write and dispatcher context. This is a real measured context, not a model of every local consumer, and its results must be interpreted independently from the headline helpers.

The historical table below measured copied `Reduce` at 831.8 ns/op on base versus 1.078 µs/op on slice: +29.61% (`p<0.001`, n=10), with identical 8.00 KiB and one allocation. The first isolated global-sink run reduced that difference to +9.23% (773.8 versus 845.1 ns/op, `p=0.009`, n=10), and a separate 5 s profile run was +5.6%. A fresh independent rerun using 12 alternating 500 ms fresh-process pairs does not reproduce either slowdown: base is 817.8 ns/op versus slice at 828.4 ns/op, +1.30% and statistically neutral (`p=0.355`, n=12), again with identical 8.00 KiB and one allocation. The headline local-helper row is likewise neutral at -2.56% (`p=0.589`, n=6).

The implementation-shaping experiment is inconsistent for the same reason. One run put an inlined internal-style helper 2.11% ahead of the current slice method (`p=0.033`, n=12); the fresh rerun puts the current method at 847.2 ns/op and the helper at 771.2 ns/op, but the wider difference is not significant (`p=0.078`, n=12). The base and slice inlined clone/reduce loops are instruction-identical, while the copied benchmark samples are bimodal and CPU profiles are GC/assist-heavy. Function placement changes without changing the loop instructions. The earlier +29.61% and +9.23% measurements are therefore retained as measured context/layout-sensitive outliers, but they do not establish a reproducible slice regression or justify a production code change. The table's `REGRESSION` marker mechanically reflects its historical medians, not the focused rerun conclusion.

### Typed global sink: borrowed/view

| Operation | Base time ±CI | Slice time ±CI | Δ time | p | Base/Slice B | Δ B | Base/Slice allocs | Δ allocs | Flag |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---|
| Pipeline F→M→T→R | 776.8 ns ±6% | 752.1 ns ±0% | -3.19% | p<0.001 n=10 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| All | 264.5 ns ±6% | 263.9 ns ±0% | -0.25% | p=0.285 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Any | 265.9 ns ±7% | 264.8 ns ±1% | -0.41% | p=0.036 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| None | 265.4 ns ±6% | 263.9 ns ±0% | -0.55% | p=0.007 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| First | 0.512 ns ±6% | 0.512 ns ±0% | +0.06% | p=0.985 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Last | 0.511 ns ±6% | 0.511 ns ±0% | -0.06% | p=0.517 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| IndexWhere | 265.5 ns ±6% | 264.6 ns ±1% | -0.30% | p=0.271 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Each | 264.4 ns ±6% | 263.9 ns ±0% | -0.17% | p=0.071 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Map | 373.7 ns ±6% | 347.2 ns ±0% | -7.09% | p<0.001 n=10 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Reduce (sum) | 263.8 ns ±6% | 263.5 ns ±0% | -0.09% | p=0.115 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Filter | 604.6 ns ±5% | 561.9 ns ±8% | -7.08% | p=0.007 n=10 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Chunk | 218.1 ns ±14% | 217.0 ns ±12% | -0.50% | p=0.481 n=10 | 1.25 KiB / 1.25 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Take | 15.7 ns ±25% | 1.401 ns ±5% | -91.07% | p<0.001 n=10 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Contains | 265.8 ns ±6% | 265.1 ns ±1% | -0.24% | p=0.145 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| FirstWhere | 265.4 ns ±6% | 265.1 ns ±2% | -0.11% | p=0.225 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| GroupBy† | 9.159 µs ±8% | 9.185 µs ±9% | +0.29% | p=0.912 n=10 | 20.85 KiB / 20.85 KiB | +0.00% | 85 / 85 | +0.00% | — |
| CountBy | 8.055 µs ±6% | 7.945 µs ±2% | -1.36% | p=0.190 n=10 | 9.32 KiB / 9.32 KiB | +0.00% | 13 / 13 | +0.00% | — |
| CountByValue | 8.065 µs ±7% | 8.042 µs ±2% | -0.29% | p=0.481 n=10 | 9.32 KiB / 9.32 KiB | +0.00% | 13 / 13 | +0.00% | — |
| Skip | 15.0 ns ±26% | 0.579 ns ±1% | -96.15% | p<0.001 n=10 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| SkipLast | 14.9 ns ±26% | 0.513 ns ±1% | -96.56% | p<0.001 n=10 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Reverse | 223.8 ns ±6% | 218.7 ns ±1% | -2.28% | p<0.001 n=10 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Shuffle | 3.789 µs ±5% | 3.676 µs ±3% | -2.98% | p=0.003 n=10 | 24 B / 0 B | -100.00% | 1 / 0 | -100.00% | — |
| Zip | 1.200 µs ±10% | 1.018 µs ±14% | -15.17% | p=0.001 n=10 | 16.02 KiB / 16.00 KiB | -0.15% | 2 / 1 | -50.00% | — |
| ZipWith | 945.4 ns ±11% | 758.0 ns ±13% | -19.83% | p<0.001 n=10 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Unique | 6.277 µs ±8% | 6.233 µs ±5% | -0.70% | p=0.101 n=10 | 44.10 KiB / 44.08 KiB | -0.05% | 7 / 6 | -14.29% | — |
| UniqueBy | 6.516 µs ±7% | 6.421 µs ±4% | -1.47% | p=0.063 n=10 | 44.10 KiB / 44.08 KiB | -0.05% | 7 / 6 | -14.29% | — |
| Union | 16.696 µs ±8% | 16.613 µs ±3% | -0.49% | p=0.325 n=10 | 88.18 KiB / 88.16 KiB | -0.03% | 11 / 10 | -9.09% | — |
| Intersect | 12.119 µs ±7% | 12.009 µs ±2% | -0.91% | p=0.105 n=10 | 11.09 KiB / 11.12 KiB | +0.28% | 17 / 18 | +5.88% | REGRESSION: allocs |
| Difference | 19.642 µs ±7% | 19.722 µs ±8% | +0.41% | p=0.670 n=10 | 80.18 KiB / 80.16 KiB | -0.03% | 12 / 11 | -8.33% | — |
| ToMap | 6.968 µs ±8% | 6.954 µs ±3% | -0.20% | p=0.393 n=10 | 36.12 KiB / 36.12 KiB | +0.00% | 6 / 6 | +0.00% | — |
| Sum | 264.0 ns ±7% | 263.9 ns ±0% | -0.06% | p=0.540 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Min | 510.8 ns ±6% | 510.7 ns ±0% | -0.03% | p=0.566 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |
| Max | 511.6 ns ±6% | 512.2 ns ±0% | +0.12% | p=0.782 n=10 | 0 B / 0 B | +0.00% | 0 / 0 | +0.00% | — |

### Typed global sink: copied/isolation

| Operation | Base time ±CI | Slice time ±CI | Δ time | p | Base/Slice B | Δ B | Base/Slice allocs | Δ allocs | Flag |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---|
| Pipeline F→M→T→R | 1.079 µs ±13% | 1.033 µs ±10% | -4.22% | p=0.043 n=10 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| All | 857.5 ns ±14% | 849.7 ns ±13% | -0.91% | p=0.089 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Any | 851.8 ns ±13% | 840.6 ns ±13% | -1.31% | p=0.123 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| None | 850.3 ns ±14% | 837.7 ns ±13% | -1.48% | p=0.063 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| First | 626.9 ns ±17% | 618.1 ns ±20% | -1.40% | p=0.448 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Last | 627.3 ns ±19% | 615.7 ns ±18% | -1.85% | p=0.105 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| IndexWhere | 835.1 ns ±14% | 827.9 ns ±13% | -0.87% | p=0.481 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Each | 851.5 ns ±14% | 859.3 ns ±13% | +0.91% | p=0.912 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Map | 958.1 ns ±13% | 829.6 ns ±15% | -13.41% | p=0.001 n=10 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Reduce (sum) | 831.8 ns ±14% | 1.078 µs ±10% | +29.61% | p<0.001 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | REGRESSION: time |
| Filter | 1.214 µs ±12% | 1.064 µs ±10% | -12.36% | p<0.001 n=10 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Chunk | 867.8 ns ±18% | 793.1 ns ±21% | -8.61% | p=0.005 n=10 | 9.25 KiB / 9.25 KiB | +0.00% | 2 / 2 | +0.00% | — |
| Take | 656.0 ns ±18% | 556.4 ns ±21% | -15.18% | p=0.005 n=10 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Contains | 844.8 ns ±14% | 834.5 ns ±13% | -1.22% | p=0.190 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| FirstWhere | 844.4 ns ±13% | 834.2 ns ±13% | -1.21% | p=0.165 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| GroupBy† | 9.904 µs ±10% | 9.806 µs ±9% | -0.99% | p=0.280 n=10 | 28.85 KiB / 28.85 KiB | +0.00% | 86 / 86 | +0.00% | — |
| CountBy | 8.757 µs ±7% | 8.726 µs ±3% | -0.35% | p=0.631 n=10 | 17.32 KiB / 17.32 KiB | +0.00% | 14 / 14 | +0.00% | — |
| CountByValue | 8.757 µs ±6% | 8.651 µs ±2% | -1.20% | p=0.165 n=10 | 17.32 KiB / 17.32 KiB | +0.00% | 14 / 14 | +0.00% | — |
| Skip | 651.4 ns ±18% | 552.7 ns ±22% | -15.16% | p=0.005 n=10 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| SkipLast | 645.6 ns ±22% | 544.7 ns ±21% | -15.62% | p=0.005 n=10 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Reverse | 810.9 ns ±18% | 696.7 ns ±18% | -14.08% | p=0.005 n=10 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Shuffle | 4.333 µs ±6% | 4.270 µs ±3% | -1.47% | p=0.075 n=10 | 8.02 KiB / 8.00 KiB | -0.29% | 2 / 1 | -50.00% | — |
| Zip | 2.543 µs ±16% | 2.321 µs ±15% | -8.75% | p=0.005 n=10 | 32.02 KiB / 32.00 KiB | -0.07% | 4 / 3 | -25.00% | — |
| ZipWith | 2.186 µs ±16% | 1.985 µs ±17% | -9.17% | p=0.005 n=10 | 24.02 KiB / 24.00 KiB | -0.10% | 4 / 3 | -25.00% | — |
| Unique | 7.093 µs ±7% | 6.960 µs ±5% | -1.88% | p=0.225 n=10 | 52.10 KiB / 52.08 KiB | -0.04% | 8 / 7 | -12.50% | — |
| UniqueBy | 7.246 µs ±6% | 7.236 µs ±6% | -0.14% | p=0.579 n=10 | 52.10 KiB / 52.08 KiB | -0.04% | 8 / 7 | -12.50% | — |
| Union | 18.020 µs ±8% | 18.014 µs ±4% | -0.04% | p=0.796 n=10 | 104.18 KiB / 104.16 KiB | -0.02% | 13 / 12 | -7.69% | — |
| Intersect | 13.466 µs ±8% | 13.408 µs ±2% | -0.43% | p=0.393 n=10 | 27.09 KiB / 27.12 KiB | +0.12% | 19 / 20 | +5.26% | REGRESSION: allocs |
| Difference | 20.987 µs ±8% | 21.126 µs ±3% | +0.66% | p=0.838 n=10 | 96.18 KiB / 96.16 KiB | -0.02% | 14 / 13 | -7.14% | — |
| ToMap | 7.770 µs ±6% | 7.724 µs ±3% | -0.59% | p=0.529 n=10 | 44.12 KiB / 44.12 KiB | +0.00% | 7 / 7 | +0.00% | — |
| Sum | 877.9 ns ±14% | 846.8 ns ±14% | -3.54% | p=0.029 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Min | 1.116 µs ±11% | 1.091 µs ±10% | -2.28% | p=0.271 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |
| Max | 1.119 µs ±10% | 1.095 µs ±10% | -2.19% | p=0.045 n=10 | 8.00 KiB / 8.00 KiB | +0.00% | 1 / 1 | +0.00% | — |

## Diagnostic matrices and artifacts

The direct-local `b.Loop` matrix has ten 500 ms samples and is preserved in `local-benchstat.txt`; the literal generator/DCE replication has six 200 ms samples in `replicated-benchstat.txt`. Neither is used for headline conclusions. Cold helper results are in `helper-cold-benchstat.txt`. The independent Reduce rerun is in `reduce-rerun-{base,slice}.txt` and `reduce-rerun-benchstat.txt`; its helper-shaping companion is in `reduce-helper-rerun-{original,shaped}.txt` and `reduce-helper-rerun-benchstat.txt`. Raw benchmark text, CSV comparisons, compiler escape output, disassembly, CPU/allocation profiles, harness sources, environment data, and correctness logs are colocated with this report.

Recommended next experiments: randomize or pad binary function placement and collect hardware counters for copy `Reduce`; collect ten or more helper samples for copy-mode operations with wide confidence intervals; add a repository benchmark helper pattern that preserves production inlining; and add an allocation regression test for isolated Intersect and view-returning operations. Avoid publishing results from the current generator until it uses live outputs and static constructors.
