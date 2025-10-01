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

// Test that Shuffle changes the order of cards
func TestShuffleChangesOrder(t *testing.T) {
	deck := NewDeck()
	original := make([]Card, len(deck.Cards))
	copy(original, deck.Cards)

	deck.Shuffle()

	// Check that at least some cards changed position
	sameCount := 0
	for i := range deck.Cards {
		if deck.Cards[i] == original[i] {
			sameCount++
		}
	}

	// It's extremely unlikely all 52 cards remain in same position after shuffle
	if sameCount == 52 {
		t.Error("Shuffle did not change card order")
	}
}

// Test that Shuffle preserves all cards (no cards lost or added)
func TestShufflePreservesAllCards(t *testing.T) {
	deck := NewDeck()

	// Count cards before shuffle
	beforeCounts := make(map[Card]int)
	for _, card := range deck.Cards {
		beforeCounts[card]++
	}

	deck.Shuffle()

	// Count cards after shuffle
	afterCounts := make(map[Card]int)
	for _, card := range deck.Cards {
		afterCounts[card]++
	}

	// Verify same cards exist
	if len(deck.Cards) != 52 {
		t.Errorf("Shuffle changed deck size: expected 52, got %d", len(deck.Cards))
	}

	for card, count := range beforeCounts {
		if afterCounts[card] != count {
			t.Errorf("Card %v count changed: before=%d, after=%d", card, count, afterCounts[card])
		}
	}
}

// Test that multiple shuffles produce different orders
func TestShuffleProducesDifferentOrders(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()

	deck1.Shuffle()
	deck2.Shuffle()

	// Count how many positions are identical
	samePositions := 0
	for i := range deck1.Cards {
		if deck1.Cards[i] == deck2.Cards[i] {
			samePositions++
		}
	}

	// It's statistically extremely unlikely that two shuffles produce identical results
	// We expect at most a few cards to be in the same position by chance
	if samePositions > 10 {
		t.Errorf("Two shuffles too similar: %d/52 cards in same positions", samePositions)
	}
}
