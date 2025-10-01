package poker

import (
	"testing"
)

func TestCombinations5From5(t *testing.T) {
	// Test edge case: selecting 5 from exactly 5 cards
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Ten, Suit: Hearts},
	}

	result := Combinations(cards, 5)

	// Should return exactly 1 combination (the original hand)
	if len(result) != 1 {
		t.Errorf("Combinations(5, 5) = %d combinations, want 1", len(result))
	}

	// Verify it's the same 5 cards
	if len(result[0]) != 5 {
		t.Errorf("Combination has %d cards, want 5", len(result[0]))
	}
}

func TestCombinations5From6(t *testing.T) {
	// Create 6 distinct cards
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Ten, Suit: Hearts},
		{Rank: Nine, Suit: Hearts},
	}

	result := Combinations(cards, 5)

	// C(6,5) = 6
	if len(result) != 6 {
		t.Errorf("Combinations(6, 5) = %d combinations, want 6", len(result))
	}

	// Verify each combination has exactly 5 cards
	for i, combo := range result {
		if len(combo) != 5 {
			t.Errorf("Combination %d has %d cards, want 5", i, len(combo))
		}
	}
}

func TestCombinations5From7(t *testing.T) {
	// Create 7 distinct cards
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Ten, Suit: Hearts},
		{Rank: Nine, Suit: Hearts},
		{Rank: Eight, Suit: Hearts},
	}

	result := Combinations(cards, 5)

	// C(7,5) = 21
	if len(result) != 21 {
		t.Errorf("Combinations(7, 5) = %d combinations, want 21", len(result))
	}

	// Verify each combination has exactly 5 cards
	for i, combo := range result {
		if len(combo) != 5 {
			t.Errorf("Combination %d has %d cards, want 5", i, len(combo))
		}
	}
}

func TestCombinationsUniqueness(t *testing.T) {
	// Create 6 cards to test uniqueness
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Jack, Suit: Spades},
		{Rank: Ten, Suit: Hearts},
		{Rank: Nine, Suit: Diamonds},
	}

	result := Combinations(cards, 5)

	// Check that all combinations are unique by comparing each pair
	seen := make(map[string]bool)
	for _, combo := range result {
		// Create a unique signature for this combination
		signature := ""
		for _, card := range combo {
			signature += card.String()
		}

		if seen[signature] {
			t.Errorf("Found duplicate combination: %v", combo)
		}
		seen[signature] = true
	}

	// Should have 6 unique combinations
	if len(seen) != 6 {
		t.Errorf("Found %d unique combinations, want 6", len(seen))
	}
}

func TestCombinationsAllCardsValid(t *testing.T) {
	// Test that all cards in combinations come from the original set
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Ten, Suit: Hearts},
		{Rank: Nine, Suit: Hearts},
		{Rank: Eight, Suit: Hearts},
	}

	result := Combinations(cards, 5)

	// Create a map of valid cards
	validCards := make(map[string]bool)
	for _, card := range cards {
		validCards[card.String()] = true
	}

	// Check each combination
	for i, combo := range result {
		for j, card := range combo {
			if !validCards[card.String()] {
				t.Errorf("Combination %d contains invalid card at position %d: %v", i, j, card)
			}
		}
	}
}

func TestCombinationsEdgeCase3From4(t *testing.T) {
	// Test with different k value (not just 5)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
	}

	result := Combinations(cards, 3)

	// C(4,3) = 4
	if len(result) != 4 {
		t.Errorf("Combinations(4, 3) = %d combinations, want 4", len(result))
	}

	// Verify each combination has exactly 3 cards
	for i, combo := range result {
		if len(combo) != 3 {
			t.Errorf("Combination %d has %d cards, want 3", i, len(combo))
		}
	}
}
