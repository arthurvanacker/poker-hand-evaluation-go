package main

import "testing"

// Test isFlush returns true when all 5 cards are the same suit
func TestIsFlushAllHearts(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Ten, Suit: Hearts},
	}

	if !isFlush(cards) {
		t.Error("Expected flush with all hearts, got false")
	}
}

// Test isFlush returns true for all suits
func TestIsFlushAllSuits(t *testing.T) {
	tests := []struct {
		name string
		suit Suit
	}{
		{"all hearts", Hearts},
		{"all diamonds", Diamonds},
		{"all clubs", Clubs},
		{"all spades", Spades},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards := []Card{
				{Rank: Ace, Suit: tt.suit},
				{Rank: King, Suit: tt.suit},
				{Rank: Queen, Suit: tt.suit},
				{Rank: Jack, Suit: tt.suit},
				{Rank: Ten, Suit: tt.suit},
			}

			if !isFlush(cards) {
				t.Errorf("Expected flush with %s, got false", tt.name)
			}
		})
	}
}

// Test isFlush returns false for mixed suits
func TestIsFlushMixedSuits(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Ten, Suit: Hearts},
	}

	if isFlush(cards) {
		t.Error("Expected false for mixed suits, got true")
	}
}

// Test isFlush with 4 matching and 1 different
func TestIsFlushFourMatching(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected bool
	}{
		{
			"first card different",
			[]Card{
				{Rank: Ace, Suit: Diamonds},
				{Rank: King, Suit: Hearts},
				{Rank: Queen, Suit: Hearts},
				{Rank: Jack, Suit: Hearts},
				{Rank: Ten, Suit: Hearts},
			},
			false,
		},
		{
			"last card different",
			[]Card{
				{Rank: Ace, Suit: Hearts},
				{Rank: King, Suit: Hearts},
				{Rank: Queen, Suit: Hearts},
				{Rank: Jack, Suit: Hearts},
				{Rank: Ten, Suit: Clubs},
			},
			false,
		},
		{
			"middle card different",
			[]Card{
				{Rank: Ace, Suit: Spades},
				{Rank: King, Suit: Spades},
				{Rank: Queen, Suit: Hearts},
				{Rank: Jack, Suit: Spades},
				{Rank: Ten, Suit: Spades},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isFlush(tt.cards)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Test isFlush with all different suits
func TestIsFlushAllDifferent(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Jack, Suit: Spades},
		{Rank: Ten, Suit: Hearts},
	}

	if isFlush(cards) {
		t.Error("Expected false for all different suits, got true")
	}
}

// Test isStraight returns true for low straight (2-3-4-5-6)
func TestIsStraightLow(t *testing.T) {
	cards := []Card{
		{Rank: Two, Suit: Hearts},
		{Rank: Three, Suit: Diamonds},
		{Rank: Four, Suit: Clubs},
		{Rank: Five, Suit: Spades},
		{Rank: Six, Suit: Hearts},
	}

	isStraight, highCard := isStraight(cards)
	if !isStraight {
		t.Error("Expected straight for 2-3-4-5-6, got false")
	}
	if highCard != Six {
		t.Errorf("Expected high card Six (6), got %v (%d)", highCard, highCard)
	}
}

// Test isStraight returns true for broadway straight (10-J-Q-K-A)
func TestIsStraightBroadway(t *testing.T) {
	cards := []Card{
		{Rank: Ten, Suit: Hearts},
		{Rank: Jack, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: King, Suit: Spades},
		{Rank: Ace, Suit: Hearts},
	}

	isStraight, highCard := isStraight(cards)
	if !isStraight {
		t.Error("Expected straight for 10-J-Q-K-A, got false")
	}
	if highCard != Ace {
		t.Errorf("Expected high card Ace (14), got %v (%d)", highCard, highCard)
	}
}

// Test isStraight returns true for wheel straight (A-2-3-4-5)
// Critical: Ace acts as low (value 1), returns high card of Five
func TestIsStraightWheel(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: Two, Suit: Diamonds},
		{Rank: Three, Suit: Clubs},
		{Rank: Four, Suit: Spades},
		{Rank: Five, Suit: Hearts},
	}

	isStraight, highCard := isStraight(cards)
	if !isStraight {
		t.Error("Expected straight for A-2-3-4-5 wheel, got false")
	}
	if highCard != Five {
		t.Errorf("Expected high card Five (5) for wheel, got %v (%d)", highCard, highCard)
	}
}

// Test isStraight returns true for middle straight (7-8-9-10-J)
func TestIsStraightMiddle(t *testing.T) {
	cards := []Card{
		{Rank: Seven, Suit: Hearts},
		{Rank: Eight, Suit: Diamonds},
		{Rank: Nine, Suit: Clubs},
		{Rank: Ten, Suit: Spades},
		{Rank: Jack, Suit: Hearts},
	}

	isStraight, highCard := isStraight(cards)
	if !isStraight {
		t.Error("Expected straight for 7-8-9-10-J, got false")
	}
	if highCard != Jack {
		t.Errorf("Expected high card Jack (11), got %v (%d)", highCard, highCard)
	}
}

// Test isStraight returns false for cards with gap
func TestIsStraightWithGap(t *testing.T) {
	cards := []Card{
		{Rank: Two, Suit: Hearts},
		{Rank: Three, Suit: Diamonds},
		{Rank: Four, Suit: Clubs},
		{Rank: Six, Suit: Spades}, // Gap: missing Five
		{Rank: Seven, Suit: Hearts},
	}

	isStraight, _ := isStraight(cards)
	if isStraight {
		t.Error("Expected false for cards with gap (2-3-4-6-7), got true")
	}
}

// Test isStraight returns false for non-sequential cards
func TestIsStraightNonSequential(t *testing.T) {
	cards := []Card{
		{Rank: Two, Suit: Hearts},
		{Rank: Five, Suit: Diamonds},
		{Rank: Eight, Suit: Clubs},
		{Rank: Jack, Suit: Spades},
		{Rank: King, Suit: Hearts},
	}

	isStraight, _ := isStraight(cards)
	if isStraight {
		t.Error("Expected false for non-sequential cards, got true")
	}
}

// Test isStraight returns false for pair in sequence
func TestIsStraightWithPair(t *testing.T) {
	cards := []Card{
		{Rank: Five, Suit: Hearts},
		{Rank: Five, Suit: Diamonds}, // Duplicate
		{Rank: Six, Suit: Clubs},
		{Rank: Seven, Suit: Spades},
		{Rank: Eight, Suit: Hearts},
	}

	isStraight, _ := isStraight(cards)
	if isStraight {
		t.Error("Expected false for cards with duplicate rank, got true")
	}
}

// Test isStraight handles unordered input (9-7-10-8-J)
func TestIsStraightUnordered(t *testing.T) {
	cards := []Card{
		{Rank: Nine, Suit: Hearts},
		{Rank: Seven, Suit: Diamonds},
		{Rank: Ten, Suit: Clubs},
		{Rank: Eight, Suit: Spades},
		{Rank: Jack, Suit: Hearts},
	}

	isStraight, highCard := isStraight(cards)
	if !isStraight {
		t.Error("Expected straight for unordered 7-8-9-10-J, got false")
	}
	if highCard != Jack {
		t.Errorf("Expected high card Jack (11), got %v (%d)", highCard, highCard)
	}
}

// Test isStraight returns false for almost-wheel (A-2-3-4-6)
func TestIsStraightAlmostWheel(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: Two, Suit: Diamonds},
		{Rank: Three, Suit: Clubs},
		{Rank: Four, Suit: Spades},
		{Rank: Six, Suit: Hearts}, // Missing Five
	}

	isStraight, _ := isStraight(cards)
	if isStraight {
		t.Error("Expected false for A-2-3-4-6 (missing 5), got true")
	}
}
