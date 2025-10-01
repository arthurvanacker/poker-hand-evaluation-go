package main

import "testing"

// Test that Rank constants exist and have correct numeric values
func TestRankValues(t *testing.T) {
	tests := []struct {
		rank     Rank
		expected int
		name     string
	}{
		{Two, 2, "Two"},
		{Three, 3, "Three"},
		{Four, 4, "Four"},
		{Five, 5, "Five"},
		{Six, 6, "Six"},
		{Seven, 7, "Seven"},
		{Eight, 8, "Eight"},
		{Nine, 9, "Nine"},
		{Ten, 10, "Ten"},
		{Jack, 11, "Jack"},
		{Queen, 12, "Queen"},
		{King, 13, "King"},
		{Ace, 14, "Ace"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.rank) != tt.expected {
				t.Errorf("%s should be %d, got %d", tt.name, tt.expected, int(tt.rank))
			}
		})
	}
}

// Test that Suit constants exist
func TestSuitConstants(t *testing.T) {
	suits := []struct {
		suit Suit
		name string
	}{
		{Hearts, "Hearts"},
		{Diamonds, "Diamonds"},
		{Clubs, "Clubs"},
		{Spades, "Spades"},
	}

	for _, s := range suits {
		t.Run(s.name, func(t *testing.T) {
			// Just verify the constant exists and is of type Suit
			var _ Suit = s.suit
		})
	}
}

// Test Card struct creation and field access
func TestCardCreation(t *testing.T) {
	tests := []struct {
		name string
		card Card
		rank Rank
		suit Suit
	}{
		{"Ace of Hearts", Card{Rank: Ace, Suit: Hearts}, Ace, Hearts},
		{"King of Diamonds", Card{Rank: King, Suit: Diamonds}, King, Diamonds},
		{"Two of Clubs", Card{Rank: Two, Suit: Clubs}, Two, Clubs},
		{"Ten of Spades", Card{Rank: Ten, Suit: Spades}, Ten, Spades},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.card.Rank != tt.rank {
				t.Errorf("Expected rank %v, got %v", tt.rank, tt.card.Rank)
			}
			if tt.card.Suit != tt.suit {
				t.Errorf("Expected suit %v, got %v", tt.suit, tt.card.Suit)
			}
		})
	}
}

// Test Rank.String() method
func TestRankString(t *testing.T) {
	tests := []struct {
		rank     Rank
		expected string
	}{
		{Two, "2"},
		{Three, "3"},
		{Four, "4"},
		{Five, "5"},
		{Six, "6"},
		{Seven, "7"},
		{Eight, "8"},
		{Nine, "9"},
		{Ten, "T"},
		{Jack, "J"},
		{Queen, "Q"},
		{King, "K"},
		{Ace, "A"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if tt.rank.String() != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.rank.String())
			}
		})
	}
}

// Test Suit.String() method
func TestSuitString(t *testing.T) {
	tests := []struct {
		suit     Suit
		expected string
	}{
		{Hearts, "h"},
		{Diamonds, "d"},
		{Clubs, "c"},
		{Spades, "s"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if tt.suit.String() != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.suit.String())
			}
		})
	}
}

// Test Card.String() method - full card notation
func TestCardString(t *testing.T) {
	tests := []struct {
		card     Card
		expected string
	}{
		{Card{Rank: Ace, Suit: Hearts}, "Ah"},
		{Card{Rank: King, Suit: Diamonds}, "Kd"},
		{Card{Rank: Queen, Suit: Clubs}, "Qc"},
		{Card{Rank: Jack, Suit: Spades}, "Js"},
		{Card{Rank: Ten, Suit: Hearts}, "Th"},
		{Card{Rank: Two, Suit: Clubs}, "2c"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if tt.card.String() != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.card.String())
			}
		})
	}
}

// Test ParseCard function - basic valid inputs
func TestParseCard(t *testing.T) {
	tests := []struct {
		input    string
		expected Card
	}{
		{"Ah", Card{Rank: Ace, Suit: Hearts}},
		{"Kd", Card{Rank: King, Suit: Diamonds}},
		{"Qc", Card{Rank: Queen, Suit: Clubs}},
		{"Js", Card{Rank: Jack, Suit: Spades}},
		{"Th", Card{Rank: Ten, Suit: Hearts}},
		{"9d", Card{Rank: Nine, Suit: Diamonds}},
		{"8c", Card{Rank: Eight, Suit: Clubs}},
		{"7s", Card{Rank: Seven, Suit: Spades}},
		{"6h", Card{Rank: Six, Suit: Hearts}},
		{"5d", Card{Rank: Five, Suit: Diamonds}},
		{"4c", Card{Rank: Four, Suit: Clubs}},
		{"3s", Card{Rank: Three, Suit: Spades}},
		{"2h", Card{Rank: Two, Suit: Hearts}},
		{"2c", Card{Rank: Two, Suit: Clubs}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseCard(tt.input)
			if err != nil {
				t.Errorf("ParseCard(%q) returned error: %v", tt.input, err)
			}
			if got != tt.expected {
				t.Errorf("ParseCard(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// Test ParseCard with "10" notation for Ten
func TestParseCardTenNotation(t *testing.T) {
	tests := []struct {
		input    string
		expected Card
	}{
		{"10h", Card{Rank: Ten, Suit: Hearts}},
		{"10d", Card{Rank: Ten, Suit: Diamonds}},
		{"10c", Card{Rank: Ten, Suit: Clubs}},
		{"10s", Card{Rank: Ten, Suit: Spades}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseCard(tt.input)
			if err != nil {
				t.Errorf("ParseCard(%q) returned error: %v", tt.input, err)
			}
			if got != tt.expected {
				t.Errorf("ParseCard(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// Test ParseCard with case-insensitive suits
func TestParseCardCaseInsensitive(t *testing.T) {
	tests := []struct {
		input    string
		expected Card
	}{
		{"AH", Card{Rank: Ace, Suit: Hearts}},
		{"aH", Card{Rank: Ace, Suit: Hearts}},
		{"Ah", Card{Rank: Ace, Suit: Hearts}},
		{"ah", Card{Rank: Ace, Suit: Hearts}},
		{"KD", Card{Rank: King, Suit: Diamonds}},
		{"kd", Card{Rank: King, Suit: Diamonds}},
		{"QC", Card{Rank: Queen, Suit: Clubs}},
		{"qc", Card{Rank: Queen, Suit: Clubs}},
		{"JS", Card{Rank: Jack, Suit: Spades}},
		{"js", Card{Rank: Jack, Suit: Spades}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseCard(tt.input)
			if err != nil {
				t.Errorf("ParseCard(%q) returned error: %v", tt.input, err)
			}
			if got != tt.expected {
				t.Errorf("ParseCard(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// Test ParseCard error cases
func TestParseCardErrors(t *testing.T) {
	tests := []struct {
		input       string
		description string
	}{
		{"", "empty string"},
		{"A", "missing suit"},
		{"h", "missing rank"},
		{"Ax", "invalid suit"},
		{"Xh", "invalid rank"},
		{"11h", "invalid rank (11)"},
		{"1h", "invalid rank (1)"},
		{"AAh", "too long"},
		{"Ahh", "too long"},
		{"10", "missing suit with 10"},
		{"10x", "invalid suit with 10"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got, err := ParseCard(tt.input)
			if err == nil {
				t.Errorf("ParseCard(%q) expected error for %s, got %v", tt.input, tt.description, got)
			}
		})
	}
}

// Test invalid Rank.String() returns "?" for unknown ranks
func TestRankStringInvalid(t *testing.T) {
	invalidRank := Rank(99)
	result := invalidRank.String()
	if result != "?" {
		t.Errorf("Expected '?' for invalid rank, got %q", result)
	}
}

// Test invalid Suit.String() returns "?" for unknown suits
func TestSuitStringInvalid(t *testing.T) {
	invalidSuit := Suit(99)
	result := invalidSuit.String()
	if result != "?" {
		t.Errorf("Expected '?' for invalid suit, got %q", result)
	}
}
