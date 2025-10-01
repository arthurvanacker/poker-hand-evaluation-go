package poker

// Combinations generates all k-card combinations from the given cards.
// Uses a recursive algorithm to generate all possible selections.
// For example, Combinations(7 cards, 5) returns 21 combinations (C(7,5) = 21).
func Combinations(cards []Card, k int) [][]Card {
	var result [][]Card

	// Base case: if k is 0, return empty combination
	if k == 0 {
		return [][]Card{{}}
	}

	// Base case: if k equals the number of cards, return all cards as one combination
	if k == len(cards) {
		return [][]Card{cards}
	}

	// Base case: if k > len(cards), no valid combinations
	if k > len(cards) {
		return result
	}

	// Recursive case: generate combinations
	// For each position, we either include the card at that position or skip it
	generate(cards, k, 0, []Card{}, &result)

	return result
}

// generate is a helper function that recursively builds combinations
// cards: the source cards to choose from
// k: number of cards still needed in the current combination
// start: starting index in cards array to avoid duplicates
// current: the combination being built
// result: pointer to the result slice to accumulate all combinations
func generate(cards []Card, k int, start int, current []Card, result *[][]Card) {
	// Base case: if we've selected k cards, add this combination to result
	if k == 0 {
		// Make a copy of current to avoid mutation issues
		combo := make([]Card, len(current))
		copy(combo, current)
		*result = append(*result, combo)
		return
	}

	// Recursive case: try including each remaining card
	for i := start; i <= len(cards)-k; i++ {
		// Include cards[i] in the current combination
		generate(cards, k-1, i+1, append(current, cards[i]), result)
	}
}
