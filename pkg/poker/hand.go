package poker

import "fmt"

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

// Hand represents a poker hand with its cards, category, and tiebreakers.
// Tiebreakers are ranks in descending order of importance for comparing
// hands of the same category.
type Hand struct {
	Cards       []Card       // The 5 cards in the hand
	Category    HandCategory // The hand category (Royal Flush, etc.)
	Tiebreakers []Rank       // Ranks for tiebreaker comparison
}

// NewHand creates a new Hand from the given cards.
// Returns an error if the number of cards is not exactly 5.
func NewHand(cards []Card) (*Hand, error) {
	if len(cards) != 5 {
		return nil, fmt.Errorf("hand must contain exactly 5 cards, got %d", len(cards))
	}

	// Create a copy of the cards slice to avoid external modification
	cardsCopy := make([]Card, 5)
	copy(cardsCopy, cards)

	return &Hand{
		Cards:       cardsCopy,
		Category:    HighCard, // Default category, will be set by evaluator
		Tiebreakers: []Rank{}, // Will be populated by evaluator
	}, nil
}
