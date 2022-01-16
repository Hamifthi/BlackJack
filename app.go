package Blackjack

import (
	"github.com/Hamifthi/deck"
	"strings"
)

type Player struct {
	Cards      []deck.Card
	MinScore   int
	Score      int
	HiddenCard deck.Card
}

func (p Player) String() string {
	strs := make([]string, len(p.Cards))
	for index, card := range p.Cards {
		strs[index] = card.String()
	}
	return strings.Join(strs, ", ")
}

func (p *Player) FlipHiddenCard() {
	if p.HiddenCard != *new(deck.Card) {
		p.Cards = append(p.Cards, p.HiddenCard)
	}
	p.HiddenCard = *new(deck.Card)
}

func (p *Player) CalculateMinScore() {
	p.MinScore = 0
	for _, card := range p.Cards {
		var score int
		switch card.Rank {
		case deck.Jack, deck.Queen, deck.King:
			score = 10
		default:
			score = int(card.Rank)
		}
		p.MinScore += score
	}
}

func (p *Player) CalculateScore() {
	p.CalculateMinScore()
	if p.MinScore > 11 {
		p.Score = p.MinScore
	} else {
		for _, card := range p.Cards {
			if card.Rank == deck.Ace {
				p.MinScore = p.MinScore + 10
			}
		}
		p.Score = p.MinScore
	}
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
	StateNull
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Player
	Dealer Player
}

func (gs *GameState) CurrentPlayer() *Player {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("It's in invalid state.")
	}
}
