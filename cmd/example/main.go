package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/puupa/poker-hand-evaluation/pkg/poker"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Texas Hold'em Hand Evaluator Demo ===")
	fmt.Println()

	// Create a new deck
	deck := poker.NewDeck()
	fmt.Printf("Created new deck with %d cards\n\n", len(deck.Cards))

	// Shuffle the deck
	shuffleDeck(deck)
	fmt.Println("Deck shuffled")
	fmt.Println()

	// Deal cards for Player 1
	player1Cards, err := deck.Deal(2)
	if err != nil {
		fmt.Printf("Error dealing cards: %v\n", err)
		return
	}
	fmt.Printf("Player 1 hole cards: %s %s\n", player1Cards[0], player1Cards[1])

	// Deal cards for Player 2
	player2Cards, err := deck.Deal(2)
	if err != nil {
		fmt.Printf("Error dealing cards: %v\n", err)
		return
	}
	fmt.Printf("Player 2 hole cards: %s %s\n\n", player2Cards[0], player2Cards[1])

	// Deal community cards (flop, turn, river)
	communityCards, err := deck.Deal(5)
	if err != nil {
		fmt.Printf("Error dealing community cards: %v\n", err)
		return
	}
	fmt.Printf("Community cards: %s %s %s %s %s\n\n",
		communityCards[0], communityCards[1], communityCards[2],
		communityCards[3], communityCards[4])

	// Combine hole cards + community cards for each player
	player1Hand := make([]poker.Card, 0, 7)
	player1Hand = append(player1Hand, player1Cards...)
	player1Hand = append(player1Hand, communityCards...)

	player2Hand := make([]poker.Card, 0, 7)
	player2Hand = append(player2Hand, player2Cards...)
	player2Hand = append(player2Hand, communityCards...)

	fmt.Printf("Player 1's 7 cards: ")
	for _, card := range player1Hand {
		fmt.Printf("%s ", card)
	}
	fmt.Println()

	fmt.Printf("Player 2's 7 cards: ")
	for _, card := range player2Hand {
		fmt.Printf("%s ", card)
	}
	fmt.Println()
	fmt.Println()

	// Evaluate best hands using FindBestHand
	fmt.Println("=== Hand Analysis ===")
	best1 := poker.FindBestHand(player1Hand)
	best2 := poker.FindBestHand(player2Hand)

	fmt.Printf("Player 1's best hand: %s\n", best1.Category)
	fmt.Printf("Player 2's best hand: %s\n", best2.Category)

	// Compare hands and determine winner
	result := poker.CompareHands(best1, best2)
	fmt.Println()
	if result > 0 {
		fmt.Println("🏆 Player 1 WINS!")
	} else if result < 0 {
		fmt.Println("🏆 Player 2 WINS!")
	} else {
		fmt.Println("🤝 SPLIT POT - Hands are equal!")
	}

	fmt.Printf("\nRemaining cards in deck: %d\n", len(deck.Cards))
}

// shuffleDeck shuffles the deck using Fisher-Yates algorithm
func shuffleDeck(d *poker.Deck) {
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
