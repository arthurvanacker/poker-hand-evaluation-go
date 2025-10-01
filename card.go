package main

// Rank represents the rank of a playing card (2-14, where Ace=14)
type Rank int

// Rank constants with numeric values for easy comparison
const (
	Two   Rank = 2
	Three Rank = 3
	Four  Rank = 4
	Five  Rank = 5
	Six   Rank = 6
	Seven Rank = 7
	Eight Rank = 8
	Nine  Rank = 9
	Ten   Rank = 10
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
	Ace   Rank = 14
)

// Suit represents the suit of a playing card
type Suit int

// Suit constants
const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

// Card represents a playing card with a rank and suit
type Card struct {
	Rank Rank
	Suit Suit
}

// String returns the string representation of a Rank (e.g., "A", "K", "Q", "J", "T", "9", ...)
func (r Rank) String() string {
	switch r {
	case Ace:
		return "A"
	case King:
		return "K"
	case Queen:
		return "Q"
	case Jack:
		return "J"
	case Ten:
		return "T"
	case Nine:
		return "9"
	case Eight:
		return "8"
	case Seven:
		return "7"
	case Six:
		return "6"
	case Five:
		return "5"
	case Four:
		return "4"
	case Three:
		return "3"
	case Two:
		return "2"
	default:
		return "?"
	}
}

// String returns the string representation of a Suit (e.g., "h", "d", "c", "s")
func (s Suit) String() string {
	switch s {
	case Hearts:
		return "h"
	case Diamonds:
		return "d"
	case Clubs:
		return "c"
	case Spades:
		return "s"
	default:
		return "?"
	}
}

// String returns the card notation (e.g., "Ah" for Ace of Hearts)
func (c Card) String() string {
	return c.Rank.String() + c.Suit.String()
}

