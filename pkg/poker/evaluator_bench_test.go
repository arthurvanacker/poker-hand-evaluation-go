package poker

import (
	"testing"
)

// BenchmarkEvaluateHand measures the performance of evaluating a single 5-card hand.
// Uses a realistic mid-strength hand (pair of kings) to avoid bias.
func BenchmarkEvaluateHand(b *testing.B) {
	// Realistic test case: pair of kings with queen, jack, nine kickers
	cards := []Card{
		{Rank: King, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Jack, Suit: Spades},
		{Rank: Nine, Suit: Hearts},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EvaluateHand(cards)
	}
}

// BenchmarkFindBestHand5Cards measures the baseline performance with exactly 5 cards.
// This is the fastest case as no combinations need to be generated.
func BenchmarkFindBestHand5Cards(b *testing.B) {
	// Realistic test case: flush (5 hearts with mixed ranks)
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: Ten, Suit: Hearts},
		{Rank: Seven, Suit: Hearts},
		{Rank: Five, Suit: Hearts},
		{Rank: Three, Suit: Hearts},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FindBestHand(cards)
	}
}

// BenchmarkFindBestHand6Cards measures performance with 6 cards (6 combinations).
// Represents typical scenarios like Omaha or early Texas Hold'em streets.
func BenchmarkFindBestHand6Cards(b *testing.B) {
	// Realistic test case: 6 cards with multiple possible hands
	// Best hand is a straight (9-8-7-6-5)
	cards := []Card{
		{Rank: Nine, Suit: Hearts},
		{Rank: Eight, Suit: Diamonds},
		{Rank: Seven, Suit: Clubs},
		{Rank: Six, Suit: Spades},
		{Rank: Five, Suit: Hearts},
		{Rank: Two, Suit: Diamonds},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FindBestHand(cards)
	}
}

// BenchmarkFindBestHand7Cards measures performance with 7 cards (21 combinations).
// This is the most common case in Texas Hold'em (2 hole cards + 5 community cards).
func BenchmarkFindBestHand7Cards(b *testing.B) {
	// Realistic test case: 7 cards with multiple possible hands
	// Best hand is three of a kind (three jacks)
	cards := []Card{
		{Rank: Jack, Suit: Hearts},
		{Rank: Jack, Suit: Diamonds},
		{Rank: Jack, Suit: Clubs},
		{Rank: King, Suit: Spades},
		{Rank: Queen, Suit: Hearts},
		{Rank: Eight, Suit: Diamonds},
		{Rank: Three, Suit: Clubs},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FindBestHand(cards)
	}
}

// BenchmarkCombinations measures the performance of generating 5-card combinations.
// Tests the core combination generation algorithm with 7 cards (21 combinations).
func BenchmarkCombinations(b *testing.B) {
	// Use 7 cards for realistic Texas Hold'em scenario
	cards := []Card{
		{Rank: Ace, Suit: Hearts},
		{Rank: King, Suit: Diamonds},
		{Rank: Queen, Suit: Clubs},
		{Rank: Jack, Suit: Spades},
		{Rank: Ten, Suit: Hearts},
		{Rank: Nine, Suit: Diamonds},
		{Rank: Eight, Suit: Clubs},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Combinations(cards, 5)
	}
}
