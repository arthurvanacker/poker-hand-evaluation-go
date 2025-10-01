package main

import (
	"fmt"
	"strings"
)

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
	default:
		return string('0' + byte(r))
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

// ParseCard parses a card string (e.g., "Ah", "Kd", "10s") into a Card struct
func ParseCard(s string) (Card, error) {
	if len(s) < 2 || len(s) > 3 {
		return Card{}, fmt.Errorf("invalid card string: %q (must be 2-3 characters)", s)
	}

	// Parse rank (first 1 or 2 characters)
	var rank Rank
	var suitChar string

	if len(s) == 3 {
		// Could be "10s" format
		if s[:2] == "10" {
			rank = Ten
			suitChar = s[2:]
		} else {
			return Card{}, fmt.Errorf("invalid card string: %q", s)
		}
	} else {
		// 2-character format like "Ah" or "9d"
		rankChar := s[0]
		suitChar = s[1:]

		switch rankChar {
		case 'A', 'a':
			rank = Ace
		case 'K', 'k':
			rank = King
		case 'Q', 'q':
			rank = Queen
		case 'J', 'j':
			rank = Jack
		case 'T', 't':
			rank = Ten
		case '9':
			rank = Nine
		case '8':
			rank = Eight
		case '7':
			rank = Seven
		case '6':
			rank = Six
		case '5':
			rank = Five
		case '4':
			rank = Four
		case '3':
			rank = Three
		case '2':
			rank = Two
		default:
			return Card{}, fmt.Errorf("invalid rank: %q", rankChar)
		}
	}

	// Parse suit (last character, case-insensitive)
	var suit Suit
	switch strings.ToLower(suitChar) {
	case "h":
		suit = Hearts
	case "d":
		suit = Diamonds
	case "c":
		suit = Clubs
	case "s":
		suit = Spades
	default:
		return Card{}, fmt.Errorf("invalid suit: %q", suitChar)
	}

	return Card{Rank: rank, Suit: suit}, nil
}
