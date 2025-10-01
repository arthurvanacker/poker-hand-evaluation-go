# Poker Hand Evaluation Algorithm - Complete Beginner's Guide

**Author's Note**: This document is written for developers who are new to both poker and algorithms. I'll explain everything from scratch, assuming zero prior knowledge. By the end, you'll understand not just *what* the code does, but *why* it works this way.

---

## Table of Contents

1. [What Problem Are We Solving?](#what-problem-are-we-solving)
2. [Poker Basics (The Minimum You Need to Know)](#poker-basics)
3. [High-Level Algorithm Overview](#high-level-algorithm-overview)
4. [Data Structures Explained](#data-structures-explained)
5. [The Algorithm Step-by-Step](#the-algorithm-step-by-step)
6. [Helper Functions Deep-Dive](#helper-functions-deep-dive)
7. [Detection Functions Explained](#detection-functions-explained)
8. [Combination Generation Algorithm](#combination-generation-algorithm)
9. [Performance Analysis](#performance-analysis)
10. [Common Pitfalls and Edge Cases](#common-pitfalls-and-edge-cases)
11. [Why This Design?](#why-this-design)

---

## What Problem Are We Solving?

### The Real-World Problem

In Texas Hold'em poker, you have **7 cards total**:
- 2 "hole cards" (private, only you see them)
- 5 "community cards" (shared with all players)

But your poker hand is made from **only 5 cards**. So the question is:

> **Out of my 7 cards, which 5-card combination makes the strongest poker hand?**

### Example Scenario

```
Your hole cards:  Ah Kh
Community cards:  Qh Jh Th 9d 2c

All 7 cards:      Ah Kh Qh Jh Th 9d 2c
```

**Question**: What's your best hand?

**Answer**: If you pick `Ah Kh Qh Jh Th`, you have a **Royal Flush** (the best possible hand).

But there are **21 different ways** to pick 5 cards from 7 cards (this is a math concept called "combinations"). Our algorithm needs to:
1. Check all 21 combinations
2. Evaluate each one to see what poker hand it makes
3. Return the strongest one

---

## Poker Basics

### The 10 Hand Categories (Strongest to Weakest)

You don't need to memorize poker rules, but here's a quick reference:

| Rank | Name            | Example           | What It Means                                    |
|------|-----------------|-------------------|--------------------------------------------------|
| 10   | Royal Flush     | `Ah Kh Qh Jh Th` | 10-Jack-Queen-King-Ace, all same suit            |
| 9    | Straight Flush  | `9h 8h 7h 6h 5h` | 5 cards in sequence, all same suit               |
| 8    | Four of a Kind  | `Ah Ad Ac As Kh` | 4 cards with same rank (like four Aces)          |
| 7    | Full House      | `Ah Ad Ac Kh Kd` | 3 of one rank + 2 of another (three Aces + two Kings) |
| 6    | Flush           | `Ah Kh Qh Jh 9h` | 5 cards, all same suit (but not in sequence)     |
| 5    | Straight        | `9h 8d 7c 6s 5h` | 5 cards in sequence (but mixed suits)            |
| 4    | Three of a Kind | `Ah Ad Ac Kh Qd` | 3 cards with same rank                           |
| 3    | Two Pair        | `Ah Ad Kh Kd Qc` | Two different pairs                              |
| 2    | One Pair        | `Ah Ad Kh Qd Jc` | Two cards with same rank                         |
| 1    | High Card       | `Ah Kd Qh Jc 9s` | No matching cards (weakest hand)                 |

### What Are "Tiebreakers"?

If two players both have "One Pair", who wins? We need **tiebreakers**:

```
Player 1: Ah Ad Kh Qd Jc  (Pair of Aces, with K-Q-J kickers)
Player 2: Ah Ad Kh Qd 9c  (Pair of Aces, with K-Q-9 kickers)
```

Both have a pair of Aces, both have King kicker, both have Queen kicker. But Player 1 has a Jack (11) and Player 2 has a 9. **Player 1 wins** because Jack > 9.

Our algorithm stores these tiebreakers in an array: `[14, 13, 12, 11]` for Player 1 vs `[14, 13, 12, 9]` for Player 2 (where 14=Ace, 13=King, 12=Queen, 11=Jack).

---

## High-Level Algorithm Overview

### The Big Picture (10,000-foot view)

```
INPUT:  7 cards (e.g., Ah Kh Qh Jh Th 9d 2c)
OUTPUT: Best 5-card hand (e.g., Royal Flush: Ah Kh Qh Jh Th)

ALGORITHM:
1. Generate all possible 5-card combinations from the 7 cards (21 total)
2. For each combination:
   a. Check if it's a Royal Flush
   b. If not, check if it's a Straight Flush
   c. If not, check if it's Four of a Kind
   d. ... (continue down the list)
   e. If nothing matches, it's High Card
3. Compare all 21 evaluated hands
4. Return the strongest one
```

### Why This Approach?

**Alternative**: We could use a "lookup table" with pre-computed rankings for all 2.6 million possible 5-card hands. This would be faster (O(1) lookup), but:
- More complex to implement
- Requires 10-20 MB of memory
- Harder to understand and debug

**Our approach**: Brute force evaluation
- Simpler to understand
- Only ~210 operations for 7 cards (very fast for modern computers)
- Easy to test and verify correctness
- **Prioritizes code clarity over raw performance**

---

## Data Structures Explained

### 1. Rank (What Number Is the Card?)

```go
type Rank int

const (
    Two   Rank = 2
    Three Rank = 3
    Four  Rank = 4
    Five  Rank = 5
    Six   Rank = 6
    Seven Rank = 7
    Eight Rank = 8
    Nine  Rank = 9
    Ten   Rank = 10
    Jack  Rank = 11
    Queen Rank = 12
    King  Rank = 13
    Ace   Rank = 14
)
```

**What is this?** A `Rank` represents the "number" on a card.

**Why use numbers instead of names?**
- We can compare them easily: `Ace (14) > King (13)` is just `14 > 13`
- Detecting straights becomes simple: check if ranks are consecutive numbers
- Sorting is built-in: Go can sort numbers automatically

**Why is Ace = 14?**
- In poker, Ace is the highest card (beats King)
- Exception: In a "wheel straight" (A-2-3-4-5), Ace acts as 1 (we handle this specially)

### 2. Suit (What Symbol Is the Card?)

```go
type Suit int

const (
    Hearts   Suit = iota // 0
    Diamonds             // 1
    Clubs                // 2
    Spades               // 3
)
```

**What is this?** The suit is the symbol: ♥ ♦ ♣ ♠

**Why use numbers?**
- Suits don't have a "rank" in poker (Hearts isn't "better" than Clubs)
- We just need to check if suits **match** (for flushes)
- Numbers work fine: `card1.Suit == card2.Suit` checks if they're the same

**What is `iota`?**
- It's a Go keyword that auto-numbers constants: 0, 1, 2, 3...
- `Hearts = 0`, `Diamonds = 1`, etc.
- We don't care about the actual numbers, just that each suit gets a unique ID

### 3. Card (A Single Playing Card)

```go
type Card struct {
    Rank Rank
    Suit Suit
}
```

**What is this?** A card combines a rank and a suit.

**Example**:
```go
card := Card{Rank: Ace, Suit: Hearts}
// This represents the Ace of Hearts (Ah)
```

**Why a struct?**
- We need both pieces of information together
- A struct groups related data
- Think of it like a "record" or "object" in other languages

### 4. HandCategory (Type of Poker Hand)

```go
type HandCategory int

const (
    HighCard      HandCategory = 1
    OnePair       HandCategory = 2
    TwoPair       HandCategory = 3
    ThreeOfAKind  HandCategory = 4
    Straight      HandCategory = 5
    Flush         HandCategory = 6
    FullHouse     HandCategory = 7
    FourOfAKind   HandCategory = 8
    StraightFlush HandCategory = 9
    RoyalFlush    HandCategory = 10
)
```

**What is this?** An enum of all 10 poker hand types.

**Why numbered 1-10?**
- Higher number = stronger hand
- We can compare: `RoyalFlush (10) > Flush (6)` is just `10 > 6`
- Makes the comparison logic trivial

### 5. Hand (An Evaluated 5-Card Hand)

```go
type Hand struct {
    Cards       []Card       // The 5 cards
    Category    HandCategory // What type of hand (Flush, Straight, etc.)
    Tiebreakers []Rank       // Ranks used to break ties
}
```

**What is this?** The result of evaluating 5 cards.

**Example**:
```go
hand := &Hand{
    Cards:       [5]Card{Ah, Ad, Kh, Qd, Jc},
    Category:    OnePair,
    Tiebreakers: []Rank{14, 13, 12, 11},  // Pair of Aces, K-Q-J kickers
}
```

**Why Tiebreakers?**
- When two hands have the same category, we need a way to compare them
- Tiebreakers are stored in **order of importance**
- Example for One Pair: `[pair rank, kicker1, kicker2, kicker3]`

---

## The Algorithm Step-by-Step

### Step 1: Finding the Best Hand (FindBestHand)

```go
func FindBestHand(cards []Card) *Hand {
    if len(cards) < 5 {
        return nil
    }

    // Optimization: if exactly 5 cards, evaluate directly
    if len(cards) == 5 {
        return EvaluateHand(cards)
    }

    // Generate all 5-card combinations
    combinations := Combinations(cards, 5)

    var bestHand *Hand

    // Evaluate each combination and keep the best
    for _, combo := range combinations {
        hand := EvaluateHand(combo)
        if bestHand == nil || CompareHands(hand, bestHand) > 0 {
            bestHand = hand
        }
    }

    return bestHand
}
```

**Line-by-line explanation**:

1. **Lines 2-4**: If we have fewer than 5 cards, we can't make a hand. Return `nil` (null/nothing).

2. **Lines 6-9**: **Optimization trick**: If we already have exactly 5 cards, there's only one possible hand. No need to generate combinations. Just evaluate those 5 cards directly.

3. **Line 12**: Generate all possible 5-card combinations. If we have 7 cards, this creates 21 different 5-card groups.

4. **Line 14**: Create a variable to track the best hand we've found so far. Initially, it's `nil` (empty).

5. **Lines 17-21**: Loop through each combination:
   - **Line 18**: Evaluate this 5-card combination (figure out what hand it makes)
   - **Line 19**: If this is our first hand (`bestHand == nil`) OR this hand beats our current best hand, update `bestHand`
   - `CompareHands(hand, bestHand) > 0` means "hand is stronger than bestHand"

6. **Line 24**: Return the best hand we found.

**Example walkthrough**:

```
Input: 7 cards [Ah Kh Qh Jh Th 9d 2c]

Combination 1: [Ah Kh Qh Jh Th] → Evaluates to Royal Flush
Combination 2: [Ah Kh Qh Jh 9d] → Evaluates to Flush (5 hearts? No, 4 hearts + 1 diamond = High Card Ace)
Combination 3: [Ah Kh Qh Jh 2c] → Evaluates to High Card Ace
... (18 more combinations)

Best hand: Combination 1 (Royal Flush)
```

### Step 2: Evaluating a 5-Card Hand (EvaluateHand)

```go
func EvaluateHand(cards []Card) *Hand {
    if len(cards) != 5 {
        return nil
    }

    // Check Royal Flush
    if detectRoyalFlush(cards) {
        return &Hand{
            Cards:       cards,
            Category:    RoyalFlush,
            Tiebreakers: []Rank{},
        }
    }

    // Check Straight Flush
    if found, highCard := detectStraightFlush(cards); found {
        return &Hand{
            Cards:       cards,
            Category:    StraightFlush,
            Tiebreakers: []Rank{highCard},
        }
    }

    // ... (continue for all 10 categories)

    // Default to High Card
    _, tiebreakers := detectHighCard(cards)
    return &Hand{
        Cards:       cards,
        Category:    HighCard,
        Tiebreakers: tiebreakers,
    }
}
```

**The strategy**: Check hand categories from **strongest to weakest**.

**Why this order?**
- A Royal Flush is also a Straight Flush, Flush, AND Straight
- By checking strongest first, we avoid misclassifying hands
- Once we find a match, we return immediately (no need to check weaker hands)

**Example**:
```
Cards: Ah Kh Qh Jh Th

Check Royal Flush? YES → Return immediately
(Never reaches Straight Flush, Flush, or Straight checks)
```

---

## Helper Functions Deep-Dive

### isFlush: Are All 5 Cards the Same Suit?

```go
func isFlush(cards []Card) bool {
    if len(cards) != 5 {
        return false
    }

    // Use the first card's suit as reference
    firstSuit := cards[0].Suit

    // Check if all remaining cards match the first suit
    for i := 1; i < len(cards); i++ {
        if cards[i].Suit != firstSuit {
            return false
        }
    }

    return true
}
```

**What does this do?** Checks if all 5 cards have the same suit (all Hearts, all Spades, etc.).

**Algorithm**:
1. Take the first card's suit as the "reference" suit
2. Loop through the remaining 4 cards
3. If any card has a different suit, return `false`
4. If we make it through the loop without finding a mismatch, return `true`

**Example**:
```
Input: [Ah Kh Qh Jh Th]

firstSuit = Hearts (from Ah)
Check Kh: Hearts == Hearts? ✓
Check Qh: Hearts == Hearts? ✓
Check Jh: Hearts == Hearts? ✓
Check Th: Hearts == Hearts? ✓

Result: true (it's a flush)
```

**Complexity**:
- **Time**: O(n) where n = number of cards (5 in this case)
- **Space**: O(1) - we only store one variable (`firstSuit`)

### isStraight: Are the Cards in Sequence?

```go
func isStraight(cards []Card) (bool, Rank) {
    if len(cards) != 5 {
        return false, 0
    }

    // Extract and sort ranks in descending order
    ranks := make([]Rank, 5)
    for i, card := range cards {
        ranks[i] = card.Rank
    }

    // Sort ranks in descending order using insertion sort
    sortRanksDescending(ranks)

    // Check for wheel straight: A-2-3-4-5
    if ranks[0] == Ace && ranks[1] == Five && ranks[2] == Four && ranks[3] == Three && ranks[4] == Two {
        return true, Five  // Ace acts as low, high card is Five
    }

    // Check for regular straight
    for i := 1; i < 5; i++ {
        if ranks[i] != ranks[i-1]-1 {
            return false, 0
        }
    }

    // Regular straight found
    return true, ranks[0]
}

// sortRanksDescending sorts a slice of Rank values in descending order using insertion sort.
// Optimized for small arrays (n ≤ 10) with zero allocations.
func sortRanksDescending(ranks []Rank) {
    for i := 1; i < len(ranks); i++ {
        key := ranks[i]
        j := i - 1
        // Shift elements smaller than key to the right (descending order)
        for j >= 0 && ranks[j] < key {
            ranks[j+1] = ranks[j]
            j--
        }
        ranks[j+1] = key
    }
}
```

**What does this do?** Checks if 5 cards form a sequence (like 9-8-7-6-5).

**Returns**:
- `bool`: true if it's a straight, false otherwise
- `Rank`: the highest card in the straight (used for tiebreakers)

**Algorithm**:

**Step 1: Extract ranks** (lines 8-11)
```
Input cards:  [9h 8d 7c 6s 5h]
Extract ranks: [9, 8, 7, 6, 5]
```

**Step 2: Sort ranks descending** (lines 13-19)
```
Before sort: [9, 8, 7, 6, 5]
After sort:  [9, 8, 7, 6, 5]  (already sorted in this example)
```

We use **insertion sort** here. While insertion sort has O(n²) time complexity like bubble sort, it's **29% faster** for small arrays (n=5) due to better cache locality and fewer comparisons.

**How insertion sort works** (for beginners):
```
Build sorted array one element at a time by inserting each element into its correct position:

Start:   [9, 8, 7, 6, 5]
Step 1:  [9] | 8, 7, 6, 5        (first element already "sorted")
Step 2:  [9, 8] | 7, 6, 5        (8 < 9, already correct position)
Step 3:  [9, 8, 7] | 6, 5        (7 < 8, already correct position)
Step 4:  [9, 8, 7, 6] | 5        (6 < 7, already correct position)
Step 5:  [9, 8, 7, 6, 5]         (5 < 6, already correct position)

Result: [9, 8, 7, 6, 5]
```

**Why insertion sort instead of stdlib sort.Slice?**
- For tiny arrays (n ≤ 10), insertion sort is faster than Go's `sort.Slice`
- `sort.Slice` uses reflection and creates allocations (3+ per call)
- Insertion sort is zero-allocation and has better constant factors for small n
- Benchmarks showed 29% performance improvement over bubble sort
- See `SORTING_ANALYSIS.md` for detailed algorithm comparison

**Step 3: Check for wheel straight** (lines 21-24)

The "wheel" is a special case: **A-2-3-4-5**

When sorted descending, this looks like: `[14, 5, 4, 3, 2]` (Ace=14, Five=5, etc.)

**Why special?** In a wheel, the Ace acts as a "low" card (value 1, not 14). So the high card of the straight is **5**, not Ace.

**Step 4: Check for regular straight** (lines 26-31)

A straight means each rank is **exactly 1 less** than the previous:

```
ranks = [9, 8, 7, 6, 5]

Check: 8 == 9-1? ✓ (8 == 8)
Check: 7 == 8-1? ✓ (7 == 7)
Check: 6 == 7-1? ✓ (6 == 6)
Check: 5 == 6-1? ✓ (5 == 5)

All checks passed → It's a straight!
High card = ranks[0] = 9
```

### rankCounts: How Many of Each Rank?

```go
func rankCounts(cards []Card) map[Rank]int {
    counts := make(map[Rank]int)

    for _, card := range cards {
        counts[card.Rank]++
    }

    return counts
}
```

**What does this do?** Counts how many cards of each rank we have.

**Example**:
```
Input:  [Ah Ad Kh Qd Jc]
        (Ace, Ace, King, Queen, Jack)

Output: map[Rank]int{
    Ace:   2,  // Two Aces
    King:  1,  // One King
    Queen: 1,  // One Queen
    Jack:  1,  // One Jack
}
```

**Why is this useful?**
- Detecting pairs: look for rank with count = 2
- Detecting trips: look for rank with count = 3
- Detecting quads: look for rank with count = 4
- Detecting full house: look for count=3 AND count=2

**How it works**:

1. Create an empty map (dictionary/hash table)
2. Loop through each card
3. Increment the count for that card's rank

**Line-by-line**:
```go
counts := make(map[Rank]int)
// Creates an empty map: {} (no entries yet)

for _, card := range cards {
// Loop through each card. The underscore _ means "I don't care about the index"

    counts[card.Rank]++
// If Ace already in map: increment count
// If Ace NOT in map: create entry with count=1 (Go does this automatically)
}
```

**Complexity**:
- **Time**: O(n) - we loop through n cards once
- **Space**: O(k) - where k is number of unique ranks (max 5)

---

## Detection Functions Explained

Each detection function checks for a specific hand category. Let's examine a few in detail.

### detectRoyalFlush: Is It 10-J-Q-K-A of the Same Suit?

```go
func detectRoyalFlush(cards []Card) bool {
    if len(cards) != 5 {
        return false
    }

    // Check if all cards are the same suit
    if !isFlush(cards) {
        return false
    }

    // Check for the exact ranks: 10, J, Q, K, A
    requiredRanks := map[Rank]bool{
        Ten:   true,
        Jack:  true,
        Queen: true,
        King:  true,
        Ace:   true,
    }

    // Verify all required ranks are present
    for _, card := range cards {
        if !requiredRanks[card.Rank] {
            return false  // Found a card that's not Ten/Jack/Queen/King/Ace
        }
        delete(requiredRanks, card.Rank)  // Mark this rank as "seen"
    }

    // All required ranks should be consumed (map should be empty)
    return len(requiredRanks) == 0
}
```

**Strategy**:
1. Must be a flush (all same suit)
2. Must have exactly these 5 ranks: Ten, Jack, Queen, King, Ace

**Clever trick** (lines 20-26):
- Create a map of required ranks
- As we loop through cards, **delete** ranks we've seen from the map
- If the map is empty at the end, we saw all 5 required ranks

**Why delete from the map?**
- Prevents duplicates: if we had two Tens somehow (impossible with 5 cards, but safe)
- After processing 5 cards with the correct ranks, the map will be empty
- If any required rank is missing, the map won't be empty

**Example**:
```
Input: [Ah Kh Qh Jh Th]

Check isFlush: true ✓

requiredRanks = {Ten: true, Jack: true, Queen: true, King: true, Ace: true}

Process Ah:  Ace is in requiredRanks? Yes → delete Ace
             requiredRanks = {Ten: true, Jack: true, Queen: true, King: true}

Process Kh:  King is in requiredRanks? Yes → delete King
             requiredRanks = {Ten: true, Jack: true, Queen: true}

Process Qh:  Queen is in requiredRanks? Yes → delete Queen
             requiredRanks = {Ten: true, Jack: true}

Process Jh:  Jack is in requiredRanks? Yes → delete Jack
             requiredRanks = {Ten: true}

Process Th:  Ten is in requiredRanks? Yes → delete Ten
             requiredRanks = {} (empty!)

len(requiredRanks) == 0? true → It's a Royal Flush!
```

### detectFourOfAKind: Do We Have 4 Cards of the Same Rank?

```go
func detectFourOfAKind(cards []Card) (bool, []Rank) {
    if len(cards) != 5 {
        return false, []Rank{}
    }

    counts := rankCounts(cards)

    var quadRank Rank
    var kicker Rank

    // Find the rank that appears 4 times
    for rank, count := range counts {
        if count == 4 {
            quadRank = rank
        } else if count == 1 {
            kicker = rank
        }
    }

    // If we found a quad rank, we have four of a kind
    if quadRank != 0 {
        return true, []Rank{quadRank, kicker}
    }

    return false, []Rank{}
}
```

**Returns**:
- `bool`: true if four of a kind, false otherwise
- `[]Rank`: tiebreakers `[quad rank, kicker rank]`

**Algorithm**:

1. **Count ranks** (line 6): Get map of rank → count
   ```
   Input: [Ah Ad Ac As Kh]
   counts = {Ace: 4, King: 1}
   ```

2. **Find quad and kicker** (lines 11-18):
   - Loop through the counts
   - If count == 4, that's our quad rank (the four-of-a-kind)
   - If count == 1, that's our kicker (the 5th card)

3. **Return result** (lines 21-25):
   - If we found a quad rank (not zero), return true + tiebreakers
   - Otherwise, return false

**Why `quadRank != 0`?**
- `Rank` is an int type, default value is 0
- The lowest actual rank is `Two = 2`
- So if `quadRank` is still 0, we never found a quad

**Tiebreaker format**: `[quad rank, kicker]`
```
Example: [Ah Ad Ac As Kh]
Tiebreakers: [14, 13]  (Quad Aces with King kicker)
```

### detectTwoPair: Do We Have Two Different Pairs?

```go
func detectTwoPair(cards []Card) (bool, []Rank) {
    if len(cards) != 5 {
        return false, nil
    }

    counts := rankCounts(cards)

    pairs := make([]Rank, 0, 2)
    hasTrips := false

    for rank, count := range counts {
        if count == 2 {
            pairs = append(pairs, rank)
        } else if count >= 3 {
            hasTrips = true
        }
    }

    // Must have exactly two pairs and no trips/quads
    if len(pairs) != 2 || hasTrips {
        return false, nil
    }

    // Sort pairs descending (high pair first)
    if pairs[0] < pairs[1] {
        pairs[0], pairs[1] = pairs[1], pairs[0]
    }

    // Find the kicker
    var kicker Rank
    for rank, count := range counts {
        if count == 1 {
            kicker = rank
            break
        }
    }

    return true, []Rank{pairs[0], pairs[1], kicker}
}
```

**Strategy**:
1. Count ranks
2. Find all ranks with count == 2 (these are pairs)
3. Make sure we have **exactly 2** pairs (not 0, not 1, not 3)
4. Make sure we have **no trips or quads** (if count >= 3, it's a full house, not two pair)
5. Sort pairs so higher pair comes first
6. Find the kicker (the 5th card)

**Edge case handling** (line 19-21):
```go
if len(pairs) != 2 || hasTrips {
    return false, nil
}
```

Why check `hasTrips`?
```
Example: [Ah Ad Ac Kh Kd]
counts = {Ace: 3, King: 2}
pairs = [King]  (only one pair found)
hasTrips = true (Ace appears 3 times)

This is a FULL HOUSE, not two pair!
Return false here, so the full house detector can catch it instead.
```

**Sorting pairs** (lines 24-27):
```go
if pairs[0] < pairs[1] {
    pairs[0], pairs[1] = pairs[1], pairs[0]
}
```

This ensures the **higher pair comes first** in our tiebreakers.

```
Example: pairs = [King, Ace]  (King=13, Ace=14)
13 < 14? true → swap them
Result:  pairs = [Ace, King]
```

**Tiebreaker format**: `[high pair, low pair, kicker]`
```
Example: [Ah Ad Kh Kd Qc]
Tiebreakers: [14, 13, 12]  (Aces and Kings, Queen kicker)
```

---

## Combination Generation Algorithm

### How Do We Generate All 5-Card Combinations from 7 Cards?

This is the trickiest part. We use **recursion** (a function that calls itself).

```go
func Combinations(cards []Card, k int) [][]Card {
    var result [][]Card

    if k == 0 {
        return [][]Card{{}}
    }

    if k == len(cards) {
        return [][]Card{cards}
    }

    if k > len(cards) {
        return result
    }

    generate(cards, k, 0, []Card{}, &result)

    return result
}
```

**Parameters**:
- `cards`: The full set of cards (e.g., 7 cards)
- `k`: How many cards to pick (e.g., 5)

**Returns**: All possible k-card combinations

**Base cases** (lines 4-13):
1. If k == 0 (pick 0 cards): return one empty combination `[[]]`
2. If k == len(cards) (pick all cards): return all cards as one combination
3. If k > len(cards) (impossible): return empty result

**Recursive case** (line 15): Call the helper function `generate`

### The Generate Helper Function (Recursion Explained)

```go
func generate(cards []Card, k int, start int, current []Card, result *[][]Card) {
    // Base case: if we've selected k cards, add this combination to result
    if k == 0 {
        combo := make([]Card, len(current))
        copy(combo, current)
        *result = append(*result, combo)
        return
    }

    // Recursive case: try including each remaining card
    for i := start; i <= len(cards)-k; i++ {
        generate(cards, k-1, i+1, append(current, cards[i]), result)
    }
}
```

**Parameters**:
- `cards`: Full set of cards
- `k`: How many more cards we need to pick
- `start`: Where to start picking from (avoids duplicates)
- `current`: The combination we're building
- `result`: Pointer to the final result (we append combinations here)

**How recursion works** (for absolute beginners):

Imagine you're choosing pizza toppings. You have 7 toppings and need to pick 5.

**Recursive thinking**:
1. Pick the first topping (or don't)
2. Now you have a smaller problem: pick 4 more toppings from the remaining 6
3. Repeat until you've picked 5 toppings

**Code walkthrough with a small example**:

```
Choose 2 cards from [A, K, Q]

Call: generate([A,K,Q], k=2, start=0, current=[], result=&[])

Step 1: Pick A
  Call: generate([A,K,Q], k=1, start=1, current=[A], result=&[])

  Step 1a: Pick K
    Call: generate([A,K,Q], k=0, start=2, current=[A,K], result=&[])
    k==0 → Add [A,K] to result
    result = [[A,K]]

  Step 1b: Pick Q
    Call: generate([A,K,Q], k=0, start=3, current=[A,Q], result=&[[A,K]])
    k==0 → Add [A,Q] to result
    result = [[A,K], [A,Q]]

Step 2: Pick K (skip A)
  Call: generate([A,K,Q], k=1, start=2, current=[K], result=&[[A,K],[A,Q]])

  Step 2a: Pick Q
    Call: generate([A,K,Q], k=0, start=3, current=[K,Q], result=&[[A,K],[A,Q]])
    k==0 → Add [K,Q] to result
    result = [[A,K], [A,Q], [K,Q]]

Final result: [[A,K], [A,Q], [K,Q]]
```

**Key insight**: The `start` parameter prevents duplicates.
- When we pick `A` first, we only look at `K` and `Q` next (not `A` again)
- When we pick `K` first, we only look at `Q` next (not `A` or `K`)
- This ensures we generate `[A,K]` but never `[K,A]` (they're the same combination)

**Why copy the combination?** (lines 4-6)
```go
combo := make([]Card, len(current))
copy(combo, current)
*result = append(*result, combo)
```

Without copying:
```go
*result = append(*result, current)  // BUG!
```

This would append the **same slice** multiple times. In Go, slices are references. When we modify `current` later, it would change all the previous combinations we added!

By copying, we create a **new, independent slice** for each combination.

---

## Performance Analysis

### Time Complexity

**FindBestHand with 7 cards**:

1. **Generate combinations**: C(7,5) = 21 combinations
   - Time: O(C(n,k)) where n=7, k=5
   - For 7 cards: O(21) = O(1) (constant, since input size is fixed)

2. **Evaluate each combination**: ~10 checks per combination (worst case)
   - Time: O(10) per combination
   - Total: 21 × 10 = 210 operations

3. **Compare hands**: Compare 21 hands
   - Time: O(21) comparisons

**Total**: O(21 + 210 + 21) = O(252) = **O(1)** (constant time for fixed input size)

**Why is this O(1)?**
- We always have at most 7 cards
- C(7,5) is always 21
- 21 × 10 = 210 is constant
- **For a fixed maximum input size, the algorithm is constant time**

**But doesn't it scale poorly?**
- If we had 10 cards: C(10,5) = 252 combinations → 2,520 operations
- If we had 13 cards: C(13,5) = 1,287 combinations → 12,870 operations
- For poker, we never have more than 7 cards, so this doesn't matter

### Space Complexity

**Memory usage**:

1. **Combinations array**: 21 combinations × 5 cards = 105 card references
   - Each Card is 16 bytes (8 bytes for Rank + 8 bytes for Suit)
   - Total: 105 × 16 = 1,680 bytes ≈ **1.6 KB**

2. **Temporary variables**: negligible (a few hundred bytes)

**Total**: O(C(n,k) × k) = O(21 × 5) = **O(105) = O(1)** (constant space)

### Comparison with Lookup Table Approach

| Approach       | Time      | Space   | Complexity | Use Case                |
|----------------|-----------|---------|------------|-------------------------|
| Our Algorithm  | O(1)      | O(1)    | Low        | Educational, readable   |
| Lookup Table   | O(1)      | O(10MB) | High       | High-performance poker bots |
| Bit Manipulation | O(1)    | O(1)    | Very High  | Extreme optimization    |

**Our choice**: Clarity over performance
- For 7 cards, 252 operations is **negligible** on modern CPUs (~1 microsecond)
- Only relevant for high-frequency poker simulations (millions of hands per second)
- For learning and most applications, our approach is perfect

### Actual Performance Benchmarks

Here are **real benchmark results** from the codebase after Phase 06 optimization (insertion sort), measured on an AMD Ryzen 5 3600 (6-core) processor:

```
BenchmarkEvaluateHand-12          	 1581703	       747 ns/op	     216 B/op	       9 allocs/op
BenchmarkFindBestHand5Cards-12    	 2633970	       455 ns/op	     112 B/op	       2 allocs/op
BenchmarkFindBestHand6Cards-12    	  201056	      5823 ns/op	    3824 B/op	      77 allocs/op
BenchmarkFindBestHand7Cards-12    	   64371	     18778 ns/op	   11208 B/op	     229 allocs/op
BenchmarkCombinations-12          	  337502	      3407 ns/op	    6888 B/op	      67 allocs/op
```

**What do these numbers mean?**

| Benchmark | Time per operation | Memory per op | Allocations | Operations/sec |
|-----------|-------------------|---------------|-------------|----------------|
| **EvaluateHand** (5 cards) | 747 ns | 216 bytes | 9 | ~1,339,000 |
| **FindBestHand** (5 cards) | 455 ns | 112 bytes | 2 | ~2,198,000 |
| **FindBestHand** (6 cards) | 5,823 ns (~6 μs) | 3,824 bytes | 77 | ~171,700 |
| **FindBestHand** (7 cards) | 18,778 ns (~19 μs) | 11,208 bytes | 229 | ~53,250 |
| **Combinations** (7→5) | 3,407 ns (~3 μs) | 6,888 bytes | 67 | ~293,500 |

**Key Insights:**

1. **7-card evaluation is ~41x slower than 5-card** (18,778 ns vs 455 ns)
   - This makes sense: 21 combinations vs 1 combination
   - Still incredibly fast: ~19 microseconds = 0.000019 seconds

2. **6-card evaluation is ~13x slower than 5-card** (5,823 ns vs 455 ns)
   - 6 combinations vs 1 combination
   - Matches theoretical expectations

3. **Combination generation is ~18% of total 7-card time**
   - 3,407 ns out of 18,778 ns total
   - Remaining time is evaluation + comparison

4. **Real-world throughput** (after Phase 06 optimization):
   - **7 cards**: Can evaluate ~53,250 hands per second (single-threaded)
   - **5 cards**: Can evaluate ~2.2 million hands per second
   - For a typical poker game with 10 players: can simulate ~5,325 complete hands/sec
   - **29% faster** than pre-optimization baseline

5. **Memory efficiency**:
   - 7-card evaluation uses only ~11 KB per operation
   - 229 allocations per evaluation (future optimization target)
   - Most allocations come from combination generation (67) and evaluation (9 per combo)

6. **Phase 06 Optimization Impact** (Insertion Sort):
   - **Before**: 1,017 ns/op for EvaluateHand
   - **After**: 747 ns/op for EvaluateHand
   - **Improvement**: 29% faster, zero memory overhead
   - Achieved by replacing `sort.Slice` with custom insertion sort (eliminates reflection overhead)

**Performance is "Excellent":**
- For interactive applications (poker games, hand analyzers): **excellent performance**
- For batch simulations (Monte Carlo analysis): **excellent** (53K hands/sec covers most use cases)
- For high-frequency poker bots: **good** (lookup tables would be 100-1000x faster if needed)

**Optimization Journey:**
- ✅ **Phase 06 completed**: Insertion sort optimization (+29% speed)
- 🔄 **Future potential**: Further optimizations could add another 2-5x improvement
- 🔄 **Concurrency**: Multiple cores could add 4-6x speedup
- Total potential: ~10-40x faster than original baseline with full optimization

**Bottom Line:** The current implementation achieves **excellent performance** while maintaining code clarity. After Phase 06 optimizations, it's fast enough for 99%+ of use cases. If you need to evaluate millions of hands per second, consider lookup tables (at the cost of significant complexity).

---

## Common Pitfalls and Edge Cases

### 1. The Wheel Straight (A-2-3-4-5)

**Problem**: In poker, Ace can be high OR low.
- High: 10-J-Q-K-A (Ace = 14)
- Low: A-2-3-4-5 (Ace = 1)

**Our solution** (in `isStraight`):
```go
if ranks[0] == Ace && ranks[1] == Five && ranks[2] == Four && ranks[3] == Three && ranks[4] == Two {
    return true, Five  // High card is Five, not Ace!
}
```

**Why return `Five` as the high card?**
- A wheel straight (A-2-3-4-5) is the **lowest possible straight**
- It loses to 6-5-4-3-2, which loses to 7-6-5-4-3, etc.
- By returning `Five` (5), we ensure proper comparison
- Otherwise, returning `Ace` (14) would make it seem stronger than K-Q-J-10-9!

### 2. Multiple Flush Suits in 7 Cards

**Scenario**: You have 7 cards with 4 hearts and 3 diamonds.

```
Cards: [Ah Kh Qh Jh 9d 8d 7d]

Possible combinations:
- [Ah Kh Qh Jh 9d] → Not a flush (4 hearts + 1 diamond)
- [Ah Kh Qh 9d 8d] → Not a flush (3 hearts + 2 diamonds)
- [Ah Kh Qh 9d 7d] → Not a flush (3 hearts + 2 diamonds)
... etc.
```

**How our algorithm handles this**:
- It tries all 21 combinations
- Each combination is checked independently for flush
- If a 5-card combination has all hearts OR all diamonds, `isFlush` returns true
- If a combination mixes suits, `isFlush` returns false

**No special handling needed!** The brute-force approach naturally handles this.

### 3. Tiebreaker Edge Case: Identical Kickers

**Scenario**: Both players have the exact same hand.

```
Player 1: [Ah Ad Kh Qd Jc]  (Pair of Aces, K-Q-J kickers)
Player 2: [As Ac Ks Qc Js]  (Pair of Aces, K-Q-J kickers)

Tiebreakers: [14, 13, 12, 11] vs [14, 13, 12, 11]

Result: Split pot (0 from CompareHands)
```

**Our comparison logic**:
```go
for i := 0; i < minLen; i++ {
    if hand1.Tiebreakers[i] > hand2.Tiebreakers[i] {
        return 1
    }
    if hand1.Tiebreakers[i] < hand2.Tiebreakers[i] {
        return -1
    }
}
return 0  // All tiebreakers equal
```

If we compare all tiebreakers and they're all equal, we return 0 (tie).

### 4. Straight Flush vs. Royal Flush

**Question**: Why check Royal Flush separately? Isn't it just a Straight Flush?

**Answer**: Yes, a Royal Flush IS a Straight Flush. But:
- Royal Flush has **no tiebreakers** (all Royal Flushes are equal)
- Straight Flush has **one tiebreaker** (the high card)

```go
// Royal Flush
return &Hand{
    Cards:       cards,
    Category:    RoyalFlush,
    Tiebreakers: []Rank{},  // Empty!
}

// Straight Flush
return &Hand{
    Cards:       cards,
    Category:    StraightFlush,
    Tiebreakers: []Rank{highCard},  // Has a high card
}
```

By checking Royal Flush first, we avoid the question "what's the high card of a Royal Flush?" (Answer: always Ace, but we don't need to store it since all Royal Flushes are equal).

---

## Why This Design?

### Design Principle 1: Clarity Over Performance (with Pragmatic Optimization)

**We chose**:
- Brute-force combination generation
- Sequential checking of hand categories
- Insertion sort for 5 elements (optimized from bubble sort)

**We could have chosen**:
- Lookup tables (O(1) but complex)
- Bit manipulation (fast but unreadable)
- Hash-based perfect hashing (maximum performance)

**Why our choice?**
- **Learning**: You can read the code and understand what it does
- **Debugging**: When something goes wrong, you can trace the logic
- **Testing**: Easy to write tests for each function
- **Maintenance**: Future developers can modify it without specialized knowledge
- **Pragmatic optimization**: When profiling revealed sorting as a bottleneck, we optimized with insertion sort (29% faster) while maintaining readability

**Performance is "excellent"**:
- ~250 operations for 7 cards
- Modern CPUs: billions of operations per second
- Our algorithm: ~19 microseconds per 7-card hand (after Phase 06 optimization)
- Fast enough for real-time poker games and most simulations
- Only matters if you're simulating millions of hands per second

### Design Principle 2: Separation of Concerns

Each layer has a single responsibility:

1. **Foundation**: Define what a card is
2. **Helpers**: Basic checks (flush, straight, counting)
3. **Detectors**: Recognize each hand category
4. **Evaluation**: Orchestrate detectors
5. **Combination**: Generate possibilities
6. **Comparison**: Determine winners

**Benefits**:
- Easy to test each part independently
- Easy to understand each part in isolation
- Changes to one part don't break others

### Design Principle 3: Fail-Safe Defaults

**Examples**:
- `EvaluateHand` always returns a hand (falls back to High Card)
- `CompareHands` returns 0 for nil hands (safe default)
- `FindBestHand` returns nil for invalid input (explicit failure)

**Why?**
- Defensive programming: handle edge cases gracefully
- No crashes from unexpected input
- Clear contracts: functions specify what they return for all cases

### Design Principle 4: Test-Driven Development

Every function was built using TDD:
1. Write tests first (red)
2. Implement to pass tests (green)
3. Refactor for clarity (refactor)

**Result**:
- 92.7% code coverage
- Comprehensive edge case testing
- Confidence in correctness

---

## Summary

You now understand:

✅ **What**: A poker hand evaluator that finds the best 5-card hand from 5-7 cards

✅ **How**:
- Generate all 5-card combinations
- Evaluate each combination (check hand categories strongest→weakest)
- Compare hands (category first, then tiebreakers)
- Return the strongest hand

✅ **Why**:
- Clarity over performance (but still fast enough)
- Separation of concerns (each function has one job)
- Defensive programming (handle edge cases)
- Test-driven development (high confidence in correctness)

✅ **Data structures**:
- `Rank` (2-14, Ace high)
- `Suit` (0-3, no rank order)
- `Card` (rank + suit)
- `HandCategory` (1-10, higher is stronger)
- `Hand` (cards + category + tiebreakers)

✅ **Key algorithms**:
- `isFlush`: Check all cards same suit
- `isStraight`: Check cards sequential (handles wheel)
- `rankCounts`: Count duplicates (for pairs, trips, quads)
- `Combinations`: Recursive combination generation
- `EvaluateHand`: Cascade of detectors
- `CompareHands`: Category-first comparison

✅ **Edge cases**:
- Wheel straight (A-2-3-4-5)
- Royal Flush (special case of straight flush)
- Multiple flush suits in 7 cards
- Identical hands (split pot)

---

## Next Steps

Now that you understand the algorithm:

1. **Read the actual code** (`pkg/poker/*.go`) - it should make sense now
2. **Run the tests** (`go test -v ./...`) - see the algorithm in action
3. **Try the example** (`go run cmd/example/main.go`) - play with it
4. **Modify it** - change detection logic, add logging, experiment
5. **Optimize it** - now that you understand it, you can make it faster (see Phase 06 issues)

**Remember**: Understanding comes before optimization. You now have a solid foundation.

---

**Questions?** Re-read sections that are unclear. The best way to learn is to trace through examples by hand with pen and paper. Pick a 7-card hand and manually generate all 21 combinations, then evaluate each one. This will cement your understanding.

**Happy coding!** 🎰♠️♥️♣️♦️
