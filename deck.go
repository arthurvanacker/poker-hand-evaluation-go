package main

import "math/rand"

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

// Shuffle randomizes the order of cards in the deck using Fisher-Yates algorithm
func (d *Deck) Shuffle() {
	n := len(d.Cards)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
