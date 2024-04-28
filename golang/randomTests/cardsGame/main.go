package main

import (
	"fmt"
	"math/rand"
)

/*
REQUEST:

- Given a deck of cards, distribute 5 random cards (without repetition of value+suit)
- Example:
	"Ace of Spades"
	"2 of Clubs"
	"8 of Diamonds"
	"5 of Spades"
	"Queen of Hearts"

Notes:
- a deck has 13 cards: Ace, 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King
- a deck has 4 suits: Clubs, Diamonds, Hearts, and Spades

RESOLUTION:

1. create a map for the deck of cards, where the key represents the card value+suit and the value is not relevant
2. randomly pick a card from 1 to 13 and assign it to the proper card name (from "Ace" to "King")
3. randomly pick a suit from 1 to 4 and assign it the proper suit name
4. if the selected card value+suit is already part of the deck of cards map, select a new card value+suit
5. else, add the selected card value+suit as a new key to the deck of cards map

*/

func main() {

	selectedCards := make(map[string]string)

	cardNames := map[int]string{
		1:  "Ace",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "10",
		11: "Jack",
		12: "Queen",
		13: "King",
	}

	cardSuits := map[int]string{
		1: "Clubs",
		2: "Diamonds",
		3: "Hearts",
		4: "Spades",
	}

	for i := 0; i < 5; {

		randomCard := rand.Intn(13) + 1 // from 1 to 13
		randomSuit := rand.Intn(4) + 1  // from 1 to 4
		candidateCard := fmt.Sprintf("%s of %s", cardNames[randomCard], cardSuits[randomSuit])

		if _, ok := selectedCards[candidateCard]; ok {
			continue // card already selected => select another card
		}

		selectedCards[candidateCard] = ""
		i++
	}

	fmt.Printf("The selected cards are:\n\n")

	i := 1
	for key := range selectedCards {
		fmt.Printf("%v: %v\n", i, key)
		i++
	}
}
