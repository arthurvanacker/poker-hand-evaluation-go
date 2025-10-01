// Package main implements a Texas Hold'em poker hand evaluator
// that identifies the best 5-card hand from 5, 6, or 7 cards.
package main

import "fmt"

// Deck represents a collection of playing cards
type Deck struct {
	Cards []Card
}

// NewDeck creates and returns a new deck containing all 52 standard playing cards
func NewDeck() *Deck {
	deck := &Deck{
		Cards: make([]Card, 0, 52),
	}

	// Generate all combinations of ranks and suits
	ranks := []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	suits := []Suit{Hearts, Diamonds, Clubs, Spades}

	for _, suit := range suits {
		for _, rank := range ranks {
			deck.Cards = append(deck.Cards, Card{Rank: rank, Suit: suit})
		}
	}

	return deck
}

// Deal removes and returns the top n cards from the deck.
// Returns an error if n is greater than the number of available cards.
func (d *Deck) Deal(n int) ([]Card, error) {
	if n > len(d.Cards) {
		return nil, fmt.Errorf("cannot deal %d cards, only %d available", n, len(d.Cards))
	}

	// Take n cards from the top of the deck
	dealt := d.Cards[:n]

	// Remove those cards from the deck
	d.Cards = d.Cards[n:]

	return dealt, nil
}
