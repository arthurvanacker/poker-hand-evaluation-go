package main

import "testing"

// Test that HandCategory enum has all 10 values
func TestHandCategoryValues(t *testing.T) {
	tests := []struct {
		name     string
		category HandCategory
		expected int
	}{
		{"HighCard", HighCard, 1},
		{"OnePair", OnePair, 2},
		{"TwoPair", TwoPair, 3},
		{"ThreeOfAKind", ThreeOfAKind, 4},
		{"Straight", Straight, 5},
		{"Flush", Flush, 6},
		{"FullHouse", FullHouse, 7},
		{"FourOfAKind", FourOfAKind, 8},
		{"StraightFlush", StraightFlush, 9},
		{"RoyalFlush", RoyalFlush, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.category) != tt.expected {
				t.Errorf("%s should have value %d, got %d", tt.name, tt.expected, int(tt.category))
			}
		})
	}
}

// Test that categories are ordered by strength
func TestHandCategoryOrdering(t *testing.T) {
	if HighCard >= OnePair {
		t.Error("HighCard should be less than OnePair")
	}
	if OnePair >= TwoPair {
		t.Error("OnePair should be less than TwoPair")
	}
	if TwoPair >= ThreeOfAKind {
		t.Error("TwoPair should be less than ThreeOfAKind")
	}
	if ThreeOfAKind >= Straight {
		t.Error("ThreeOfAKind should be less than Straight")
	}
	if Straight >= Flush {
		t.Error("Straight should be less than Flush")
	}
	if Flush >= FullHouse {
		t.Error("Flush should be less than FullHouse")
	}
	if FullHouse >= FourOfAKind {
		t.Error("FullHouse should be less than FourOfAKind")
	}
	if FourOfAKind >= StraightFlush {
		t.Error("FourOfAKind should be less than StraightFlush")
	}
	if StraightFlush >= RoyalFlush {
		t.Error("StraightFlush should be less than RoyalFlush")
	}
}

// Test that RoyalFlush is the strongest hand
func TestRoyalFlushIsStrongest(t *testing.T) {
	categories := []HandCategory{
		HighCard, OnePair, TwoPair, ThreeOfAKind, Straight,
		Flush, FullHouse, FourOfAKind, StraightFlush,
	}

	for _, category := range categories {
		if category >= RoyalFlush {
			t.Errorf("%v should be weaker than RoyalFlush", category)
		}
	}
}

// Test that HighCard is the weakest hand
func TestHighCardIsWeakest(t *testing.T) {
	categories := []HandCategory{
		OnePair, TwoPair, ThreeOfAKind, Straight, Flush,
		FullHouse, FourOfAKind, StraightFlush, RoyalFlush,
	}

	for _, category := range categories {
		if category <= HighCard {
			t.Errorf("%v should be stronger than HighCard", category)
		}
	}
}

// Test HandCategory String() method
func TestHandCategoryString(t *testing.T) {
	tests := []struct {
		category HandCategory
		expected string
	}{
		{HighCard, "High Card"},
		{OnePair, "One Pair"},
		{TwoPair, "Two Pair"},
		{ThreeOfAKind, "Three of a Kind"},
		{Straight, "Straight"},
		{Flush, "Flush"},
		{FullHouse, "Full House"},
		{FourOfAKind, "Four of a Kind"},
		{StraightFlush, "Straight Flush"},
		{RoyalFlush, "Royal Flush"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if tt.category.String() != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.category.String())
			}
		})
	}
}
