package main

import "testing"

// Test that NewDeck creates a deck with exactly 52 cards
func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.Cards) != 52 {
		t.Errorf("NewDeck should create 52 cards, got %d", len(deck.Cards))
	}
}

// Test that deck contains all 13 ranks
func TestDeckContainsAllRanks(t *testing.T) {
	deck := NewDeck()
	rankCounts := make(map[Rank]int)

	for _, card := range deck.Cards {
		rankCounts[card.Rank]++
	}

	expectedRanks := []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	for _, rank := range expectedRanks {
		if rankCounts[rank] != 4 {
			t.Errorf("Rank %v should appear 4 times (once per suit), got %d", rank, rankCounts[rank])
		}
	}
}

// Test that deck contains all 4 suits
func TestDeckContainsAllSuits(t *testing.T) {
	deck := NewDeck()
	suitCounts := make(map[Suit]int)

	for _, card := range deck.Cards {
		suitCounts[card.Suit]++
	}

	expectedSuits := []Suit{Hearts, Diamonds, Clubs, Spades}
	for _, suit := range expectedSuits {
		if suitCounts[suit] != 13 {
			t.Errorf("Suit %v should appear 13 times (once per rank), got %d", suit, suitCounts[suit])
		}
	}
}

// Test that deck has no duplicate cards
func TestDeckNoDuplicates(t *testing.T) {
	deck := NewDeck()
	seen := make(map[Card]bool)

	for _, card := range deck.Cards {
		if seen[card] {
			t.Errorf("Duplicate card found: %v", card)
		}
		seen[card] = true
	}
}
