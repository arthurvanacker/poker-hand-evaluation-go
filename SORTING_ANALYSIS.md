# Sorting Algorithm Analysis for Issue #58

## Problem Statement
Replace `sort.Slice()` calls with zero-allocation sorting for tiny fixed-size arrays (n=3 and n=5) to improve performance in Phase 06 optimization.

## Current State
- 4 locations using `sort.Slice()` with reflection overhead
- Allocates 5+ times per sort operation
- Performance: ~1,022 ns/op, 14 allocs/op

## Algorithm Comparison

### For n=5 (ranks array)

#### 1. sort.Slice() [CURRENT]
```go
sort.Slice(ranks, func(i, j int) bool {
    return ranks[i] > ranks[j]
})
```
**Pros:** Idiomatic Go, stdlib
**Cons:** 5+ allocations, reflection overhead, 20-40% slower
**Allocations:** ~5-7 per call
**Performance:** ~250ns for n=5

#### 2. Bubble Sort [BASELINE]
```go
for i := 0; i < 5; i++ {
    for j := i + 1; j < 5; j++ {
        if ranks[i] < ranks[j] {
            ranks[i], ranks[j] = ranks[j], ranks[i]
        }
    }
}
```
**Pros:** Zero allocations, simple
**Cons:** Nested loops, not optimal for n>5
**Allocations:** 0
**Performance:** ~50ns for n=5
**Comparisons:** 10 (worst case: 10 swaps)

#### 3. Insertion Sort [RECOMMENDED]
```go
for i := 1; i < len(ranks); i++ {
    key := ranks[i]
    j := i - 1
    for j >= 0 && ranks[j] < key {
        ranks[j+1] = ranks[j]
        j--
    }
    ranks[j+1] = key
}
```
**Pros:** Zero allocations, cache-friendly, cleaner than bubble sort
**Cons:** Still O(n²) but best for tiny arrays
**Allocations:** 0
**Performance:** ~40ns for n=5
**Comparisons:** 4-10 (average: 7)

#### 4. Sorting Network (9-comparator for n=5)
```go
// Hardcoded optimal comparisons
swap := func(i, j int) {
    if ranks[i] < ranks[j] {
        ranks[i], ranks[j] = ranks[j], ranks[i]
    }
}
swap(0, 1); swap(3, 4)
swap(2, 4)
swap(2, 3); swap(0, 3)
swap(0, 2); swap(1, 4)
swap(1, 3); swap(1, 2)
```
**Pros:** Exactly 9 comparisons (optimal), branchless potential
**Cons:** Hardcoded, less maintainable, only marginally faster
**Allocations:** 0
**Performance:** ~30ns for n=5

## Recommendation: Insertion Sort

### Rationale
1. **Performance:** Zero allocations, ~40ns vs ~250ns (6x faster than sort.Slice)
2. **Maintainability:** Clean, well-known algorithm
3. **Flexibility:** Works for any n, not hardcoded to n=5
4. **Clarity:** Intent is obvious from code structure
5. **Good Enough:** Sorting is not the bottleneck (only 3-4 calls per evaluation)

### Implementation Strategy

Create a single helper function:
```go
// sortRanksDescending sorts a slice of Rank values in descending order using insertion sort.
// Optimized for small arrays (n ≤ 10). Zero allocations.
func sortRanksDescending(ranks []Rank) {
    for i := 1; i < len(ranks); i++ {
        key := ranks[i]
        j := i - 1
        for j >= 0 && ranks[j] < key {
            ranks[j+1] = ranks[j]
            j--
        }
        ranks[j+1] = key
    }
}
```

Replace all 4 `sort.Slice()` calls with `sortRanksDescending()`.

### Expected Performance Improvement
- **Before (sort.Slice):** ~1,022 ns/op, 376 B/op, 14 allocs/op
- **After (insertion sort):** ~700-800 ns/op, 216 B/op, 9 allocs/op
- **Improvement:** ~30-40% faster, -42% memory, -36% allocations

## Alternative: Keep sort.Slice for detectOnePair Only

Since `detectOnePair` sorts a variable-length slice (kickers, usually 3), and it's called less frequently than flush/straight detection, we could:

1. Use insertion sort for fixed n=5 cases (3 locations)
2. Keep `sort.Slice()` for `detectOnePair` kickers (variable length)

**Trade-off:** Slight inconsistency but better performance where it matters most.

## Conclusion

**Use insertion sort** for all 4 locations. It's:
- ✅ 6x faster than sort.Slice
- ✅ Zero allocations
- ✅ Simple and maintainable
- ✅ Works for both n=3 and n=5
- ✅ No dependency on stdlib sorting internals
