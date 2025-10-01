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
