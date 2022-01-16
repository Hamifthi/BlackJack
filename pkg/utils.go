package pkg

import (
	"Blackjack"
	"fmt"
	"github.com/Hamifthi/deck"
)

func DrawCard(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func DetermineWinner(player, dealer Blackjack.Player) string {
	switch {
	case player.Score > 21:
		return "Player busted."
	case dealer.Score > 21:
		return "Dealer busted."
	case player.Score > dealer.Score:
		return "Player wins."
	case dealer.Score > player.Score:
		return "Dealer wins."
	case player.Score == dealer.Score:
		return "draw."
	}
	return ""
}

func Clone(gs Blackjack.GameState) Blackjack.GameState {
	newGameState := Blackjack.GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: *new(Blackjack.Player),
		Dealer: *new(Blackjack.Player),
	}
	copy(gs.Deck, newGameState.Deck)
	copy(gs.Player.Cards, newGameState.Player.Cards)
	copy(gs.Dealer.Cards, newGameState.Dealer.Cards)
	return newGameState
}

func Shuffle(gs Blackjack.GameState) Blackjack.GameState {
	state := Clone(gs)
	state.Deck = deck.NewDeck(deck.MultipleDeck(3), deck.ShuffleDeck())
	return state
}

func Deal(gs Blackjack.GameState) Blackjack.GameState {
	state := Clone(gs)
	state.Player.Cards = make([]deck.Card, 0, 5)
	state.Dealer.Cards = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		if i == 0 {
			card, state.Deck = DrawCard(state.Deck)
			state.Dealer.HiddenCard = card
			card, state.Deck = DrawCard(state.Deck)
			state.Player.Cards = append(state.Player.Cards, card)
		} else {
			card, state.Deck = DrawCard(state.Deck)
			state.Player.Cards = append(state.Player.Cards, card)
			card, state.Deck = DrawCard(state.Deck)
			state.Dealer.Cards = append(state.Dealer.Cards, card)
		}
	}
	return state
}

func Stay(gs Blackjack.GameState) Blackjack.GameState {
	state := Clone(gs)
	state.State++
	return state
}

func Hit(gs Blackjack.GameState) Blackjack.GameState {
	state := Clone(gs)
	player := gs.CurrentPlayer()
	var card deck.Card
	card, state.Deck = DrawCard(state.Deck)
	player.Cards = append(player.Cards, card)
	player.CalculateScore()
	if player.Score > 21 {
		return Stay(state)
	}
	return state
}

func End(gs Blackjack.GameState) Blackjack.GameState {
	state := Clone(gs)
	gs.Player.CalculateScore()
	gs.Dealer.CalculateScore()
	fmt.Println("Player:", state.Player)
	fmt.Println("Dealer:", state.Dealer)
	fmt.Println(DetermineWinner(state.Player, state.Dealer))
	return state
}
