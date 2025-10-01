// Package main implements a Texas Hold'em poker hand evaluator
// that identifies the best 5-card hand from 5, 6, or 7 cards.
package main

// isFlush checks if all 5 cards are the same suit.
// Returns true if all cards share the same suit, false otherwise.
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

// isStraight checks if 5 cards form a sequence.
// Returns (true, highRank) if cards form a straight, (false, 0) otherwise.
// Special case: wheel straight (A-2-3-4-5) returns (true, Five).
func isStraight(cards []Card) (bool, Rank) {
	if len(cards) != 5 {
		return false, 0
	}

	// Extract and sort ranks in descending order
	ranks := make([]Rank, 5)
	for i, card := range cards {
		ranks[i] = card.Rank
	}

	// Bubble sort descending (simple for 5 elements)
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if ranks[i] < ranks[j] {
				ranks[i], ranks[j] = ranks[j], ranks[i]
			}
		}
	}

	// Check for wheel straight: A-2-3-4-5 (14-5-4-3-2 when sorted descending)
	if ranks[0] == Ace && ranks[1] == Five && ranks[2] == Four && ranks[3] == Three && ranks[4] == Two {
		return true, Five // Ace acts as low, high card is Five
	}

	// Check for regular straight: each rank should be exactly 1 less than previous
	for i := 1; i < 5; i++ {
		if ranks[i] != ranks[i-1]-1 {
			return false, 0
		}
	}

	// Regular straight found, highest rank is first element
	return true, ranks[0]
}

// rankCounts counts how many cards of each rank exist in the hand.
// Returns a map where keys are Rank values and values are occurrence counts.
// Used for detecting pairs, trips, quads, and full houses.
func rankCounts(cards []Card) map[Rank]int {
	counts := make(map[Rank]int)

	for _, card := range cards {
		counts[card.Rank]++
	}

	return counts
}

// detectRoyalFlush checks if the given 5 cards form a royal flush.
// A royal flush is 10-J-Q-K-A all of the same suit.
// Returns true if the hand is a royal flush, false otherwise.
func detectRoyalFlush(cards []Card) bool {
	if len(cards) != 5 {
		return false
	}

	// Check if all cards are the same suit
	if !isFlush(cards) {
		return false
	}

	// Check for the exact ranks: 10, J, Q, K, A (Ten=10, Jack=11, Queen=12, King=13, Ace=14)
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
			return false
		}
		// Remove the rank to ensure no duplicates
		delete(requiredRanks, card.Rank)
	}

	// All required ranks should be consumed (map should be empty)
	return len(requiredRanks) == 0
}

// detectStraightFlush checks if the given 5 cards form a straight flush.
// A straight flush is 5 sequential cards all of the same suit (excluding royal flush).
// Returns true and the high card rank if a straight flush is detected.
// For the wheel straight flush (Ah-2h-3h-4h-5h), returns Five as the high card.
func detectStraightFlush(cards []Card) (bool, Rank) {
	// Check if all cards are the same suit
	if !isFlush(cards) {
		return false, 0
	}

	// Check if cards form a straight sequence
	isStraight, highCard := isStraight(cards)
	if !isStraight {
		return false, 0
	}

	// Both conditions met: it's a straight flush
	return true, highCard
}

// detectFourOfAKind checks if the given 5 cards contain four of a kind.
// Returns true and tiebreakers [quad rank, kicker] if four of a kind is found.
// Returns false and empty slice if no four of a kind exists.
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
