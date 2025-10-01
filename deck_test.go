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

// Test that Deal returns the correct number of cards
func TestDealReturnsCorrectNumberOfCards(t *testing.T) {
	deck := NewDeck()
	cards, err := deck.Deal(5)
	if err != nil {
		t.Errorf("Deal(5) should not return error, got: %v", err)
	}
	if len(cards) != 5 {
		t.Errorf("Deal(5) should return 5 cards, got %d", len(cards))
	}
}

// Test that Deal removes cards from the deck
func TestDealRemovesCardsFromDeck(t *testing.T) {
	deck := NewDeck()
	initialCount := len(deck.Cards)
	_, err := deck.Deal(5)
	if err != nil {
		t.Errorf("Deal(5) should not return error, got: %v", err)
	}
	if len(deck.Cards) != initialCount-5 {
		t.Errorf("After Deal(5), deck should have %d cards, got %d", initialCount-5, len(deck.Cards))
	}
}

// Test that dealing more cards than available returns error
func TestDealMoreThanAvailableReturnsError(t *testing.T) {
	deck := NewDeck()
	_, err := deck.Deal(53)
	if err == nil {
		t.Error("Deal(53) should return error when deck has only 52 cards")
	}
}

// Test that Deal(0) returns empty slice without error
func TestDealZeroCards(t *testing.T) {
	deck := NewDeck()
	cards, err := deck.Deal(0)
	if err != nil {
		t.Errorf("Deal(0) should not return error, got: %v", err)
	}
	if len(cards) != 0 {
		t.Errorf("Deal(0) should return empty slice, got %d cards", len(cards))
	}
}

// Test that Deal from 52-card deck leaves 47 cards
func TestDealFiveFromFullDeckLeavesFortySevenCards(t *testing.T) {
	deck := NewDeck()
	_, err := deck.Deal(5)
	if err != nil {
		t.Errorf("Deal(5) should not return error, got: %v", err)
	}
	if len(deck.Cards) != 47 {
		t.Errorf("After Deal(5) from 52-card deck, should have 47 cards remaining, got %d", len(deck.Cards))
	}
}
