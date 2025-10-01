// Package main implements a Texas Hold'em poker hand evaluator
// that identifies the best 5-card hand from 5, 6, or 7 cards.
package main

import "sort"

// isFlush checks if all 5 cards are the same suit.
// Returns true if all cards share the same suit, false otherwise.
func isFlush(cards []Card) bool {
	if len(cards) != 5 {
		return false
	}

	// Use the first card's suit as reference
	firstSuit := cards[0].Suit

	// Check if all remaining cards match the first suit
	for i := 1; i < len(cards); i++ {
		if cards[i].Suit != firstSuit {
			return false
		}
	}

	return true
}

// isStraight checks if 5 cards form a sequence.
// Returns (true, highRank) if cards form a straight, (false, 0) otherwise.
// Special case: wheel straight (A-2-3-4-5) returns (true, Five).
func isStraight(cards []Card) (bool, Rank) {
	if len(cards) != 5 {
		return false, 0
	}

	// Extract and sort ranks in descending order
	ranks := make([]Rank, 5)
	for i, card := range cards {
		ranks[i] = card.Rank
	}

	// Bubble sort descending (simple for 5 elements)
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if ranks[i] < ranks[j] {
				ranks[i], ranks[j] = ranks[j], ranks[i]
			}
		}
	}

	// Check for wheel straight: A-2-3-4-5 (14-5-4-3-2 when sorted descending)
	if ranks[0] == Ace && ranks[1] == Five && ranks[2] == Four && ranks[3] == Three && ranks[4] == Two {
		return true, Five // Ace acts as low, high card is Five
	}

	// Check for regular straight: each rank should be exactly 1 less than previous
	for i := 1; i < 5; i++ {
		if ranks[i] != ranks[i-1]-1 {
			return false, 0
		}
	}

	// Regular straight found, highest rank is first element
	return true, ranks[0]
}

// rankCounts counts how many cards of each rank exist in the hand.
// Returns a map where keys are Rank values and values are occurrence counts.
// Used for detecting pairs, trips, quads, and full houses.
func rankCounts(cards []Card) map[Rank]int {
	counts := make(map[Rank]int)

	for _, card := range cards {
		counts[card.Rank]++
	}

	return counts
}

// detectRoyalFlush checks if the given 5 cards form a royal flush.
// A royal flush is 10-J-Q-K-A all of the same suit.
// Returns true if the hand is a royal flush, false otherwise.
func detectRoyalFlush(cards []Card) bool {
	if len(cards) != 5 {
		return false
	}

	// Check if all cards are the same suit
	if !isFlush(cards) {
		return false
	}

	// Check for the exact ranks: 10, J, Q, K, A (Ten=10, Jack=11, Queen=12, King=13, Ace=14)
	requiredRanks := map[Rank]bool{
		Ten:   true,
		Jack:  true,
		Queen: true,
		King:  true,
		Ace:   true,
	}

	// Verify all required ranks are present
	for _, card := range cards {
		if !requiredRanks[card.Rank] {
			return false
		}
		// Remove the rank to ensure no duplicates
		delete(requiredRanks, card.Rank)
	}

	// All required ranks should be consumed (map should be empty)
	return len(requiredRanks) == 0
}

// detectStraightFlush checks if the given 5 cards form a straight flush.
// A straight flush is 5 sequential cards all of the same suit (excluding royal flush).
// Returns true and the high card rank if a straight flush is detected.
// For the wheel straight flush (Ah-2h-3h-4h-5h), returns Five as the high card.
func detectStraightFlush(cards []Card) (bool, Rank) {
	// Check if all cards are the same suit
	if !isFlush(cards) {
		return false, 0
	}

	// Check if cards form a straight sequence
	isStraight, highCard := isStraight(cards)
	if !isStraight {
		return false, 0
	}

	// Both conditions met: it's a straight flush
	return true, highCard
}

// detectFourOfAKind checks if the given 5 cards contain four of a kind.
// Returns true and tiebreakers [quad rank, kicker] if four of a kind is found.
// Returns false and empty slice if no four of a kind exists.
func detectFourOfAKind(cards []Card) (bool, []Rank) {
	if len(cards) != 5 {
		return false, []Rank{}
	}

	counts := rankCounts(cards)

	var quadRank Rank
	var kicker Rank

	// Find the rank that appears 4 times
	for rank, count := range counts {
		if count == 4 {
			quadRank = rank
		} else if count == 1 {
			kicker = rank
		}
	}

	// If we found a quad rank, we have four of a kind
	if quadRank != 0 {
		return true, []Rank{quadRank, kicker}
	}

	return false, []Rank{}
}

// detectFullHouse checks if the given 5 cards contain a full house.
// A full house is three of a kind plus a pair.
// Returns true and tiebreakers [trip rank, pair rank] if full house is found.
// Returns false and empty slice if no full house exists.
func detectFullHouse(cards []Card) (bool, []Rank) {
	if len(cards) != 5 {
		return false, []Rank{}
	}

	counts := rankCounts(cards)

	var tripRank Rank
	var pairRank Rank

	// Find the rank that appears 3 times and 2 times
	for rank, count := range counts {
		if count == 3 {
			tripRank = rank
		} else if count == 2 {
			pairRank = rank
		}
	}

	// If we found both trip and pair, we have a full house
	if tripRank != 0 && pairRank != 0 {
		return true, []Rank{tripRank, pairRank}
	}

	return false, []Rank{}
}

// detectFlush detects a flush (5 suited cards, not sequential)
// Returns true and ranks in descending order if flush, false and nil otherwise
// Returns false for straight flushes (they are handled separately)
func detectFlush(cards []Card) (bool, []Rank) {
	// Must be a flush
	if !isFlush(cards) {
		return false, nil
	}

	// Must NOT be a straight (to exclude straight flushes)
	if isStraight, _ := isStraight(cards); isStraight {
		return false, nil
	}

	// Extract ranks and sort in descending order
	ranks := make([]Rank, len(cards))
	for i, card := range cards {
		ranks[i] = card.Rank
	}
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i] > ranks[j]
	})

	return true, ranks
}

// detectStraight checks if the given 5 cards form a straight (5 sequential cards with mixed suits).
// Returns (true, highRank) if cards form a non-flush straight, (false, 0) otherwise.
// Returns false for straight flushes (those should be detected separately).
// Special case: wheel straight (A-2-3-4-5) returns (true, Five).
func detectStraight(cards []Card) (bool, Rank) {
	if len(cards) != 5 {
		return false, 0
	}

	// Check if it's a flush - if so, it's not a plain straight
	if isFlush(cards) {
		return false, 0
	}

	// Use isStraight helper to check for sequential ranks
	return isStraight(cards)
}

// detectThreeOfAKind checks if the given 5 cards contain three of a kind.
// Returns true and tiebreakers [trip rank, kicker1, kicker2] if three of a kind is found.
// Returns false and empty slice if no three of a kind exists.
// Full houses (three of a kind + pair) return false as they should be detected separately.
func detectThreeOfAKind(cards []Card) (bool, []Rank) {
	if len(cards) != 5 {
		return false, []Rank{}
	}

	counts := rankCounts(cards)

	var tripRank Rank
	var kickers []Rank

	// Find the rank that appears exactly 3 times
	for rank, count := range counts {
		if count == 3 {
			tripRank = rank
		} else if count == 1 {
			kickers = append(kickers, rank)
		}
	}

	// Must have exactly one trip and exactly two kickers (excludes full houses)
	if tripRank != 0 && len(kickers) == 2 {
		// Sort kickers in descending order
		if kickers[0] < kickers[1] {
			kickers[0], kickers[1] = kickers[1], kickers[0]
		}

		return true, []Rank{tripRank, kickers[0], kickers[1]}
	}

	return false, []Rank{}
}

// detectTwoPair checks if the given 5 cards contain exactly two pairs.
// Returns (true, [high pair, low pair, kicker]) if found,
// or (false, nil) if not two pair or if it's a full house.
func detectTwoPair(cards []Card) (bool, []Rank) {
	if len(cards) != 5 {
		return false, nil
	}

	counts := rankCounts(cards)

	// Find ranks with exactly 2 cards and ranks with 3+ cards
	pairs := make([]Rank, 0, 2)
	hasTrips := false

	for rank, count := range counts {
		if count == 2 {
			pairs = append(pairs, rank)
		} else if count >= 3 {
			hasTrips = true
		}
	}

	// Must have exactly two pairs and no trips/quads (if trips exist, it's a full house)
	if len(pairs) != 2 || hasTrips {
		return false, nil
	}

	// Sort pairs descending (high pair first)
	if pairs[0] < pairs[1] {
		pairs[0], pairs[1] = pairs[1], pairs[0]
	}

	// Find the kicker (the single remaining card)
	var kicker Rank
	for rank, count := range counts {
		if count == 1 {
			kicker = rank
			break
		}
	}

	// Return [high pair, low pair, kicker]
	return true, []Rank{pairs[0], pairs[1], kicker}
}

// detectOnePair checks if the given 5 cards contain exactly one pair.
// Returns (true, tiebreakers) if exactly one pair exists, where tiebreakers
// contains [pair rank, kicker1, kicker2, kicker3] in descending order.
// Returns (false, nil) if no pair, two pairs, trips, or quads are present.
func detectOnePair(cards []Card) (bool, []Rank) {
	if len(cards) != 5 {
		return false, nil
	}

	counts := rankCounts(cards)

	// Count how many ranks appear exactly twice (pairs)
	pairCount := 0
	var pairRank Rank
	var kickers []Rank

	for rank, count := range counts {
		if count == 2 {
			pairCount++
			pairRank = rank
		} else if count == 1 {
			kickers = append(kickers, rank)
		} else if count > 2 {
			// Trips or quads detected, not one pair
			return false, nil
		}
	}

	// Must have exactly one pair (not zero, not two)
	if pairCount != 1 {
		return false, nil
	}

	// Sort kickers in descending order
	for i := 0; i < len(kickers); i++ {
		for j := i + 1; j < len(kickers); j++ {
			if kickers[i] < kickers[j] {
				kickers[i], kickers[j] = kickers[j], kickers[i]
			}
		}
	}

	// Build tiebreakers: [pair rank, kicker1, kicker2, kicker3]
	tiebreakers := []Rank{pairRank}
	tiebreakers = append(tiebreakers, kickers...)

	return true, tiebreakers
}

// detectHighCard checks if the given 5 cards form a high card hand.
// Returns true (always, as it's the fallback category) and all 5 ranks in descending order as tiebreakers.
func detectHighCard(cards []Card) (bool, []Rank) {
	if len(cards) != 5 {
		return false, []Rank{}
	}

	// Extract all ranks
	ranks := make([]Rank, 5)
	for i, card := range cards {
		ranks[i] = card.Rank
	}

	// Sort ranks in descending order (bubble sort for 5 elements)
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if ranks[i] < ranks[j] {
				ranks[i], ranks[j] = ranks[j], ranks[i]
			}
		}
	}

	// Always returns true with all 5 ranks as tiebreakers
	return true, ranks
}

// EvaluateHand evaluates a 5-card poker hand and returns the best hand category with tiebreakers.
// Checks hand categories from strongest (Royal Flush) to weakest (High Card).
// Returns nil if the input is not exactly 5 cards.
func EvaluateHand(cards []Card) *Hand {
	if len(cards) != 5 {
		return nil
	}

	// Check Royal Flush
	if detectRoyalFlush(cards) {
		return &Hand{
			Cards:       cards,
			Category:    RoyalFlush,
			Tiebreakers: []Rank{}, // Royal flush has no tiebreakers (all equal)
		}
	}

	// Check Straight Flush
	if found, highCard := detectStraightFlush(cards); found {
		return &Hand{
			Cards:       cards,
			Category:    StraightFlush,
			Tiebreakers: []Rank{highCard},
		}
	}

	// Check Four of a Kind
	if found, tiebreakers := detectFourOfAKind(cards); found {
		return &Hand{
			Cards:       cards,
			Category:    FourOfAKind,
			Tiebreakers: tiebreakers,
		}
	}

	// Check Full House
	if found, tiebreakers := detectFullHouse(cards); found {
		return &Hand{
			Cards:       cards,
			Category:    FullHouse,
			Tiebreakers: tiebreakers,
		}
	}

	// Check Flush
	if found, tiebreakers := detectFlush(cards); found {
		return &Hand{
			Cards:       cards,
			Category:    Flush,
			Tiebreakers: tiebreakers,
		}
	}

	// Check Straight
	if found, highCard := detectStraight(cards); found {
		return &Hand{
			Cards:       cards,
			Category:    Straight,
			Tiebreakers: []Rank{highCard},
		}
	}

	// Check Three of a Kind
	if found, tiebreakers := detectThreeOfAKind(cards); found {
		return &Hand{
			Cards:       cards,
			Category:    ThreeOfAKind,
			Tiebreakers: tiebreakers,
		}
	}

	// Check Two Pair
	if found, tiebreakers := detectTwoPair(cards); found {
		return &Hand{
			Cards:       cards,
			Category:    TwoPair,
			Tiebreakers: tiebreakers,
		}
	}

	// Check One Pair
	if found, tiebreakers := detectOnePair(cards); found {
		return &Hand{
			Cards:       cards,
			Category:    OnePair,
			Tiebreakers: tiebreakers,
		}
	}

	// Default to High Card (always succeeds)
	_, tiebreakers := detectHighCard(cards)
	return &Hand{
		Cards:       cards,
		Category:    HighCard,
		Tiebreakers: tiebreakers,
	}
}

// CompareHands compares two poker hands and returns:
// 1 if hand1 is stronger, -1 if hand2 is stronger, 0 if equal.
// Compares by category first, then by tiebreakers element-by-element.
func CompareHands(hand1, hand2 *Hand) int {
	if hand1 == nil || hand2 == nil {
		return 0
	}

	// Compare by category first (higher category wins)
	if hand1.Category > hand2.Category {
		return 1
	}
	if hand1.Category < hand2.Category {
		return -1
	}

	// Same category - compare tiebreakers element-by-element
	minLen := len(hand1.Tiebreakers)
	if len(hand2.Tiebreakers) < minLen {
		minLen = len(hand2.Tiebreakers)
	}

	for i := 0; i < minLen; i++ {
		if hand1.Tiebreakers[i] > hand2.Tiebreakers[i] {
			return 1
		}
		if hand1.Tiebreakers[i] < hand2.Tiebreakers[i] {
			return -1
		}
	}

	// All tiebreakers equal
	return 0
}

// FindBestHand finds the best 5-card poker hand from 5, 6, or 7 cards.
// Generates all possible 5-card combinations, evaluates each, and returns the strongest.
// Returns nil if fewer than 5 cards are provided.
func FindBestHand(cards []Card) *Hand {
	if len(cards) < 5 {
		return nil
	}

	// Optimization: if exactly 5 cards, evaluate directly
	if len(cards) == 5 {
		return EvaluateHand(cards)
	}

	// Generate all 5-card combinations
	combinations := Combinations(cards, 5)

	var bestHand *Hand

	// Evaluate each combination and keep the best
	for _, combo := range combinations {
		hand := EvaluateHand(combo)
		if bestHand == nil || CompareHands(hand, bestHand) > 0 {
			bestHand = hand
		}
	}

	return bestHand
}
