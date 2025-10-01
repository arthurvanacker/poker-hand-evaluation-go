# Poker Hand Evaluator

A high-quality Texas Hold'em poker hand evaluator written in Go that identifies the best 5-card hand from 5, 6, or 7 cards. Supports all standard poker hand categories from Royal Flush to High Card with proper tiebreaker resolution.

[![Go Version](https://img.shields.io/badge/Go-1.24.7-blue.svg)](https://golang.org/dl/)
[![Test Coverage](https://img.shields.io/badge/coverage-92.7%25-brightgreen.svg)](https://github.com/puupa/poker-hand-evaluation)

## Features

- **Complete Hand Evaluation**: Recognizes all 10 poker hand categories
- **Multi-Card Support**: Evaluates 5, 6, or 7 cards (perfect for Texas Hold'em)
- **Accurate Tiebreaker Logic**: Properly resolves ties within the same hand category
- **Special Cases Handled**: Wheel straight (A-2-3-4-5), royal flush detection
- **Standard 52-Card Deck**: Full deck implementation with shuffle and deal operations
- **Card Parsing**: Flexible card notation parser supporting "Ah", "10s", "Td" formats
- **High Test Coverage**: 92.7% code coverage with comprehensive edge case testing
- **Clean API**: Simple, idiomatic Go interfaces following SOLID principles

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Usage Guide](#usage-guide)
  - [Creating Cards](#creating-cards)
  - [Working with Decks](#working-with-decks)
  - [Evaluating Hands](#evaluating-hands)
  - [Comparing Hands](#comparing-hands)
  - [Finding the Best Hand](#finding-the-best-hand)
- [API Reference](#api-reference)
  - [Core Types](#core-types)
  - [Card Operations](#card-operations)
  - [Deck Operations](#deck-operations)
  - [Hand Evaluation](#hand-evaluation)
  - [Hand Comparison](#hand-comparison)
- [Architecture](#architecture)
- [Hand Categories](#hand-categories)
- [Development](#development)
- [Testing](#testing)
- [License](#license)

## Installation

### Prerequisites

- Go 1.24.7 or higher

### Install Package

```bash
go get github.com/puupa/poker-hand-evaluation
```

### Clone Repository

```bash
git clone https://github.com/puupa/poker-hand-evaluation.git
cd poker-hand-evaluation
```

### Verify Installation

```bash
# Run tests
go test -v ./...

# Run example program
go run cmd/example/main.go
```

## Quick Start

Here's a simple example demonstrating the core functionality:

```go
package main

import (
    "fmt"
    "github.com/puupa/poker-hand-evaluation/pkg/poker"
)

func main() {
    // Create cards using ParseCard
    cards := []poker.Card{}
    cardStrings := []string{"Ah", "Kh", "Qh", "Jh", "Th"}

    for _, s := range cardStrings {
        card, _ := poker.ParseCard(s)
        cards = append(cards, card)
    }

    // Evaluate the hand
    hand := poker.EvaluateHand(cards)
    fmt.Printf("Hand: %s\n", hand.Category) // Output: Hand: Royal Flush
}
```

## Usage Guide

### Creating Cards

Cards can be created using the `Card` struct or parsed from string notation:

```go
import "github.com/puupa/poker-hand-evaluation/pkg/poker"

// Method 1: Direct construction
card1 := poker.Card{Rank: poker.Ace, Suit: poker.Hearts}

// Method 2: Parse from string (recommended)
card2, err := poker.ParseCard("Ah")  // Ace of Hearts
card3, err := poker.ParseCard("10d") // Ten of Diamonds
card4, err := poker.ParseCard("Ts")  // Ten of Spades (alternative notation)

if err != nil {
    fmt.Printf("Parse error: %v\n", err)
}

// Card string representation
fmt.Println(card1) // Output: Ah
```

#### Supported Card Notations

- **Ranks**: `A` (Ace), `K` (King), `Q` (Queen), `J` (Jack), `T` or `10` (Ten), `9`, `8`, `7`, `6`, `5`, `4`, `3`, `2`
- **Suits**: `h` (Hearts), `d` (Diamonds), `c` (Clubs), `s` (Spades)
- **Examples**: `Ah`, `Kd`, `10s`, `Ts`, `2c`

### Working with Decks

Create and manipulate a standard 52-card deck:

```go
import (
    "math/rand"
    "time"
    "github.com/puupa/poker-hand-evaluation/pkg/poker"
)

// Create a new deck
deck := poker.NewDeck()
fmt.Printf("Deck has %d cards\n", len(deck.Cards)) // Output: 52

// Shuffle the deck (Fisher-Yates algorithm)
rand.Seed(time.Now().UnixNano())
for i := len(deck.Cards) - 1; i > 0; i-- {
    j := rand.Intn(i + 1)
    deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
}

// Deal cards
holeCards, err := deck.Deal(2)
if err != nil {
    fmt.Printf("Error: %v\n", err)
}

communityCards, err := deck.Deal(5)
fmt.Printf("Remaining: %d cards\n", len(deck.Cards)) // Output: 45
```

### Evaluating Hands

Evaluate exactly 5 cards to determine hand category and tiebreakers:

```go
import "github.com/puupa/poker-hand-evaluation/pkg/poker"

// Parse 5 cards
cardStrings := []string{"Ah", "Kh", "Qh", "Jh", "Th"}
cards := []poker.Card{}

for _, s := range cardStrings {
    card, _ := poker.ParseCard(s)
    cards = append(cards, card)
}

// Evaluate the hand
hand := poker.EvaluateHand(cards)

// Access hand information
fmt.Printf("Category: %s\n", hand.Category)           // Royal Flush
fmt.Printf("Tiebreakers: %v\n", hand.Tiebreakers)     // []
fmt.Printf("Cards: %v\n", hand.Cards)                 // [Ah Kh Qh Jh Th]
```

### Comparing Hands

Compare two hands to determine the winner:

```go
import "github.com/puupa/poker-hand-evaluation/pkg/poker"

// Create two hands
cards1, _ := parseCards([]string{"Ah", "Ad", "Kh", "Kd", "Qh"})
cards2, _ := parseCards([]string{"Ks", "Kc", "Qs", "Qc", "Jh"})

hand1 := poker.EvaluateHand(cards1)
hand2 := poker.EvaluateHand(cards2)

// Compare hands
result := poker.CompareHands(hand1, hand2)

if result > 0 {
    fmt.Println("Hand 1 wins")
} else if result < 0 {
    fmt.Println("Hand 2 wins")
} else {
    fmt.Println("Split pot - hands are equal")
}

// Helper function
func parseCards(strs []string) ([]poker.Card, error) {
    cards := []poker.Card{}
    for _, s := range strs {
        card, err := poker.ParseCard(s)
        if err != nil {
            return nil, err
        }
        cards = append(cards, card)
    }
    return cards, nil
}
```

### Finding the Best Hand

Automatically find the best 5-card hand from 5, 6, or 7 cards:

```go
import "github.com/puupa/poker-hand-evaluation/pkg/poker"

// Texas Hold'em scenario: 2 hole cards + 5 community cards
holeCards, _ := parseCards([]string{"Ah", "Kh"})
communityCards, _ := parseCards([]string{"Qh", "Jh", "Th", "9d", "2c"})

// Combine all 7 cards
allCards := append(holeCards, communityCards...)

// Find best possible 5-card hand
bestHand := poker.FindBestHand(allCards)

fmt.Printf("Best hand: %s\n", bestHand.Category) // Royal Flush

// Works with 5 or 6 cards too
fiveCards, _ := parseCards([]string{"Ah", "Kh", "Qh", "Jh", "Th"})
best5 := poker.FindBestHand(fiveCards) // Evaluates directly (optimization)
```

## API Reference

### Core Types

#### `Rank`

Represents card rank with numeric values for easy comparison.

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

**Methods:**
- `String() string` - Returns single-character rank representation (`"A"`, `"K"`, `"Q"`, `"J"`, `"T"`, `"9"`, etc.)

#### `Suit`

Represents card suit.

```go
type Suit int

const (
    Hearts   Suit = iota // 0
    Diamonds             // 1
    Clubs                // 2
    Spades               // 3
)
```

**Methods:**
- `String() string` - Returns single-character suit representation (`"h"`, `"d"`, `"c"`, `"s"`)

#### `Card`

Represents a playing card.

```go
type Card struct {
    Rank Rank
    Suit Suit
}
```

**Methods:**
- `String() string` - Returns card notation (e.g., `"Ah"`, `"Kd"`, `"10s"`)

#### `HandCategory`

Enumeration of poker hand categories ordered by strength.

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

**Methods:**
- `String() string` - Returns human-readable category name (e.g., `"Royal Flush"`, `"Full House"`)

#### `Hand`

Represents an evaluated poker hand.

```go
type Hand struct {
    Cards       []Card       // The 5 cards in the hand
    Category    HandCategory // The hand category
    Tiebreakers []Rank       // Ranks for tiebreaker comparison
}
```

**Methods:**
- `NewHand(cards []Card) (*Hand, error)` - Creates a new hand (requires exactly 5 cards)

#### `Deck`

Represents a collection of playing cards.

```go
type Deck struct {
    Cards []Card
}
```

### Card Operations

#### `ParseCard`

Parses a card string into a `Card` struct.

```go
func ParseCard(s string) (Card, error)
```

**Parameters:**
- `s` - Card string (e.g., `"Ah"`, `"10d"`, `"Ts"`)

**Returns:**
- `Card` - Parsed card
- `error` - Error if parsing fails

**Examples:**
```go
card, err := poker.ParseCard("Ah")   // Ace of Hearts
card, err := poker.ParseCard("10d")  // Ten of Diamonds (using "10")
card, err := poker.ParseCard("Ts")   // Ten of Spades (using "T")
card, err := poker.ParseCard("2c")   // Two of Clubs
```

**Error Cases:**
- String too short (< 2 characters)
- Invalid rank
- Invalid suit
- Invalid length (not 2 or 3 characters)

### Deck Operations

#### `NewDeck`

Creates a new standard 52-card deck.

```go
func NewDeck() *Deck
```

**Returns:**
- `*Deck` - Pointer to new deck containing all 52 cards

**Card Order:**
Cards are ordered by suit (Hearts, Diamonds, Clubs, Spades), then by rank (2-A) within each suit.

**Example:**
```go
deck := poker.NewDeck()
fmt.Println(len(deck.Cards)) // Output: 52
```

#### `Deal`

Removes and returns the top n cards from the deck.

```go
func (d *Deck) Deal(n int) ([]Card, error)
```

**Parameters:**
- `n` - Number of cards to deal

**Returns:**
- `[]Card` - Slice of dealt cards
- `error` - Error if insufficient cards available

**Example:**
```go
deck := poker.NewDeck()
cards, err := deck.Deal(5)
if err != nil {
    // Handle error: not enough cards
}
fmt.Println(len(deck.Cards)) // Output: 47
```

### Hand Evaluation

#### `EvaluateHand`

Evaluates a 5-card poker hand and returns the best hand category with tiebreakers.

```go
func EvaluateHand(cards []Card) *Hand
```

**Parameters:**
- `cards` - Exactly 5 cards to evaluate

**Returns:**
- `*Hand` - Evaluated hand with category and tiebreakers, or `nil` if not exactly 5 cards

**Evaluation Order:**
Checks categories from strongest to weakest: Royal Flush → Straight Flush → Four of a Kind → Full House → Flush → Straight → Three of a Kind → Two Pair → One Pair → High Card

**Tiebreaker Format by Category:**

| Category        | Tiebreakers                                  |
|-----------------|----------------------------------------------|
| Royal Flush     | `[]` (empty - all equal)                     |
| Straight Flush  | `[high card]`                                |
| Four of a Kind  | `[quad rank, kicker]`                        |
| Full House      | `[trip rank, pair rank]`                     |
| Flush           | `[card1, card2, card3, card4, card5]` (desc) |
| Straight        | `[high card]`                                |
| Three of a Kind | `[trip rank, kicker1, kicker2]` (desc)       |
| Two Pair        | `[high pair, low pair, kicker]`              |
| One Pair        | `[pair rank, k1, k2, k3]` (desc)             |
| High Card       | `[card1, card2, card3, card4, card5]` (desc) |

**Example:**
```go
cards, _ := parseCards([]string{"Ah", "Ad", "Kh", "Kd", "Qh"})
hand := poker.EvaluateHand(cards)

fmt.Println(hand.Category)     // Two Pair
fmt.Println(hand.Tiebreakers)  // [14 13 12] (Aces, Kings, Queen kicker)
```

**Special Cases:**
- **Wheel Straight** (A-2-3-4-5): Returns high card of `Five` (5), not `Ace` (14)
- **Royal Flush**: Must be exactly 10-J-Q-K-A suited

### Hand Comparison

#### `CompareHands`

Compares two poker hands and determines the winner.

```go
func CompareHands(hand1, hand2 *Hand) int
```

**Parameters:**
- `hand1` - First hand to compare
- `hand2` - Second hand to compare

**Returns:**
- `1` - `hand1` is stronger
- `-1` - `hand2` is stronger
- `0` - Hands are equal (split pot)

**Comparison Logic:**
1. Compare by category first (higher category wins)
2. If categories equal, compare tiebreakers element-by-element
3. First differing tiebreaker determines winner

**Example:**
```go
// Aces over Kings vs Kings over Queens
cards1, _ := parseCards([]string{"Ah", "Ad", "Kh", "Kd", "Qh"})
cards2, _ := parseCards([]string{"Ks", "Kc", "Qs", "Qc", "Jh"})

hand1 := poker.EvaluateHand(cards1)
hand2 := poker.EvaluateHand(cards2)

result := poker.CompareHands(hand1, hand2)
fmt.Println(result) // Output: 1 (hand1 wins - higher two pair)
```

#### `FindBestHand`

Finds the best 5-card poker hand from 5, 6, or 7 cards.

```go
func FindBestHand(cards []Card) *Hand
```

**Parameters:**
- `cards` - 5, 6, or 7 cards to evaluate

**Returns:**
- `*Hand` - Best possible 5-card hand, or `nil` if fewer than 5 cards

**Algorithm:**
- **5 cards**: Evaluates directly (optimization)
- **6 cards**: Generates 6 combinations (C(6,5) = 6), evaluates each
- **7 cards**: Generates 21 combinations (C(7,5) = 21), evaluates each

**Performance:**
- 7 cards: ~210 operations (21 combinations × ~10 checks each)
- Fast enough for typical use cases
- Future optimization: lookup tables if needed

**Example:**
```go
// Texas Hold'em: 2 hole cards + 5 community cards
allCards, _ := parseCards([]string{
    "Ah", "Kh",           // Hole cards
    "Qh", "Jh", "Th",     // Flop
    "9d", "2c",           // Turn, River
})

bestHand := poker.FindBestHand(allCards)
fmt.Println(bestHand.Category) // Royal Flush
```

#### `Combinations`

Generates all k-card combinations from the given cards.

```go
func Combinations(cards []Card, k int) [][]Card
```

**Parameters:**
- `cards` - Source cards to choose from
- `k` - Number of cards per combination

**Returns:**
- `[][]Card` - All possible k-card combinations

**Combination Counts:**
- C(7,5) = 21 combinations
- C(6,5) = 6 combinations
- C(5,5) = 1 combination

**Example:**
```go
cards, _ := parseCards([]string{"Ah", "Kh", "Qh", "Jh", "Th", "9d", "2c"})
combos := poker.Combinations(cards, 5)

fmt.Println(len(combos)) // Output: 21
```

## Architecture

The codebase follows a clean layered architecture:

### Phase 1: Foundation Layer
- **Files**: `card.go`, `deck.go`
- **Purpose**: Card representation, deck operations
- **Key Concepts**:
  - Rank values 2-14 (Ace=14) enable direct comparison
  - Card notation uses single characters (e.g., "Ah", "Td")
  - Standard 52-card deck with Fisher-Yates shuffle

### Phase 2: Evaluation Core
- **Files**: `hand.go`, `evaluator.go`
- **Purpose**: Hand categories, helper functions
- **Key Components**:
  - `HandCategory` enum ordered 1-10 for comparison
  - `isFlush()`, `isStraight()`, `rankCounts()` helpers
  - Wheel straight handling (A-2-3-4-5 returns high rank 5)

### Phase 3: Detection Layer
- **Files**: `evaluator.go`
- **Purpose**: Ten detection functions for each hand category
- **Functions**: `detectRoyalFlush()`, `detectStraightFlush()`, `detectFourOfAKind()`, etc.
- **Tiebreaker System**: Each detector returns `[]Rank` in descending importance

### Phase 4: Integration Layer
- **Files**: `evaluator.go`, `combinations.go`
- **Purpose**: High-level evaluation and comparison
- **Functions**:
  - `EvaluateHand()` - Checks categories strongest to weakest
  - `Combinations()` - Generates all k-card combinations
  - `FindBestHand()` - Evaluates all combinations, returns strongest
  - `CompareHands()` - Category-first comparison with tiebreakers

### Phase 5: Examples
- **Files**: `cmd/example/main.go`
- **Purpose**: Demonstrates complete workflow
- **Shows**: Deck creation, shuffling, dealing, evaluation, comparison

## Hand Categories

### Royal Flush (10)
- **Definition**: 10-J-Q-K-A all of the same suit
- **Tiebreakers**: None (all royal flushes are equal)
- **Example**: `Ah Kh Qh Jh Th`

### Straight Flush (9)
- **Definition**: 5 sequential cards, all same suit (excluding royal flush)
- **Tiebreakers**: `[high card]`
- **Example**: `9h 8h 7h 6h 5h` (high card: 9)
- **Special Case**: Wheel straight flush `5h 4h 3h 2h Ah` (high card: 5)

### Four of a Kind (8)
- **Definition**: 4 cards of the same rank
- **Tiebreakers**: `[quad rank, kicker]`
- **Example**: `Ah Ad Ac As Kh` (quad Aces, King kicker)

### Full House (7)
- **Definition**: 3 cards of one rank + 2 cards of another rank
- **Tiebreakers**: `[trip rank, pair rank]`
- **Example**: `Ah Ad Ac Kh Kd` (Aces over Kings)

### Flush (6)
- **Definition**: 5 cards of the same suit, not sequential
- **Tiebreakers**: `[card1, card2, card3, card4, card5]` (descending)
- **Example**: `Ah Kh Qh Jh 9h`

### Straight (5)
- **Definition**: 5 sequential cards with mixed suits
- **Tiebreakers**: `[high card]`
- **Example**: `9h 8d 7c 6s 5h` (high card: 9)
- **Special Case**: Wheel straight `5h 4d 3c 2s Ah` (high card: 5)

### Three of a Kind (4)
- **Definition**: 3 cards of the same rank (no pair)
- **Tiebreakers**: `[trip rank, kicker1, kicker2]` (descending)
- **Example**: `Ah Ad Ac Kh Qd` (trip Aces, K-Q kickers)

### Two Pair (3)
- **Definition**: 2 cards of one rank + 2 cards of another rank
- **Tiebreakers**: `[high pair, low pair, kicker]`
- **Example**: `Ah Ad Kh Kd Qc` (Aces and Kings, Queen kicker)

### One Pair (2)
- **Definition**: 2 cards of the same rank
- **Tiebreakers**: `[pair rank, kicker1, kicker2, kicker3]` (descending)
- **Example**: `Ah Ad Kh Qd Jc` (pair of Aces, K-Q-J kickers)

### High Card (1)
- **Definition**: No matching cards
- **Tiebreakers**: `[card1, card2, card3, card4, card5]` (descending)
- **Example**: `Ah Kd Qh Jc 9s` (Ace high)

## Development

### Project Structure

```
poker-hand-evaluation/
├── cmd/
│   └── example/
│       ├── main.go         # Example usage
│       └── main_test.go    # Example tests
├── pkg/
│   └── poker/
│       ├── card.go         # Card, Rank, Suit types
│       ├── card_test.go    # Card tests
│       ├── deck.go         # Deck operations
│       ├── deck_test.go    # Deck tests
│       ├── hand.go         # HandCategory, Hand struct
│       ├── hand_test.go    # Hand tests
│       ├── evaluator.go    # Detection functions, evaluation
│       ├── evaluator_test.go  # Evaluator tests
│       ├── combinations.go    # Combination generator
│       └── combinations_test.go  # Combination tests
├── go.mod                  # Go module definition
├── CLAUDE.md              # Project-specific instructions
└── README.md              # This file
```

### Test-Driven Development

All features follow strict Red-Green-Refactor methodology:

1. **RED**: Write failing tests first in `*_test.go`
2. **GREEN**: Write minimal code to pass tests
3. **REFACTOR**: Clean up while keeping tests green

### Code Quality Standards

- **Target Coverage**: ≥90% (currently at 92.7%)
- **Linting**: Code passes `go vet` and `gofmt`
- **Testing**: Comprehensive edge case coverage
- **Documentation**: All public functions documented with examples

## Testing

### Run All Tests

```bash
# Verbose output
go test -v ./...

# With coverage
go test -cover ./...

# Coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Run Specific Tests

```bash
# Test a specific function
go test -v -run TestCardString ./pkg/poker

# Test a specific file
go test -v ./pkg/poker/card_test.go ./pkg/poker/card.go

# Test with race detection
go test -race ./...
```

### Current Coverage

- **pkg/poker**: 92.7% coverage
- **cmd/example**: 85.2% coverage

### Test Edge Cases Covered

- ✅ Wheel straight (A-2-3-4-5)
- ✅ Wheel straight flush (Ah-2h-3h-4h-5h)
- ✅ Royal flush detection (exact 10-J-Q-K-A)
- ✅ Multiple flush suits in 7 cards
- ✅ Tiebreaker scenarios (same category, different kickers)
- ✅ Split pots (identical hands)
- ✅ Card parsing (both "T" and "10" for Ten)
- ✅ Error handling (insufficient cards, invalid input)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please ensure:

1. All tests pass (`go test ./...`)
2. Code is formatted (`gofmt -w .`)
3. No vet warnings (`go vet ./...`)
4. Test coverage remains ≥90%
5. Follow conventional commit format

## Author

**Puupa** - [GitHub Profile](https://github.com/puupa)

## Acknowledgments

- Built with strict Test-Driven Development methodology
- Follows Go best practices and SOLID principles
- Inspired by Texas Hold'em poker rules and hand evaluation algorithms

---

**Made with ❤️ and Go**
