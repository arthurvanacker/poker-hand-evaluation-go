package poker

import "testing"

// Test that HandCategory enum has all 10 values
func TestHandCategoryValues(t *testing.T) {
	tests := []struct {
		name     string
		category HandCategory
		expected int
	}{
		{"HighCard", HighCard, 1},
		{"OnePair", OnePair, 2},
		{"TwoPair", TwoPair, 3},
		{"ThreeOfAKind", ThreeOfAKind, 4},
		{"Straight", Straight, 5},
		{"Flush", Flush, 6},
		{"FullHouse", FullHouse, 7},
		{"FourOfAKind", FourOfAKind, 8},
		{"StraightFlush", StraightFlush, 9},
		{"RoyalFlush", RoyalFlush, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.category) != tt.expected {
				t.Errorf("%s should have value %d, got %d", tt.name, tt.expected, int(tt.category))
			}
		})
	}
}

// Test that categories are ordered by strength
func TestHandCategoryOrdering(t *testing.T) {
	if HighCard >= OnePair {
		t.Error("HighCard should be less than OnePair")
	}
	if OnePair >= TwoPair {
		t.Error("OnePair should be less than TwoPair")
	}
	if TwoPair >= ThreeOfAKind {
		t.Error("TwoPair should be less than ThreeOfAKind")
	}
	if ThreeOfAKind >= Straight {
		t.Error("ThreeOfAKind should be less than Straight")
	}
	if Straight >= Flush {
		t.Error("Straight should be less than Flush")
	}
	if Flush >= FullHouse {
		t.Error("Flush should be less than FullHouse")
	}
	if FullHouse >= FourOfAKind {
		t.Error("FullHouse should be less than FourOfAKind")
	}
	if FourOfAKind >= StraightFlush {
		t.Error("FourOfAKind should be less than StraightFlush")
	}
	if StraightFlush >= RoyalFlush {
		t.Error("StraightFlush should be less than RoyalFlush")
	}
}

// Test that RoyalFlush is the strongest hand
func TestRoyalFlushIsStrongest(t *testing.T) {
	categories := []HandCategory{
		HighCard, OnePair, TwoPair, ThreeOfAKind, Straight,
		Flush, FullHouse, FourOfAKind, StraightFlush,
	}

	for _, category := range categories {
		if category >= RoyalFlush {
			t.Errorf("%v should be weaker than RoyalFlush", category)
		}
	}
}

// Test that HighCard is the weakest hand
func TestHighCardIsWeakest(t *testing.T) {
	categories := []HandCategory{
		OnePair, TwoPair, ThreeOfAKind, Straight, Flush,
		FullHouse, FourOfAKind, StraightFlush, RoyalFlush,
	}

	for _, category := range categories {
		if category <= HighCard {
			t.Errorf("%v should be stronger than HighCard", category)
		}
	}
}

// Test HandCategory String() method
func TestHandCategoryString(t *testing.T) {
	tests := []struct {
		category HandCategory
		expected string
	}{
		{HighCard, "High Card"},
		{OnePair, "One Pair"},
		{TwoPair, "Two Pair"},
		{ThreeOfAKind, "Three of a Kind"},
		{Straight, "Straight"},
		{Flush, "Flush"},
		{FullHouse, "Full House"},
		{FourOfAKind, "Four of a Kind"},
		{StraightFlush, "Straight Flush"},
		{RoyalFlush, "Royal Flush"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if tt.category.String() != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.category.String())
			}
		})
	}
}

// Test Hand struct has correct fields
func TestHandStructFields(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Ten, Suit: Hearts},
	}
	tiebreakers := []Rank{Ace, King, Queen, Jack, Ten}

	hand := Hand{
		Cards:       cards,
		Category:    RoyalFlush,
		Tiebreakers: tiebreakers,
	}

	if len(hand.Cards) != 5 {
		t.Errorf("Expected 5 cards, got %d", len(hand.Cards))
	}
	if hand.Category != RoyalFlush {
		t.Errorf("Expected RoyalFlush, got %v", hand.Category)
	}
	if len(hand.Tiebreakers) != 5 {
		t.Errorf("Expected 5 tiebreakers, got %d", len(hand.Tiebreakers))
	}
}

// Test NewHand with exactly 5 cards succeeds
func TestNewHandWithFiveCards(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Ten, Suit: Hearts},
	}

	hand, err := NewHand(cards)
	if err != nil {
		t.Errorf("Expected no error with 5 cards, got: %v", err)
	}
	if hand == nil {
		t.Fatal("Expected hand to be non-nil")
	}
	if len(hand.Cards) != 5 {
		t.Errorf("Expected 5 cards, got %d", len(hand.Cards))
	}
}

// Test NewHand rejects fewer than 5 cards
func TestNewHandRejectsTooFewCards(t *testing.T) {
	tests := []struct {
		name     string
		numCards int
	}{
		{"zero cards", 0},
		{"one card", 1},
		{"two cards", 2},
		{"three cards", 3},
		{"four cards", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards := make([]Card, tt.numCards)
			for i := 0; i < tt.numCards; i++ {
				cards[i] = Card{Rank: Ace, Suit: Hearts}
			}

			hand, err := NewHand(cards)
			if err == nil {
				t.Errorf("Expected error with %d cards, got nil", tt.numCards)
			}
			if hand != nil {
				t.Errorf("Expected nil hand with %d cards, got non-nil", tt.numCards)
			}
		})
	}
}

// Test NewHand rejects more than 5 cards
func TestNewHandRejectsTooManyCards(t *testing.T) {
	tests := []struct {
		name     string
		numCards int
	}{
		{"six cards", 6},
		{"seven cards", 7},
		{"ten cards", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards := make([]Card, tt.numCards)
			for i := 0; i < tt.numCards; i++ {
				cards[i] = Card{Rank: Rank(2 + i%13), Suit: Hearts}
			}

			hand, err := NewHand(cards)
			if err == nil {
				t.Errorf("Expected error with %d cards, got nil", tt.numCards)
			}
			if hand != nil {
				t.Errorf("Expected nil hand with %d cards, got non-nil", tt.numCards)
			}
		})
	}
}

// Test NewHand preserves card order
func TestNewHandPreservesCardOrder(t *testing.T) {
	cards := []Card{
		{Rank: Two, Suit: Hearts},
		{Rank: Five, Suit: Diamonds},
		{Rank: Ace, Suit: Clubs},
		{Rank: King, Suit: Spades},
		{Rank: Three, Suit: Hearts},
	}

	hand, err := NewHand(cards)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	for i := 0; i < 5; i++ {
		if hand.Cards[i] != cards[i] {
			t.Errorf("Card at position %d: expected %v, got %v", i, cards[i], hand.Cards[i])
		}
	}
}

// Test HandCategory String() returns "Unknown" for invalid category
func TestHandCategoryStringInvalid(t *testing.T) {
	invalidCategory := HandCategory(99)
	result := invalidCategory.String()
	if result != "Unknown" {
		t.Errorf("Expected 'Unknown' for invalid category, got %q", result)
	}
}
