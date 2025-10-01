// Package main implements a Texas Hold'em poker hand evaluator
// that identifies the best 5-card hand from 5, 6, or 7 cards.
package main

// HandCategory represents the category of a poker hand, ordered by strength.
// Higher values indicate stronger hands.
type HandCategory int

const (
	HighCard      HandCategory = 1  // No matching cards
	OnePair       HandCategory = 2  // Two cards of the same rank
	TwoPair       HandCategory = 3  // Two different pairs
	ThreeOfAKind  HandCategory = 4  // Three cards of the same rank
	Straight      HandCategory = 5  // Five cards in sequence
	Flush         HandCategory = 6  // Five cards of the same suit
	FullHouse     HandCategory = 7  // Three of a kind plus a pair
	FourOfAKind   HandCategory = 8  // Four cards of the same rank
	StraightFlush HandCategory = 9  // Straight with all cards the same suit
	RoyalFlush    HandCategory = 10 // Ace-high straight flush (10-J-Q-K-A)
)

// String returns the human-readable name of the hand category.
func (hc HandCategory) String() string {
	switch hc {
	case HighCard:
		return "High Card"
	case OnePair:
		return "One Pair"
	case TwoPair:
		return "Two Pair"
	case ThreeOfAKind:
		return "Three of a Kind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of a Kind"
	case StraightFlush:
		return "Straight Flush"
	case RoyalFlush:
		return "Royal Flush"
	default:
		return "Unknown"
	}
}
