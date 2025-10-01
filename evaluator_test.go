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

// Test rankCounts returns map with counts for all unique ranks
func TestRankCountsAllUnique(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Jack, Suit: Spades},
		{Rank: Ten, Suit: Hearts},
	}

	counts := rankCounts(cards)

	// Should have exactly 5 entries (all unique)
	if len(counts) != 5 {
		t.Errorf("Expected 5 entries, got %d", len(counts))
	}

	// Each rank should appear exactly once
	expectedCounts := map[Rank]int{
		Ace:   1,
		King:  1,
		Queen: 1,
		Jack:  1,
		Ten:   1,
	}

	for rank, expectedCount := range expectedCounts {
		if counts[rank] != expectedCount {
			t.Errorf("Rank %v: expected count %d, got %d", rank, expectedCount, counts[rank])
		}
	}
}

// Test rankCounts correctly identifies a pair (count=2)
func TestRankCountsPair(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: Ace, Suit: Diamonds}, // Pair of Aces
		{Rank: King, Suit: Clubs},
		{Rank: Queen, Suit: Spades},
		{Rank: Jack, Suit: Hearts},
	}

	counts := rankCounts(cards)

	// Should have 4 entries (one pair, three singletons)
	if len(counts) != 4 {
		t.Errorf("Expected 4 entries, got %d", len(counts))
	}

	// Ace should have count of 2
	if counts[Ace] != 2 {
		t.Errorf("Expected Ace count 2, got %d", counts[Ace])
	}

	// Other ranks should have count of 1
	if counts[King] != 1 || counts[Queen] != 1 || counts[Jack] != 1 {
		t.Error("Expected other ranks to have count 1")
	}
}

// Test rankCounts correctly identifies two pairs
func TestRankCountsTwoPairs(t *testing.T) {
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: Ace, Suit: Diamonds}, // Pair of Aces
		{Rank: King, Suit: Clubs},
		{Rank: King, Suit: Spades}, // Pair of Kings
		{Rank: Queen, Suit: Hearts},
	}

	counts := rankCounts(cards)

	// Should have 3 entries (two pairs, one singleton)
	if len(counts) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(counts))
	}

	// Aces and Kings should have count of 2
	if counts[Ace] != 2 {
		t.Errorf("Expected Ace count 2, got %d", counts[Ace])
	}
	if counts[King] != 2 {
		t.Errorf("Expected King count 2, got %d", counts[King])
	}

	// Queen should have count of 1
	if counts[Queen] != 1 {
		t.Errorf("Expected Queen count 1, got %d", counts[Queen])
	}
}

// Test rankCounts correctly identifies trips (count=3)
func TestRankCountsTrips(t *testing.T) {
	cards := []Card{
		{Rank: Seven, Suit: Hearts},
		{Rank: Seven, Suit: Diamonds}, // Trip Sevens
		{Rank: Seven, Suit: Clubs},
		{Rank: King, Suit: Spades},
		{Rank: Queen, Suit: Hearts},
	}

	counts := rankCounts(cards)

	// Should have 3 entries (one trip, two singletons)
	if len(counts) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(counts))
	}

	// Seven should have count of 3
	if counts[Seven] != 3 {
		t.Errorf("Expected Seven count 3, got %d", counts[Seven])
	}

	// Other ranks should have count of 1
	if counts[King] != 1 || counts[Queen] != 1 {
		t.Error("Expected other ranks to have count 1")
	}
}

// Test rankCounts correctly identifies full house (trips + pair)
func TestRankCountsFullHouse(t *testing.T) {
	cards := []Card{
		{Rank: Seven, Suit: Hearts},
		{Rank: Seven, Suit: Diamonds}, // Trip Sevens
		{Rank: Seven, Suit: Clubs},
		{Rank: King, Suit: Spades}, // Pair of Kings
		{Rank: King, Suit: Hearts},
	}

	counts := rankCounts(cards)

	// Should have 2 entries (one trip, one pair)
	if len(counts) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(counts))
	}

	// Seven should have count of 3
	if counts[Seven] != 3 {
		t.Errorf("Expected Seven count 3, got %d", counts[Seven])
	}

	// King should have count of 2
	if counts[King] != 2 {
		t.Errorf("Expected King count 2, got %d", counts[King])
	}
}

// Test rankCounts correctly identifies quads (count=4)
func TestRankCountsQuads(t *testing.T) {
	cards := []Card{
		{Rank: Nine, Suit: Hearts},
		{Rank: Nine, Suit: Diamonds}, // Quad Nines
		{Rank: Nine, Suit: Clubs},
		{Rank: Nine, Suit: Spades},
		{Rank: Ace, Suit: Hearts}, // Kicker
	}

	counts := rankCounts(cards)

	// Should have 2 entries (one quad, one singleton)
	if len(counts) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(counts))
	}

	// Nine should have count of 4
	if counts[Nine] != 4 {
		t.Errorf("Expected Nine count 4, got %d", counts[Nine])
	}

	// Ace should have count of 1
	if counts[Ace] != 1 {
		t.Errorf("Expected Ace count 1, got %d", counts[Ace])
	}
}
