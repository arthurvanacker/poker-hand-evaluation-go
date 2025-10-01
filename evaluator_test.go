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
