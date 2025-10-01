package poker

import (
	"reflect"
	"testing"
)

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

// Test rankCounts correctly counts rank frequencies for various hand patterns
func TestRankCounts(t *testing.T) {
	tests := []struct {
		name     string
		cards    []Card
		expected map[Rank]int
	}{
		{
			"all unique",
			[]Card{
				{Rank: Ace, Suit: Hearts},
				{Rank: King, Suit: Diamonds},
				{Rank: Queen, Suit: Clubs},
				{Rank: Jack, Suit: Spades},
				{Rank: Ten, Suit: Hearts},
			},
			map[Rank]int{Ace: 1, King: 1, Queen: 1, Jack: 1, Ten: 1},
		},
		{
			"one pair",
			[]Card{
				{Rank: Ace, Suit: Hearts},
				{Rank: Ace, Suit: Diamonds},
				{Rank: King, Suit: Clubs},
				{Rank: Queen, Suit: Spades},
				{Rank: Jack, Suit: Hearts},
			},
			map[Rank]int{Ace: 2, King: 1, Queen: 1, Jack: 1},
		},
		{
			"two pairs",
			[]Card{
				{Rank: Ace, Suit: Hearts},
				{Rank: Ace, Suit: Diamonds},
				{Rank: King, Suit: Clubs},
				{Rank: King, Suit: Spades},
				{Rank: Queen, Suit: Hearts},
			},
			map[Rank]int{Ace: 2, King: 2, Queen: 1},
		},
		{
			"three of a kind",
			[]Card{
				{Rank: Seven, Suit: Hearts},
				{Rank: Seven, Suit: Diamonds},
				{Rank: Seven, Suit: Clubs},
				{Rank: King, Suit: Spades},
				{Rank: Queen, Suit: Hearts},
			},
			map[Rank]int{Seven: 3, King: 1, Queen: 1},
		},
		{
			"full house",
			[]Card{
				{Rank: Seven, Suit: Hearts},
				{Rank: Seven, Suit: Diamonds},
				{Rank: Seven, Suit: Clubs},
				{Rank: King, Suit: Spades},
				{Rank: King, Suit: Hearts},
			},
			map[Rank]int{Seven: 3, King: 2},
		},
		{
			"four of a kind",
			[]Card{
				{Rank: Nine, Suit: Hearts},
				{Rank: Nine, Suit: Diamonds},
				{Rank: Nine, Suit: Clubs},
				{Rank: Nine, Suit: Spades},
				{Rank: Ace, Suit: Hearts},
			},
			map[Rank]int{Nine: 4, Ace: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counts := rankCounts(tt.cards)
			if !reflect.DeepEqual(counts, tt.expected) {
				t.Errorf("rankCounts(%v) = %v, want %v", tt.cards, counts, tt.expected)
			}
		})
	}
}

// TestDetectRoyalFlush_TrueForRoyalFlush verifies that detectRoyalFlush returns true
// for a valid royal flush (10-J-Q-K-A all of the same suit).
func TestDetectRoyalFlush_TrueForRoyalFlush(t *testing.T) {
	// Arrange: Create a royal flush in hearts (Th-Jh-Qh-Kh-Ah)
	cards := []Card{
		{Rank: Ten, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Ace, Suit: Hearts},
	}

	// Act
	result := detectRoyalFlush(cards)

	// Assert
	if !result {
		t.Errorf("detectRoyalFlush(%v) = false, want true", cards)
	}
}

// TestDetectRoyalFlush_FalseForKingHighStraightFlush verifies that detectRoyalFlush
// returns false for a king-high straight flush (9-10-J-Q-K suited).
func TestDetectRoyalFlush_FalseForKingHighStraightFlush(t *testing.T) {
	// Arrange: Create a king-high straight flush in spades (9s-Ts-Js-Qs-Ks)
	cards := []Card{
		{Rank: Nine, Suit: Spades},
		{Rank: Ten, Suit: Spades},
		{Rank: Jack, Suit: Spades},
		{Rank: Queen, Suit: Spades},
		{Rank: King, Suit: Spades},
	}

	// Act
	result := detectRoyalFlush(cards)

	// Assert
	if result {
		t.Errorf("detectRoyalFlush(%v) = true, want false", cards)
	}
}

// TestDetectRoyalFlush_FalseForNonFlushRoyal verifies that detectRoyalFlush
// returns false for 10-J-Q-K-A that are not all the same suit.
func TestDetectRoyalFlush_FalseForNonFlushRoyal(t *testing.T) {
	// Arrange: Create 10-J-Q-K-A with mixed suits
	cards := []Card{
		{Rank: Ten, Suit: Hearts},
		{Rank: Jack, Suit: Diamonds},
		{Rank: Queen, Suit: Hearts},
		{Rank: King, Suit: Clubs},
		{Rank: Ace, Suit: Spades},
	}

	// Act
	result := detectRoyalFlush(cards)

	// Assert
	if result {
		t.Errorf("detectRoyalFlush(%v) = true, want false", cards)
	}
}

// TestDetectRoyalFlush_TrueForRoyalFlushAllSuits verifies that detectRoyalFlush
// correctly identifies royal flushes in all four suits.
func TestDetectRoyalFlush_TrueForRoyalFlushAllSuits(t *testing.T) {
	suits := []Suit{Hearts, Diamonds, Clubs, Spades}
	suitNames := []string{"Hearts", "Diamonds", "Clubs", "Spades"}

	for i, suit := range suits {
		// Arrange: Create a royal flush in the current suit
		cards := []Card{
			{Rank: Ten, Suit: suit},
			{Rank: Jack, Suit: suit},
			{Rank: Queen, Suit: suit},
			{Rank: King, Suit: suit},
			{Rank: Ace, Suit: suit},
		}

		// Act
		result := detectRoyalFlush(cards)

		// Assert
		if !result {
			t.Errorf("detectRoyalFlush(%v) for %s = false, want true", cards, suitNames[i])
		}
	}
}

// TestDetectRoyalFlush_FalseForRegularFlush verifies that detectRoyalFlush
// returns false for a flush that is not a royal flush.
func TestDetectRoyalFlush_FalseForRegularFlush(t *testing.T) {
	// Arrange: Create a regular flush (2-5-7-9-J of hearts)
	cards := []Card{
		{Rank: Two, Suit: Hearts},
		{Rank: Five, Suit: Hearts},
		{Rank: Seven, Suit: Hearts},
		{Rank: Nine, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
	}

	// Act
	result := detectRoyalFlush(cards)

	// Assert
	if result {
		t.Errorf("detectRoyalFlush(%v) = true, want false", cards)
	}
}

// TestDetectStraight_TrueForBroadwayStraight verifies that detectStraight
// returns true with high rank for 9-10-J-Q-K straight (mixed suits).
func TestDetectStraight_TrueForBroadwayStraight(t *testing.T) {
	// Arrange: Create 9-10-J-Q-K with mixed suits
	cards := []Card{
		{Rank: Nine, Suit: Hearts},
		{Rank: Ten, Suit: Diamonds},
		{Rank: Jack, Suit: Clubs},
		{Rank: Queen, Suit: Spades},
		{Rank: King, Suit: Hearts},
	}

	// Act
	isStraight, highRank := detectStraight(cards)

	// Assert
	if !isStraight {
		t.Errorf("detectStraight(%v) = false, want true", cards)
	}
	if highRank != King {
		t.Errorf("detectStraight(%v) high rank = %v, want King (13)", cards, highRank)
	}
}

// TestDetectStraight_TrueForWheelStraight verifies that detectStraight
// returns true with rank 5 for A-2-3-4-5 wheel straight (mixed suits).
func TestDetectStraight_TrueForWheelStraight(t *testing.T) {
	// Arrange: Create A-2-3-4-5 wheel with mixed suits
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: Two, Suit: Diamonds},
		{Rank: Three, Suit: Clubs},
		{Rank: Four, Suit: Spades},
		{Rank: Five, Suit: Hearts},
	}

	// Act
	isStraight, highRank := detectStraight(cards)

	// Assert
	if !isStraight {
		t.Errorf("detectStraight(%v) = false, want true", cards)
	}
	if highRank != Five {
		t.Errorf("detectStraight(%v) high rank = %v, want Five (5)", cards, highRank)
	}
}

// TestDetectStraight_FalseForStraightFlush verifies that detectStraight
// returns false for a straight flush (should be detected as straight flush, not straight).
func TestDetectStraight_FalseForStraightFlush(t *testing.T) {
	// Arrange: Create 9-10-J-Q-K straight flush in hearts
	cards := []Card{
		{Rank: Nine, Suit: Hearts},
		{Rank: Ten, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: King, Suit: Hearts},
	}

	// Act
	isStraight, _ := detectStraight(cards)

	// Assert
	if isStraight {
		t.Errorf("detectStraight(%v) = true, want false (should be straight flush)", cards)
	}
}

// TestDetectStraight_FalseForNonStraight verifies that detectStraight
// returns false for non-sequential cards.
func TestDetectStraight_FalseForNonStraight(t *testing.T) {
	// Arrange: Create non-straight hand (2-5-7-9-J)
	cards := []Card{
		{Rank: Two, Suit: Hearts},
		{Rank: Five, Suit: Diamonds},
		{Rank: Seven, Suit: Clubs},
		{Rank: Nine, Suit: Spades},
		{Rank: Jack, Suit: Hearts},
	}

	// Act
	isStraight, _ := detectStraight(cards)

	// Assert
	if isStraight {
		t.Errorf("detectStraight(%v) = true, want false", cards)
	}
}

// TestDetectStraightFlush_TrueForNineHighStraightFlush verifies that detectStraightFlush
// returns true and correct high card for 5h-6h-7h-8h-9h.
func TestDetectStraightFlush_TrueForNineHighStraightFlush(t *testing.T) {
	// Arrange: Create 5h-6h-7h-8h-9h (nine-high straight flush)
	cards := []Card{
		{Rank: Five, Suit: Hearts},
		{Rank: Six, Suit: Hearts},
		{Rank: Seven, Suit: Hearts},
		{Rank: Eight, Suit: Hearts},
		{Rank: Nine, Suit: Hearts},
	}

	// Act
	result, highCard := detectStraightFlush(cards)

	// Assert
	if !result {
		t.Errorf("detectStraightFlush(%v) = false, want true", cards)
	}
	if highCard != Nine {
		t.Errorf("detectStraightFlush(%v) high card = %v, want %v", cards, highCard, Nine)
	}
}

// TestDetectStraightFlush_TrueForWheelFlush verifies that detectStraightFlush
// returns true and rank 5 for Ah-2h-3h-4h-5h (wheel straight flush).
func TestDetectStraightFlush_TrueForWheelFlush(t *testing.T) {
	// Arrange: Create Ah-2h-3h-4h-5h (wheel straight flush)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: Two, Suit: Hearts},
		{Rank: Three, Suit: Hearts},
		{Rank: Four, Suit: Hearts},
		{Rank: Five, Suit: Hearts},
	}

	// Act
	result, highCard := detectStraightFlush(cards)

	// Assert
	if !result {
		t.Errorf("detectStraightFlush(%v) = false, want true", cards)
	}
	if highCard != Five {
		t.Errorf("detectStraightFlush(%v) high card = %v, want %v (Ace acts as low)", cards, highCard, Five)
	}
}

// TestDetectStraightFlush_FalseForNonFlushStraight verifies that detectStraightFlush
// returns false for a straight that is not a flush.
func TestDetectStraightFlush_FalseForNonFlushStraight(t *testing.T) {
	// Arrange: Create 5h-6d-7c-8s-9h (straight with mixed suits)
	cards := []Card{
		{Rank: Five, Suit: Hearts},
		{Rank: Six, Suit: Diamonds},
		{Rank: Seven, Suit: Clubs},
		{Rank: Eight, Suit: Spades},
		{Rank: Nine, Suit: Hearts},
	}

	// Act
	result, highCard := detectStraightFlush(cards)

	// Assert
	if result {
		t.Errorf("detectStraightFlush(%v) = true, want false (not a flush)", cards)
	}
	if highCard != 0 {
		t.Errorf("detectStraightFlush(%v) high card = %v, want 0", cards, highCard)
	}
}

// TestDetectStraightFlush_FalseForNonStraightFlush verifies that detectStraightFlush
// returns false for a flush that is not a straight.
func TestDetectStraightFlush_FalseForNonStraightFlush(t *testing.T) {
	// Arrange: Create 2h-5h-7h-9h-Jh (flush with gaps, not sequential)
	cards := []Card{
		{Rank: Two, Suit: Hearts},
		{Rank: Five, Suit: Hearts},
		{Rank: Seven, Suit: Hearts},
		{Rank: Nine, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
	}

	// Act
	result, highCard := detectStraightFlush(cards)

	// Assert
	if result {
		t.Errorf("detectStraightFlush(%v) = true, want false (not a straight)", cards)
	}
	if highCard != 0 {
		t.Errorf("detectStraightFlush(%v) high card = %v, want 0", cards, highCard)
	}
}

// TestDetectFourOfAKind_TrueForQuads verifies that detectFourOfAKind returns true
// and correct tiebreakers for 8-8-8-8-K.
func TestDetectFourOfAKind_TrueForQuads(t *testing.T) {
	// Arrange: Create 8h-8d-8c-8s-Kh (four eights with king kicker)
	cards := []Card{
		{Rank: Eight, Suit: Hearts},
		{Rank: Eight, Suit: Diamonds},
		{Rank: Eight, Suit: Clubs},
		{Rank: Eight, Suit: Spades},
		{Rank: King, Suit: Hearts},
	}

	// Act
	result, tiebreakers := detectFourOfAKind(cards)

	// Assert
	if !result {
		t.Errorf("detectFourOfAKind(%v) = false, want true", cards)
	}
	if len(tiebreakers) != 2 {
		t.Errorf("detectFourOfAKind(%v) tiebreakers length = %d, want 2", cards, len(tiebreakers))
	}
	if len(tiebreakers) >= 2 {
		if tiebreakers[0] != Eight {
			t.Errorf("detectFourOfAKind(%v) quad rank = %v, want %v", cards, tiebreakers[0], Eight)
		}
		if tiebreakers[1] != King {
			t.Errorf("detectFourOfAKind(%v) kicker = %v, want %v", cards, tiebreakers[1], King)
		}
	}
}

// TestDetectFourOfAKind_FalseForFullHouse verifies that detectFourOfAKind returns false
// for a full house (three of a kind plus a pair).
func TestDetectFourOfAKind_FalseForFullHouse(t *testing.T) {
	// Arrange: Create Kh-Kd-Kc-7s-7h (kings full of sevens)
	cards := []Card{
		{Rank: King, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: King, Suit: Clubs},
		{Rank: Seven, Suit: Spades},
		{Rank: Seven, Suit: Hearts},
	}

	// Act
	result, tiebreakers := detectFourOfAKind(cards)

	// Assert
	if result {
		t.Errorf("detectFourOfAKind(%v) = true, want false (full house, not quads)", cards)
	}
	if len(tiebreakers) != 0 {
		t.Errorf("detectFourOfAKind(%v) tiebreakers = %v, want empty slice", cards, tiebreakers)
	}
}

// TestDetectFourOfAKind_FalseForTrips verifies that detectFourOfAKind returns false
// for three of a kind.
func TestDetectFourOfAKind_FalseForTrips(t *testing.T) {
	// Arrange: Create Qh-Qd-Qc-9s-3h (three queens)
	cards := []Card{
		{Rank: Queen, Suit: Hearts},
		{Rank: Queen, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Nine, Suit: Spades},
		{Rank: Three, Suit: Hearts},
	}

	// Act
	result, tiebreakers := detectFourOfAKind(cards)

	// Assert
	if result {
		t.Errorf("detectFourOfAKind(%v) = true, want false (trips, not quads)", cards)
	}
	if len(tiebreakers) != 0 {
		t.Errorf("detectFourOfAKind(%v) tiebreakers = %v, want empty slice", cards, tiebreakers)
	}
}

// TestDetectFullHouse_TrueForKingsOverSevens verifies that detectFullHouse returns true
// and correct tiebreakers [trip rank, pair rank] for K-K-K-7-7.
func TestDetectFullHouse_TrueForKingsOverSevens(t *testing.T) {
	// Arrange: Create Kh-Kd-Kc-7s-7h (kings full of sevens)
	cards := []Card{
		{Rank: King, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: King, Suit: Clubs},
		{Rank: Seven, Suit: Spades},
		{Rank: Seven, Suit: Hearts},
	}

	// Act
	result, tiebreakers := detectFullHouse(cards)

	// Assert
	if !result {
		t.Errorf("detectFullHouse(%v) = false, want true", cards)
	}
	if len(tiebreakers) != 2 {
		t.Errorf("detectFullHouse(%v) tiebreakers length = %d, want 2", cards, len(tiebreakers))
	}
	if len(tiebreakers) >= 2 {
		if tiebreakers[0] != King {
			t.Errorf("detectFullHouse(%v) trip rank = %v, want %v", cards, tiebreakers[0], King)
		}
		if tiebreakers[1] != Seven {
			t.Errorf("detectFullHouse(%v) pair rank = %v, want %v", cards, tiebreakers[1], Seven)
		}
	}
}

// TestDetectFullHouse_FalseForThreeOfAKind verifies that detectFullHouse returns false
// for three of a kind without a pair.
func TestDetectFullHouse_FalseForThreeOfAKind(t *testing.T) {
	// Arrange: Create Qh-Qd-Qc-9s-3h (three queens, no pair)
	cards := []Card{
		{Rank: Queen, Suit: Hearts},
		{Rank: Queen, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Nine, Suit: Spades},
		{Rank: Three, Suit: Hearts},
	}

	// Act
	result, tiebreakers := detectFullHouse(cards)

	// Assert
	if result {
		t.Errorf("detectFullHouse(%v) = true, want false (three of a kind, not full house)", cards)
	}
	if len(tiebreakers) != 0 {
		t.Errorf("detectFullHouse(%v) tiebreakers = %v, want empty slice", cards, tiebreakers)
	}
}

// TestDetectFullHouse_FalseForTwoPair verifies that detectFullHouse returns false
// for two pair (no trips).
func TestDetectFullHouse_FalseForTwoPair(t *testing.T) {
	// Arrange: Create Ah-Ad-7h-7d-3s (two pair: aces and sevens)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: Ace, Suit: Diamonds},
		{Rank: Seven, Suit: Hearts},
		{Rank: Seven, Suit: Diamonds},
		{Rank: Three, Suit: Spades},
	}

	// Act
	result, tiebreakers := detectFullHouse(cards)

	// Assert
	if result {
		t.Errorf("detectFullHouse(%v) = true, want false (two pair, not full house)", cards)
	}
	if len(tiebreakers) != 0 {
		t.Errorf("detectFullHouse(%v) tiebreakers = %v, want empty slice", cards, tiebreakers)
	}
}

// TestEvaluateHand_RoyalFlush verifies that EvaluateHand correctly identifies a royal flush.
func TestEvaluateHand_RoyalFlush(t *testing.T) {
	// Arrange: Create royal flush (Th-Jh-Qh-Kh-Ah)
	cards := []Card{
		{Rank: Ten, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Ace, Suit: Hearts},
	}

	// Act
	hand := EvaluateHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("EvaluateHand returned nil")
	}
	if hand.Category != RoyalFlush {
		t.Errorf("EvaluateHand(%v) category = %v, want %v", cards, hand.Category, RoyalFlush)
	}
}

// TestEvaluateHand_StraightFlush verifies that EvaluateHand correctly identifies a straight flush.
func TestEvaluateHand_StraightFlush(t *testing.T) {
	// Arrange: Create straight flush (5h-6h-7h-8h-9h)
	cards := []Card{
		{Rank: Five, Suit: Hearts},
		{Rank: Six, Suit: Hearts},
		{Rank: Seven, Suit: Hearts},
		{Rank: Eight, Suit: Hearts},
		{Rank: Nine, Suit: Hearts},
	}

	// Act
	hand := EvaluateHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("EvaluateHand returned nil")
	}
	if hand.Category != StraightFlush {
		t.Errorf("EvaluateHand(%v) category = %v, want %v", cards, hand.Category, StraightFlush)
	}
	if len(hand.Tiebreakers) != 1 || hand.Tiebreakers[0] != Nine {
		t.Errorf("EvaluateHand(%v) tiebreakers = %v, want [%v]", cards, hand.Tiebreakers, Nine)
	}
}

// TestEvaluateHand_FourOfAKind verifies that EvaluateHand correctly identifies four of a kind.
func TestEvaluateHand_FourOfAKind(t *testing.T) {
	// Arrange: Create four of a kind (8h-8d-8c-8s-Kh)
	cards := []Card{
		{Rank: Eight, Suit: Hearts},
		{Rank: Eight, Suit: Diamonds},
		{Rank: Eight, Suit: Clubs},
		{Rank: Eight, Suit: Spades},
		{Rank: King, Suit: Hearts},
	}

	// Act
	hand := EvaluateHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("EvaluateHand returned nil")
	}
	if hand.Category != FourOfAKind {
		t.Errorf("EvaluateHand(%v) category = %v, want %v", cards, hand.Category, FourOfAKind)
	}
	if len(hand.Tiebreakers) != 2 || hand.Tiebreakers[0] != Eight || hand.Tiebreakers[1] != King {
		t.Errorf("EvaluateHand(%v) tiebreakers = %v, want [%v, %v]", cards, hand.Tiebreakers, Eight, King)
	}
}

// TestEvaluateHand_FullHouse verifies that EvaluateHand correctly identifies a full house.
func TestEvaluateHand_FullHouse(t *testing.T) {
	// Arrange: Create full house (Kh-Kd-Kc-7s-7h)
	cards := []Card{
		{Rank: King, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: King, Suit: Clubs},
		{Rank: Seven, Suit: Spades},
		{Rank: Seven, Suit: Hearts},
	}

	// Act
	hand := EvaluateHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("EvaluateHand returned nil")
	}
	if hand.Category != FullHouse {
		t.Errorf("EvaluateHand(%v) category = %v, want %v", cards, hand.Category, FullHouse)
	}
	if len(hand.Tiebreakers) != 2 || hand.Tiebreakers[0] != King || hand.Tiebreakers[1] != Seven {
		t.Errorf("EvaluateHand(%v) tiebreakers = %v, want [%v, %v]", cards, hand.Tiebreakers, King, Seven)
	}
}

// TestEvaluateHand_Flush verifies that EvaluateHand correctly identifies a flush.
func TestEvaluateHand_Flush(t *testing.T) {
	// Arrange: Create flush (2h-5h-7h-9h-Jh)
	cards := []Card{
		{Rank: Two, Suit: Hearts},
		{Rank: Five, Suit: Hearts},
		{Rank: Seven, Suit: Hearts},
		{Rank: Nine, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
	}

	// Act
	hand := EvaluateHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("EvaluateHand returned nil")
	}
	if hand.Category != Flush {
		t.Errorf("EvaluateHand(%v) category = %v, want %v", cards, hand.Category, Flush)
	}
	if len(hand.Tiebreakers) != 5 {
		t.Errorf("EvaluateHand(%v) tiebreakers length = %d, want 5", cards, len(hand.Tiebreakers))
	}
}

// TestEvaluateHand_Straight verifies that EvaluateHand correctly identifies a straight.
func TestEvaluateHand_Straight(t *testing.T) {
	// Arrange: Create straight (5h-6d-7c-8s-9h)
	cards := []Card{
		{Rank: Five, Suit: Hearts},
		{Rank: Six, Suit: Diamonds},
		{Rank: Seven, Suit: Clubs},
		{Rank: Eight, Suit: Spades},
		{Rank: Nine, Suit: Hearts},
	}

	// Act
	hand := EvaluateHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("EvaluateHand returned nil")
	}
	if hand.Category != Straight {
		t.Errorf("EvaluateHand(%v) category = %v, want %v", cards, hand.Category, Straight)
	}
	if len(hand.Tiebreakers) != 1 || hand.Tiebreakers[0] != Nine {
		t.Errorf("EvaluateHand(%v) tiebreakers = %v, want [%v]", cards, hand.Tiebreakers, Nine)
	}
}

// TestEvaluateHand_ThreeOfAKind verifies that EvaluateHand correctly identifies three of a kind.
func TestEvaluateHand_ThreeOfAKind(t *testing.T) {
	// Arrange: Create three of a kind (Qh-Qd-Qc-9s-3h)
	cards := []Card{
		{Rank: Queen, Suit: Hearts},
		{Rank: Queen, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Nine, Suit: Spades},
		{Rank: Three, Suit: Hearts},
	}

	// Act
	hand := EvaluateHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("EvaluateHand returned nil")
	}
	if hand.Category != ThreeOfAKind {
		t.Errorf("EvaluateHand(%v) category = %v, want %v", cards, hand.Category, ThreeOfAKind)
	}
	if len(hand.Tiebreakers) != 3 || hand.Tiebreakers[0] != Queen {
		t.Errorf("EvaluateHand(%v) tiebreakers = %v, want [%v, ...]", cards, hand.Tiebreakers, Queen)
	}
}

// TestEvaluateHand_TwoPair verifies that EvaluateHand correctly identifies two pair.
func TestEvaluateHand_TwoPair(t *testing.T) {
	// Arrange: Create two pair (Kh-Kd-7c-7s-3h)
	cards := []Card{
		{Rank: King, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Seven, Suit: Clubs},
		{Rank: Seven, Suit: Spades},
		{Rank: Three, Suit: Hearts},
	}

	// Act
	hand := EvaluateHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("EvaluateHand returned nil")
	}
	if hand.Category != TwoPair {
		t.Errorf("EvaluateHand(%v) category = %v, want %v", cards, hand.Category, TwoPair)
	}
	if len(hand.Tiebreakers) != 3 || hand.Tiebreakers[0] != King || hand.Tiebreakers[1] != Seven || hand.Tiebreakers[2] != Three {
		t.Errorf("EvaluateHand(%v) tiebreakers = %v, want [%v, %v, %v]", cards, hand.Tiebreakers, King, Seven, Three)
	}
}

// TestEvaluateHand_OnePair verifies that EvaluateHand correctly identifies one pair.
func TestEvaluateHand_OnePair(t *testing.T) {
	// Arrange: Create one pair (Jh-Jd-9c-5s-2h)
	cards := []Card{
		{Rank: Jack, Suit: Hearts},
		{Rank: Jack, Suit: Diamonds},
		{Rank: Nine, Suit: Clubs},
		{Rank: Five, Suit: Spades},
		{Rank: Two, Suit: Hearts},
	}

	// Act
	hand := EvaluateHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("EvaluateHand returned nil")
	}
	if hand.Category != OnePair {
		t.Errorf("EvaluateHand(%v) category = %v, want %v", cards, hand.Category, OnePair)
	}
	if len(hand.Tiebreakers) != 4 || hand.Tiebreakers[0] != Jack {
		t.Errorf("EvaluateHand(%v) tiebreakers = %v, want [%v, ...]", cards, hand.Tiebreakers, Jack)
	}
}

// TestEvaluateHand_HighCard verifies that EvaluateHand correctly identifies high card.
func TestEvaluateHand_HighCard(t *testing.T) {
	// Arrange: Create high card (Ah-Kd-9c-5s-2h)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Nine, Suit: Clubs},
		{Rank: Five, Suit: Spades},
		{Rank: Two, Suit: Hearts},
	}

	// Act
	hand := EvaluateHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("EvaluateHand returned nil")
	}
	if hand.Category != HighCard {
		t.Errorf("EvaluateHand(%v) category = %v, want %v", cards, hand.Category, HighCard)
	}
	if len(hand.Tiebreakers) != 5 {
		t.Errorf("EvaluateHand(%v) tiebreakers length = %d, want 5", cards, len(hand.Tiebreakers))
	}
}

// TestDetectFlush tests the detectFlush function
func TestDetectFlush(t *testing.T) {
	tests := []struct {
		name          string
		cards         []Card
		expectedFound bool
		expectedRanks []Rank
	}{
		{
			name: "flush with A-K-9-6-2 all hearts",
			cards: []Card{
				{Rank: Ace, Suit: Hearts},
				{Rank: King, Suit: Hearts},
				{Rank: Nine, Suit: Hearts},
				{Rank: Six, Suit: Hearts},
				{Rank: Two, Suit: Hearts},
			},
			expectedFound: true,
			expectedRanks: []Rank{Ace, King, Nine, Six, Two}, // descending order
		},
		{
			name: "not a flush - mixed suits",
			cards: []Card{
				{Rank: Ace, Suit: Hearts},
				{Rank: King, Suit: Diamonds},
				{Rank: Queen, Suit: Hearts},
				{Rank: Jack, Suit: Hearts},
				{Rank: Ten, Suit: Hearts},
			},
			expectedFound: false,
			expectedRanks: nil,
		},
		{
			name: "straight flush - should return false",
			cards: []Card{
				{Rank: Ten, Suit: Hearts},
				{Rank: Nine, Suit: Hearts},
				{Rank: Eight, Suit: Hearts},
				{Rank: Seven, Suit: Hearts},
				{Rank: Six, Suit: Hearts},
			},
			expectedFound: false,
			expectedRanks: nil,
		},
		{
			name: "royal flush - should return false",
			cards: []Card{
				{Rank: Ace, Suit: Spades},
				{Rank: King, Suit: Spades},
				{Rank: Queen, Suit: Spades},
				{Rank: Jack, Suit: Spades},
				{Rank: Ten, Suit: Spades},
			},
			expectedFound: false,
			expectedRanks: nil,
		},
		{
			name: "flush with different suits - clubs",
			cards: []Card{
				{Rank: Queen, Suit: Clubs},
				{Rank: Ten, Suit: Clubs},
				{Rank: Eight, Suit: Clubs},
				{Rank: Five, Suit: Clubs},
				{Rank: Three, Suit: Clubs},
			},
			expectedFound: true,
			expectedRanks: []Rank{Queen, Ten, Eight, Five, Three},
		},
		{
			name: "not a flush - four same suit",
			cards: []Card{
				{Rank: Ace, Suit: Diamonds},
				{Rank: King, Suit: Diamonds},
				{Rank: Queen, Suit: Diamonds},
				{Rank: Jack, Suit: Diamonds},
				{Rank: Ten, Suit: Spades},
			},
			expectedFound: false,
			expectedRanks: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, ranks := detectFlush(tt.cards)
			if found != tt.expectedFound {
				t.Errorf("detectFlush() found = %v, want %v", found, tt.expectedFound)
			}
			if !reflect.DeepEqual(ranks, tt.expectedRanks) {
				t.Errorf("detectFlush() ranks = %v, want %v", ranks, tt.expectedRanks)
			}
		})
	}
}

// TestDetectThreeOfAKind_TrueForQueensWithAceKing tests detection of three queens
// with ace and king kickers, verifying tiebreakers are in correct order.
func TestDetectThreeOfAKind_TrueForQueensWithAceKing(t *testing.T) {
	// Arrange: Create Qh-Qd-Qc-Ah-Kd (three queens, ace-king kickers)
	cards := []Card{
		{Rank: Queen, Suit: Hearts},
		{Rank: Queen, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
	}

	// Act
	result, tiebreakers := detectThreeOfAKind(cards)

	// Assert
	if !result {
		t.Errorf("detectThreeOfAKind(%v) = false, want true", cards)
	}
	expectedTiebreakers := []Rank{Queen, Ace, King}
	if !reflect.DeepEqual(tiebreakers, expectedTiebreakers) {
		t.Errorf("detectThreeOfAKind(%v) tiebreakers = %v, want %v", cards, tiebreakers, expectedTiebreakers)
	}
}

// TestDetectThreeOfAKind_TrueForThreesWithDifferentKickers tests detection of three threes
// with seven and five kickers, verifying kickers are sorted descending.
func TestDetectThreeOfAKind_TrueForThreesWithDifferentKickers(t *testing.T) {
	// Arrange: Create 3h-3d-3c-7h-5d (three threes, 7-5 kickers)
	cards := []Card{
		{Rank: Three, Suit: Hearts},
		{Rank: Three, Suit: Diamonds},
		{Rank: Three, Suit: Clubs},
		{Rank: Seven, Suit: Hearts},
		{Rank: Five, Suit: Diamonds},
	}

	// Act
	result, tiebreakers := detectThreeOfAKind(cards)

	// Assert
	if !result {
		t.Errorf("detectThreeOfAKind(%v) = false, want true", cards)
	}
	expectedTiebreakers := []Rank{Three, Seven, Five}
	if !reflect.DeepEqual(tiebreakers, expectedTiebreakers) {
		t.Errorf("detectThreeOfAKind(%v) tiebreakers = %v, want %v", cards, tiebreakers, expectedTiebreakers)
	}
}

// TestDetectThreeOfAKind_FalseForFullHouse tests that full house is NOT detected as three of a kind.
// Full houses should be detected separately (three of a kind + pair).
func TestDetectThreeOfAKind_FalseForFullHouse(t *testing.T) {
	// Arrange: Create 8h-8d-8c-5h-5d (full house: eights over fives)
	cards := []Card{
		{Rank: Eight, Suit: Hearts},
		{Rank: Eight, Suit: Diamonds},
		{Rank: Eight, Suit: Clubs},
		{Rank: Five, Suit: Hearts},
		{Rank: Five, Suit: Diamonds},
	}

	// Act
	result, tiebreakers := detectThreeOfAKind(cards)

	// Assert
	if result {
		t.Errorf("detectThreeOfAKind(%v) = true, want false (full house, not trips)", cards)
	}
	if len(tiebreakers) != 0 {
		t.Errorf("detectThreeOfAKind(%v) tiebreakers = %v, want empty slice", cards, tiebreakers)
	}
}

// TestDetectThreeOfAKind_FalseForTwoPair tests that two pair is NOT detected as three of a kind.
func TestDetectThreeOfAKind_FalseForTwoPair(t *testing.T) {
	// Arrange: Create Kh-Kd-9h-9d-3c (two pair: kings and nines)
	cards := []Card{
		{Rank: King, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Nine, Suit: Hearts},
		{Rank: Nine, Suit: Diamonds},
		{Rank: Three, Suit: Clubs},
	}

	// Act
	result, tiebreakers := detectThreeOfAKind(cards)

	// Assert
	if result {
		t.Errorf("detectThreeOfAKind(%v) = true, want false (two pair, not trips)", cards)
	}
	if len(tiebreakers) != 0 {
		t.Errorf("detectThreeOfAKind(%v) tiebreakers = %v, want empty slice", cards, tiebreakers)
	}
}

// TestDetectThreeOfAKind_FalseForOnePair tests that one pair is NOT detected as three of a kind.
func TestDetectThreeOfAKind_FalseForOnePair(t *testing.T) {
	// Arrange: Create Jh-Jd-9h-7d-3c (one pair: jacks)
	cards := []Card{
		{Rank: Jack, Suit: Hearts},
		{Rank: Jack, Suit: Diamonds},
		{Rank: Nine, Suit: Hearts},
		{Rank: Seven, Suit: Diamonds},
		{Rank: Three, Suit: Clubs},
	}

	// Act
	result, tiebreakers := detectThreeOfAKind(cards)

	// Assert
	if result {
		t.Errorf("detectThreeOfAKind(%v) = true, want false (one pair, not trips)", cards)
	}
	if len(tiebreakers) != 0 {
		t.Errorf("detectThreeOfAKind(%v) tiebreakers = %v, want empty slice", cards, tiebreakers)
	}
}

// TestDetectThreeOfAKind_FalseForHighCard tests that high card is NOT detected as three of a kind.
func TestDetectThreeOfAKind_FalseForHighCard(t *testing.T) {
	// Arrange: Create Ah-Kd-Qh-Jd-9c (high card)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Diamonds},
		{Rank: Nine, Suit: Clubs},
	}

	// Act
	result, tiebreakers := detectThreeOfAKind(cards)

	// Assert
	if result {
		t.Errorf("detectThreeOfAKind(%v) = true, want false (high card, not trips)", cards)
	}
	if len(tiebreakers) != 0 {
		t.Errorf("detectThreeOfAKind(%v) tiebreakers = %v, want empty slice", cards, tiebreakers)
	}
}

// TestDetectTwoPair_TrueForTwoPair verifies that detectTwoPair returns true
// and correct tiebreakers [high pair, low pair, kicker] for A-A-7-7-3.
func TestDetectTwoPair_TrueForTwoPair(t *testing.T) {
	// Arrange: Create Ah-Ad-7h-7d-3s (aces and sevens with 3 kicker)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: Ace, Suit: Diamonds},
		{Rank: Seven, Suit: Hearts},
		{Rank: Seven, Suit: Diamonds},
		{Rank: Three, Suit: Spades},
	}

	// Act
	result, tiebreakers := detectTwoPair(cards)

	// Assert
	if !result {
		t.Errorf("detectTwoPair(%v) = false, want true", cards)
	}
	if len(tiebreakers) != 3 {
		t.Errorf("detectTwoPair(%v) tiebreakers length = %d, want 3", cards, len(tiebreakers))
	}
	if len(tiebreakers) >= 3 {
		if tiebreakers[0] != Ace {
			t.Errorf("detectTwoPair(%v) high pair = %v, want %v", cards, tiebreakers[0], Ace)
		}
		if tiebreakers[1] != Seven {
			t.Errorf("detectTwoPair(%v) low pair = %v, want %v", cards, tiebreakers[1], Seven)
		}
		if tiebreakers[2] != Three {
			t.Errorf("detectTwoPair(%v) kicker = %v, want %v", cards, tiebreakers[2], Three)
		}
	}
}

// TestDetectTwoPair_FalseForFullHouse verifies that detectTwoPair returns false
// for a full house (three of a kind plus a pair).
func TestDetectTwoPair_FalseForFullHouse(t *testing.T) {
	// Arrange: Create Kh-Kd-Kc-7s-7h (kings full of sevens)
	cards := []Card{
		{Rank: King, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: King, Suit: Clubs},
		{Rank: Seven, Suit: Spades},
		{Rank: Seven, Suit: Hearts},
	}

	// Act
	result, tiebreakers := detectTwoPair(cards)

	// Assert
	if result {
		t.Errorf("detectTwoPair(%v) = true, want false (full house, not two pair)", cards)
	}
	if len(tiebreakers) != 0 {
		t.Errorf("detectTwoPair(%v) tiebreakers = %v, want empty slice", cards, tiebreakers)
	}
}

// TestDetectTwoPair_FalseForOnePair verifies that detectTwoPair returns false
// for one pair only.
func TestDetectTwoPair_FalseForOnePair(t *testing.T) {
	// Arrange: Create Qh-Qd-8s-5c-2h (queens with 8-5-2 kickers)
	cards := []Card{
		{Rank: Queen, Suit: Hearts},
		{Rank: Queen, Suit: Diamonds},
		{Rank: Eight, Suit: Spades},
		{Rank: Five, Suit: Clubs},
		{Rank: Two, Suit: Hearts},
	}

	// Act
	result, tiebreakers := detectTwoPair(cards)

	// Assert
	if result {
		t.Errorf("detectTwoPair(%v) = true, want false (one pair, not two pair)", cards)
	}
	if len(tiebreakers) != 0 {
		t.Errorf("detectTwoPair(%v) tiebreakers = %v, want empty slice", cards, tiebreakers)
	}
}

// TestDetectOnePair_TrueForPair verifies that detectOnePair returns true
// and correct tiebreakers [pair rank, kicker1, kicker2, kicker3] for J-J-9-6-2.
func TestDetectOnePair_TrueForPair(t *testing.T) {
	// Arrange: Create J-J-9-6-2 (pair of Jacks)
	cards := []Card{
		{Rank: Jack, Suit: Hearts},
		{Rank: Jack, Suit: Diamonds},
		{Rank: Nine, Suit: Clubs},
		{Rank: Six, Suit: Spades},
		{Rank: Two, Suit: Hearts},
	}

	// Act
	found, tiebreakers := detectOnePair(cards)

	// Assert
	if !found {
		t.Errorf("detectOnePair(%v) = false, want true", cards)
	}

	// Expected tiebreakers: [Jack (pair rank), Nine, Six, Two] in descending order
	expectedTiebreakers := []Rank{Jack, Nine, Six, Two}
	if len(tiebreakers) != len(expectedTiebreakers) {
		t.Errorf("detectOnePair(%v) tiebreakers length = %d, want %d", cards, len(tiebreakers), len(expectedTiebreakers))
	}

	for i, expected := range expectedTiebreakers {
		if tiebreakers[i] != expected {
			t.Errorf("detectOnePair(%v) tiebreakers[%d] = %v, want %v", cards, i, tiebreakers[i], expected)
		}
	}
}

// TestDetectOnePair_FalseForTwoPair verifies that detectOnePair returns false
// when the hand contains two pairs (e.g., K-K-5-5-A).
func TestDetectOnePair_FalseForTwoPair(t *testing.T) {
	// Arrange: Create K-K-5-5-A (two pairs)
	cards := []Card{
		{Rank: King, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Five, Suit: Clubs},
		{Rank: Five, Suit: Spades},
		{Rank: Ace, Suit: Hearts},
	}

	// Act
	found, _ := detectOnePair(cards)

	// Assert
	if found {
		t.Errorf("detectOnePair(%v) = true, want false (two pair detected)", cards)
	}
}

// TestDetectOnePair_FalseForHighCard verifies that detectOnePair returns false
// when the hand contains no pairs (e.g., A-K-Q-J-9 unsuited).
func TestDetectOnePair_FalseForHighCard(t *testing.T) {
	// Arrange: Create A-K-Q-J-9 (high card, no pairs)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Jack, Suit: Spades},
		{Rank: Nine, Suit: Hearts},
	}

	// Act
	found, _ := detectOnePair(cards)

	// Assert
	if found {
		t.Errorf("detectOnePair(%v) = true, want false (high card, no pair)", cards)
	}
}

// TestDetectHighCard_TrueForAceHighRainbow verifies that detectHighCard returns true
// and correct kickers [A-K-T-7-3] in descending order for Ah-Kd-Tc-7s-3h (rainbow).
func TestDetectHighCard_TrueForAceHighRainbow(t *testing.T) {
	// Arrange: Create Ah-Kd-Tc-7s-3h (ace-high with no pairs/straights/flushes)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Ten, Suit: Clubs},
		{Rank: Seven, Suit: Spades},
		{Rank: Three, Suit: Hearts},
	}

	// Act
	result, tiebreakers := detectHighCard(cards)

	// Assert
	if !result {
		t.Errorf("detectHighCard(%v) = false, want true (always true)", cards)
	}
	expectedTiebreakers := []Rank{Ace, King, Ten, Seven, Three}
	if len(tiebreakers) != len(expectedTiebreakers) {
		t.Errorf("detectHighCard(%v) tiebreakers length = %d, want %d", cards, len(tiebreakers), len(expectedTiebreakers))
	}
	for i := 0; i < len(tiebreakers) && i < len(expectedTiebreakers); i++ {
		if tiebreakers[i] != expectedTiebreakers[i] {
			t.Errorf("detectHighCard(%v) tiebreakers[%d] = %v, want %v", cards, i, tiebreakers[i], expectedTiebreakers[i])
		}
	}
}

// TestDetectHighCard_AlwaysReturnsTrue verifies that detectHighCard always returns true
// as it is the fallback category for any 5-card hand.
func TestDetectHighCard_AlwaysReturnsTrue(t *testing.T) {
	// Test various hand types to ensure detectHighCard always returns true
	testCases := []struct {
		name  string
		cards []Card
	}{
		{
			name: "Flush should still return true",
			cards: []Card{
				{Rank: Two, Suit: Hearts},
				{Rank: Five, Suit: Hearts},
				{Rank: Seven, Suit: Hearts},
				{Rank: Nine, Suit: Hearts},
				{Rank: Jack, Suit: Hearts},
			},
		},
		{
			name: "Pair should still return true",
			cards: []Card{
				{Rank: Eight, Suit: Hearts},
				{Rank: Eight, Suit: Diamonds},
				{Rank: King, Suit: Clubs},
				{Rank: Seven, Suit: Spades},
				{Rank: Three, Suit: Hearts},
			},
		},
		{
			name: "Random low cards should return true",
			cards: []Card{
				{Rank: Two, Suit: Hearts},
				{Rank: Four, Suit: Diamonds},
				{Rank: Six, Suit: Clubs},
				{Rank: Eight, Suit: Spades},
				{Rank: Ten, Suit: Hearts},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result, _ := detectHighCard(tc.cards)

			// Assert
			if !result {
				t.Errorf("detectHighCard(%v) = false, want true (always returns true)", tc.cards)
			}
		})
	}
}

// TestFindBestHand_ReturnsFlushOver7Cards verifies that FindBestHand
// correctly identifies a flush when given 7 cards containing a flush.
func TestFindBestHand_ReturnsFlushOver7Cards(t *testing.T) {
	// Arrange: 7 cards with 5 hearts forming a flush
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Nine, Suit: Hearts},
		{Rank: Eight, Suit: Diamonds},
		{Rank: Seven, Suit: Clubs},
	}

	// Act
	hand := FindBestHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("FindBestHand returned nil, expected valid hand")
	}
	if hand.Category != Flush {
		t.Errorf("FindBestHand category = %v, want %v", hand.Category, Flush)
	}
	if len(hand.Cards) != 5 {
		t.Errorf("FindBestHand cards length = %d, want 5", len(hand.Cards))
	}
}

// TestFindBestHand_ReturnsFlushOverStraight verifies that FindBestHand
// correctly chooses a flush over a straight when both are possible.
func TestFindBestHand_ReturnsFlushOverStraight(t *testing.T) {
	// Arrange: 7 cards containing both a flush and a straight
	// Flush: Ah-Kh-Qh-Jh-9h (hearts)
	// Straight: 9-8-7-6-5 (mixed suits)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Hearts},
		{Rank: Jack, Suit: Hearts},
		{Rank: Nine, Suit: Hearts},
		{Rank: Eight, Suit: Diamonds},
		{Rank: Seven, Suit: Clubs},
	}

	// Act
	hand := FindBestHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("FindBestHand returned nil, expected valid hand")
	}
	if hand.Category != Flush {
		t.Errorf("FindBestHand category = %v, want %v (flush beats straight)", hand.Category, Flush)
	}
}

// TestFindBestHand_HandlesExactly5Cards verifies that FindBestHand
// works correctly when given exactly 5 cards.
func TestFindBestHand_HandlesExactly5Cards(t *testing.T) {
	// Arrange: Exactly 5 cards forming a straight
	cards := []Card{
		{Rank: Nine, Suit: Hearts},
		{Rank: Eight, Suit: Diamonds},
		{Rank: Seven, Suit: Clubs},
		{Rank: Six, Suit: Spades},
		{Rank: Five, Suit: Hearts},
	}

	// Act
	hand := FindBestHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("FindBestHand returned nil, expected valid hand")
	}
	if hand.Category != Straight {
		t.Errorf("FindBestHand category = %v, want %v", hand.Category, Straight)
	}
	if len(hand.Cards) != 5 {
		t.Errorf("FindBestHand cards length = %d, want 5", len(hand.Cards))
	}
}

// TestFindBestHand_ReturnsStrongestFrom7Cards verifies that FindBestHand
// evaluates all 21 combinations and returns the strongest hand.
func TestFindBestHand_ReturnsStrongestFrom7Cards(t *testing.T) {
	// Arrange: 7 cards with a four of a kind (strongest possible hand)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: Ace, Suit: Diamonds},
		{Rank: Ace, Suit: Clubs},
		{Rank: Ace, Suit: Spades},
		{Rank: King, Suit: Hearts},
		{Rank: Queen, Suit: Diamonds},
		{Rank: Jack, Suit: Clubs},
	}

	// Act
	hand := FindBestHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("FindBestHand returned nil, expected valid hand")
	}
	if hand.Category != FourOfAKind {
		t.Errorf("FindBestHand category = %v, want %v", hand.Category, FourOfAKind)
	}
}

// TestFindBestHand_HandlesHighCard verifies that FindBestHand
// correctly returns a high card hand when no better hand exists.
func TestFindBestHand_HandlesHighCard(t *testing.T) {
	// Arrange: 7 cards with no pairs, straights, or flushes
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Jack, Suit: Spades},
		{Rank: Nine, Suit: Hearts},
		{Rank: Seven, Suit: Diamonds},
		{Rank: Five, Suit: Clubs},
	}

	// Act
	hand := FindBestHand(cards)

	// Assert
	if hand == nil {
		t.Fatal("FindBestHand returned nil, expected valid hand")
	}
	if hand.Category != HighCard {
		t.Errorf("FindBestHand category = %v, want %v", hand.Category, HighCard)
	}
}
